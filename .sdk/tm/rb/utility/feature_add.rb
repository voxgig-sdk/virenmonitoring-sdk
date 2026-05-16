# Virenmonitoring SDK utility: feature_add
module VirenmonitoringUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
