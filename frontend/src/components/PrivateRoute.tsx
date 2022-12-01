import { Navigate, Outlet, useLocation } from 'react-router-dom';

import { useAuthStore } from 'contexts/authContext';
import { Observer } from 'mobx-react';

function PrivateRoute() {
  const location = useLocation();

  const auth = useAuthStore();

  return (
    <Observer>
      {() => {
        return auth.authorized() ?
          (<Outlet />) :
          (<Navigate to="/" state={{ from: location.pathname }} />)
      }
      }
    </Observer>
  )

}

export default PrivateRoute;
