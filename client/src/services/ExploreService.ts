import { ISearchQuery } from '@/types/Explore';
import { ILandmark } from '@/types/Landmark';
import { IStop } from '@/types/Stop';
import BaseHttpService from './BaseHttpService';

class ExploreService extends BaseHttpService {

    async searchStops(query: ISearchQuery): Promise<IStop[]> {
        const response = await this.post<IStop[]>('/search-stops', query);
        if (response.ok) {
            return response.result;
        }
    }

    async searchLandmarks(query: ISearchQuery): Promise<ILandmark[]> {
        const response = await this.post<ILandmark[]>('/search-landmarks', query);
        if (response.ok) {
            return response.result;
        }
    }

}

export default new ExploreService();