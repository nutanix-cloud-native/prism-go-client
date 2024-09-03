package v3

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nutanix-cloud-native/prism-go-client/internal"
	"github.com/nutanix-cloud-native/prism-go-client/utils"
	"github.com/nutanix-cloud-native/prism-go-client/v3/models"
)

// Operations ...
type Operations struct {
	client *internal.Client
}

// Service ...
type Service interface {
	CreateVM(ctx context.Context, createRequest *models.VmIntentInput) (*models.VmIntentResponse, error)
	DeleteVM(ctx context.Context, uuid string) (*models.VmIntentResponse, error)
	GetVM(ctx context.Context, uuid string) (*models.VmIntentResponse, error)
	ListVM(ctx context.Context, getEntitiesRequest *models.VmListMetadata) (*models.VmListIntentResponse, error)
	UpdateVM(ctx context.Context, uuid string, body *models.VmIntentInput) (*models.VmIntentResponse, error)
	CreateSubnet(ctx context.Context, createRequest *models.SubnetIntentInput) (*models.SubnetIntentResponse, error)
	DeleteSubnet(ctx context.Context, uuid string) (*models.SubnetIntentResponse, error)
	GetSubnet(ctx context.Context, uuid string) (*models.SubnetIntentResponse, error)
	ListSubnet(ctx context.Context, getEntitiesRequest *models.SubnetListMetadata) (*models.SubnetListIntentResponse, error)
	UpdateSubnet(ctx context.Context, uuid string, body *models.SubnetIntentInput) (*models.SubnetIntentResponse, error)
	CreateImage(ctx context.Context, createRequest *models.ImageIntentInput) (*models.ImageIntentResponse, error)
	DeleteImage(ctx context.Context, uuid string) (*models.ImageIntentResponse, error)
	GetImage(ctx context.Context, uuid string) (*models.ImageIntentResponse, error)
	ListImage(ctx context.Context, getEntitiesRequest *models.ImageListMetadata) (*models.ImageListIntentResponse, error)
	UpdateImage(ctx context.Context, uuid string, body *models.ImageIntentInput) (*models.ImageIntentResponse, error)
	UploadImage(ctx context.Context, uuid, filepath string) error
	CreateOrUpdateCategoryKey(ctx context.Context, body *models.CategoryKey) (*models.CategoryKeyStatus, error)
	ListCategories(ctx context.Context, getEntitiesRequest *models.CategoryListMetadata) (*models.CategoryKeyListResponse, error)
	DeleteCategoryKey(ctx context.Context, name string) error
	GetCategoryKey(ctx context.Context, name string) (*models.CategoryKeyStatus, error)
	ListCategoryValues(ctx context.Context, name string, getEntitiesRequest *models.CategoryListMetadata) (*models.CategoryValueListResponse, error)
	CreateOrUpdateCategoryValue(ctx context.Context, name string, body *models.CategoryValue) (*models.CategoryValueStatus, error)
	GetCategoryValue(ctx context.Context, name string, value string) (*models.CategoryValueStatus, error)
	DeleteCategoryValue(ctx context.Context, name string, value string) error
	GetCategoryQuery(ctx context.Context, query *models.CategoryQueryInput) (*models.CategoryQueryResponse, error)
	UpdateNetworkSecurityRule(ctx context.Context, uuid string, body *models.NetworkSecurityRuleIntentInput) (*models.NetworkSecurityRuleIntentResponse, error)
	ListNetworkSecurityRule(ctx context.Context, getEntitiesRequest *models.NetworkSecurityRuleListMetadata) (*models.NetworkSecurityRuleListIntentResponse, error)
	GetNetworkSecurityRule(ctx context.Context, uuid string) (*models.NetworkSecurityRuleIntentResponse, error)
	DeleteNetworkSecurityRule(ctx context.Context, uuid string) (*models.NetworkSecurityRuleIntentResponse, error)
	CreateNetworkSecurityRule(ctx context.Context, request *models.NetworkSecurityRuleIntentInput) (*models.NetworkSecurityRuleIntentResponse, error)
	ListCluster(ctx context.Context, getEntitiesRequest *models.ClusterListMetadata) (*models.ClusterListIntentResponse, error)
	GetCluster(ctx context.Context, uuid string) (*models.ClusterIntentResponse, error)
	UpdateVolumeGroup(ctx context.Context, uuid string, body *models.VolumeGroupIntentInput) (*models.VolumeGroupIntentResponse, error)
	ListVolumeGroup(ctx context.Context, getEntitiesRequest *models.VolumeGroupListMetadata) (*models.VolumeGroupListIntentResponse, error)
	GetVolumeGroup(ctx context.Context, uuid string) (*models.VolumeGroupIntentResponse, error)
	DeleteVolumeGroup(ctx context.Context, uuid string) error
	CreateVolumeGroup(ctx context.Context, request *models.VolumeGroupIntentInput) (*models.VolumeGroupIntentResponse, error)
	ListAllVM(ctx context.Context, filter string) (*models.VmListIntentResponse, error)
	ListAllSubnet(ctx context.Context, filter string) (*models.SubnetListIntentResponse, error)
	ListAllNetworkSecurityRule(ctx context.Context, filter string) (*models.NetworkSecurityRuleListIntentResponse, error)
	ListAllImage(ctx context.Context, filter string) (*models.ImageListIntentResponse, error)
	ListAllCluster(ctx context.Context, filter string) (*models.ClusterListIntentResponse, error)
	ListAllCategoryValues(ctx context.Context, categoryName, filter string) (*models.CategoryValueListResponse, error)
	GetTask(ctx context.Context, taskUUID string) (*models.Task, error)
	GetHost(ctx context.Context, taskUUID string) (*models.HostIntentResponse, error)
	ListHost(ctx context.Context, getEntitiesRequest *models.HostListMetadata) (*models.HostListIntentResponse, error)
	ListAllHost(ctx context.Context) (*models.HostListIntentResponse, error)
	CreateProject(ctx context.Context, request *models.ProjectIntentInput) (*models.ProjectIntentResponse, error)
	GetProject(ctx context.Context, projectUUID string) (*models.ProjectIntentResponse, error)
	ListProject(ctx context.Context, getEntitiesRequest *models.ProjectListMetadata) (*models.ProjectListIntentResponse, error)
	ListAllProject(ctx context.Context, filter string) (*models.ProjectListIntentResponse, error)
	UpdateProject(ctx context.Context, uuid string, body *models.ProjectIntentInput) (*models.ProjectIntentResponse, error)
	DeleteProject(ctx context.Context, uuid string) (*models.ProjectIntentResponse, error)
	CreateAccessControlPolicy(ctx context.Context, request *models.AccessControlPolicyIntentInput) (*models.AccessControlPolicyIntentResponse, error)
	GetAccessControlPolicy(ctx context.Context, accessControlPolicyUUID string) (*models.AccessControlPolicyIntentResponse, error)
	ListAccessControlPolicy(ctx context.Context, getEntitiesRequest *models.AccessControlPolicyListMetadata) (*models.AccessControlPolicyListIntentResponse, error)
	ListAllAccessControlPolicy(ctx context.Context, filter string) (*models.AccessControlPolicyListIntentResponse, error)
	UpdateAccessControlPolicy(ctx context.Context, uuid string, body *models.AccessControlPolicyIntentInput) (*models.AccessControlPolicyIntentResponse, error)
	DeleteAccessControlPolicy(ctx context.Context, uuid string) (*models.AccessControlPolicyIntentResponse, error)
	CreateRole(ctx context.Context, request *models.RoleIntentInput) (*models.RoleIntentResponse, error)
	GetRole(ctx context.Context, uuid string) (*models.RoleIntentResponse, error)
	ListRole(ctx context.Context, getEntitiesRequest *models.RoleListMetadata) (*models.RoleListIntentResponse, error)
	ListAllRole(ctx context.Context, filter string) (*models.RoleListIntentResponse, error)
	UpdateRole(ctx context.Context, uuid string, body *models.RoleIntentInput) (*models.RoleIntentResponse, error)
	DeleteRole(ctx context.Context, uuid string) (*models.RoleIntentResponse, error)
	CreateUser(ctx context.Context, request *models.UserIntentInput) (*models.UserIntentResponse, error)
	GetUser(ctx context.Context, userUUID string) (*models.UserIntentResponse, error)
	UpdateUser(ctx context.Context, uuid string, body *models.UserIntentInput) (*models.UserIntentResponse, error)
	DeleteUser(ctx context.Context, uuid string) (*models.UserIntentResponse, error)
	ListUser(ctx context.Context, getEntitiesRequest *models.UserListMetadata) (*models.UserListIntentResponse, error)
	ListAllUser(ctx context.Context, filter string) (*models.UserListIntentResponse, error)
	GetCurrentLoggedInUser(ctx context.Context) (*models.UserIntentResponse, error)
	GetUserGroup(ctx context.Context, userUUID string) (*models.UserGroupIntentResponse, error)
	ListUserGroup(ctx context.Context, getEntitiesRequest *models.UserGroupListMetadata) (*models.UserGroupListIntentResponse, error)
	ListAllUserGroup(ctx context.Context, filter string) (*models.UserGroupListIntentResponse, error)
	GetPermission(ctx context.Context, permissionUUID string) (*models.PermissionIntentResponse, error)
	ListPermission(ctx context.Context, getEntitiesRequest *models.PermissionListMetadata) (*models.PermissionListIntentResponse, error)
	ListAllPermission(ctx context.Context, filter string) (*models.PermissionListIntentResponse, error)
	GetProtectionRule(ctx context.Context, uuid string) (*models.ProtectionRuleIntentResponse, error)
	ListProtectionRules(ctx context.Context, getEntitiesRequest *models.ProtectionRuleListMetadata) (*models.ProtectionRuleListIntentResponse, error)
	ListAllProtectionRules(ctx context.Context, filter string) (*models.ProtectionRuleListIntentResponse, error)
	CreateProtectionRule(ctx context.Context, request *models.ProtectionRuleIntentInput) (*models.ProtectionRuleIntentResponse, error)
	UpdateProtectionRule(ctx context.Context, uuid string, body *models.ProtectionRuleIntentInput) (*models.ProtectionRuleIntentResponse, error)
	DeleteProtectionRule(ctx context.Context, uuid string) (*models.ProtectionRuleIntentResponse, error)
	ProcessProtectionRule(ctx context.Context, uuid string) error
	GetRecoveryPlan(ctx context.Context, uuid string) (*models.RecoveryPlanIntentResponse, error)
	ListRecoveryPlans(ctx context.Context, getEntitiesRequest *models.RecoveryPlanListMetadata) (*models.RecoveryPlanListIntentResponse, error)
	ListAllRecoveryPlans(ctx context.Context, filter string) (*models.RecoveryPlanListIntentResponse, error)
	CreateRecoveryPlan(ctx context.Context, request *models.RecoveryPlanIntentInput) (*models.RecoveryPlanIntentResponse, error)
	UpdateRecoveryPlan(ctx context.Context, uuid string, body *models.RecoveryPlanIntentInput) (*models.RecoveryPlanIntentResponse, error)
	DeleteRecoveryPlan(ctx context.Context, uuid string) (*models.RecoveryPlanIntentResponse, error)
	GetServiceGroup(ctx context.Context, uuid string) (*models.ServiceGroupResponse, error)
	ListAllServiceGroups(ctx context.Context, filter string) (*models.ServiceGroupListResponse, error)
	CreateServiceGroup(ctx context.Context, request *models.ServiceGroup) (*models.Reference, error)
	UpdateServiceGroup(ctx context.Context, uuid string, body *models.ServiceGroup) error
	DeleteServiceGroup(ctx context.Context, uuid string) error
	GetAddressGroup(ctx context.Context, uuid string) (*models.AddressGroupResponse, error)
	ListAddressGroups(ctx context.Context, getEntitiesRequest *models.AddressGroupListMetadata) (*models.AddressGroupListResponse, error)
	ListAllAddressGroups(ctx context.Context, filter string) (*models.AddressGroupListResponse, error)
	DeleteAddressGroup(ctx context.Context, uuid string) error
	CreateAddressGroup(ctx context.Context, request *models.AddressGroup) (*models.Reference, error)
	UpdateAddressGroup(ctx context.Context, uuid string, body *models.AddressGroup) error
	GetRecoveryPlanJob(ctx context.Context, uuid string) (*models.RecoveryPlanJobIntentResponse, error)
	GetRecoveryPlanJobStatus(ctx context.Context, uuid string, status string) (*models.RecoveryPlanJobExecutionStatus, error)
	ListRecoveryPlanJobs(ctx context.Context, getEntitiesRequest *models.RecoveryPlanListMetadata) (*models.RecoveryPlanJobListIntentResponse, error)
	DeleteRecoveryPlanJob(ctx context.Context, uuid string) error
	CreateRecoveryPlanJob(ctx context.Context, request *models.RecoveryPlanJobIntentInput) (*models.RecoveryPlanJobResponse, error)
	PerformRecoveryPlanJobAction(ctx context.Context, uuid string, action string, request *models.RecoveryPlanJobActionRequest) (*models.RecoveryPlanJobResponse, error)
	GroupsGetEntities(ctx context.Context, request *models.GroupsGetEntitiesRequest) (*models.GroupsGetEntitiesResponse, error)
	GetAvailabilityZone(ctx context.Context, uuid string) (*models.AvailabilityZoneIntentResponse, error)
	GetPrismCentral(ctx context.Context) (*models.PrismCentral, error)
}

