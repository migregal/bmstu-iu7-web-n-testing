import { Box, Container, Heading, Level, Tile, Notification } from "react-bulma-components";

import StatForm from "components/Forms/StatForm";

import Page from "components/Pages/Page";

function Stats() {
  return (
    <Page>
      <Container className="contentContainer">
        <Box>
          <Level>
            <Level.Side>
              <Heading>Stats</Heading>
            </Level.Side>
          </Level>

          <Tile kind="ancestor">
            <Tile kind="parent">
              <Tile kind="child" renderAs={Notification}>
                <div className="content">
                  <Heading>Users stats</Heading>
                  <StatForm object={'users'} />
                </div>
              </Tile>
            </Tile>
            <Tile kind="parent">
              <Tile kind="child" renderAs={Notification}>
                <div className="content">
                  <Heading>Models</Heading>
                  <StatForm object={'models'} />
                </div>
              </Tile>
            </Tile>
            <Tile kind="parent">
              <Tile kind="child" renderAs={Notification}>
                <div className="content">
                  <Heading>Weights</Heading>
                  <StatForm object={'weights'} />
                </div>
              </Tile>
            </Tile>
          </Tile>
        </Box >
      </Container >
    </Page>
  );
}

export default Stats;
