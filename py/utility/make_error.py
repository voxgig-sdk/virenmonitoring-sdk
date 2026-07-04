# Virenmonitoring SDK utility: make_error

from __future__ import annotations
from core.operation import VirenmonitoringOperation
from core.result import VirenmonitoringResult
from core.control import VirenmonitoringControl
from core.error import VirenmonitoringError


def make_error_util(ctx, err):
    if ctx is None:
        from core.context import VirenmonitoringContext
        ctx = VirenmonitoringContext({}, None)

    op = ctx.op
    if op is None:
        op = VirenmonitoringOperation({})
    opname = op.name
    if opname == "" or opname == "_":
        opname = "unknown operation"

    result = ctx.result
    if result is None:
        result = VirenmonitoringResult({})
    result.ok = False

    if err is None:
        err = result.err
    if err is None:
        err = ctx.make_error("unknown", "unknown error")

    errmsg = ""
    if isinstance(err, VirenmonitoringError):
        errmsg = err.msg
    elif hasattr(err, "msg") and err.msg is not None:
        errmsg = err.msg
    elif isinstance(err, str):
        errmsg = err
    else:
        errmsg = str(err)

    msg = "VirenmonitoringSDK: " + opname + ": " + errmsg
    msg = ctx.utility.clean(ctx, msg)

    result.err = None

    spec = ctx.spec

    if ctx.ctrl.explain is not None:
        ctx.ctrl.explain["err"] = {"message": msg}

    sdk_err = VirenmonitoringError("", msg, ctx)
    sdk_err.result = ctx.utility.clean(ctx, result)
    sdk_err.spec = ctx.utility.clean(ctx, spec)

    if isinstance(err, VirenmonitoringError):
        sdk_err.code = err.code

    ctx.ctrl.err = sdk_err

    if ctx.ctrl.throw_err is False:
        return result.resdata

    raise sdk_err
