import { IStop } from './Stop';

export interface IRegion extends IStop {
    stops: IStop[];
}