# Virenmonitoring SDK utility: make_context

from virenmonitoring_sdk.core.context import VirenmonitoringContext


def make_context_util(ctxmap, basectx):
    return VirenmonitoringContext(ctxmap, basectx)
