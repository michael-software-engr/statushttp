import React from 'react';

import {
  Container, Card, List
} from 'semantic-ui-react';

const Header = () => (
  <Container textAlign="left">
    <Card raised fluid>
      <Card.Content>
        <Card.Header>HTTP Checker</Card.Header>
        <Card.Meta>Check if a server is up using HTTP</Card.Meta>
        <Card.Description>
          <List>
            <List.Item>
              Back end (Go):
              {' '}
              <a
                href="https://github.com/michael-software-engr/statushttp"
                target="_blank"
                rel="noopener noreferrer"
              >
                github.com/michael-software-engr/statushttp
              </a>
            </List.Item>

            <List.Item>
              Front end (React JS):
              {' '}
              <a
                href="https://github.com/michael-software-engr/statushttp/tree/master/app/javascript/app"
                target="_blank"
                rel="noopener noreferrer"
              >
                github.com/michael-software-engr/statushttp/tree/master/app/javascript/app
              </a>
            </List.Item>
          </List>
        </Card.Description>
      </Card.Content>
    </Card>
  </Container>
);

export default Header;
