package voxgigvirenmonitoringsdk

import (
	"github.com/voxgig-sdk/virenmonitoring-sdk/go/core"
	"github.com/voxgig-sdk/virenmonitoring-sdk/go/entity"
	"github.com/voxgig-sdk/virenmonitoring-sdk/go/feature"
	_ "github.com/voxgig-sdk/virenmonitoring-sdk/go/utility"
)

// Type aliases preserve external API.
type VirenmonitoringSDK = core.VirenmonitoringSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type VirenmonitoringEntity = core.VirenmonitoringEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type VirenmonitoringError = core.VirenmonitoringError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewDatasetMetadataEntityFunc = func(client *core.VirenmonitoringSDK, entopts map[string]any) core.VirenmonitoringEntity {
		return entity.NewDatasetMetadataEntity(client, entopts)
	}
	core.NewVirusMonitoringEntityFunc = func(client *core.VirenmonitoringSDK, entopts map[string]any) core.VirenmonitoringEntity {
		return entity.NewVirusMonitoringEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewVirenmonitoringSDK = core.NewVirenmonitoringSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewVirenmonitoringSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *VirenmonitoringSDK  { return NewVirenmonitoringSDK(nil) }
func Test() *VirenmonitoringSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
