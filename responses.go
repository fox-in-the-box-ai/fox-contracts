package contracts

// HealthResponse is the response for GET /health.
type HealthResponse struct {
	Status string `json:"status"`
}

// ReadyzCheck is a single readiness check result.
type ReadyzCheck struct {
	OK     bool   `json:"ok"`
	Detail string `json:"detail,omitempty"`
}

// ReadyzResponse is the response for GET /readyz.
type ReadyzResponse struct {
	Ready  bool                   `json:"ready"`
	Checks map[string]ReadyzCheck `json:"checks"`
}

// VersionResponse is the response for GET /version.
type VersionResponse struct {
	ContractVersion string `json:"contract_version"`
	ImageDigest     string `json:"image_digest,omitempty"`
	Runtime         string `json:"runtime,omitempty"`
	RuntimeVersion  string `json:"runtime_version,omitempty"`
	OverlayVersion  string `json:"overlay_version,omitempty"`
}

// CapabilitiesResponse is the response for GET /capabilities.
type CapabilitiesResponse struct {
	ContractVersion string          `json:"contract_version"`
	Capabilities    map[string]bool `json:"capabilities"`
}

// SkillsetResponse is the response for GET /skillset (contract v2.0).
type SkillsetResponse struct {
	Name                 string   `json:"name"`
	Version              string   `json:"version"`
	ContractVersion      string   `json:"contract_version"`
	DataSources          []string `json:"data_sources,omitempty"`
	CapabilitiesDeclared []string `json:"capabilities_declared,omitempty"`
}
