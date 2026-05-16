<?php
declare(strict_types=1);

// Virenmonitoring SDK utility: prepare_body

class VirenmonitoringPrepareBody
{
    public static function call(VirenmonitoringContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