var _ Service = &Operations{}

/*CreateVM Creates a VM
 * This operation submits a request to create a VM based on the input parameters.
 *
 * @param body
 * @return *VMIntentResponse
 */
func (op Operations) CreateVM(ctx context.Context, createRequest *models.VmIntentInput) (*models.VmIntentResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/vms", createRequest)
	if err != nil {
		return nil, err
	}

	vmIntentResponse := new(models.VmIntentResponse)
	if err := op.client.Do(ctx, req, vmIntentResponse); err != nil {
		return nil, err
	}

	return vmIntentResponse, nil
}

/*DeleteVM Deletes a VM
 * This operation submits a request to delete a op.
 *
 * @param uuid The uuid of the entity.
 * @return error
 */
func (op Operations) DeleteVM(ctx context.Context, uuid string) (*models.VmIntentResponse, error) {
	path := fmt.Sprintf("/vms/%s", uuid)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	deleteResponse := new(models.VmIntentResponse)
	if err := op.client.Do(ctx, req, deleteResponse); err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

/*GetVM Gets a VM
 * This operation gets a op.
 *
 * @param uuid The uuid of the entity.
 * @return *VMIntentResponse
 */
func (op Operations) GetVM(ctx context.Context, uuid string) (*models.VmIntentResponse, error) {
	path := fmt.Sprintf("/vms/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	vmIntentResponse := new(models.VmIntentResponse)
	if err := op.client.Do(ctx, req, vmIntentResponse); err != nil {
		return nil, err
	}

	return vmIntentResponse, nil
}

/*ListVM Get a list of VMs This operation gets a list of VMs, allowing for sorting and pagination. Note: Entities that have not been created
 * successfully are not listed.
 *
 * @param getEntitiesRequest @return *VmListIntentResponse
 */
func (op Operations) ListVM(ctx context.Context, getEntitiesRequest *models.VmListMetadata) (*models.VmListIntentResponse, error) {
	path := "/vms/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	vmListIntentResponse := new(models.VmListIntentResponse)
	if err := op.client.Do(ctx, req, vmListIntentResponse); err != nil {
		return nil, err
	}

	return vmListIntentResponse, nil
}

/*UpdateVM Updates a VM
 * This operation submits a request to update a VM based on the input parameters.
 *
 * @param uuid The uuid of the entity.
 * @param body
 * @return *VMIntentResponse
 */
func (op Operations) UpdateVM(ctx context.Context, uuid string, body *models.VmIntentInput) (*models.VmIntentResponse, error) {
	path := fmt.Sprintf("/vms/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	vmIntentResponse := new(models.VmIntentResponse)
	if err := op.client.Do(ctx, req, vmIntentResponse); err != nil {
		return nil, err
	}

	return vmIntentResponse, nil
}

/*CreateSubnet Creates a subnet
 * This operation submits a request to create a subnet based on the input parameters. A subnet is a block of IP addresses.
 *
 * @param body
 * @return *SubnetIntentResponse
 */
func (op Operations) CreateSubnet(ctx context.Context, createRequest *models.SubnetIntentInput) (*models.SubnetIntentResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/subnets", createRequest)
	if err != nil {
		return nil, err
	}

	subnetIntentResponse := new(models.SubnetIntentResponse)
	if err := op.client.Do(ctx, req, subnetIntentResponse); err != nil {
		return nil, err
	}

	return subnetIntentResponse, nil
}

/*DeleteSubnet Deletes a subnet
 * This operation submits a request to delete a subnet.
 *
 * @param uuid The uuid of the entity.
 * @return error if exist error
 */
func (op Operations) DeleteSubnet(ctx context.Context, uuid string) (*models.SubnetIntentResponse, error) {
	path := fmt.Sprintf("/subnets/%s", uuid)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	deleteResponse := new(models.SubnetIntentResponse)
	if err := op.client.Do(ctx, req, deleteResponse); err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

/*GetSubnet Gets a subnet entity
 * This operation gets a subnet.
 *
 * @param uuid The uuid of the entity.
 * @return *SubnetIntentResponse
 */
func (op Operations) GetSubnet(ctx context.Context, uuid string) (*models.SubnetIntentResponse, error) {
	path := fmt.Sprintf("/subnets/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	subnetIntentResponse := new(models.SubnetIntentResponse)
	if err := op.client.Do(ctx, req, subnetIntentResponse); err != nil {
		return nil, err
	}

	return subnetIntentResponse, nil
}

/*ListSubnet Gets a list of subnets This operation gets a list of subnets, allowing for sorting and pagination. Note: Entities that have not
 * been created successfully are not listed.
 *
 * @param getEntitiesRequest @return *SubnetListIntentResponse
 */
func (op Operations) ListSubnet(ctx context.Context, getEntitiesRequest *models.SubnetListMetadata) (*models.SubnetListIntentResponse, error) {
	path := "/subnets/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	subnetListIntentResponse := new(models.SubnetListIntentResponse)
	if err := op.client.Do(ctx, req, subnetListIntentResponse); err != nil {
		return nil, err
	}

	return subnetListIntentResponse, nil
}

/*UpdateSubnet Updates a subnet
 * This operation submits a request to update a subnet based on the input parameters.
 *
 * @param uuid The uuid of the entity.
 * @param body
 * @return *SubnetIntentResponse
 */
func (op Operations) UpdateSubnet(ctx context.Context, uuid string, body *models.SubnetIntentInput) (*models.SubnetIntentResponse, error) {
	path := fmt.Sprintf("/subnets/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	subnetIntentResponse := new(models.SubnetIntentResponse)
	if err := op.client.Do(ctx, req, subnetIntentResponse); err != nil {
		return nil, err
	}

	return subnetIntentResponse, nil
}

/*CreateImage Creates a IMAGE Images are raw ISO, QCOW2, or VMDK files that are uploaded by a user can be attached to a op. An ISO image is
 * attached as a virtual CD-ROM drive, and QCOW2 and VMDK files are attached as SCSI disks. An image has to be explicitly added to the
 * self-service catalog before users can create VMs from it.
 *
 * @param body @return *ImageIntentResponse
 */
func (op Operations) CreateImage(ctx context.Context, body *models.ImageIntentInput) (*models.ImageIntentResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/images", body)
	if err != nil {
		return nil, err
	}

	imageIntentResponse := new(models.ImageIntentResponse)
	if err := op.client.Do(ctx, req, imageIntentResponse); err != nil {
		return nil, err
	}

	return imageIntentResponse, nil
}

/*UploadImage Uplloads a Image Binary file Images are raw ISO, QCOW2, or VMDK files that are uploaded by a user can be attached to a op. An
 * ISO image is attached as a virtual CD-ROM drive, and QCOW2 and VMDK files are attached as SCSI disks. An image has to be explicitly added
 * to the self-service catalog before users can create VMs from it.
 *
 * @param uuid @param filepath
 */
func (op Operations) UploadImage(ctx context.Context, uuid, filepath string) error {
	path := fmt.Sprintf("/images/%s/file", uuid)
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("error: cannot open file: %s", err)
	}

	defer file.Close()

	req, err := op.client.NewUploadRequest(http.MethodPut, path, file)
	if err != nil {
		return fmt.Errorf("error: Creating request %s", err)
	}

	return op.client.Do(ctx, req, nil)
}

/*DeleteImage deletes a IMAGE
 * This operation submits a request to delete a IMAGE.
 *
 * @param uuid The uuid of the entity.
 * @return error if error exists
 */
func (op Operations) DeleteImage(ctx context.Context, uuid string) (*models.ImageIntentResponse, error) {
	path := fmt.Sprintf("/images/%s", uuid)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	deleteResponse := new(models.ImageIntentResponse)
	if err := op.client.Do(ctx, req, deleteResponse); err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

/*GetImage gets a IMAGE
 * This operation gets a IMAGE.
 *
 * @param uuid The uuid of the entity.
 * @return *ImageIntentResponse
 */
func (op Operations) GetImage(ctx context.Context, uuid string) (*models.ImageIntentResponse, error) {
	path := fmt.Sprintf("/images/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	imageIntentResponse := new(models.ImageIntentResponse)
	if err := op.client.Do(ctx, req, imageIntentResponse); err != nil {
		return nil, err
	}

	return imageIntentResponse, nil
}

/*ListImage gets a list of IMAGEs This operation gets a list of IMAGEs, allowing for sorting and pagination. Note: Entities that have not
 * been created successfully are not listed.
 *
 * @param getEntitiesRequest @return *ImageListIntentResponse
 */
func (op Operations) ListImage(ctx context.Context, getEntitiesRequest *models.ImageListMetadata) (*models.ImageListIntentResponse, error) {
	path := "/images/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	imageListIntentResponse := new(models.ImageListIntentResponse)
	if err := op.client.Do(ctx, req, imageListIntentResponse); err != nil {
		return nil, err
	}

	return imageListIntentResponse, nil
}

/*UpdateImage updates a IMAGE
 * This operation submits a request to update a IMAGE based on the input parameters.
 *
 * @param uuid The uuid of the entity.
 * @param body
 * @return *ImageIntentResponse
 */
func (op Operations) UpdateImage(ctx context.Context, uuid string, body *models.ImageIntentInput) (*models.ImageIntentResponse, error) {
	path := fmt.Sprintf("/images/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	imageIntentResponse := new(models.ImageIntentResponse)
	if err := op.client.Do(ctx, req, imageIntentResponse); err != nil {
		return nil, err
	}

	return imageIntentResponse, nil
}

/*GetCluster gets a CLUSTER
 * This operation gets a CLUSTER.
 *
 * @param uuid The uuid of the entity.
 * @return *ImageIntentResponse
 */
func (op Operations) GetCluster(ctx context.Context, uuid string) (*models.ClusterIntentResponse, error) {
	path := fmt.Sprintf("/clusters/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	clusterIntentResponse := new(models.ClusterIntentResponse)
	if err := op.client.Do(ctx, req, clusterIntentResponse); err != nil {
		return nil, err
	}

	return clusterIntentResponse, nil
}

/*ListCluster gets a list of CLUSTERS This operation gets a list of CLUSTERS, allowing for sorting and pagination. Note: Entities that have
 * not been created successfully are not listed.
 *
 * @param getEntitiesRequest @return *ClusterListIntentResponse
 */
func (op Operations) ListCluster(ctx context.Context, getEntitiesRequest *models.ClusterListMetadata) (*models.ClusterListIntentResponse, error) {
	path := "/clusters/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	clusterList := new(models.ClusterListIntentResponse)
	if err := op.client.Do(ctx, req, clusterList); err != nil {
		return nil, err
	}

	return clusterList, nil
}

// CreateOrUpdateCategoryKey ...
func (op Operations) CreateOrUpdateCategoryKey(ctx context.Context, body *models.CategoryKey) (*models.CategoryKeyStatus, error) {
	path := fmt.Sprintf("/categories/%s", body.Name)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	categoryKeyResponse := new(models.CategoryKeyStatus)
	if err := op.client.Do(ctx, req, categoryKeyResponse); err != nil {
		return nil, err
	}

	return categoryKeyResponse, nil
}

/*ListCategories gets a list of Categories This operation gets a list of Categories, allowing for sorting and pagination. Note: Entities
 * that have not been created successfully are not listed.
 *
 * @param getEntitiesRequest @return *ImageListIntentResponse
 */
func (op Operations) ListCategories(ctx context.Context, getEntitiesRequest *models.CategoryListMetadata) (*models.CategoryKeyListResponse, error) {
	path := "/categories/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	categoryKeyListResponse := new(models.CategoryKeyListResponse)
	if err := op.client.Do(ctx, req, categoryKeyListResponse); err != nil {
		return nil, err
	}

	return categoryKeyListResponse, nil
}

/*DeleteCategoryKey Deletes a Category
 * This operation submits a request to delete a op.
 *
 * @param name The name of the entity.
 * @return error
 */
func (op Operations) DeleteCategoryKey(ctx context.Context, name string) error {
	path := fmt.Sprintf("/categories/%s", name)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return op.client.Do(ctx, req, nil)
}

/*GetCategoryKey gets a Category
 * This operation gets a Category.
 *
 * @param name The name of the entity.
 * @return *CategoryKeyStatus
 */
func (op Operations) GetCategoryKey(ctx context.Context, name string) (*models.CategoryKeyStatus, error) {
	path := fmt.Sprintf("/categories/%s", name)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	categoryKeyStatusResponse := new(models.CategoryKeyStatus)
	if err := op.client.Do(ctx, req, categoryKeyStatusResponse); err != nil {
		return nil, err
	}

	return categoryKeyStatusResponse, nil
}

/*ListCategoryValues gets a list of Category values for a specific key This operation gets a list of Categories, allowing for sorting and
 * pagination. Note: Entities that have not been created successfully are not listed.
 *
 * @param name @param getEntitiesRequest @return *CategoryValueListResponse
 */
func (op Operations) ListCategoryValues(ctx context.Context, name string, getEntitiesRequest *models.CategoryListMetadata) (*models.CategoryValueListResponse, error) {
	path := fmt.Sprintf("/categories/%s/list", name)
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	categoryValueListResponse := new(models.CategoryValueListResponse)
	if err := op.client.Do(ctx, req, categoryValueListResponse); err != nil {
		return nil, err
	}

	return categoryValueListResponse, nil
}

// CreateOrUpdateCategoryValue ...
func (op Operations) CreateOrUpdateCategoryValue(ctx context.Context, name string, body *models.CategoryValue) (*models.CategoryValueStatus, error) {
	path := fmt.Sprintf("/categories/%s/%s", name, utils.StringValue(body.Value))
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	categoryValueResponse := new(models.CategoryValueStatus)
	if err := op.client.Do(ctx, req, categoryValueResponse); err != nil {
		return nil, err
	}

	return categoryValueResponse, nil
}

/*GetCategoryValue gets a Category Value
 * This operation gets a Category Value.
 *
 * @param name The name of the entity.
 * @params value the value of entity that belongs to category key
 * @return *CategoryValueStatus
 */
func (op Operations) GetCategoryValue(ctx context.Context, name string, value string) (*models.CategoryValueStatus, error) {
	path := fmt.Sprintf("/categories/%s/%s", name, value)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	categoryValueStatusResponse := new(models.CategoryValueStatus)
	if err := op.client.Do(ctx, req, categoryValueStatusResponse); err != nil {
		return nil, err
	}

	return categoryValueStatusResponse, nil
}

/*DeleteCategoryValue Deletes a Category Value
 * This operation submits a request to delete a op.
 *
 * @param name The name of the entity.
 * @params value the value of entity that belongs to category key
 * @return error
 */
func (op Operations) DeleteCategoryValue(ctx context.Context, name string, value string) error {
	path := fmt.Sprintf("/categories/%s/%s", name, value)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return op.client.Do(ctx, req, nil)
}

/*GetCategoryQuery gets list of entities attached to categories or policies in which categories are used as defined by the filter criteria.
 *
 * @param query Categories query input object.
 * @return *CategoryQueryResponse
 */
func (op Operations) GetCategoryQuery(ctx context.Context, query *models.CategoryQueryInput) (*models.CategoryQueryResponse, error) {
	path := "/category/query"
	req, err := op.client.NewRequest(http.MethodPost, path, query)
	categoryQueryResponse := new(models.CategoryQueryResponse)

	if err != nil {
		return nil, err
	}

	return categoryQueryResponse, op.client.Do(ctx, req, categoryQueryResponse)
}

/*CreateNetworkSecurityRule Creates a Network security rule
 * This operation submits a request to create a Network security rule based on the input parameters.
 *
 * @param request
 * @return *NetworkSecurityRuleIntentResponse
 */
func (op Operations) CreateNetworkSecurityRule(ctx context.Context, request *models.NetworkSecurityRuleIntentInput) (*models.NetworkSecurityRuleIntentResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/network_security_rules", request)
	if err != nil {
		return nil, err
	}

	networkSecurityRuleIntentResponse := new(models.NetworkSecurityRuleIntentResponse)
	if err := op.client.Do(ctx, req, networkSecurityRuleIntentResponse); err != nil {
		return nil, err
	}

	return networkSecurityRuleIntentResponse, nil
}

/*DeleteNetworkSecurityRule Deletes a Network security rule
 * This operation submits a request to delete a Network security rule.
 *
 * @param uuid The uuid of the entity.
 * @return void
 */
func (op Operations) DeleteNetworkSecurityRule(ctx context.Context, uuid string) (*models.NetworkSecurityRuleIntentResponse, error) {
	path := fmt.Sprintf("/network_security_rules/%s", uuid)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	deleteResponse := new(models.NetworkSecurityRuleIntentResponse)
	if err := op.client.Do(ctx, req, deleteResponse); err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

/*GetNetworkSecurityRule Gets a Network security rule
 * This operation gets a Network security rule.
 *
 * @param uuid The uuid of the entity.
 * @return *NetworkSecurityRuleIntentResponse
 */
func (op Operations) GetNetworkSecurityRule(ctx context.Context, uuid string) (*models.NetworkSecurityRuleIntentResponse, error) {
	path := fmt.Sprintf("/network_security_rules/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	networkSecurityRuleIntentResponse := new(models.NetworkSecurityRuleIntentResponse)
	if err := op.client.Do(ctx, req, networkSecurityRuleIntentResponse); err != nil {
		return nil, err
	}

	return networkSecurityRuleIntentResponse, nil
}

/*ListNetworkSecurityRule Gets all network security rules This operation gets a list of Network security rules, allowing for sorting and
 * pagination. Note: Entities that have not been created successfully are not listed.
 *
 * @param getEntitiesRequest @return *NetworkSecurityRuleListIntentResponse
 */
func (op Operations) ListNetworkSecurityRule(ctx context.Context, getEntitiesRequest *models.NetworkSecurityRuleListMetadata) (*models.NetworkSecurityRuleListIntentResponse, error) {
	path := "/network_security_rules/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	networkSecurityRuleListIntentResponse := new(models.NetworkSecurityRuleListIntentResponse)
	if err := op.client.Do(ctx, req, networkSecurityRuleListIntentResponse); err != nil {
		return nil, err
	}

	return networkSecurityRuleListIntentResponse, nil
}

/*UpdateNetworkSecurityRule Updates a Network security rule
 * This operation submits a request to update a Network security rule based on the input parameters.
 *
 * @param uuid The uuid of the entity.
 * @param body
 * @return void
 */
func (op Operations) UpdateNetworkSecurityRule(ctx context.Context, uuid string, body *models.NetworkSecurityRuleIntentInput) (*models.NetworkSecurityRuleIntentResponse, error) {
	path := fmt.Sprintf("/network_security_rules/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	networkSecurityRuleIntentResponse := new(models.NetworkSecurityRuleIntentResponse)
	if err := op.client.Do(ctx, req, networkSecurityRuleIntentResponse); err != nil {
		return nil, err
	}

	return networkSecurityRuleIntentResponse, nil
}

/*CreateVolumeGroup Creates a Volume group
 * This operation submits a request to create a Volume group based on the input parameters.
 *
 * @param request
 * @return *VolumeGroupResponse
 */
func (op Operations) CreateVolumeGroup(ctx context.Context, request *models.VolumeGroupIntentInput) (*models.VolumeGroupIntentResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/volume_groups", request)
	if err != nil {
		return nil, err
	}

	networkSecurityRuleResponse := new(models.VolumeGroupIntentResponse)
	if err := op.client.Do(ctx, req, networkSecurityRuleResponse); err != nil {
		return nil, err
	}

	return networkSecurityRuleResponse, nil
}

/*DeleteVolumeGroup Deletes a Volume group
 * This operation submits a request to delete a Volume group.
 *
 * @param uuid The uuid of the entity.
 * @return void
 */
func (op Operations) DeleteVolumeGroup(ctx context.Context, uuid string) error {
	path := fmt.Sprintf("/volume_groups/%s", uuid)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return op.client.Do(ctx, req, nil)
}

/*GetVolumeGroup Gets a Volume group
 * This operation gets a Volume group.
 *
 * @param uuid The uuid of the entity.
 * @return *VolumeGroupResponse
 */
func (op Operations) GetVolumeGroup(ctx context.Context, uuid string) (*models.VolumeGroupIntentResponse, error) {
	path := fmt.Sprintf("/volume_groups/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	networkSecurityRuleResponse := new(models.VolumeGroupIntentResponse)
	if err := op.client.Do(ctx, req, networkSecurityRuleResponse); err != nil {
		return nil, err
	}

	return networkSecurityRuleResponse, nil
}

/*ListVolumeGroup Gets all network security rules This operation gets a list of Volume groups, allowing for sorting and pagination. Note:
 * Entities that have not been created successfully are not listed.
 *
 * @param getEntitiesRequest @return *VolumeGroupListResponse
 */
func (op Operations) ListVolumeGroup(ctx context.Context, getEntitiesRequest *models.VolumeGroupListMetadata) (*models.VolumeGroupListIntentResponse, error) {
	path := "/volume_groups/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	networkSecurityRuleListResponse := new(models.VolumeGroupListIntentResponse)
	if err := op.client.Do(ctx, req, networkSecurityRuleListResponse); err != nil {
		return nil, err
	}

	return networkSecurityRuleListResponse, nil
}

/*UpdateVolumeGroup Updates a Volume group
 * This operation submits a request to update a Volume group based on the input parameters.
 *
 * @param uuid The uuid of the entity.
 * @param body
 * @return void
 */
func (op Operations) UpdateVolumeGroup(ctx context.Context, uuid string, body *models.VolumeGroupIntentInput) (*models.VolumeGroupIntentResponse, error) {
	path := fmt.Sprintf("/volume_groups/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	networkSecurityRuleResponse := new(models.VolumeGroupIntentResponse)
	if err := op.client.Do(ctx, req, networkSecurityRuleResponse); err != nil {
		return nil, err
	}

	return networkSecurityRuleResponse, nil
}

const itemsPerPage int32 = 100

func hasNext(ri *int32) bool {
	*ri -= itemsPerPage
	return *ri >= (0 - itemsPerPage)
}

// ListAllVM ...
func (op Operations) ListAllVM(ctx context.Context, filter string) (*models.VmListIntentResponse, error) {
	entities := make([]models.VmIntentResource, 0)

	resp, err := op.ListVM(ctx, &models.VmListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("vm"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListVM(ctx, &models.VmListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("vm"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
		}

		resp.Entities = entities
	}

	return resp, nil
}

// ListAllSubnet ...
func (op Operations) ListAllSubnet(ctx context.Context, filter string) (*models.SubnetListIntentResponse, error) {
	entities := make([]models.SubnetIntentResource, 0)

	resp, err := op.ListSubnet(ctx, &models.SubnetListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("subnet"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListSubnet(ctx, &models.SubnetListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("subnet"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

// ListAllNetworkSecurityRule ...
func (op Operations) ListAllNetworkSecurityRule(ctx context.Context, filter string) (*models.NetworkSecurityRuleListIntentResponse, error) {
	entities := make([]models.NetworkSecurityRuleIntentResource, 0)

	resp, err := op.ListNetworkSecurityRule(ctx, &models.NetworkSecurityRuleListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("network_security_rule"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListNetworkSecurityRule(ctx, &models.NetworkSecurityRuleListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("network_security_rule"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

// ListAllImage ...
func (op Operations) ListAllImage(ctx context.Context, filter string) (*models.ImageListIntentResponse, error) {
	entities := make([]models.ImageIntentResource, 0)

	resp, err := op.ListImage(ctx, &models.ImageListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("image"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListImage(ctx, &models.ImageListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("image"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

// ListAllCluster ...
func (op Operations) ListAllCluster(ctx context.Context, filter string) (*models.ClusterListIntentResponse, error) {
	entities := make([]models.ClusterIntentResource, 0)

	resp, err := op.ListCluster(ctx, &models.ClusterListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("cluster"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListCluster(ctx, &models.ClusterListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("cluster"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

// ListAllCategoryValues ...
func (op Operations) ListAllCategoryValues(ctx context.Context, categoryName, filter string) (*models.CategoryValueListResponse, error) {
	entities := make([]models.CategoryValueStatus, 0)

	resp, err := op.ListCategoryValues(ctx, categoryName, &models.CategoryListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("category"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	// CategoryValueListResponse.Metadata.TotalMatches does not exist so just use the length of the entities
	// to determine the likelihood of more entities
	totalEntities := int32(len(resp.Entities))
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities >= itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListCategoryValues(ctx, categoryName, &models.CategoryListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("category"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

// GetTask ...
func (op Operations) GetTask(ctx context.Context, taskUUID string) (*models.Task, error) {
	path := fmt.Sprintf("/tasks/%s", taskUUID)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	tasksTesponse := new(models.Task)
	if err := op.client.Do(ctx, req, tasksTesponse); err != nil {
		return nil, err
	}

	return tasksTesponse, nil
}

// GetHost ...
func (op Operations) GetHost(ctx context.Context, hostUUID string) (*models.HostIntentResponse, error) {
	path := fmt.Sprintf("/hosts/%s", hostUUID)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	host := new(models.HostIntentResponse)
	if err := op.client.Do(ctx, req, host); err != nil {
		return nil, err
	}

	return host, nil
}

// ListHost ...
func (op Operations) ListHost(ctx context.Context, getEntitiesRequest *models.HostListMetadata) (*models.HostListIntentResponse, error) {
	path := "/hosts/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	hostList := new(models.HostListIntentResponse)
	if err := op.client.Do(ctx, req, hostList); err != nil {
		return nil, err
	}

	return hostList, nil
}

// ListAllHost ...
func (op Operations) ListAllHost(ctx context.Context) (*models.HostListIntentResponse, error) {
	entities := make([]models.HostIntentResource, 0)

	resp, err := op.ListHost(ctx, &models.HostListMetadata{
		Kind:   utils.StringPtr("host"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListHost(ctx, &models.HostListMetadata{
				Kind:   utils.StringPtr("cluster"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

/*CreateProject creates a project
 * This operation submits a request to create a project based on the input parameters.
 *
 * @param request *Project
 * @return *Project
 */
func (op Operations) CreateProject(ctx context.Context, request *models.ProjectIntentInput) (*models.ProjectIntentResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/projects", request)
	if err != nil {
		return nil, err
	}

	projectResponse := new(models.ProjectIntentResponse)
	if err := op.client.Do(ctx, req, projectResponse); err != nil {
		return nil, err
	}

	return projectResponse, nil
}

/*GetProject This operation gets a project.
 *
 * @param uuid The prject uuid - string.
 * @return *Project
 */
func (op Operations) GetProject(ctx context.Context, projectUUID string) (*models.ProjectIntentResponse, error) {
	path := fmt.Sprintf("/projects/%s", projectUUID)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	project := new(models.ProjectIntentResponse)
	if err := op.client.Do(ctx, req, project); err != nil {
		return nil, err
	}

	return project, nil
}

/*ListProject gets a list of projects.
 *
 * @param metadata allows create filters to get specific data - *DSMetadata.
 * @return *ProjectListResponse
 */
func (op Operations) ListProject(ctx context.Context, getEntitiesRequest *models.ProjectListMetadata) (*models.ProjectListIntentResponse, error) {
	path := "/projects/list"

	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	projectList := new(models.ProjectListIntentResponse)
	if err := op.client.Do(ctx, req, projectList); err != nil {
		return nil, err
	}

	return projectList, nil
}

/*ListAllProject gets a list of projects
 * This operation gets a list of Projects, allowing for sorting and pagination.
 * Note: Entities that have not been created successfully are not listed.
 * @return *ProjectListResponse
 */
func (op Operations) ListAllProject(ctx context.Context, filter string) (*models.ProjectListIntentResponse, error) {
	entities := make([]models.ProjectIntentResource, 0)

	resp, err := op.ListProject(ctx, &models.ProjectListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("project"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListProject(ctx, &models.ProjectListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("project"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

/*UpdateProject Updates a project
 * This operation submits a request to update a existing Project based on the input parameters
 * @param uuid The uuid of the entity - string.
 * @param body - *Project
 * @return *Project, error
 */
func (op Operations) UpdateProject(ctx context.Context, uuid string, body *models.ProjectIntentInput) (*models.ProjectIntentResponse, error) {
	path := fmt.Sprintf("/projects/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	response := new(models.ProjectIntentResponse)
	if err := op.client.Do(ctx, req, response); err != nil {
		return nil, err
	}

	return response, nil
}

/*DeleteProject Deletes a project
 * This operation submits a request to delete a existing Project.
 *
 * @param uuid The uuid of the entity.
 * @return void
 */
func (op Operations) DeleteProject(ctx context.Context, uuid string) (*models.ProjectIntentResponse, error) {
	path := fmt.Sprintf("/projects/%s", uuid)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	deleteResponse := new(models.ProjectIntentResponse)
	if err := op.client.Do(ctx, req, deleteResponse); err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

/*CreateAccessControlPolicy creates a access policy
 * This operation submits a request to create a access policy based on the input parameters.
 *
 * @param request *Access Policy
 * @return *Access Policy
 */
func (op Operations) CreateAccessControlPolicy(ctx context.Context, request *models.AccessControlPolicyIntentInput) (*models.AccessControlPolicyIntentResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/access_control_policies", request)
	if err != nil {
		return nil, err
	}

	accessControlPolicyResponse := new(models.AccessControlPolicyIntentResponse)
	if err := op.client.Do(ctx, req, accessControlPolicyResponse); err != nil {
		return nil, err
	}

	return accessControlPolicyResponse, nil
}

/*GetAccessControlPolicy This operation gets a AccessControlPolicy.
 *
 * @param uuid The access policy uuid - string.
 * @return *AccessControlPolicy
 */
func (op Operations) GetAccessControlPolicy(ctx context.Context, accessControlPolicyUUID string) (*models.AccessControlPolicyIntentResponse, error) {
	path := fmt.Sprintf("/access_control_policies/%s", accessControlPolicyUUID)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	accessControlPolicy := new(models.AccessControlPolicyIntentResponse)
	if err := op.client.Do(ctx, req, accessControlPolicy); err != nil {
		return nil, err
	}

	return accessControlPolicy, nil
}

/*ListAccessControlPolicy gets a list of AccessControlPolicys.
 *
 * @param metadata allows create filters to get specific data - *DSMetadata.
 * @return *AccessControlPolicyListResponse
 */
func (op Operations) ListAccessControlPolicy(ctx context.Context, getEntitiesRequest *models.AccessControlPolicyListMetadata) (*models.AccessControlPolicyListIntentResponse, error) {
	path := "/access_control_policies/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	accessControlPolicyList := new(models.AccessControlPolicyListIntentResponse)
	if err := op.client.Do(ctx, req, accessControlPolicyList); err != nil {
		return nil, err
	}

	return accessControlPolicyList, nil
}

/*ListAllAccessControlPolicy gets a list of AccessControlPolicys
 * This operation gets a list of AccessControlPolicys, allowing for sorting and pagination.
 * Note: Entities that have not been created successfully are not listed.
 * @return *AccessControlPolicyListResponse
 */
func (op Operations) ListAllAccessControlPolicy(ctx context.Context, filter string) (*models.AccessControlPolicyListIntentResponse, error) {
	entities := make([]models.AccessControlPolicyIntentResource, 0)

	resp, err := op.ListAccessControlPolicy(ctx, &models.AccessControlPolicyListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("access_control_policy"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListAccessControlPolicy(ctx, &models.AccessControlPolicyListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("access_control_policy"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

/*UpdateAccessControlPolicy Updates a AccessControlPolicy
 * This operation submits a request to update a existing AccessControlPolicy based on the input parameters
 * @param uuid The uuid of the entity - string.
 * @param body - *AccessControlPolicy
 * @return *AccessControlPolicy, error
 */
func (op Operations) UpdateAccessControlPolicy(ctx context.Context, uuid string, body *models.AccessControlPolicyIntentInput) (*models.AccessControlPolicyIntentResponse, error) {
	path := fmt.Sprintf("/access_control_policies/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	accessControlPolicyInput := new(models.AccessControlPolicyIntentResponse)
	if err := op.client.Do(ctx, req, accessControlPolicyInput); err != nil {
		return nil, err
	}

	return accessControlPolicyInput, nil
}

/*DeleteAccessControlPolicy Deletes a AccessControlPolicy
 * This operation submits a request to delete a existing AccessControlPolicy.
 *
 * @param uuid The uuid of the entity.
 * @return void
 */
func (op Operations) DeleteAccessControlPolicy(ctx context.Context, uuid string) (*models.AccessControlPolicyIntentResponse, error) {
	path := fmt.Sprintf("/access_control_policies/%s", uuid)

	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	deleteResponse := new(models.AccessControlPolicyIntentResponse)
	if err := op.client.Do(ctx, req, deleteResponse); err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

/*CreateRole creates a role
 * This operation submits a request to create a role based on the input parameters.
 *
 * @param request *Role
 * @return *Role
 */
func (op Operations) CreateRole(ctx context.Context, request *models.RoleIntentInput) (*models.RoleIntentResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/roles", request)
	if err != nil {
		return nil, err
	}

	roleResponse := new(models.RoleIntentResponse)
	if err := op.client.Do(ctx, req, roleResponse); err != nil {
		return nil, err
	}

	return roleResponse, nil
}

/*GetRole This operation gets a role.
 *
 * @param uuid The role uuid - string.
 * @return *Role
 */
func (op Operations) GetRole(ctx context.Context, roleUUID string) (*models.RoleIntentResponse, error) {
	path := fmt.Sprintf("/roles/%s", roleUUID)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	role := new(models.RoleIntentResponse)
	if err := op.client.Do(ctx, req, role); err != nil {
		return nil, err
	}

	return role, nil
}

/*ListRole gets a list of roles.
 *
 * @param metadata allows create filters to get specific data - *DSMetadata.
 * @return *RoleListResponse
 */
func (op Operations) ListRole(ctx context.Context, getEntitiesRequest *models.RoleListMetadata) (*models.RoleListIntentResponse, error) {
	path := "/roles/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	roleList := new(models.RoleListIntentResponse)
	if err := op.client.Do(ctx, req, roleList); err != nil {
		return nil, err
	}

	return roleList, nil
}

/*ListAllRole gets a list of Roles
 * This operation gets a list of Roles, allowing for sorting and pagination.
 * Note: Entities that have not been created successfully are not listed.
 * @return *RoleListResponse
 */
func (op Operations) ListAllRole(ctx context.Context, filter string) (*models.RoleListIntentResponse, error) {
	entities := make([]models.RoleIntentResource, 0)

	resp, err := op.ListRole(ctx, &models.RoleListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("role"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListRole(ctx, &models.RoleListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("role"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

/*UpdateRole Updates a role
 * This operation submits a request to update a existing role based on the input parameters
 * @param uuid The uuid of the entity - string.
 * @param body - *Role
 * @return *Role, error
 */
func (op Operations) UpdateRole(ctx context.Context, uuid string, body *models.RoleIntentInput) (*models.RoleIntentResponse, error) {
	path := fmt.Sprintf("/roles/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	roleInput := new(models.RoleIntentResponse)
	if err := op.client.Do(ctx, req, roleInput); err != nil {
		return nil, err
	}

	return roleInput, nil
}

/*DeleteRole Deletes a role
 * This operation submits a request to delete a existing role.
 *
 * @param uuid The uuid of the entity.
 * @return void
 */
func (op Operations) DeleteRole(ctx context.Context, uuid string) (*models.RoleIntentResponse, error) {
	path := fmt.Sprintf("/roles/%s", uuid)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	deleteResponse := new(models.RoleIntentResponse)
	if err := op.client.Do(ctx, req, deleteResponse); err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

/*CreateUser creates a User
 * This operation submits a request to create a userbased on the input parameters.
 *
 * @param request *VMIntentInput
 * @return *UserIntentResponse
 */
func (op Operations) CreateUser(ctx context.Context, request *models.UserIntentInput) (*models.UserIntentResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/users", request)
	if err != nil {
		return nil, err
	}

	userIntentResponse := new(models.UserIntentResponse)
	if err := op.client.Do(ctx, req, userIntentResponse); err != nil {
		return nil, err
	}

	return userIntentResponse, nil
}

/*GetUser This operation gets a User.
 *
 * @param uuid The user uuid - string.
 * @return *User
 */
func (op Operations) GetUser(ctx context.Context, userUUID string) (*models.UserIntentResponse, error) {
	path := fmt.Sprintf("/users/%s", userUUID)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	user := new(models.UserIntentResponse)
	if err := op.client.Do(ctx, req, user); err != nil {
		return nil, err
	}

	return user, nil
}

/*UpdateUser Updates a User
 * This operation submits a request to update a existing User based on the input parameters
 * @param uuid The uuid of the entity - string.
 * @param body - *User
 * @return *User, error
 */
func (op Operations) UpdateUser(ctx context.Context, uuid string, body *models.UserIntentInput) (*models.UserIntentResponse, error) {
	path := fmt.Sprintf("/users/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	userInput := new(models.UserIntentResponse)
	if err := op.client.Do(ctx, req, userInput); err != nil {
		return nil, err
	}

	return userInput, nil
}

/*DeleteUser Deletes a User
 * This operation submits a request to delete a existing User.
 *
 * @param uuid The uuid of the entity.
 * @return void
 */
func (op Operations) DeleteUser(ctx context.Context, uuid string) (*models.UserIntentResponse, error) {
	path := fmt.Sprintf("/users/%s", uuid)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	deleteResponse := new(models.UserIntentResponse)
	if err := op.client.Do(ctx, req, deleteResponse); err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

/*ListUser gets a list of Users.
 *
 * @param metadata allows create filters to get specific data - *DSMetadata.
 * @return *UserListResponse
 */
func (op Operations) ListUser(ctx context.Context, getEntitiesRequest *models.UserListMetadata) (*models.UserListIntentResponse, error) {
	path := "/users/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	userList := new(models.UserListIntentResponse)
	if err := op.client.Do(ctx, req, userList); err != nil {
		return nil, err
	}

	return userList, nil
}

// ListAllUser ...
func (op Operations) ListAllUser(ctx context.Context, filter string) (*models.UserListIntentResponse, error) {
	entities := make([]models.UserIntentResource, 0)

	resp, err := op.ListUser(ctx, &models.UserListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("user"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListUser(ctx, &models.UserListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("user"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
		}

		resp.Entities = entities
	}

	return resp, nil
}

/*GetCurrentLoggedInUser This operation gets the user info for the currently logged in User.
 *
 * @param context
 * @return *User
 */
func (op Operations) GetCurrentLoggedInUser(ctx context.Context) (*models.UserIntentResponse, error) {
	path := "/users/me"
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	user := new(models.UserIntentResponse)
	if err := op.client.Do(ctx, req, user); err != nil {
		return nil, err
	}

	return user, nil
}

/*GetUserGroup This operation gets a User.
 *
 * @param uuid The user uuid - string.
 * @return *User
 */
func (op Operations) GetUserGroup(ctx context.Context, userGroupUUID string) (*models.UserGroupIntentResponse, error) {
	path := fmt.Sprintf("/user_groups/%s", userGroupUUID)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	user := new(models.UserGroupIntentResponse)
	if err := op.client.Do(ctx, req, user); err != nil {
		return nil, err
	}

	return user, nil
}

/*ListUserGroup gets a list of UserGroups.
 *
 * @param metadata allows create filters to get specific data - *DSMetadata.
 * @return *UserGroupListResponse
 */
func (op Operations) ListUserGroup(ctx context.Context, getEntitiesRequest *models.UserGroupListMetadata) (*models.UserGroupListIntentResponse, error) {
	path := "/user_groups/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	userGroupList := new(models.UserGroupListIntentResponse)
	if err := op.client.Do(ctx, req, userGroupList); err != nil {
		return nil, err
	}

	return userGroupList, nil
}

// ListAllUserGroup ...
func (op Operations) ListAllUserGroup(ctx context.Context, filter string) (*models.UserGroupListIntentResponse, error) {
	entities := make([]models.UserGroupIntentResource, 0)

	resp, err := op.ListUserGroup(ctx, &models.UserGroupListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("user_group"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListUserGroup(ctx, &models.UserGroupListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("user"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
		}

		resp.Entities = entities
	}

	return resp, nil
}

/*GePermission This operation gets a Permission.
 *
 * @param uuid The permission uuid - string.
 * @return *PermissionIntentResponse
 */
func (op Operations) GetPermission(ctx context.Context, permissionUUID string) (*models.PermissionIntentResponse, error) {
	path := fmt.Sprintf("/permissions/%s", permissionUUID)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	permission := new(models.PermissionIntentResponse)
	if err := op.client.Do(ctx, req, permission); err != nil {
		return nil, err
	}

	return permission, nil
}

/*ListPermission gets a list of Permissions.
 *
 * @param metadata allows create filters to get specific data - *DSMetadata.
 * @return *PermissionListResponse
 */
func (op Operations) ListPermission(ctx context.Context, getEntitiesRequest *models.PermissionListMetadata) (*models.PermissionListIntentResponse, error) {
	path := "/permissions/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	permissionList := new(models.PermissionListIntentResponse)
	if err := op.client.Do(ctx, req, permissionList); err != nil {
		return nil, err
	}

	return permissionList, nil
}

// ListAllPermission ...
func (op Operations) ListAllPermission(ctx context.Context, filter string) (*models.PermissionListIntentResponse, error) {
	entities := make([]models.PermissionIntentResource, 0)

	resp, err := op.ListPermission(ctx, &models.PermissionListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("permission"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListPermission(ctx, &models.PermissionListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("permission"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
		}

		resp.Entities = entities
	}

	return resp, nil
}

// GetProtectionRule ...
func (op Operations) GetProtectionRule(ctx context.Context, uuid string) (*models.ProtectionRuleIntentResponse, error) {
	path := fmt.Sprintf("/protection_rules/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	protectionRule := new(models.ProtectionRuleIntentResponse)
	if err := op.client.Do(ctx, req, protectionRule); err != nil {
		return nil, err
	}

	return protectionRule, nil
}

// ListProtectionRules ...
func (op Operations) ListProtectionRules(ctx context.Context, getEntitiesRequest *models.ProtectionRuleListMetadata) (*models.ProtectionRuleListIntentResponse, error) {
	path := "/protection_rules/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	list := new(models.ProtectionRuleListIntentResponse)
	if err := op.client.Do(ctx, req, list); err != nil {
		return nil, err
	}

	return list, nil
}

// ListAllProtectionRules ...
func (op Operations) ListAllProtectionRules(ctx context.Context, filter string) (*models.ProtectionRuleListIntentResponse, error) {
	entities := make([]models.ProtectionRuleIntentResource, 0)

	resp, err := op.ListProtectionRules(ctx, &models.ProtectionRuleListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("protection_rule"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListProtectionRules(ctx, &models.ProtectionRuleListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("protection_rule"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

// CreateProtectionRule ...
func (op Operations) CreateProtectionRule(ctx context.Context, createRequest *models.ProtectionRuleIntentInput) (*models.ProtectionRuleIntentResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/protection_rules", createRequest)
	if err != nil {
		return nil, err
	}

	protectionRuleResponse := new(models.ProtectionRuleIntentResponse)
	if err := op.client.Do(ctx, req, protectionRuleResponse); err != nil {
		return nil, err
	}

	return protectionRuleResponse, nil
}

// UpdateProtectionRule ...
func (op Operations) UpdateProtectionRule(ctx context.Context, uuid string, body *models.ProtectionRuleIntentInput) (*models.ProtectionRuleIntentResponse, error) {
	path := fmt.Sprintf("/protection_rules/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	protectionRuleResponse := new(models.ProtectionRuleIntentResponse)
	if err := op.client.Do(ctx, req, protectionRuleResponse); err != nil {
		return nil, err
	}

	return protectionRuleResponse, nil
}

// DeleteProtectionRule ...
func (op Operations) DeleteProtectionRule(ctx context.Context, uuid string) (*models.ProtectionRuleIntentResponse, error) {
	path := fmt.Sprintf("/protection_rules/%s", uuid)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	deleteResponse := new(models.ProtectionRuleIntentResponse)
	if err := op.client.Do(ctx, req, deleteResponse); err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

/*ProcessProtectionRule triggers the evaluation of a processing rule
 * immediately.
 *
 * @param uuid is the uuid of the protection rule to process.
 */
func (op Operations) ProcessProtectionRule(ctx context.Context, uuid string) error {
	path := fmt.Sprintf("/protection_rules/%s/process", uuid)

	req, err := op.client.NewRequest(http.MethodPost, path, nil)
	if err != nil {
		return err
	}

	return op.client.Do(ctx, req, nil)
}

// GetRecoveryPlan ...
func (op Operations) GetRecoveryPlan(ctx context.Context, uuid string) (*models.RecoveryPlanIntentResponse, error) {
	path := fmt.Sprintf("/recovery_plans/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	recoveryPlan := new(models.RecoveryPlanIntentResponse)
	if err := op.client.Do(ctx, req, recoveryPlan); err != nil {
		return nil, err
	}

	return recoveryPlan, nil
}

// ListRecoveryPlans ...
func (op Operations) ListRecoveryPlans(ctx context.Context, getEntitiesRequest *models.RecoveryPlanListMetadata) (*models.RecoveryPlanListIntentResponse, error) {
	path := "/recovery_plans/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	list := new(models.RecoveryPlanListIntentResponse)
	if err := op.client.Do(ctx, req, list); err != nil {
		return nil, err
	}

	return list, nil
}

// ListAllRecoveryPlans ...
func (op Operations) ListAllRecoveryPlans(ctx context.Context, filter string) (*models.RecoveryPlanListIntentResponse, error) {
	entities := make([]models.RecoveryPlanIntentResource, 0)

	resp, err := op.ListRecoveryPlans(ctx, &models.RecoveryPlanListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("recovery_plan"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int32Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListRecoveryPlans(ctx, &models.RecoveryPlanListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("recovery_plan"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

// CreateRecoveryPlan ...
func (op Operations) CreateRecoveryPlan(ctx context.Context, createRequest *models.RecoveryPlanIntentInput) (*models.RecoveryPlanIntentResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/recovery_plans", createRequest)
	if err != nil {
		return nil, err
	}

	recoveryPlanResponse := new(models.RecoveryPlanIntentResponse)
	if err := op.client.Do(ctx, req, recoveryPlanResponse); err != nil {
		return nil, err
	}

	return recoveryPlanResponse, nil
}

// UpdateRecoveryPlan ...
func (op Operations) UpdateRecoveryPlan(ctx context.Context, uuid string, body *models.RecoveryPlanIntentInput) (*models.RecoveryPlanIntentResponse, error) {
	path := fmt.Sprintf("/recovery_plans/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	recoveryPlanResponse := new(models.RecoveryPlanIntentResponse)
	if err := op.client.Do(ctx, req, recoveryPlanResponse); err != nil {
		return nil, err
	}

	return recoveryPlanResponse, nil
}

// DeleteRecoveryPlan ...
func (op Operations) DeleteRecoveryPlan(ctx context.Context, uuid string) (*models.RecoveryPlanIntentResponse, error) {
	path := fmt.Sprintf("/recovery_plans/%s", uuid)
	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	deleteResponse := new(models.RecoveryPlanIntentResponse)
	if err := op.client.Do(ctx, req, deleteResponse); err != nil {
		return nil, err
	}

	return deleteResponse, nil
}

func (op Operations) GetServiceGroup(ctx context.Context, uuid string) (*models.ServiceGroupResponse, error) {
	path := fmt.Sprintf("/service_groups/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	serviceGroup := new(models.ServiceGroupResponse)
	if err := op.client.Do(ctx, req, serviceGroup); err != nil {
		return nil, err
	}

	return serviceGroup, nil
}

func (op Operations) CreateServiceGroup(ctx context.Context, request *models.ServiceGroup) (*models.Reference, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/service_groups", request)
	if err != nil {
		return nil, err
	}

	serviceGroup := new(models.Reference)
	if err := op.client.Do(ctx, req, serviceGroup); err != nil {
		return nil, err
	}

	return serviceGroup, nil
}

func (op Operations) DeleteServiceGroup(ctx context.Context, uuid string) error {
	path := fmt.Sprintf("/service_groups/%s", uuid)

	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return op.client.Do(ctx, req, nil)
}

func (op Operations) ListAllServiceGroups(ctx context.Context, filter string) (*models.ServiceGroupListResponse, error) {
	entities := make([]models.ServiceGroupResponseResource, 0)

	resp, err := op.listServiceGroups(ctx, &models.ServiceGroupListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("service_group"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	// ServiceGroupListResponse.Metadata.TotalMatches does not exist so just use the length of the entities
	// to determine the likelihood of more entities
	totalEntities := int32(len(resp.Entities))
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities >= itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.listServiceGroups(ctx, &models.ServiceGroupListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("service_group"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

func (op Operations) listServiceGroups(ctx context.Context, getEntitiesRequest *models.ServiceGroupListMetadata) (*models.ServiceGroupListResponse, error) {
	path := "/service_groups/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	list := new(models.ServiceGroupListResponse)
	if err := op.client.Do(ctx, req, list); err != nil {
		return nil, err
	}

	return list, nil
}

func (op Operations) UpdateServiceGroup(ctx context.Context, uuid string, body *models.ServiceGroup) error {
	path := fmt.Sprintf("/service_groups/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return err
	}

	return op.client.Do(ctx, req, nil)
}

func (op Operations) GetAddressGroup(ctx context.Context, uuid string) (*models.AddressGroupResponse, error) {
	path := fmt.Sprintf("/address_groups/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	addressGroup := new(models.AddressGroupResponse)
	if err := op.client.Do(ctx, req, addressGroup); err != nil {
		return nil, err
	}

	return addressGroup, nil
}

func (op Operations) ListAllAddressGroups(ctx context.Context, filter string) (*models.AddressGroupListResponse, error) {
	entities := make([]models.AddressGroupResponseResource, 0)

	resp, err := op.ListAddressGroups(ctx, &models.AddressGroupListMetadata{
		Filter: &filter,
		Kind:   utils.StringPtr("address_group"),
		Length: utils.Int32Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	// AddressGroupListResponse.Metadata.TotalMatches does not exist so just use the length of the entities
	// to determine the likelihood of more entities
	totalEntities := int32(len(resp.Entities))
	remaining := totalEntities
	offset := utils.Int32Value(resp.Metadata.Offset)

	if totalEntities >= itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListAddressGroups(ctx, &models.AddressGroupListMetadata{
				Filter: &filter,
				Kind:   utils.StringPtr("address_group"),
				Length: utils.Int32Ptr(itemsPerPage),
				Offset: utils.Int32Ptr(offset),
			})
			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

func (op Operations) ListAddressGroups(ctx context.Context, getEntitiesRequest *models.AddressGroupListMetadata) (*models.AddressGroupListResponse, error) {
	path := "/address_groups/list"
	req, err := op.client.NewRequest(http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	list := new(models.AddressGroupListResponse)
	if err := op.client.Do(ctx, req, list); err != nil {
		return nil, err
	}

	return list, nil
}

func (op Operations) DeleteAddressGroup(ctx context.Context, uuid string) error {
	path := fmt.Sprintf("/address_groups/%s", uuid)

	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return op.client.Do(ctx, req, nil)
}

func (op Operations) CreateAddressGroup(ctx context.Context, request *models.AddressGroup) (*models.Reference, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/address_groups", request)
	if err != nil {
		return nil, err
	}

	addressGroup := new(models.Reference)
	if err := op.client.Do(ctx, req, addressGroup); err != nil {
		return nil, err
	}

	return addressGroup, nil
}

func (op Operations) UpdateAddressGroup(ctx context.Context, uuid string, body *models.AddressGroup) error {
	path := fmt.Sprintf("/address_groups/%s", uuid)
	req, err := op.client.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return err
	}

	return op.client.Do(ctx, req, nil)
}

/*Creates a recovery plan job.
 * This operation creates a new recovery plan job based on the inputs in the 'request'.
 *
 * @param request Pointer to a specification of type RecoveryPlanJobIntentInput.
 */
func (op Operations) CreateRecoveryPlanJob(ctx context.Context, request *models.RecoveryPlanJobIntentInput) (*models.RecoveryPlanJobResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/recovery_plan_jobs", request)
	if err != nil {
		return nil, err
	}

	response := new(models.RecoveryPlanJobResponse)
	if err := op.client.Do(ctx, req, response); err != nil {
		return nil, err
	}

	return response, nil
}

/*Deletes a recovery plan job.
 * This operation deletes the new recovery plan job identified by 'uuid'.
 *
 * @param uuid UUID of the recovery plan job to be deleted.
 */
func (op Operations) DeleteRecoveryPlanJob(ctx context.Context, uuid string) error {
	path := fmt.Sprintf("/recovery_plan_jobs/%s", uuid)

	req, err := op.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return op.client.Do(ctx, req, nil)
}

/*Perform a recovery plan job action.
 * This operation initiates the 'action' on the recovery plan job identified
 * by 'uuid' and governed by the specification in 'request'.
 *
 * @param uuid UUID of the recovery plan job.
 * @param action one of {'cleanup', 'rerun'}.
 * @param request pointer to the specification of type RecoveryPlanJobActionRequest.
 */
func (op Operations) PerformRecoveryPlanJobAction(ctx context.Context, uuid string, action string, request *models.RecoveryPlanJobActionRequest) (*models.RecoveryPlanJobResponse, error) {
	path := fmt.Sprintf("/recovery_plan_jobs/%s/%s", uuid, action)
	req, err := op.client.NewRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	response := new(models.RecoveryPlanJobResponse)
	if err := op.client.Do(ctx, req, response); err != nil {
		return nil, err
	}

	return response, nil
}

/*Get a recovery plan job.
 * This operation gets the recovery plan job identified by 'uuid'.
 *
 * @param uuid UUID of the recovery plan job.
 */
func (op Operations) GetRecoveryPlanJob(ctx context.Context, uuid string) (*models.RecoveryPlanJobIntentResponse, error) {
	path := fmt.Sprintf("/recovery_plan_jobs/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new(models.RecoveryPlanJobIntentResponse)
	if err := op.client.Do(ctx, req, response); err != nil {
		return nil, err
	}

	return response, nil
}

/*Get the status of a recovery plan job.
 * This operation gets the execution status of a recovery plan job identified by 'uuid'.
 *
 * @param uuid UUID of the recovery plan job.
 * @param status is one of {'execution_status', 'cleanup_status'}
 */
func (op Operations) GetRecoveryPlanJobStatus(ctx context.Context, uuid string, status string) (*models.RecoveryPlanJobExecutionStatus, error) {
	path := fmt.Sprintf("/recovery_plan_jobs/%s/%s", uuid, status)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	executionStatus := new(models.RecoveryPlanJobExecutionStatus)
	if err := op.client.Do(ctx, req, executionStatus); err != nil {
		return nil, err
	}

	return executionStatus, nil
}

/*List recovery plan jobs.
 * This Operations lists the recovery plan jobs matching the criteria specified in 'request'.
 *
 * @param request pointer to specification of type DSMetadata.
 */
func (op Operations) ListRecoveryPlanJobs(ctx context.Context, request *models.RecoveryPlanListMetadata) (*models.RecoveryPlanJobListIntentResponse, error) {
	path := "/recovery_plan_jobs/list"
	req, err := op.client.NewRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	list := new(models.RecoveryPlanJobListIntentResponse)
	if err := op.client.Do(ctx, req, list); err != nil {
		return nil, err
	}

	return list, nil
}

/*Get a projection of the attributes of entities of certain type.
 * This operation returns the attributes of the entities that match the
 * filter criteria specified in 'request'.
 *
 * @param request pointer to the specification of type GroupsGetEntitiesRequest
 */
func (op Operations) GroupsGetEntities(ctx context.Context, request *models.GroupsGetEntitiesRequest) (*models.GroupsGetEntitiesResponse, error) {
	req, err := op.client.NewRequest(http.MethodPost, "/groups", request)
	if err != nil {
		return nil, err
	}

	response := new(models.GroupsGetEntitiesResponse)
	if err := op.client.Do(ctx, req, response); err != nil {
		return nil, err
	}

	return response, nil
}

/*Get information about an availability zone (AZ).
 * This operation gets the information about an AZ identified by 'uuid'.
 *
 * @param uuid UUID of the AZ.
 */
func (op Operations) GetAvailabilityZone(ctx context.Context, uuid string) (*models.AvailabilityZoneIntentResponse, error) {
	path := fmt.Sprintf("/availability_zones/%s", uuid)
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new(models.AvailabilityZoneIntentResponse)
	if err := op.client.Do(ctx, req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetPrismCentral gets the information about the Prism Central
func (op Operations) GetPrismCentral(ctx context.Context) (*models.PrismCentral, error) {
	path := "/prism_central"
	req, err := op.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new(models.PrismCentral)
	if err := op.client.Do(ctx, req, response); err != nil {
		return nil, err
	}

	return response, nil
}
