package shared

import (
"testing"

shared "github.com/resource-ownership/go-common/pkg/common"
)

// MockSchemaProvider implements shared.SchemaProvider for testing
type MockSchemaProvider struct {
	schema shared.QueryServiceSchema
}

func (m *MockSchemaProvider) GetQuerySchema() shared.QueryServiceSchema {
	return m.schema
}

func TestGetQueryServiceRegistry_Singleton(t *testing.T) {
	registry1 := shared.GetQueryServiceRegistry()
	registry2 := shared.GetQueryServiceRegistry()

	if registry1 != registry2 {
		t.Error("expected GetQueryServiceRegistry to return same instance")
	}
}
