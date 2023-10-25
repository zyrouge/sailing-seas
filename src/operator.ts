import { App } from "./app";

export interface AppOperator {
    operate(app: App): Promise<void>;
}

export const defineAppOperator = (operator: AppOperator) => operator;
