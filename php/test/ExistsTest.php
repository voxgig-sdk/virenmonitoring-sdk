<?php
declare(strict_types=1);

// Virenmonitoring SDK exists test

require_once __DIR__ . '/../virenmonitoring_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = VirenmonitoringSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
