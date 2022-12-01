import { Button, Container, Heading, Hero, Navbar } from "react-bulma-components";
import { Link } from "react-router-dom";

import { ReactComponent as Logo404 } from './not_found.svg';

function NotFound() {
  return (
    <Hero className='page'>
      <Hero.Header p={4}>
        <Navbar>
          <Navbar.Container align="right">
            <Button rounded color="primary" renderAs={Link} to='/'>
              Let's go home
            </Button>
          </Navbar.Container>
        </Navbar>

        <Heading p={6} style={{ margin: "auto", textAlign: "center" }} >
          Seems like you are lost...
        </Heading>

      </Hero.Header>

      <Hero.Body style={{ maxWidth: "90%", margin: 'auto' }}>
        <Logo404 style={{ maxWidth: "100%" }}  />
      </Hero.Body>

      <Hero.Footer style={{ marginBottom: "4em", textAlign: "center" }}>
        <Container style={{ margin: "auto", textAlign: "center" }}>
          Made by migregal
        </Container>
      </Hero.Footer>
    </Hero >
  )
}

export default NotFound;
