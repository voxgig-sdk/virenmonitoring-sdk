<?php
declare(strict_types=1);

// Virenmonitoring SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class VirenmonitoringMakeContext
{
    public static function call(array $ctxmap, ?VirenmonitoringContext $basectx): VirenmonitoringContext
    {
        return new VirenmonitoringContext($ctxmap, $basectx);
    }
}
