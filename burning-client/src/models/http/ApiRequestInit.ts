import { IApiRequestOptions } from 'http';
import Cookies from 'js-cookie';
import config from '@/config';

export default class ApiRequestInit implements RequestInit {
    readonly body?: any;
    readonly method: string;
    readonly creadentials: RequestCredentials = 'same-origin';
    readonly mode: RequestMode = 'cors';
    options: IApiRequestOptions;
    headers: HeadersInit;
    url: string;

    constructor(url: string, method: string, data: any = null, options?: IApiRequestOptions) {
        let token = Cookies.get(config.JWT_NAME);
        if (token) token = config.JWT_NAME + " " + token;
        this.url = url;
        this.method = method;
        this.options = options || {};
        this.headers = {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'Authorization': token ? token : null,
            'X-Requested-With': 'XMLHttpRequest'
        };
        if (method !== 'GET' && data != null) {
            this.body = JSON.stringify(data);
        }
    }

}