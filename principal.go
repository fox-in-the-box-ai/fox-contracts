package contracts

// FoxPrincipal identifies an authenticated user across the Fox platform.
// Produced by the edge gateway from OIDC JWT claims and forwarded in the
// X-Fox-Principal header as JSON.
type FoxPrincipal struct {
	UserID     string         `json:"user_id"`
	TenantID   string         `json:"tenant_id"`
	Groups     []string       `json:"groups"`
	AuthMethod string         `json:"auth_method"`
	Claims     map[string]any `json:"claims,omitempty"`
}
