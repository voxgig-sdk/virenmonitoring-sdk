<?php
declare(strict_types=1);

// Virenmonitoring SDK utility: feature_add

class VirenmonitoringFeatureAdd
{
    public static function call(VirenmonitoringContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
