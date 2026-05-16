<?php
declare(strict_types=1);

// Virenmonitoring SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class VirenmonitoringFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new VirenmonitoringBaseFeature();
            case "test":
                return new VirenmonitoringTestFeature();
            default:
                return new VirenmonitoringBaseFeature();
        }
    }
}
