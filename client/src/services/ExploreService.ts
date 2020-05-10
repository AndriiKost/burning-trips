import { ISearchQuery } from '@/types/Explore';
import { IStop } from '@/types/Stop';
import BaseHttpService from './BaseHttpService';

class ExploreService extends BaseHttpService {

    async searchStops(query: ISearchQuery): Promise<IStop[]> {
        const response = await this.post<IStop[]>('/search-stops', query);
        if (response.ok) {
            return response.result;
        }
    }

}

export default new ExploreService();