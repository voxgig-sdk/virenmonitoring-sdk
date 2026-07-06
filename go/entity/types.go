// Typed models for the Virenmonitoring SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// DatasetMetadata is the typed data model for the dataset_metadata entity.
type DatasetMetadata struct {
	Description *string `json:"description,omitempty"`
	Label *string `json:"label,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// DatasetMetadataListMatch is the typed request payload for DatasetMetadata.ListTyped.
type DatasetMetadataListMatch struct {
	Description *string `json:"description,omitempty"`
	Label *string `json:"label,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// VirusMonitoring is the typed data model for the virus_monitoring entity.
type VirusMonitoring struct {
	Datasetid *string `json:"datasetid,omitempty"`
	Field *map[string]any `json:"field,omitempty"`
	RecordTimestamp *string `json:"record_timestamp,omitempty"`
	Recordid *string `json:"recordid,omitempty"`
}

// VirusMonitoringListMatch is the typed request payload for VirusMonitoring.ListTyped.
type VirusMonitoringListMatch struct {
	Datasetid *string `json:"datasetid,omitempty"`
	Field *map[string]any `json:"field,omitempty"`
	RecordTimestamp *string `json:"record_timestamp,omitempty"`
	Recordid *string `json:"recordid,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
