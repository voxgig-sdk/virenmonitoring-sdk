import { VirenmonitoringEntityBase } from '../VirenmonitoringEntityBase';
import type { VirenmonitoringSDK } from '../VirenmonitoringSDK';
import type { Control } from '../types';
import type { DatasetMetadata, DatasetMetadataListMatch } from '../VirenmonitoringTypes';
declare class DatasetMetadataEntity extends VirenmonitoringEntityBase<DatasetMetadata> {
    constructor(client: VirenmonitoringSDK, entopts: any);
    make(this: DatasetMetadataEntity): DatasetMetadataEntity;
    list(this: any, reqmatch?: DatasetMetadataListMatch, ctrl?: Control): Promise<DatasetMetadataEntity[]>;
}
export { DatasetMetadataEntity };
