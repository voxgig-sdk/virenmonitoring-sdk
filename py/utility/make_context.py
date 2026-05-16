# Virenmonitoring SDK utility: make_context

from core.context import VirenmonitoringContext


def make_context_util(ctxmap, basectx):
    return VirenmonitoringContext(ctxmap, basectx)
