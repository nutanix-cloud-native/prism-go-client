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

// K8sClusterCounterMapping Cluster counters for this k8s cluster. This allows setting up multiple values from a single key.
type K8sClusterCounterMapping map[string]string

// K8sClusterAddonCounter Addons counters for this k8s cluster. This allows setting up multiple values from a single key.
type K8sClusterAddonCounter map[string]string

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

// K8sUpdateClusterRegistrationCounterRequest k8s update cluster registration counter request
type K8sUpdateClusterRegistrationCounterRequest struct {
	// cluster counter
	// Required: true
	ClusterCounter K8sClusterCounterMapping `json:"cluster_counter"`
}

// K8sUpdateClusterRegistrationCounterResponse k8s update cluster registration counter response
type K8sUpdateClusterRegistrationCounterResponse struct {
	// cluster name
	ClusterName string `json:"cluster_name,omitempty"`
	// cluster uuid
	ClusterUUID string `json:"cluster_uuid,omitempty"`
}

// K8sUpdateClusterRegistrationAddonCounterRequest k8s update cluster registration addon counter request
type K8sUpdateClusterRegistrationAddonCounterRequest struct {
	// cluster addon counter
	// Required: true
	ClusterAddonCounter K8sClusterAddonCounter `json:"cluster_addon_counter"`
}

// K8sUpdateClusterRegistrationAddonCounterResponse k8s update cluster registration addon counter response
type K8sUpdateClusterRegistrationAddonCounterResponse struct {
	// cluster name
	ClusterName string `json:"cluster_name,omitempty"`
	// cluster uuid
	ClusterUUID string `json:"cluster_uuid,omitempty"`
}
