# frozen_string_literal: true

# Typed models for the Virenmonitoring SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# DatasetMetadata entity data model.
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] label
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
DatasetMetadata = Struct.new(
  :description,
  :label,
  :name,
  :type,
  keyword_init: true
)

# Request payload for DatasetMetadata#list.
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] label
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
DatasetMetadataListMatch = Struct.new(
  :description,
  :label,
  :name,
  :type,
  keyword_init: true
)

# VirusMonitoring entity data model.
#
# @!attribute [rw] datasetid
#   @return [String, nil]
#
# @!attribute [rw] field
#   @return [Hash, nil]
#
# @!attribute [rw] record_timestamp
#   @return [String, nil]
#
# @!attribute [rw] recordid
#   @return [String, nil]
VirusMonitoring = Struct.new(
  :datasetid,
  :field,
  :record_timestamp,
  :recordid,
  keyword_init: true
)

# Request payload for VirusMonitoring#list.
#
# @!attribute [rw] datasetid
#   @return [String, nil]
#
# @!attribute [rw] field
#   @return [Hash, nil]
#
# @!attribute [rw] record_timestamp
#   @return [String, nil]
#
# @!attribute [rw] recordid
#   @return [String, nil]
VirusMonitoringListMatch = Struct.new(
  :datasetid,
  :field,
  :record_timestamp,
  :recordid,
  keyword_init: true
)

