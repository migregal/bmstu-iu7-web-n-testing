import Page from 'components/Pages/Page';
import { Container, Hero } from 'react-bulma-components';

import { ReactComponent as HomeLogo } from './programming.svg';

function Home() {
  return (
    <Page>
      <Container className="contentContainer">
        <Hero style={{ maxWidth: "90%", margin: 'auto' }}>
          <Hero.Body>
            <HomeLogo style={{ maxWidth: "100%" }} />
          </Hero.Body>
        </Hero >
      </Container>
    </Page>
  );
}

export default Home;
