<?php
declare(strict_types=1);

// Virenmonitoring SDK utility: result_headers

class VirenmonitoringResultHeaders
{
    public static function call(VirenmonitoringContext $ctx): ?VirenmonitoringResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
