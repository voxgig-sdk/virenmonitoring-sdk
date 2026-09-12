import { Context } from './Context';
declare class VirenmonitoringError extends Error {
    isVirenmonitoringError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { VirenmonitoringError };
