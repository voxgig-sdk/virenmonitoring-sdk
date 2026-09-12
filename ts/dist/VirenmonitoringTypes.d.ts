export interface DatasetMetadata {
    description?: string;
    label?: string;
    name?: string;
    type?: string;
}
export interface DatasetMetadataListMatch {
    lang?: string;
}
export interface VirusMonitoring {
    datasetid?: string;
    fields?: Record<string, any>;
    record_timestamp?: string;
    recordid?: string;
}
export interface VirusMonitoringListMatch {
    dataset?: string;
    exclude?: string;
    lang?: string;
    q?: string;
    refine?: string;
    row?: number;
    sort?: string;
    start?: number;
}
