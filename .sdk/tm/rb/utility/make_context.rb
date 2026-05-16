# Virenmonitoring SDK utility: make_context
require_relative '../core/context'
module VirenmonitoringUtilities
  MakeContext = ->(ctxmap, basectx) {
    VirenmonitoringContext.new(ctxmap, basectx)
  }
end
