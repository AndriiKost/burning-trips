import { IApiRequestOptions } from 'http';
import Cookies from 'js-cookie';
import config from '@/config';

export default class ApiRequestInit implements RequestInit {
    readonly body?: any;
    readonly method: string;
    readonly creadentials: RequestCredentials;
    readonly mode: RequestMode = 'cors';
    options: IApiRequestOptions;
    headers: HeadersInit;
    url: string;

    constructor(url: string, method: string, data: any = null, options?: IApiRequestOptions) {
        const token = Cookies.get(config.JWT_NAME);
        this.url = url;
        this.method = method;
        this.options = options || {};
        this.headers = {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'Authorization': token ? `Bearer ${token}` : null,
            'X-Requested-With': 'XMLHttpRequest'
        };
        if (method !== 'GET' && data != null) {
            this.body = JSON.stringify(data);
        }
    }

}