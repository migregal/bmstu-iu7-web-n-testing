import { DEFAULT_CORS } from "contstants/api";

import { ApiCollection, ApiMember, ApiPaginatedCollection, ApiRequest, Collection, Member, PaginatedCollection, PaginatedCollectionParams } from "types/api";

export interface ApiCallMethods<T> {
    all: ApiPaginatedCollection<T>;
    create: ApiCollection<T>;
    one: ApiMember<T>;
    update: ApiMember<T>;
    destroy: ApiMember<void>;
    stat: ApiMember<T>;
}

async function apiCall<T>({
    endpoint, options, method, query, payload,
}: ApiRequest): Promise<T> {
    const headers: { Authorization: string, "Content-Type"?: string } = {
        'Authorization': "",
    }

    if (options?.auth) {
        headers['Authorization'] = "Token " + options.auth
    }

    let body = null

    if (payload) {
        if (Object.entries(payload).some((v) => v[1] instanceof File)) {
            const formData = new FormData();
            Object.entries(payload).forEach((v) => {
                formData.append(v[0], v[1] instanceof File ? v[1] : v[1].toString());
            })
            body = formData
        } else {
            body = JSON.stringify({ ...payload });
            headers["Content-Type"] = "application/json"
        }
    }

    const mode = options?.cors ? options.cors : DEFAULT_CORS;

    if (options?.suffix) {
        endpoint += '/' + options.suffix
    }

    const q = Object
        .entries(query ?? {})
        .filter((r) => r[1] !== undefined)
        .map((r) => r as string[])

    const response = await fetch(
        `${endpoint}?${new URLSearchParams(q)}`,
        { method: method, headers: headers, body: body, mode: mode });

    if (!response.ok) {
        throw new Error(`Error: ${response.status}`);
    }

    return options?.resp == 'json' ? response.json() : null as T;
}

export function useApiCall<T>({
    endpoint,
    options: defaultOption,
}: Pick<ApiRequest, 'endpoint' | 'options'>): ApiCallMethods<T> {
    const apiCallWrapper = <T>({
        endpoint,
        method,
        payload,
        query,
        options,
    }: ApiRequest): Promise<T> => {
        options = { ...defaultOption, ...options }
        return apiCall<T>({
            endpoint,
            method,
            payload,
            query,
            options: options,
        })
    };

    const all = ({ page, perPage, query, options }: PaginatedCollectionParams) => {
        return apiCallWrapper<PaginatedCollection<T>>({
            endpoint: endpoint,
            method: 'GET',
            query: { ...query, page: page?.toString(), 'per_page': perPage?.toString() },
            options,
        });
    };

    const create = ({ payload, options }: Collection): Promise<T> => {
        return apiCallWrapper<T>({ endpoint, method: 'POST', payload, options: options });
    };

    const one = ({ resourceID, options }: Member): Promise<T> => {
        return apiCallWrapper<T>({
            endpoint: `${endpoint}/${resourceID}`,
            method: 'GET',
            options: options,
        });
    };

    const update = ({ resourceID, options, payload }: Member): Promise<T> => {
        return apiCallWrapper<T>({
            endpoint: `${endpoint}/${resourceID}`,
            method: 'PATCH',
            options: options,
            payload: payload,
        });
    };

    const destroy = ({ resourceID, options }: Member): Promise<void> => {
        return apiCallWrapper<void>({
            endpoint: `${endpoint}/${resourceID}`,
            method: 'DELETE',
            options: options,
        });
    };

    const stat = ({ query, options }: Member): Promise<T> => {
        return apiCallWrapper<T>({
            endpoint: `${endpoint}`,
            method: 'GET',
            query: {...query},
            options: options,
        });
    };

    return {
        all,
        create,
        one,
        update,
        destroy,
        stat,
    };
}
