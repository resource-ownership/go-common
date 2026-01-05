package shared

type IntendedAudienceKey uint8

type VisibilityTypeKey uint8

const (
	// semantic aliases
	ALLOW = true
	DENY  = false
)

const (
	PublicVisibilityTypeKey VisibilityTypeKey = 1 << iota
	RestrictedVisibilityTypeKey
	PrivateVisibilityTypeKey
	CustomVisibilityTypeKey
)

const (
	UserAudienceIDKey              IntendedAudienceKey = 1 << iota // 1
	GroupAudienceIDKey                                             // 2
	ClientApplicationAudienceIDKey                                 // 4
	TenantAudienceIDKey                                            // 8
)

type ResourceType string
