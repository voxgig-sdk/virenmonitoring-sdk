package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Virenmonitoring",
			"slug": "virenmonitoring",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://data.bs.ch/api",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"dataset_metadata": map[string]any{},
				"virus_monitoring": map[string]any{},
			},
		},
		"entity": map[string]any{
			"dataset_metadata": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "description",
						"short": "Field description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "label",
						"short": "Field label",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"short": "Field name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"short": "Field data type",
						"type": "`$STRING`",
					},
				},
				"name": "dataset_metadata",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "de",
											"kind": "query",
											"name": "lang",
											"orig": "lang",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/datasets/1.0/100304/",
								"segments": []any{
									map[string]any{
										"lit": "datasets",
									},
									map[string]any{
										"lit": "1.0",
									},
									map[string]any{
										"lit": "100304",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"lang",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"datasets",
									"1.0",
									"100304",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"virus_monitoring": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "datasetid",
						"short": "Dataset identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fields",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"format": "date-time",
						"name": "record_timestamp",
						"short": "Timestamp when the record was created/updated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "recordid",
						"short": "Unique record identifier",
						"type": "`$STRING`",
					},
				},
				"name": "virus_monitoring",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "100304",
											"kind": "query",
											"name": "dataset",
											"orig": "dataset",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "exclude",
											"orig": "exclude",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "de",
											"kind": "query",
											"name": "lang",
											"orig": "lang",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "refine",
											"orig": "refine",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 10,
											"kind": "query",
											"name": "row",
											"orig": "row",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "start",
											"orig": "start",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/records/1.0/search/",
								"segments": []any{
									map[string]any{
										"lit": "records",
									},
									map[string]any{
										"lit": "1.0",
									},
									map[string]any{
										"lit": "search",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"dataset",
										"exclude",
										"lang",
										"q",
										"refine",
										"row",
										"sort",
										"start",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"records",
									"1.0",
									"search",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
