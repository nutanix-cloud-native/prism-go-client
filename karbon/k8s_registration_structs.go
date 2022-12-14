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
