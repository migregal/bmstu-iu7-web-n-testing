import { createContext, useContext } from 'react'

import { useLocalObservable } from 'mobx-react';

import { useAuthStore } from './authContext';
import { useApiCall } from 'hooks/apiCall';
import { Block, BlocksStore, createBlocksStore } from 'stores/blocksStore';

const BlocksContext = createContext({} as BlocksStore)

export const BlocksProvider = ({ children }: { children: JSX.Element }) => {
  const auth = useAuthStore();

  const api = useApiCall<Block>({
    endpoint: `${process.env.REACT_APP_API_URL}/blocks/users/`
  });
  const blocksStore = useLocalObservable(() => createBlocksStore(auth, api))

  return (
    <BlocksContext.Provider value={blocksStore}>{children}</BlocksContext.Provider>
  )
}

export const useBlocksStore = () => useContext(BlocksContext)
