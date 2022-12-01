import { Navigate, useLocation } from 'react-router';

import { Observer } from 'mobx-react-lite';

import 'bulma/css/bulma.min.css';
import { Box, Container } from 'react-bulma-components';

import { ToastContainer } from 'react-toastify';

import { useAuthStore } from 'contexts/authContext';
import LoginForm from 'components/Forms/LoginForm';
import { Link } from 'react-router-dom';

function Login() {
  const location = useLocation();

  const auth = useAuthStore();

  return (
    <Observer>
      {() => {
        return auth.authorized() ?
          (<Navigate to={location.state?.from ?? "/"} state={{ from: location }} />) : (
            <Container
              display="flex"
              flexDirection="column"
              style={{
                justifyContent: 'center',
                alignContent: 'center',
                height: '100vh',
                maxWidth: "25vw",
              }}
            >
              <ToastContainer />

              <Box style={{ width: '25vw', margin: 'auto' }}>
                <LoginForm />
                <div style={{
                  marginTop: '1vh',
                  textAlign: 'center'
                }}>
                  Don't have account? <Link to='/signup'>Sign up</Link>
                </div>
              </Box>
            </Container >
          )
      }}
    </Observer >
  )
}

export default Login;
