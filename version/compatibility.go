package version

import (
	"fmt"
)

// NamespaceRequirements defines the minimum version requirements for a given API namespace
type NamespaceRequirements struct {
	MinAOSVersion  string
	MinPCVersion   string
	AdditionalReqs map[string]string // For additional product requirements like Files, Objects, Flow, etc.
	IsGA           bool              // Whether the namespace is GA or RC
}

// Namespace represents a Nutanix v4 API namespace
type Namespace string

// V4 API namespaces
const (
	Aiops          Namespace = "aiops"
	Clustermgmt    Namespace = "clustermgmt"
	Datapolicies   Namespace = "datapolicies"
	Dataprotection Namespace = "dataprotection"
	Files          Namespace = "files"
	IAM            Namespace = "iam"
	Licensing      Namespace = "licensing"
	Lifecycle      Namespace = "lifecycle"
	Microseg       Namespace = "microseg"
	Monitoring     Namespace = "monitoring"
	Networking     Namespace = "networking"
	Objects        Namespace = "objects"
	Opsmgmt        Namespace = "opsmgmt"
	Prism          Namespace = "prism"
	Security       Namespace = "security"
	VMM            Namespace = "vmm"
	Volumes        Namespace = "volumes"
)

// V4APIRequirements defines the version requirements for all v4 API namespaces
var V4APIRequirements = map[Namespace]NamespaceRequirements{
	Aiops: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		IsGA:          true,
	},
	Clustermgmt: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		IsGA:          true,
	},
	Datapolicies: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		IsGA:          true,
	},
	Dataprotection: {
		MinAOSVersion: "6.8.0",
		MinPCVersion:  "2024.3.0",
		IsGA:          true,
	},
	Files: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		AdditionalReqs: map[string]string{
			"FileServer":   "5.1.0",
			"FilesManager": "5.1.0",
		},
		IsGA: true,
	},
	IAM: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		IsGA:          true,
	},
	Licensing: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		IsGA:          true,
	},
	Lifecycle: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		AdditionalReqs: map[string]string{
			"LCM": "3.1.0",
		},
		IsGA: true,
	},
	Microseg: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		AdditionalReqs: map[string]string{
			"Flow": "4.2.0",
		},
		IsGA: true,
	},
	Monitoring: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		AdditionalReqs: map[string]string{
			"NCC": "5.1.0",
		},
		IsGA: true,
	},
	Networking: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		IsGA:          true,
	},
	Objects: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		AdditionalReqs: map[string]string{
			"ObjectsManager": "5.1.0",
		},
		IsGA: false, // RC status
	},
	Opsmgmt: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		IsGA:          true,
	},
	Prism: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		IsGA:          true,
	},
	Security: {
		MinAOSVersion: "7.0.0",
		MinPCVersion:  "2024.3.0",
		IsGA:          false, // RC status
	},
	VMM: {
		MinAOSVersion: "6.8.0", // For AHV
		MinPCVersion:  "2024.3.0",
		IsGA:          true,
	},
	Volumes: {
		MinAOSVersion: "6.1.0",
		MinPCVersion:  "2024.3.0",
		IsGA:          true,
	},
}

// CompatibilityResult represents a compatibility check result
type CompatibilityResult struct {
	Compatible      bool
	Namespace       Namespace
	Reason          string
	IsGA            bool
	MissingVersions []string
}

