<?php
declare(strict_types=1);

// Virenmonitoring SDK base feature

class VirenmonitoringBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(VirenmonitoringContext $ctx, array $options): void {}
    public function PostConstruct(VirenmonitoringContext $ctx): void {}
    public function PostConstructEntity(VirenmonitoringContext $ctx): void {}
    public function SetData(VirenmonitoringContext $ctx): void {}
    public function GetData(VirenmonitoringContext $ctx): void {}
    public function GetMatch(VirenmonitoringContext $ctx): void {}
    public function SetMatch(VirenmonitoringContext $ctx): void {}
    public function PrePoint(VirenmonitoringContext $ctx): void {}
    public function PreSpec(VirenmonitoringContext $ctx): void {}
    public function PreRequest(VirenmonitoringContext $ctx): void {}
    public function PreResponse(VirenmonitoringContext $ctx): void {}
    public function PreResult(VirenmonitoringContext $ctx): void {}
    public function PreDone(VirenmonitoringContext $ctx): void {}
    public function PreUnexpected(VirenmonitoringContext $ctx): void {}
}
