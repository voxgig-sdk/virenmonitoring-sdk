# Virenmonitoring SDK feature factory

from feature.base_feature import VirenmonitoringBaseFeature
from feature.test_feature import VirenmonitoringTestFeature


def _make_feature(name):
    features = {
        "base": lambda: VirenmonitoringBaseFeature(),
        "test": lambda: VirenmonitoringTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
