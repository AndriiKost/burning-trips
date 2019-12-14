import BaseHttpService from './BaseHttpService';

class UploadService extends BaseHttpService {

    async getPresignedUploadUrl(): Promise<string> {
        const res = await this.get<string>(`/file-upload/get-presigned-url`);
        if (!res.ok || !res) return '';
        return res.result;
    }
    
}

export default new UploadService();