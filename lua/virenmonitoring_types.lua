-- Typed models for the Virenmonitoring SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class DatasetMetadata
---@field description? string
---@field label? string
---@field name? string
---@field type? string

---@class DatasetMetadataListMatch
---@field description? string
---@field label? string
---@field name? string
---@field type? string

---@class VirusMonitoring
---@field datasetid? string
---@field fields? table
---@field record_timestamp? string
---@field recordid? string

---@class VirusMonitoringListMatch
---@field datasetid? string
---@field fields? table
---@field record_timestamp? string
---@field recordid? string

local M = {}

return M
