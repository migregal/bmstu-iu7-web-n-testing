import { createContext, useContext } from 'react'

import { useLocalObservable } from 'mobx-react';

import { useAuthStore } from './authContext';
import { useApiCall } from 'hooks/apiCall';

import { createStatsStore, Stat, StatsStore } from 'stores/statsStore';


const StatsContext = createContext({} as StatsStore)

export const StatsProvider = ({ children }: { children: JSX.Element }) => {
  const auth = useAuthStore();

  const api = useApiCall<Stat>({
    endpoint: `${process.env.REACT_APP_API_URL}`
  });
  const statsStore = useLocalObservable(() => createStatsStore(auth, api))

  return (
    <StatsContext.Provider value={statsStore}>{children}</StatsContext.Provider>
  )
}

export const useStatsStore = () => useContext(StatsContext)
