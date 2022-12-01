import { Navigate, useLocation } from 'react-router';

import 'bulma/css/bulma.min.css';
import { Container, Box } from 'react-bulma-components';

import { ToastContainer } from 'react-toastify';

import { useAuthStore } from 'contexts/authContext';
import { Observer } from 'mobx-react-lite';

import RegistrationForm from 'components/Forms/RegistrationForm';
import { Link } from 'react-router-dom';

function Registration() {
  const location = useLocation();

  const auth = useAuthStore();

  return (
    <Observer>
      {() => {
        return auth.authorized() ?
          (<Navigate to={location.state?.path ?? "/"} state={{ from: location }} />) : (
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
                <RegistrationForm />
                <div style={{
                  marginTop: '1vh',
                  textAlign: 'center'
                }}>
                  Alreay have account? <Link to='/signin'>Sign in</Link>
                </div>
              </Box>
            </Container >
          )
      }}
    </Observer >
  )
}

export default Registration;
