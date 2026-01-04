package common

import (
	"testing"
	"time"
)

func TestValidateSearchParameters(t *testing.T) {
	invalidDuration := time.Duration(-1)
	invalidDuration2 := time.Duration(0)

	validDuration := time.Duration(1)
	validDuration2 := time.Duration(2)

	// timeNow := time.Now()

	tests := []struct {
		name              string
		searchParams      []SearchAggregation
		queryableFields   map[string]bool
		expectedError     string
		maxRecursiveDepth int
	}{
		{
			name: "Valid Single Value Param",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							ValueParams: []SearchableValue{
								{Field: "GameID", Values: []interface{}{"123"}},
							},
						},
					},
				},
			},
			queryableFields: map[string]bool{"GameID": true},
			expectedError:   "",
		},
		{
			name: "Invalid Single Value Param",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							ValueParams: []SearchableValue{
								{Field: "InvalidField", Values: []interface{}{"123"}},
							},
						},
					},
				},
			},
			queryableFields: map[string]bool{"GameID": true},
			expectedError:   "filtering on ValueParams field 'InvalidField' is not permitted",
		},
		{
			name: "Valid Wildcard Value Param",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							ValueParams: []SearchableValue{
								{Field: "Header.*", Values: []interface{}{"123"}},
							},
						},
					},
				},
			},
			queryableFields: map[string]bool{"Header.Filestamp": true},
			expectedError:   "",
		},
		{
			name: "Invalid Wildcard Value Param",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							ValueParams: []SearchableValue{
								{Field: "InvalidPrefix.*", Values: []interface{}{"123"}},
							},
						},
					},
				},
			},
			queryableFields: map[string]bool{"Header.Filestamp": true},
			expectedError:   "filtering on ValueParams fields matching 'InvalidPrefix.*' is not permitted",
		},
		{
			name: "Valid Date Param",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							DateParams: []SearchableDateRange{
								{Field: "Timestamp", Min: &time.Time{}, Max: &time.Time{}},
							},
						},
					},
				},
			},
			queryableFields: map[string]bool{"Timestamp": true},
			expectedError:   "",
		},
		{
			name: "Valid Duration Param",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							DurationParams: []SearchableDurationRange{
								{Field: "Duration", Min: &validDuration, Max: &validDuration2},
							},
						},
					},
				},
			},
			queryableFields: map[string]bool{"Duration": true},
			expectedError:   "",
		},
		{
			name: "Invalid Duration Param",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							DurationParams: []SearchableDurationRange{
								{Field: "InvalidField", Min: &invalidDuration, Max: &invalidDuration2},
							},
						},
					},
				},
			},
			queryableFields: map[string]bool{"Duration": true},
			expectedError:   "filtering on DurationParams field 'InvalidField' is not permitted",
		},
		{
			name: "Valid Recursive Depth",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							ValueParams: []SearchableValue{
								{Field: "Header.Filestamp", Values: []interface{}{"HLTV"}},
							},
						},
					},
				},
			},
			queryableFields:   map[string]bool{"Header.Filestamp": true},
			expectedError:     "",
			maxRecursiveDepth: 1,
		},
		{
			name: "Invalid Recursive Depth",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							ValueParams: []SearchableValue{
								{Field: "Header.Filestamp", Values: []interface{}{"123"}},
							},
						},
					},
				},
			},
			queryableFields:   map[string]bool{"Header.Filestamp": false},
			expectedError:     "filtering on ValueParams field 'Header.Filestamp' is not permitted",
			maxRecursiveDepth: 0,
		},
		{
			name: "Invalid Wildcard in Nested Field (Date)",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							DateParams: []SearchableDateRange{
								{Field: "Header.InvalidSubField.*", Min: &time.Time{}, Max: &time.Time{}},
							},
						},
					},
				},
			},
			queryableFields: map[string]bool{"Header.Filestamp": true},
			expectedError:   "filtering on DateParams fields matching 'Header.InvalidSubField.*' is not permitted",
		},
		{
			name: "Invalid Wildcard in Nested Field (Duration)",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							DurationParams: []SearchableDurationRange{
								{Field: "Header.InvalidSubField.*", Min: &validDuration, Max: &validDuration2},
							},
						},
					},
				},
			},
			queryableFields: map[string]bool{"Header.Filestamp": true},
			expectedError:   "filtering on DurationParams fields matching 'Header.InvalidSubField.*' is not permitted",
		},
		{
			name: "Disallowed Field (ValueParam) with Wildcard Allowed",
			searchParams: []SearchAggregation{
				{
					Params: []SearchParameter{
						{
							ValueParams: []SearchableValue{
								{Field: "Header.Filestamp", Values: []interface{}{"123"}},
							},
						},
					},
				},
			},
			queryableFields: map[string]bool{"Header.*": true, "Header.Filestamp": false}, // Filestamp disallowed specifically
			expectedError:   "filtering on ValueParams field 'Header.Filestamp' is not permitted",
		},

		// // Invalid Date Range - Start After End
		// {
		// 	name: "Invalid Date Param - Start After End",
		// 	searchParams: []SearchAggregation{
		// 		{
		// 			Params: []SearchParameter{
		// 				{
		// 					DateParams: []SearchableDateRange{
		// 						{Field: "Timestamp", Min: func() *time.Time {
		// 							t := timeNow.Add(time.Microsecond * 1)
		// 							return &t
		// 						}(), Max: &timeNow,
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		queryableFields: map[string]bool{"Timestamp": true},
		// 		expectedError:   "",
		// 	},

		// 	// Invalid Duration Range - Start After End
		// 	{
		// 		name: "Invalid Duration Param - Start After End",
		// 		searchParams: []SearchAggregation{
		// 			{
		// 				Params: []SearchParameter{
		// 					{
		// 						DurationParams: []SearchableDurationRange{
		// 							{Field: "Duration", Min: &validDuration2, Max: &validDuration},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		queryableFields: map[string]bool{"Duration": true},
		// 		expectedError:   "",
		// 	},

		// 	{
		// 		name: "Missing 'Min' or 'Max' in Date Range",
		// 		searchParams: []SearchAggregation{
		// 			{
		// 				Params: []SearchParameter{
		// 					{
		// 						DateParams: []SearchableDateRange{
		// 							{Field: "Timestamp", Min: nil, Max: timeNow}, // Valid
		// 							{Field: "Timestamp", Min: timeNow, Max: nil}, // Valid
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		queryableFields: map[string]bool{"Timestamp": true},
		// 		expectedError:   "",
		// 	},

		// 	// Missing 'Min' or 'Max' in Duration Range - Should be valid
		// 	{
		// 		name: "Missing 'Min' or 'Max' in Duration Range",
		// 		searchParams: []SearchAggregation{
		// 			{
		// 				Params: []SearchParameter{
		// 					{
		// 						DurationParams: []SearchableDurationRange{
		// 							{Field: "Duration", Min: nil, Max: &validDuration},
		// 							{Field: "Duration", Min: &validDuration, Max: nil},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		queryableFields: map[string]bool{"Duration": true},
		// 		expectedError:   "",
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := ValidateSearchParameters(tt.searchParams, tt.queryableFields)
			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("Expected error '%s', but got no error", tt.expectedError)
				} else if err.Error() != tt.expectedError {
					t.Errorf("Expected error '%s', but got '%s'", tt.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got '%s'", err.Error())
				}
			}
		})
	}
}

