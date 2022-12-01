import { Route, Routes } from 'react-router-dom'

import PrivateRoute from "components/PrivateRoute";

import Home from "pages/Home/Home";
import Login from 'pages/Login/Login';
import NotFound from 'pages/NotFound/NotFound';
import Registration from 'pages/Registration/Registration';
import Users from 'pages/Users/Users';
import Models from 'pages/Models/Models';
import Stats from 'pages/Stats/Stats';

import { UsersProvider } from 'contexts/usersContext';
import { ModelsProvider } from 'contexts/modelsContext';
import { WeightsProvider } from 'contexts/weightsContext';
import { BlocksProvider } from 'contexts/blocksContext';
import { StatsProvider } from 'contexts/statsContext';

const routes = [
  { path: '/signin', element: <Login /> },
  { path: '/signup', element: <Registration /> },
  { path: '/users', element: <BlocksProvider><UsersProvider><Users /></UsersProvider></BlocksProvider>, private: true },
  {
    path: '/models', element:
      <UsersProvider>
        <ModelsProvider>
          <WeightsProvider>
            <Models />
          </WeightsProvider>
        </ModelsProvider>
      </UsersProvider>, private: true
  },
  { path: '/stats', element: <StatsProvider><Stats /></StatsProvider>, private: true },
  { path: '/', element: <Home /> },
]

function Router() {
  return (
    <Routes>
      {routes.map((route) => (
        route.private ?
          <Route key={route.path} path={route.path} element={<PrivateRoute />}>
            <Route path={route.path} element={route.element} />
          </Route> :
          <Route key={route.path} path={route.path} element={route.element} />
      ))}
      <Route path="*" element={<NotFound />} />
    </Routes>
  );
}

export default Router;
