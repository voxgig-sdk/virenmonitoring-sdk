package = "voxgig-sdk-virenmonitoring"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/virenmonitoring-sdk.git"
}
description = {
  summary = "Virenmonitoring SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["virenmonitoring_sdk"] = "virenmonitoring_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
