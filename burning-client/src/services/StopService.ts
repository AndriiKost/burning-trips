import BaseHttpService from './BaseHttpService';
import { IStop } from '../types/Stop';

class StopService extends BaseHttpService {

    async createStop(stop: IStop) {
        const res = await this.post<IStop>('/stops', stop);
        if (!res.ok || !res) return null
        return res.result;
    }
    
}

export default new StopService();