import BaseHttpService from './BaseHttpService';

class UploadService extends BaseHttpService {

    async getPresignedUploadUrl(fileNameToUpload: string): Promise<string> {
        const res = await this.get<string>(`/file-upload/presign/${fileNameToUpload}`);
        if (!res.ok || !res) return '';
        return res.result;
    }

}

export default new UploadService();