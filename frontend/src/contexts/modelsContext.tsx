import { createContext, useContext } from 'react'

import { useLocalObservable } from 'mobx-react';

import { useAuthStore } from './authContext';

import { createModelsStore, Model, ModelsStore } from 'stores/modelsStore';
import { useApiCall } from 'hooks/apiCall';
import { useUsersStore } from './usersContext';

const ModelsContext = createContext({} as ModelsStore)

export const ModelsProvider = ({ children }: { children: JSX.Element }) => {
    const auth = useAuthStore();

    const users = useUsersStore();

    const api = useApiCall<Model>({
        endpoint: `${process.env.REACT_APP_API_URL}/models`,
    })

    const modelsStore = useLocalObservable(() => createModelsStore(auth, users, api))

    return (
        <ModelsContext.Provider value={modelsStore}>{children}</ModelsContext.Provider>
    )
}

export const useModelsStore = () => useContext(ModelsContext)
