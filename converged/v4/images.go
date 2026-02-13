package v4

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	vmmPrismModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	imageModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

const (
	// defaultObjectsBucket is the default S3 bucket name for VMM images
	defaultObjectsBucket = "vmm-images"
	// defaultContentType is the default content type for image uploads
	defaultContentType = "application/octet-stream"
	// defaultAWSRegion is the default AWS region for Objects Lite
	defaultAWSRegion = "us-east-1"
)

// ImagesService provides implementation for all Images interface methods.
type ImagesService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewImagesService creates a new ImagesService instance.
func NewImagesService(client *v4prismGoClient.Client) *ImagesService {
	return &ImagesService{client: client, entitiesName: "image"}
}

// Get returns the image for the given UUID.
func (s *ImagesService) Get(ctx context.Context, uuid string) (*imageModels.Image, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*imageModels.GetImageApiResponse, imageModels.Image](
		func() (*imageModels.GetImageApiResponse, error) {
			return s.client.ImagesApiInstance.GetImageById(&uuid)
		},
		s.entitiesName,
	)
}

// List returns a list of images.
func (s *ImagesService) List(ctx context.Context, opts ...converged.ODataOption) ([]imageModels.Image, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams.Expand != nil || myParams.Apply != nil {
		return nil, fmt.Errorf("expand and apply are not supported for listing Images")
	}

	return GenericListEntities[*imageModels.ListImagesApiResponse, imageModels.Image](
		func(reqParams *V4ODataParams) (*imageModels.ListImagesApiResponse, error) {
			return s.client.ImagesApiInstance.ListImages(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing images.
func (s *ImagesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[imageModels.Image] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*imageModels.ListImagesApiResponse, imageModels.Image](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*imageModels.ListImagesApiResponse, error) {
			return s.client.ImagesApiInstance.ListImages(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// Create creates a new image.
func (s *ImagesService) Create(ctx context.Context, image *imageModels.Image) (*imageModels.Image, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	operation, err := s.CreateAsync(ctx, image)
	if err != nil {
		return nil, fmt.Errorf("failed to create image: %w", err)
	}

	result, err := operation.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create image: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("failed to create image: operation completed but no image returned")
	}

	return result[0], nil
}

// Delete deletes an existing image.
func (s *ImagesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}
	operation, err := s.DeleteAsync(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete image: %w", err)
	}

	_, err = operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete image: %w", err)
	}

	return nil
}

// GetFile downloads the image file for the given UUID.
func (s *ImagesService) GetFile(ctx context.Context, uuid string) (*imageModels.FileDetail, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*imageModels.GetImageFileApiResponse, imageModels.FileDetail](
		s.client.ImagesApiInstance.GetFileByImageId(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get image file: %w", err)
	}

	return &result, nil
}

// Upload uploads the image file to the given UUID.
func (s *ImagesService) Upload(ctx context.Context, uuid, filepath string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open image file %q: %w", filepath, err)
	}
	defer func() { _ = file.Close() }()

	if err := s.uploadToObjects(ctx, uuid, file); err != nil {
		return err
	}

	image, err := s.imageFromObjectsLite(uuid, filepath)
	if err != nil {
		return err
	}

	if _, err := s.Create(ctx, image); err != nil {
		return fmt.Errorf("failed to create image from Objects: %w", err)
	}

	return nil
}

// uploadToObjects uploads a file to Objects Lite S3 storage
func (s *ImagesService) uploadToObjects(ctx context.Context, uuid string, file *os.File) error {
	awsConfig, endpoint, err := s.awsConfig(ctx)
	if err != nil {
		return err
	}

	s3Client := s3.NewFromConfig(awsConfig, func(options *s3.Options) {
		options.UsePathStyle = true
		options.BaseEndpoint = aws.String(endpoint)
	})

	uploader := manager.NewUploader(s3Client)
	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(defaultObjectsBucket),
		Key:         aws.String(uuid),
		Body:        file,
		ContentType: aws.String(defaultContentType),
	})
	if err != nil {
		return fmt.Errorf("failed to upload image file to Objects: %w", err)
	}

	return nil
}

func (s *ImagesService) imageFromObjectsLite(objectKey, sourcePath string) (*imageModels.Image, error) {
	name := filepath.Base(sourcePath)
	imageType := imageModels.IMAGETYPE_DISK_IMAGE
	if strings.EqualFold(filepath.Ext(sourcePath), ".iso") {
		imageType = imageModels.IMAGETYPE_ISO_IMAGE
	}

	objectsSource := imageModels.NewObjectsLiteSource()
	objectsSource.Key = &objectKey
	source := imageModels.NewOneOfImageSource()
	if err := source.SetValue(*objectsSource); err != nil {
		return nil, fmt.Errorf("failed to set Objects Lite source: %w", err)
	}

	image := imageModels.NewImage()
	image.Name = &name
	image.Type = imageType.Ref()
	image.Source = source
	return image, nil
}

func (s *ImagesService) awsConfig(ctx context.Context) (aws.Config, string, error) {
	apiClient := s.client.ImagesApiInstance.ApiClient

	// Use ApiClient's host and port for the endpoint
	endpoint := fmt.Sprintf("https://%s:%d", apiClient.Host, apiClient.Port)

	// Use ApiClient's username and password
	username := strings.TrimSpace(apiClient.Username)
	password := strings.TrimSpace(apiClient.Password)
	if username == "" || password == "" {
		return aws.Config{}, "", fmt.Errorf("username and password are required for Objects Lite auth")
	}

	// AWS region from environment variable, defaults to us-east-1
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = defaultAWSRegion
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	loadOptions := []func(*config.LoadOptions) error{
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(encoded, encoded, "")),
	}
	if !apiClient.VerifySSL {
		loadOptions = append(loadOptions, config.WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}))
	}

	awsConfig, err := config.LoadDefaultConfig(ctx, loadOptions...)
	if err != nil {
		return aws.Config{}, "", fmt.Errorf("failed to load AWS config: %w", err)
	}
	return awsConfig, endpoint, nil
}

// CreateAsync creates a new image asynchronously.
func (s *ImagesService) CreateAsync(ctx context.Context, image *imageModels.Image) (converged.Operation[imageModels.Image], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	s.client.ImagesApiInstance.ApiClient.AddDefaultHeader("Ntnx-Request-Id", uuid.NewString())
	taskRef, err := CallAPI[*imageModels.CreateImageApiResponse, vmmPrismModels.TaskReference](
		s.client.ImagesApiInstance.CreateImage(image),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create image: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created image")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}

// DeleteAsync deletes an existing image asynchronously.
func (s *ImagesService) DeleteAsync(ctx context.Context, uuid string) (converged.Operation[converged.NoEntity], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	taskRef, err := CallAPI[*imageModels.DeleteImageApiResponse, vmmPrismModels.TaskReference](
		s.client.ImagesApiInstance.DeleteImageById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete image: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for deleted image")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		func(ctx context.Context, uuid string) (*converged.NoEntity, error) {
			return converged.NoEntityGetter(ctx, uuid)
		},
	), nil
}
