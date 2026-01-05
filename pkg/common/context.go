package shared

type ContextKey string

const (
	// Tenancy (internal)
	TenantIDKey ContextKey = "tenant_id"
	ClientIDKey ContextKey = "client_id"
	GroupIDKey  ContextKey = "group_id"
	UserIDKey   ContextKey = "user_id"

	// Parameters
	ResourceIDKey   ContextKey = "resource_id"
	GameIDParamKey  ContextKey = "game_id"

	AudienceKey            ContextKey = "aud"
	AuthenticatedKey       ContextKey = "auth"
	IntendedAudienceCtxKey ContextKey = "intended_audience"
)
