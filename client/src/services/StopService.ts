import { IStop } from '../types/Stop';
import BaseHttpService from './BaseHttpService';

class StopService extends BaseHttpService {

    async createStop(stop: IStop) {
        const res = await this.post<IStop>('/stop', stop);
        if (!res.ok || !res) return null
        return res.result;
    }

    async getAllStops() {
        const res = await this.get<IStop[]>('/stop');
        if (!res.ok || !res) return [];
        return res.result;
    }

    async getStop(id: number) {
        const res = await this.get<IStop>(`/stop/${id}`);
        if (!res.ok || !res) return null;
        return res.result;
    }

    async getUploadUrl() {
        const res = await this.get<any>(`/file-upload/get-presigned-url`);
        if (!res.ok || !res) return null;
        return res.result;
    }
    
}

export default new StopService();
