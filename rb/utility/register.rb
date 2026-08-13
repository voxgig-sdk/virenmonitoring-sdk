# Virenmonitoring SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

VirenmonitoringUtility.registrar = ->(u) {
  u.clean = VirenmonitoringUtilities::Clean
  u.done = VirenmonitoringUtilities::Done
  u.make_error = VirenmonitoringUtilities::MakeError
  u.feature_add = VirenmonitoringUtilities::FeatureAdd
  u.feature_hook = VirenmonitoringUtilities::FeatureHook
  u.feature_init = VirenmonitoringUtilities::FeatureInit
  u.fetcher = VirenmonitoringUtilities::Fetcher
  u.make_fetch_def = VirenmonitoringUtilities::MakeFetchDef
  u.make_context = VirenmonitoringUtilities::MakeContext
  u.make_options = VirenmonitoringUtilities::MakeOptions
  u.make_request = VirenmonitoringUtilities::MakeRequest
  u.make_response = VirenmonitoringUtilities::MakeResponse
  u.make_result = VirenmonitoringUtilities::MakeResult
  u.make_point = VirenmonitoringUtilities::MakePoint
  u.make_spec = VirenmonitoringUtilities::MakeSpec
  u.make_url = VirenmonitoringUtilities::MakeUrl
  u.param = VirenmonitoringUtilities::Param
  u.prepare_auth = VirenmonitoringUtilities::PrepareAuth
  u.prepare_body = VirenmonitoringUtilities::PrepareBody
  u.prepare_headers = VirenmonitoringUtilities::PrepareHeaders
  u.prepare_method = VirenmonitoringUtilities::PrepareMethod
  u.prepare_params = VirenmonitoringUtilities::PrepareParams
  u.prepare_path = VirenmonitoringUtilities::PreparePath
  u.prepare_query = VirenmonitoringUtilities::PrepareQuery
  u.graphql_body = VirenmonitoringUtilities::GraphqlBody
  u.graphql_errors = VirenmonitoringUtilities::GraphqlErrors
  u.result_basic = VirenmonitoringUtilities::ResultBasic
  u.result_body = VirenmonitoringUtilities::ResultBody
  u.result_headers = VirenmonitoringUtilities::ResultHeaders
  u.transform_request = VirenmonitoringUtilities::TransformRequest
  u.transform_response = VirenmonitoringUtilities::TransformResponse
}
