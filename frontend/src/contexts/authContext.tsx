import { createContext, useContext } from 'react'

import { useLocalObservable } from 'mobx-react';

import { useApiCall } from 'hooks/apiCall';

import { createAuthStore, AuthStore, Auth } from "stores/authStore";

const AuthContext = createContext({} as AuthStore)

export const AuthProvider = ({ children }: { children: JSX.Element }) => {
    const api = useApiCall<Auth>({
        endpoint: `${process.env.REACT_APP_API_URL}`,
    })

    const authStore = useLocalObservable(() => createAuthStore(api))

    return (
        <AuthContext.Provider value={authStore}>{children}</AuthContext.Provider>
    )
}

export const useAuthStore = () => useContext(AuthContext)
