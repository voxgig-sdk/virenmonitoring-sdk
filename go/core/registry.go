package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewDatasetMetadataEntityFunc func(client *VirenmonitoringSDK, entopts map[string]any) VirenmonitoringEntity

var NewVirusMonitoringEntityFunc func(client *VirenmonitoringSDK, entopts map[string]any) VirenmonitoringEntity

