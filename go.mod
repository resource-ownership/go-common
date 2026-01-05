module github.com/resource-ownership/go-common

go 1.25.4

require (
	github.com/google/uuid v1.6.0
	github.com/replay-api/replay-common v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/replay-api/replay-common => ../../replay-common
