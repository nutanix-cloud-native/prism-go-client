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

// K8sClusterInfo K8s cluster info
type K8sClusterInfo struct {
	// k8s distribution
	// Enum: [Openshift NKE CAPX]
	K8sDistribution string `json:"k8s_distribution,omitempty"`

	// k8s version
	K8sVersion string `json:"k8s_version,omitempty"`
}

// K8sClusterAddonInfo K8s cluster addon info
type K8sClusterAddonInfo struct {
	// addon name
	AddonName string `json:"addon_name,omitempty"`

	// addon version
	AddonVersion string `json:"addon_version,omitempty"`
}

// K8sClusterRegistration K8s cluster registration details.
type K8sClusterRegistration struct {
	// addons info
	AddonsInfo []*K8sClusterAddonInfo `json:"addons_info,omitempty"`

	// Categories for this k8s cluster. This allows setting up multiple values from a single key.
	CategoriesMapping map[string]string `json:"categories_mapping,omitempty"`

	// cluster info
	ClusterInfo *K8sClusterInfo `json:"cluster_info,omitempty"`

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

type K8sClusterRegistrationList []*K8sClusterRegistration

type K8sCreateClusterRegistrationRequest struct {
	// Categories for this k8s cluster. This allows setting up multiple values from a single key.
	CategoriesMapping map[string]string `json:"categories_mapping,omitempty"`

	// identity
	Identity *K8sIdentity `json:"identity,omitempty"`

	// metadata
	// Required: true
	Metadata *Metadata `json:"metadata"`

	// Unique name of the k8s cluster.
	// Required: true
	// Max Length: 40
	// Min Length: 1
	// Pattern: [a-z0-9]([-a-z0-9]*[a-z0-9])?
	Name *string `json:"name"`

	// The universally unique identifier (UUID) of the k8s cluster.
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	UUID string `json:"uuid,omitempty"`
}

type K8sIdentity struct {
	// kind
	// Required: true
	// Enum: [user]
	Kind *string `json:"kind"`

	// name
	// Max Length: 1024
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// uuid
	// Required: true
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
	ClusterInfo *K8sClusterInfo `json:"cluster_info,omitempty"`
}

// K8sUpdateClusterRegistrationInfoResponse k8s update cluster registration info response
type K8sUpdateClusterRegistrationInfoResponse struct {
	// cluster name
	// Required: true
	ClusterName *string `json:"cluster_name"`

	// cluster uuid
	// Required: true
	ClusterUUID *string `json:"cluster_uuid"`
}

// K8sUpdateClusterRegistrationAddonInfoRequest k8s update cluster registration addon info request
type K8sUpdateClusterRegistrationAddonInfoRequest struct {
	// cluster addon info
	ClusterAddonInfo *K8sClusterAddonInfo `json:"cluster_addon_info,omitempty"`
}

// K8sUpdateClusterRegistrationAddonInfoResponse k8s update cluster registration addon info response
type K8sUpdateClusterRegistrationAddonInfoResponse struct {
	// cluster name
	// Required: true
	ClusterName *string `json:"cluster_name"`

	// cluster uuid
	// Required: true
	ClusterUUID *string `json:"cluster_uuid"`
}
