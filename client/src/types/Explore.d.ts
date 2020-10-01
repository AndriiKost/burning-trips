export interface ISearchQuery {
    longitude: number;
    latitude: number;
    keyword: string;
}

export interface ISearchResult<T> {
    result: Array<T>;
}