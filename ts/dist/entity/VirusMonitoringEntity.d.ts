import { VirenmonitoringEntityBase } from '../VirenmonitoringEntityBase';
import type { VirenmonitoringSDK } from '../VirenmonitoringSDK';
import type { Control } from '../types';
import type { VirusMonitoring, VirusMonitoringListMatch } from '../VirenmonitoringTypes';
declare class VirusMonitoringEntity extends VirenmonitoringEntityBase<VirusMonitoring> {
    constructor(client: VirenmonitoringSDK, entopts: any);
    make(this: VirusMonitoringEntity): VirusMonitoringEntity;
    list(this: any, reqmatch?: VirusMonitoringListMatch, ctrl?: Control): Promise<VirusMonitoringEntity[]>;
}
export { VirusMonitoringEntity };
