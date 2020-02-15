import { IStory } from '@/types/Story';
import BaseHttpService from './BaseHttpService';

class StoryService extends BaseHttpService {

    async createStory(story: IStory) {
        const res = await this.post<IStory>('/story', story);
        if (res && res.ok) {
            return res.result;
        }
        return null
    }

    async getAllStories() {
        const res = await this.get<IStory[]>('/story');
        if (res && res.ok) {
            return res.result;
        }
        return [];
    }

    async getStory(id: number) {
        const res = await this.get<IStory>(`/story/${id}`);
        if (res && res.ok) {
            return res.result;
        }
        return null;
    }

    async getUploadUrl() {
        const res = await this.get<any>(`/file-upload/get-presigned-url`);
        if (res && res.ok) {
            return res.result;
        }
        return null;
    }
    
}

export default new StoryService();