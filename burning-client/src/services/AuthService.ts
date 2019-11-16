import BaseHttpService from './BaseHttpService';
import { ISignInRequest } from 'http';
import Cookies from 'js-cookie';
import config from '@/config';

class AuthService extends BaseHttpService {
    
    async signIn(signInRequest: ISignInRequest) {
        const res = await this.post('/login', signInRequest)
        if (res.ok) return res.result;
    }

    private addJWTCookie(jwt: string) {
        Cookies.set(config.JWT_NAME, jwt)
    }
    
}

export default new AuthService();