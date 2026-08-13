# Virenmonitoring SDK exists test

import pytest
from virenmonitoring_sdk import VirenmonitoringSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = VirenmonitoringSDK.test(None, None)
        assert testsdk is not None
