import config from '@/config';
import { IUser } from '@/types/User';
import { ISignInRequest } from 'http';
import Cookies from 'js-cookie';
import BaseHttpService from './BaseHttpService';

class AuthService extends BaseHttpService {
    
    async signIn(signInRequest: ISignInRequest) {
        const res = await this.post<string>('/auth/login', signInRequest)
        if (!res || !res.ok) return null;
        const token = res.result;
        this.addJWTCookie(token);
        const user = await this.getUserData();
        if (!user) return null;
        return user;
    }

    private addJWTCookie(jwt: string) {
        Cookies.set(config.JWT_NAME, jwt)
    }

    async getUserData(): Promise<IUser> {
        const res = await this.get<IUser>('/auth/loggedInUser')
        if (!res || !res.ok) return null;
        return res.result;
    }
    
}

export default new AuthService();