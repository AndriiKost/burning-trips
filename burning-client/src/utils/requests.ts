import { isPropertyAccessExpression } from "typescript";

export default class Requests {
    static buildUrl() {
        return process.env.API_URL;
    }
}