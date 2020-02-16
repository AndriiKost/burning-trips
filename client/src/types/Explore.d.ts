export interface ISearchQuery {
    longtitude: number;
    latitude: number;
}

export interface ISearchResult<T> {
    result: Array<T>;
}