import { createContext, useContext } from 'react'

import { useLocalObservable } from 'mobx-react';

import { useApiCall } from 'hooks/apiCall';

import { useAuthStore } from './authContext';

import { createWeightsStore, Weight, WeightsStore } from 'stores/weightsStore';

const WeightsContext = createContext({} as WeightsStore)

export const WeightsProvider = ({ children }: { children: JSX.Element }) => {
    const auth = useAuthStore();

    const api = useApiCall<Weight>({
        endpoint: `${process.env.REACT_APP_API_URL}/weights`,
    })

    const modelsStore = useLocalObservable(() => createWeightsStore(auth, api))

    return (
        <WeightsContext.Provider value={modelsStore}>{children}</WeightsContext.Provider>
    )
}

export const useWeightsStore = () => useContext(WeightsContext)
