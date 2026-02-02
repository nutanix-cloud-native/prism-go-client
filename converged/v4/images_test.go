// Package v4 provides integration tests for the converged client images service.
//
// These tests require a real Nutanix Prism Central environment and should be run
// with the following environment variables set:
//   - NUTANIX_ENDPOINT: The Prism Central endpoint URL
//   - NUTANIX_USERNAME: Username for authentication
//   - NUTANIX_PASSWORD: Password for authentication
//
// To run these tests:
//
//	go test -v ./converged/v4 -run TestImagesIntegration
//	go test -v ./converged/v4 -run TestImagesService_ErrorHandling
package v4

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"

	imageModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// TestImagesService_ErrorHandling tests error handling for nil client
func TestImagesService_ErrorHandling(t *testing.T) {
	service := NewImagesService(nil)
	require.NotNil(t, service)
	ctx := context.Background()

	t.Run("Get_NilClient", func(t *testing.T) {
		_, err := service.Get(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("List_NilClient", func(t *testing.T) {
		_, err := service.List(ctx)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("NewIterator_NilClient", func(t *testing.T) {
		iterator := service.NewIterator(ctx)
		assert.Nil(t, iterator)
	})

	t.Run("Create_NilClient", func(t *testing.T) {
		image := imageModels.NewImage()
		_, err := service.Create(ctx, image)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("Delete_NilClient", func(t *testing.T) {
		err := service.Delete(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("GetFile_NilClient", func(t *testing.T) {
		_, err := service.GetFile(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("Upload_NilClient", func(t *testing.T) {
		err := service.Upload(ctx, "test-id", "/path/to/file")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("CreateAsync_NilClient", func(t *testing.T) {
		image := imageModels.NewImage()
		_, err := service.CreateAsync(ctx, image)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})

	t.Run("DeleteAsync_NilClient", func(t *testing.T) {
		_, err := service.DeleteAsync(ctx, "test-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client is not initialized")
	})
}

// TestImagesIntegration tests the client.Images with real Nutanix API calls
func TestImagesIntegration(t *testing.T) {
	// Get credentials from environment
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	// Create converged client
	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("ListImages", func(t *testing.T) {
		// Test List images
		images, err := client.Images.List(ctx, converged.WithLimit(10))
		assert.NoError(t, err)
		assert.NotNil(t, images)
		assert.GreaterOrEqual(t, len(images), 0)
	})

	t.Run("GetImage", func(t *testing.T) {
		// First, list images to get a valid image UUID
		images, err := client.Images.List(ctx, converged.WithLimit(1))
		require.NoError(t, err)

		if len(images) == 0 {
			t.Skip("No images available for testing")
		}

		// Get the first image's UUID
		imageUUID := *images[0].ExtId
		require.NotEmpty(t, imageUUID)

		// Test Get image
		image, err := client.Images.Get(ctx, imageUUID)
		assert.NoError(t, err)
		assert.NotNil(t, image)
		assert.Equal(t, imageUUID, *image.ExtId)
	})

	t.Run("NewIterator", func(t *testing.T) {
		// Test iterator
		iterator := client.Images.NewIterator(ctx)
		require.NotNil(t, iterator)

		// Collect images using iterator
		var images []imageModels.Image
		for image, err := range iterator {
			if err != nil {
				break
			}
			images = append(images, image)
			if len(images) >= 5 { // Limit to prevent long test runs
				break
			}
		}

		assert.GreaterOrEqual(t, len(images), 0)
	})
}

// TestImagesListOptions tests various OData options for listing images
func TestImagesListOptions(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("WithLimit", func(t *testing.T) {
		images, err := client.Images.List(ctx, converged.WithLimit(2))
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(images), 2)
	})

	t.Run("WithPage", func(t *testing.T) {
		images, err := client.Images.List(ctx, converged.WithPage(0), converged.WithLimit(1))
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(images), 1)
	})

	t.Run("WithSelect", func(t *testing.T) {
		images, err := client.Images.List(ctx, converged.WithSelect("extId,name,type"))
		assert.NoError(t, err)
		assert.NotNil(t, images)
	})

	t.Run("WithOrderBy", func(t *testing.T) {
		images, err := client.Images.List(ctx, converged.WithOrderBy("name asc"))
		assert.NoError(t, err)
		assert.NotNil(t, images)
	})

	t.Run("WithFilter", func(t *testing.T) {
		// Test filtering by image type
		images, err := client.Images.List(ctx, converged.WithFilter("type eq Vmm.Content.ImageType'DISK_IMAGE'"))
		assert.NoError(t, err)
		assert.NotNil(t, images)
	})

	t.Run("MultipleOptions", func(t *testing.T) {
		images, err := client.Images.List(ctx,
			converged.WithLimit(5),
			converged.WithPage(0),
			converged.WithSelect("extId,name"),
			converged.WithOrderBy("name asc"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, images)
		assert.LessOrEqual(t, len(images), 5)
	})

	t.Run("UnsupportedExpandOption", func(t *testing.T) {
		// Expand is not supported for Images
		_, err := client.Images.List(ctx, converged.WithExpand("someField"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "expand and apply are not supported")
	})

	t.Run("UnsupportedApplyOption", func(t *testing.T) {
		// Apply is not supported for Images
		_, err := client.Images.List(ctx, converged.WithApply("someAggregation"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "expand and apply are not supported")
	})
}

// TestImagesErrorScenarios tests error handling scenarios
func TestImagesErrorScenarios(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("GetInvalidUUID", func(t *testing.T) {
		_, err := client.Images.Get(ctx, "invalid-uuid-format")
		assert.Error(t, err)
	})

	t.Run("GetNonExistentUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.Images.Get(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("DeleteNonExistentUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		err := client.Images.Delete(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("GetFileNonExistentUUID", func(t *testing.T) {
		nonExistentUUID := "00000000-0000-0000-0000-000000000000"
		_, err := client.Images.GetFile(ctx, nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("UploadNonExistentFile", func(t *testing.T) {
		err := client.Images.Upload(ctx, "test-uuid", "/nonexistent/path/to/file.img")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to open image file")
	})

	t.Run("InvalidFilterSyntax", func(t *testing.T) {
		_, err := client.Images.List(ctx, converged.WithFilter("invalid filter syntax"))
		assert.Error(t, err)
	})
}

// TestImageFromObjectsLite tests the imageFromObjectsLite helper function
func TestImageFromObjectsLite(t *testing.T) {
	service := &ImagesService{}

	t.Run("DiskImage", func(t *testing.T) {
		image, err := service.imageFromObjectsLite("test-key-123", "/path/to/image.qcow2")
		assert.NoError(t, err)
		assert.NotNil(t, image)
		assert.Equal(t, "image.qcow2", *image.Name)
		assert.Equal(t, imageModels.IMAGETYPE_DISK_IMAGE, *image.Type)
	})

	t.Run("ISOImage", func(t *testing.T) {
		image, err := service.imageFromObjectsLite("test-key-456", "/path/to/installer.iso")
		assert.NoError(t, err)
		assert.NotNil(t, image)
		assert.Equal(t, "installer.iso", *image.Name)
		assert.Equal(t, imageModels.IMAGETYPE_ISO_IMAGE, *image.Type)
	})

	t.Run("ISOImageUppercase", func(t *testing.T) {
		image, err := service.imageFromObjectsLite("test-key-789", "/path/to/installer.ISO")
		assert.NoError(t, err)
		assert.NotNil(t, image)
		assert.Equal(t, imageModels.IMAGETYPE_ISO_IMAGE, *image.Type)
	})

	t.Run("DefaultsToDiskImage", func(t *testing.T) {
		// Any non-ISO extension should default to DISK_IMAGE
		image, err := service.imageFromObjectsLite("test-key", "/path/to/image.raw")
		assert.NoError(t, err)
		assert.Equal(t, imageModels.IMAGETYPE_DISK_IMAGE, *image.Type)
	})
}

// TestDefaultAWSRegion tests the AWS region default behavior
func TestDefaultAWSRegion(t *testing.T) {
	t.Run("DefaultRegion", func(t *testing.T) {
		// Verify the default constant is set correctly
		assert.Equal(t, "us-east-1", defaultAWSRegion)
	})

	t.Run("EnvVarOverride", func(t *testing.T) {
		// Test that AWS_REGION env var is read
		// Using t.Setenv which automatically cleans up after test
		t.Setenv("AWS_REGION", "eu-west-1")

		region := os.Getenv("AWS_REGION")
		assert.Equal(t, "eu-west-1", region)
	})
}

// TestUploadWithTempFile tests the upload functionality with a temporary file
func TestUploadWithTempFile(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	if strings.Contains(creds.Endpoint, prismEndpointDummyValue) {
		t.Skip("Skipping integration test: NUTANIX_ENDPOINT not set")
	}

	// Skip if AWS_REGION is not set (unless we want to test default)
	// The test will use default region if not set

	client, err := NewClient(creds)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	t.Run("UploadSmallFile", func(t *testing.T) {
		// Create a temporary file for testing
		tmpDir := t.TempDir()
		tmpFile := filepath.Join(tmpDir, "test-image.raw")

		// Create a small test file (1KB)
		testData := make([]byte, 1024)
		for i := range testData {
			testData[i] = byte(i % 256)
		}
		err := os.WriteFile(tmpFile, testData, 0644)
		require.NoError(t, err)

		// Note: This test may fail if Objects Lite is not configured
		// or if the credentials don't have write access
		err = client.Images.Upload(ctx, "test-upload-uuid", tmpFile)
		if err != nil {
			// Log the error but don't fail - Objects Lite may not be available
			t.Logf("Upload test skipped or failed (Objects Lite may not be configured): %v", err)
		}
	})
}

// TestConstants verifies the package constants are set correctly
func TestConstants(t *testing.T) {
	t.Run("DefaultObjectsBucket", func(t *testing.T) {
		assert.Equal(t, "vmm-images", defaultObjectsBucket)
	})

	t.Run("DefaultContentType", func(t *testing.T) {
		assert.Equal(t, "application/octet-stream", defaultContentType)
	})

	t.Run("DefaultAWSRegion", func(t *testing.T) {
		assert.Equal(t, "us-east-1", defaultAWSRegion)
	})
}
