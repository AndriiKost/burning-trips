import { isPropertyAccessExpression } from "typescript";

export default class Requests {
    static buildUrl() {
        return process.env.VUE_APP_API_URL;
    }
}