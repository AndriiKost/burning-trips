declare module 'http' {
    export interface IApiResponse<T> extends Response {
        result: T;
    }

    export interface IApiRequestOptions {
        supressErrors?: boolean;
        supressSpinner?: boolean;
    }

    export interface ISignInRequest {
        username: string;
        password: string;
    }
    
}