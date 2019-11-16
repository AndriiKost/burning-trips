import BaseHttpService from './BaseHttpService';
import { ISignInRequest } from 'http';

class AuthService extends BaseHttpService {
    
    async signIn(signInRequest: ISignInRequest) {
        const res = await this.post('/login', signInRequest)
        console.log(res)
    }
    
}

export default new AuthService();