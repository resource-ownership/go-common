package shared

import "github.com/google/uuid"

var (
	// Default TenantID for TeamPRO (random and valid UUID)
	TeamPROTenantID = uuid.MustParse("a3a80810-f91c-4391-9eff-6d47a13bebde")

	// Default ClientID for TeamPRO (random and valid UUID)
	TeamPROAppClientID = uuid.MustParse("ff96c01f-a741-4429-a0cd-2868d408c42f")

	ServerClientID = uuid.MustParse("ff96c01f-a741-4429-a0cd-2868d408c42f")
)

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