// CheckNamespaceCompatibility checks if a given namespace is compatible with the provided versions
func CheckNamespaceCompatibility(namespace Namespace, aosVersion, pcVersion string, additionalVersions map[string]string) CompatibilityResult {
	result := CompatibilityResult{
		Compatible: true,
		Namespace:  namespace,
		IsGA:       V4APIRequirements[namespace].IsGA,
	}

	// Check if the namespace exists in our requirements
	reqs, exists := V4APIRequirements[namespace]
	if !exists {
		result.Compatible = false
		result.Reason = fmt.Sprintf("unknown namespace: %s", namespace)
		return result
	}

	// Check AOS version
	if aosVersion != "" {
		aosV, err := ParseVersion(aosVersion)
		if err != nil {
			result.Compatible = false
			result.Reason = fmt.Sprintf("invalid AOS version format: %s", aosVersion)
			result.MissingVersions = append(result.MissingVersions, fmt.Sprintf("AOS %s", reqs.MinAOSVersion))
			return result
		}

		minAOSV, err := ParseVersion(reqs.MinAOSVersion)
		if err != nil {
			result.Compatible = false
			result.Reason = "internal error: invalid minimum AOS version format"
			return result
		}

		isGreaterOrEqual, err := aosV.GreaterThanOrEqual(minAOSV)
		if err != nil || !isGreaterOrEqual {
			result.Compatible = false
			result.Reason = fmt.Sprintf("AOS version %s is lower than required version %s", aosVersion, reqs.MinAOSVersion)
			result.MissingVersions = append(result.MissingVersions, fmt.Sprintf("AOS %s", reqs.MinAOSVersion))
		}
	} else {
		result.Compatible = false
		result.Reason = "AOS version not provided"
		result.MissingVersions = append(result.MissingVersions, fmt.Sprintf("AOS %s", reqs.MinAOSVersion))
	}

	// Check PC version
	if pcVersion != "" {
		pcV, err := ParseVersion(pcVersion, WithComponentType(TypePC))
		if err != nil {
			result.Compatible = false
			result.Reason = fmt.Sprintf("invalid PC version format: %s", pcVersion)
			result.MissingVersions = append(result.MissingVersions, fmt.Sprintf("PC %s", reqs.MinPCVersion))
			return result
		}

		minPCV, err := ParseVersion(reqs.MinPCVersion, WithComponentType(TypePC))
		if err != nil {
			result.Compatible = false
			result.Reason = "internal error: invalid minimum PC version format"
			return result
		}

		isGreaterOrEqual, err := pcV.GreaterThanOrEqual(minPCV)
		if err != nil || !isGreaterOrEqual {
			result.Compatible = false
			result.Reason = fmt.Sprintf("PC version %s is lower than required version %s", pcVersion, reqs.MinPCVersion)
			result.MissingVersions = append(result.MissingVersions, fmt.Sprintf("PC %s", reqs.MinPCVersion))
		}
	} else {
		result.Compatible = false
		result.Reason = "PC version not provided"
		result.MissingVersions = append(result.MissingVersions, fmt.Sprintf("PC %s", reqs.MinPCVersion))
	}

	// Check additional requirements
	for product, minVersion := range reqs.AdditionalReqs {
		if version, exists := additionalVersions[product]; exists {
			productV, err := ParseVersion(version)
			if err != nil {
				result.Compatible = false
				result.Reason = fmt.Sprintf("invalid %s version format: %s", product, version)
				result.MissingVersions = append(result.MissingVersions, fmt.Sprintf("%s %s", product, minVersion))
				continue
			}

			minProductV, err := ParseVersion(minVersion)
			if err != nil {
				result.Compatible = false
				result.Reason = fmt.Sprintf("internal error: invalid minimum %s version format", product)
				continue
			}

			isGreaterOrEqual, err := productV.GreaterThanOrEqual(minProductV)
			if err != nil || !isGreaterOrEqual {
				result.Compatible = false
				result.Reason = fmt.Sprintf("%s version %s is lower than required version %s", product, version, minVersion)
				result.MissingVersions = append(result.MissingVersions, fmt.Sprintf("%s %s", product, minVersion))
			}
		} else {
			result.Compatible = false
			result.Reason = fmt.Sprintf("required %s version not provided", product)
			result.MissingVersions = append(result.MissingVersions, fmt.Sprintf("%s %s", product, minVersion))
		}
	}

	return result
}

// GetAllCompatibleNamespaces returns all compatible namespaces for the given versions
func GetAllCompatibleNamespaces(aosVersion, pcVersion string, additionalVersions map[string]string) map[Namespace]CompatibilityResult {
	results := make(map[Namespace]CompatibilityResult)

	for namespace := range V4APIRequirements {
		results[namespace] = CheckNamespaceCompatibility(namespace, aosVersion, pcVersion, additionalVersions)
	}

	return results
}

// GetCompatibilityError returns a formatted error message for incompatible namespaces
func GetCompatibilityError(result CompatibilityResult) string {
	if result.Compatible {
		return ""
	}

	if !result.IsGA {
		return fmt.Sprintf("Namespace '%s' is in Release Candidate (RC) status and not ready for production use", result.Namespace)
	}

	return fmt.Sprintf("Namespace '%s' is not compatible: %s. Required versions: %s",
		result.Namespace,
		result.Reason,
		formatMissingVersions(result.MissingVersions))
}

// formatMissingVersions formats the missing versions list
func formatMissingVersions(versions []string) string {
	result := ""
	for i, v := range versions {
		if i > 0 {
			result += ", "
		}
		result += v
	}
	return result
}
