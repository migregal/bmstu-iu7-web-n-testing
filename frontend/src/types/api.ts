type Option = {
    auth?: string | null,
    waitMS?: number;
    cors?: 'cors' | 'no-cors' | 'same-origin';
    resp?: 'json';
    suffix?: string;
};

export interface ApiRequest {
    endpoint: string;
    method: 'GET' | 'PATCH' | 'POST' | 'DELETE';
    payload?: Record<string, string | Date | File>;
    query?: Record<string, string | undefined>;
    options?: Option;
}

type Resource = { resourceID?: string | number };
type Collection = Pick<ApiRequest, 'payload' | 'options'>;
type PaginatedCollectionParams = Collection & { page?: number, perPage?: number, query?: Record<string, string> };
type Member = Collection & Resource & { query?: Record<string, string> };
type ApiMember<T> = ({ resourceID, payload }: Member) => Promise<T>;

type PaginatedCollection<T> = { infos: T[], total: number }
type ApiCollection<T> = ({ payload }: Collection) => Promise<T>;
type ApiPaginatedCollection<T> = ({ page, perPage, options }: PaginatedCollectionParams) => Promise<PaginatedCollection<T>>;

export type { Collection, PaginatedCollectionParams, PaginatedCollection, Member, ApiMember, ApiCollection, ApiPaginatedCollection };
