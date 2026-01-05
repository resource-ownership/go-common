package shared

import (
	"testing"
)

// MockSchemaProvider implements SchemaProvider for testing
type MockSchemaProvider struct {
	schema QueryServiceSchema
}

func (m *MockSchemaProvider) GetQuerySchema() QueryServiceSchema {
	return m.schema
}

func TestGetQueryServiceRegistry_Singleton(t *testing.T) {
	registry1 := GetQueryServiceRegistry()
	registry2 := GetQueryServiceRegistry()

	if registry1 != registry2 {
		t.Error("expected GetQueryServiceRegistry to return same instance")
	}
}
