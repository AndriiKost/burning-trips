import { IFeaturedStops } from '@/types/FeaturedStops';
import BaseHttpService from './BaseHttpService';


export default class ViewModelService {

    private readonly _httpSvc: BaseHttpService;

    constructor(httpSvc: BaseHttpService = new BaseHttpService()) {
        this._httpSvc = httpSvc;
    }

    async getStopViewModel(stopId: number): Promise<IFeaturedStops> {
        const res = await this._httpSvc.get<IFeaturedStops>('/view-models/featured?id=', stopId);
        if (!res.ok || !res) return null;
        return res.result;
    }
}