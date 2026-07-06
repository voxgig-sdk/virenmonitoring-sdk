// Typed models for the Virenmonitoring SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface DatasetMetadata {
  description?: string
  label?: string
  name?: string
  type?: string
}

export interface DatasetMetadataListMatch {
  description?: string
  label?: string
  name?: string
  type?: string
}

export interface VirusMonitoring {
  datasetid?: string
  field?: Record<string, any>
  record_timestamp?: string
  recordid?: string
}

export interface VirusMonitoringListMatch {
  datasetid?: string
  field?: Record<string, any>
  record_timestamp?: string
  recordid?: string
}

