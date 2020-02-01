import config from '@/config';
import ApiRequestInit from '@/models/http/ApiRequestInit';
import { IApiRequestOptions, IApiResponse } from 'http';

export default class BaseHttpService {
	
	async get<T>(url: string, options?: IApiRequestOptions): Promise<IApiResponse<T>> {
		return this._makeRequest(url, 'GET', null, options);
	}

	async post<T>(url: string, data: any, options?: IApiRequestOptions): Promise<IApiResponse<T>> {
		return this._makeRequest(url, 'POST', data, options);
	}

	async put<T>(url: string, data: any, options?: IApiRequestOptions): Promise<IApiResponse<T>> {
		return this._makeRequest(url, 'PUT', data, options);
	}

	private async _makeRequest<T>(url: string, method: string, data?: any, options?: IApiRequestOptions): Promise<IApiResponse<T>> {
		if (url.indexOf('://') === -1) url = config.API_URL + url;
		let requestInit = new ApiRequestInit(url, method, data, options);
		let request = new Request(url, requestInit);
		let response: IApiResponse<T>;
		options = options || requestInit.options;
		try {
			let res = await fetch(request);
			response = res as IApiResponse<T>;
			if (res && res.ok) response.result = await res.json();
			else this._handleBadResponse<T>(requestInit, res, options.supressErrors);
			return response;
		} catch(err) {
			this._handleError(err, options.supressErrors);
		}
	}

	private _handleError(err: Error, supressErrors: boolean) {
		console.error('the error => ', err);
	}

	private async _handleBadResponse<T>(request: ApiRequestInit, response: Response, supressErrors: boolean): Promise<IApiResponse<T>> {
		let res = response as IApiResponse<T>;
		console.error('received non 200 response with body => ', res.body);
		if (supressErrors) return res;
		try {
			let error = await res.json();
			error.request = request;
		} catch (err) {
			err.request = request;
			this._handleError(err, supressErrors);
		}
	}
}