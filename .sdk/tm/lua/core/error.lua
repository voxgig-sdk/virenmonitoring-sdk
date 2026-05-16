-- Virenmonitoring SDK error

local VirenmonitoringError = {}
VirenmonitoringError.__index = VirenmonitoringError


function VirenmonitoringError.new(code, msg, ctx)
  local self = setmetatable({}, VirenmonitoringError)
  self.is_sdk_error = true
  self.sdk = "Virenmonitoring"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function VirenmonitoringError:error()
  return self.msg
end


function VirenmonitoringError:__tostring()
  return self.msg
end


return VirenmonitoringError
