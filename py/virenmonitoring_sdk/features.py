# Virenmonitoring SDK feature factory

from virenmonitoring_sdk.feature.base_feature import VirenmonitoringBaseFeature
from virenmonitoring_sdk.feature.test_feature import VirenmonitoringTestFeature


def _make_feature(name):
    features = {
        "base": lambda: VirenmonitoringBaseFeature(),
        "test": lambda: VirenmonitoringTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
