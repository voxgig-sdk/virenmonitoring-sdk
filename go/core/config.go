package core

func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Virenmonitoring",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
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
						"active": true,
						"name": "description",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "label",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "type",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
				},
				"name": "dataset_metadata",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "de",
											"kind": "query",
											"name": "lang",
											"orig": "lang",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/datasets/1.0/100304/",
								"parts": []any{
									"datasets",
									"1.0",
									"100304",
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
								"index$": 0,
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
						"active": true,
						"name": "datasetid",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "fields",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "record_timestamp",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "recordid",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
				},
				"name": "virus_monitoring",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "100304",
											"kind": "query",
											"name": "dataset",
											"orig": "dataset",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "exclude",
											"orig": "exclude",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "de",
											"kind": "query",
											"name": "lang",
											"orig": "lang",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "refine",
											"orig": "refine",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": 10,
											"kind": "query",
											"name": "row",
											"orig": "row",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": 0,
											"kind": "query",
											"name": "start",
											"orig": "start",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/records/1.0/search/",
								"parts": []any{
									"records",
									"1.0",
									"search",
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
								"index$": 0,
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
