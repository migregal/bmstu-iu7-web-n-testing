import 'bulma/css/bulma.min.css';
import { Button, Footer, Navbar } from 'react-bulma-components';

import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from 'react-toastify';

import { useAuthStore } from 'contexts/authContext';

import { ReactComponent as Logo } from './logo.svg';
import { Observer } from 'mobx-react';
import { Link, useLocation } from 'react-router-dom';

function Page({ children }: { children: JSX.Element }) {
  const location = useLocation();

  const auth = useAuthStore();

  return (
    <div className='page'>
      <Observer>
        {() => (
          <Navbar p={4}>
            <Navbar.Brand >
              <Logo style={{ width: '4em', height: '4em', marginRight: '1em' }} />
            </Navbar.Brand>
            {auth.authorized() && <Navbar.Item href="/">Home</Navbar.Item>}
            {(auth.isregular() || auth.isadmin()) && <Navbar.Item href="/users">Users</Navbar.Item>}
            {(auth.isregular() || auth.isadmin()) && <Navbar.Item href="/models">Models</Navbar.Item>}
            {auth.isstat() && <Navbar.Item href="/stats">Stats</Navbar.Item>}

            <Navbar.Container align="right" style={{ marginTop: '1vh' }}>
              {auth.authorized() && <Button rounded onClick={() => { auth.unautorize() }}>Sign Out</Button>}
              {!auth.authorized() &&
                <Link to='/signin' state={{ from: location.state?.from }}>
                  <Button rounded color="primary">Sign in</Button>
                </Link>
              }
            </Navbar.Container>
          </Navbar>
        )}
      </Observer>

      <ToastContainer />

      {children}

      <Footer style={{ textAlign: "center" }}>
        Made by migregal
      </Footer>
    </div>
  );
}

export default Page;
