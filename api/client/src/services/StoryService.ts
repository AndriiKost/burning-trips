import { IStory } from '@/types/Story';
import BaseHttpService from './BaseHttpService';

class StoryService extends BaseHttpService {

    async createStory(story: IStory) {
        const res = await this.post<IStory>('/story', story);
        if (!res.ok || !res) return null
        return res.result;
    }

    async getAllStories() {
        const res = await this.get<IStory[]>('/story');
        if (!res.ok || !res) return [];
        return res.result;
    }

    async getStory(id: number) {
        const res = await this.get<IStory>(`/story/${id}`);
        if (!res.ok || !res) return null;
        return res.result;
    }

    async getUploadUrl() {
        const res = await this.get<any>(`/file-upload/get-presigned-url`);
        if (!res.ok || !res) return null;
        return res.result;
    }
    
}

export default new StoryService();