func TestValidateResultOptions(t *testing.T) {
	tests := []struct {
		name           string
		resultOptions  SearchResultOptions
		readableFields map[string]bool
		expectedError  string
		maxPageSize    uint
	}{
		{
			name:           "Valid Result Options",
			resultOptions:  SearchResultOptions{Skip: 0, Limit: 5},
			readableFields: map[string]bool{"GameID": true},
			expectedError:  "",
		},
		{
			name:           "Invalid Limit",
			resultOptions:  SearchResultOptions{Skip: 0, Limit: 0},
			readableFields: map[string]bool{"GameID": true},
			expectedError:  "limit must be a positive integer",
		},
		{
			name:           "Invalid Pick Field with Wildcard Allowed",
			resultOptions:  SearchResultOptions{PickFields: []string{"Header.NonExistentField"}},
			readableFields: map[string]bool{"Header.*": true},
			expectedError:  "returning field 'Header.NonExistentField' is not permitted (1)",
		},
		{
			name:           "Invalid Omit Field with Wildcard Allowed",
			resultOptions:  SearchResultOptions{OmitFields: []string{"Header.NonExistentField"}},
			readableFields: map[string]bool{"Header.*": true},
			expectedError:  "omitting field 'Header.NonExistentField' is not permitted",
		},
		{
			name:           "Disallowed Pick Field Even with Wildcard",
			resultOptions:  SearchResultOptions{PickFields: []string{"Header.FileStamp"}},
			readableFields: map[string]bool{"Header.*": true, "Header.FileStamp": false},
			expectedError:  "returning field 'Header.FileStamp' is strictly forbidden (2)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := ValidateResultOptions(tt.resultOptions, tt.readableFields)
			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("Expected error '%s', but got no error", tt.expectedError)
				} else if err.Error() != tt.expectedError {
					t.Errorf("Expected error '%s', but got '%s'", tt.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got '%s'", err.Error())
				}
			}
		})
	}
}
