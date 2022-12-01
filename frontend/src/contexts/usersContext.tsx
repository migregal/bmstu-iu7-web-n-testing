import { createContext, useContext } from 'react'

import { useLocalObservable } from 'mobx-react';

import { useAuthStore } from './authContext';
import { useApiCall } from 'hooks/apiCall';

import { createUsersStore, User, UsersStore } from 'stores/usersStore';
import { useBlocksStore } from './blocksContext';

const UsersContext = createContext({} as UsersStore)

export const UsersProvider = ({ children }: { children: JSX.Element }) => {
  const auth = useAuthStore();

  const blocks = useBlocksStore();

  const api = useApiCall<User>({
    endpoint: `${process.env.REACT_APP_API_URL}/users`
  });
  const userStore = useLocalObservable(() => createUsersStore(auth, blocks, api))

  return (
    <UsersContext.Provider value={userStore}>{children}</UsersContext.Provider>
  )
}

export const useUsersStore = () => useContext(UsersContext)
