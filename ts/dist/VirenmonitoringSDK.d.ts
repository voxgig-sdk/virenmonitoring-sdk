import { DatasetMetadataEntity } from './entity/DatasetMetadataEntity';
import { VirusMonitoringEntity } from './entity/VirusMonitoringEntity';
export type * from './VirenmonitoringTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { VirenmonitoringEntityBase } from './VirenmonitoringEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class VirenmonitoringSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    DatasetMetadata(entopts?: Record<string, any>): DatasetMetadataEntity;
    VirusMonitoring(entopts?: Record<string, any>): VirusMonitoringEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): VirenmonitoringSDK;
    tester(testopts?: any, sdkopts?: any): VirenmonitoringSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof VirenmonitoringSDK;
export { stdutil, config, BaseFeature, VirenmonitoringEntityBase, VirenmonitoringSDK, SDK, };
