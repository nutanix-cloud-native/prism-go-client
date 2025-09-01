package karbon

type K8sCreateClusterRegistrationResponse struct {
	// cluster name
	// Required: true
	ClusterName *string `json:"cluster_name"`

	// cluster uuid
	// Required: true
	ClusterUUID *string `json:"cluster_uuid"`

	// task uuid
	// Required: true
	TaskUUID *string `json:"task_uuid"`
}

type K8sClusterRegistrationDeleteResponse struct {
	// Name of the k8s cluster being deleted.
	ClusterName string `json:"cluster_name,omitempty"`

	// UUID of the k8s cluster being deleted.
	ClusterUUID string `json:"cluster_uuid,omitempty"`

	// UUID of the task tracking the k8s cluster deletion.
	// Required: true
	TaskUUID *string `json:"task_uuid"`
}

// K8sClusterAddonInfo Addons information for this k8s cluster. This allows setting up multiple values from a single key.
type K8sClusterAddonInfo map[string]string

// K8sClusterRegistration K8s cluster registration details.
type K8sClusterRegistration struct {
	// addons info
	AddonsInfo K8sClusterAddonInfoMapping `json:"addons_info,omitempty"`

	// categories mapping
	CategoriesMapping K8sClusterCategoriesMapping `json:"categories_mapping,omitempty"`

	// cluster info
	ClusterInfo K8sClusterInfoMapping `json:"cluster_info,omitempty"`

	// The type of 3rd party k8s cluster.
	// Enum: [Openshift NKE CAPX]
	K8sDistribution string `json:"k8s_distribution,omitempty"`

	// K8s cluster name.
	// Required: true
	// Max Length: 40
	// Min Length: 1
	Name *string `json:"name"`

	// K8s cluster registration status.
	Status string `json:"status,omitempty"`

	// The universally unique identifier (UUID) of the k8s cluster.
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	UUID string `json:"uuid,omitempty"`
}

// K8sClusterAddonInfoMapping Addons information for this k8s cluster. This allows setting up multiple values from a single key.
type K8sClusterAddonInfoMapping map[string]K8sClusterAddonInfo

// K8sClusterCategoriesMapping Categories for this k8s cluster. This allows setting up multiple values from a single key.
type K8sClusterCategoriesMapping map[string]string

// K8sClusterInfoMapping Cluster information for this k8s cluster. This allows setting up multiple values from a single key.
type K8sClusterInfoMapping map[string]string

type K8sClusterRegistrationList []*K8sClusterRegistration

type K8sCreateClusterRegistrationRequest struct {
	// categories mapping
	// Required: true
	CategoriesMapping K8sClusterCategoriesMapping `json:"categories_mapping"`

	// The type of 3rd party k8s cluster.
	// Required: true
	// Enum: [Openshift NKE CAPX]
	K8sDistribution *string `json:"k8s_distribution"`

	// metadata
	// Required: true
	Metadata *Metadata `json:"metadata"`

	// Unique name of the k8s cluster.
	// Example: prod-cluster
	// Required: true
	// Max Length: 40
	// Min Length: 1
	// Pattern: [a-z0-9]([-a-z0-9]*[a-z0-9])?
	Name *string `json:"name"`

	// The universally unique identifier (UUID) of the k8s cluster.
	// Required: true
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	UUID *string `json:"uuid"`
}

type Metadata struct {
	// Karbon API version.
	// Required: true
	// Pattern: ^v?(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$
	APIVersion *string `json:"api_version"`
}

// K8sUpdateClusterRegistrationInfoRequest k8s update cluster registration info request
type K8sUpdateClusterRegistrationInfoRequest struct {
	// cluster info
	// Required: true
	ClusterInfo K8sClusterInfoMapping `json:"cluster_info"`
}

// K8sUpdateClusterRegistrationInfoResponse k8s update cluster registration info response
type K8sUpdateClusterRegistrationInfoResponse struct {
	// cluster name
	ClusterName string `json:"cluster_name,omitempty"`
	// cluster uuid
	ClusterUUID string `json:"cluster_uuid,omitempty"`
}

// K8sUpdateClusterRegistrationAddonInfoRequest k8s update cluster registration addon info request
type K8sUpdateClusterRegistrationAddonInfoRequest struct {
	// cluster addon info
	// Required: true
	ClusterAddonInfo K8sClusterAddonInfo `json:"cluster_addon_info"`
}

// K8sUpdateClusterRegistrationAddonInfoResponse k8s update cluster registration addon info response
type K8sUpdateClusterRegistrationAddonInfoResponse struct {
	// cluster name
	ClusterName string `json:"cluster_name,omitempty"`
	// cluster uuid
	ClusterUUID string `json:"cluster_uuid,omitempty"`
}

// K8sUpdateClusterRegistrationMetricsRequest k8s update cluster registration metrics request
type K8sUpdateClusterRegistrationMetricsRequest struct {
	// cluster metrics
	ClusterMetrics K8sClusterMetrics `json:"cluster_metrics"`
}

// K8sUpdateClusterRegistrationMetricsResponse k8s update cluster registration metrics response
type K8sUpdateClusterRegistrationMetricsResponse struct {
	// cluster name
	ClusterName string `json:"cluster_name,omitempty"`
	// cluster uuid
	ClusterUUID string `json:"cluster_uuid,omitempty"`
}

// K8sUpdateClusterRegistrationAddonMetricsRequest k8s update cluster registration addon metrics request
type K8sUpdateClusterRegistrationAddonMetricsRequest struct {
	// cluster addon metrics
	ClusterAddonMetrics K8sClusterAddonMetrics `json:"cluster_addon_metrics"`
}

// K8sUpdateClusterRegistrationAddonMetricsResponse k8s update cluster registration addon metrics response
type K8sUpdateClusterRegistrationAddonMetricsResponse struct {
	// cluster name
	ClusterName string `json:"cluster_name,omitempty"`
	// cluster uuid
	ClusterUUID string `json:"cluster_uuid,omitempty"`
}

// K8sClusterMetrics metrics information gathered for resources under k8s cluster.
type K8sClusterMetrics map[string]K8sClusterResourceList

// K8sClusterAddonMetrics addon metrics information gathered for resources under k8s cluster
type K8sClusterAddonMetrics map[string]K8sClusterResourceList

// K8sClusterResourceList list of the resources gathered under k8s cluster
type K8sClusterResourceList []*K8sClusterResource

// K8sClusterResource k8s cluster resource includes name, UUID, metadata and children resources
type K8sClusterResource struct {
	// child resource
	ChildResource map[string]K8sClusterResourceList `json:"ChildResource,omitempty"`
	// metadata
	Metadata map[string]string `json:"Metadata,omitempty"`
	// name
	Name string `json:"Name,omitempty"`
	// UUID
	UUID string `json:"UUID,omitempty"`
}

// K8sClusterKubeconfigResponse k8s cluster kubeconfig response
type K8sClusterKubeconfigResponse struct {

	// SHA-256 hash of the kubernetes cluster's kubeconfig content
	// Required: true
	KubeconfigChecksum *string `json:"kubeconfig_checksum"`

	// Kubernetes cluster kubeconfig update status.
	TaskStatus string `json:"task_status,omitempty"`

	// The UUID of the task tracking the kubeconfig update.
	TaskUUID string `json:"task_uuid,omitempty"`
}
