package shared

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBaseEntity_GetID(t *testing.T) {
	id := uuid.New()
	entity := BaseEntity{ID: id}

	assert.Equal(t, id, entity.GetID())
}

func TestNewEntity(t *testing.T) {
	resourceOwner := NewResourceOwner(uuid.New(), uuid.New(), uuid.New(), uuid.New())

	entity := NewEntity(resourceOwner)

	assert.NotEqual(t, uuid.Nil, entity.ID)
	assert.Equal(t, ClientApplicationAudienceIDKey, entity.VisibilityLevel)
	assert.Equal(t, CustomVisibilityTypeKey, entity.VisibilityType)
	assert.Equal(t, resourceOwner, entity.ResourceOwner)
	assert.True(t, entity.CreatedAt.After(time.Now().Add(-time.Second)))
	assert.True(t, entity.UpdatedAt.After(time.Now().Add(-time.Second)))
	assert.True(t, entity.CreatedAt.Equal(entity.UpdatedAt))
}

func TestNewUnrestrictedEntity(t *testing.T) {
	resourceOwner := NewResourceOwner(uuid.New(), uuid.New(), uuid.New(), uuid.New())

	entity := NewUnrestrictedEntity(resourceOwner)

	assert.NotEqual(t, uuid.Nil, entity.ID)
	assert.Equal(t, TenantAudienceIDKey, entity.VisibilityLevel)
	assert.Equal(t, PublicVisibilityTypeKey, entity.VisibilityType)
	assert.Equal(t, resourceOwner, entity.ResourceOwner)
	assert.True(t, entity.CreatedAt.After(time.Now().Add(-time.Second)))
	assert.True(t, entity.UpdatedAt.After(time.Now().Add(-time.Second)))
	assert.True(t, entity.CreatedAt.Equal(entity.UpdatedAt))
}

func TestBaseEntity_StructFields(t *testing.T) {
	entity := BaseEntity{
		ID:              uuid.New(),
		VisibilityLevel: UserAudienceIDKey,
		VisibilityType:  PrivateVisibilityTypeKey,
		ResourceOwner:   NewResourceOwner(uuid.New(), uuid.New(), uuid.New(), uuid.New()),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Includes:        make(map[string]interface{}),
	}

	assert.NotEqual(t, uuid.Nil, entity.ID)
	assert.Equal(t, UserAudienceIDKey, entity.VisibilityLevel)
	assert.Equal(t, PrivateVisibilityTypeKey, entity.VisibilityType)
	assert.NotEqual(t, uuid.Nil, entity.ResourceOwner.TenantID)
	assert.NotNil(t, entity.Includes)
}