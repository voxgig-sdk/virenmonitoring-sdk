<?php
declare(strict_types=1);

// Virenmonitoring SDK utility: result_body

class VirenmonitoringResultBody
{
    public static function call(VirenmonitoringContext $ctx): ?VirenmonitoringResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
