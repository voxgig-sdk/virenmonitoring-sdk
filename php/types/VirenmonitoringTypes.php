<?php
declare(strict_types=1);

// Typed models for the Virenmonitoring SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** DatasetMetadata entity data model. */
class DatasetMetadata
{
    public ?string $description = null;
    public ?string $label = null;
    public ?string $name = null;
    public ?string $type = null;
}

/** Request payload for DatasetMetadata#list. */
class DatasetMetadataListMatch
{
    public ?string $description = null;
    public ?string $label = null;
    public ?string $name = null;
    public ?string $type = null;
}

/** VirusMonitoring entity data model. */
class VirusMonitoring
{
    public ?string $datasetid = null;
    public ?array $fields = null;
    public ?string $record_timestamp = null;
    public ?string $recordid = null;
}

/** Request payload for VirusMonitoring#list. */
class VirusMonitoringListMatch
{
    public ?string $datasetid = null;
    public ?array $fields = null;
    public ?string $record_timestamp = null;
    public ?string $recordid = null;
}

