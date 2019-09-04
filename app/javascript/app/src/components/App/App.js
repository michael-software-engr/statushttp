import React from 'react';

import {
  Container, Grid,
  Divider
} from 'semantic-ui-react';
import fetch from 'cross-fetch';
import humps from 'humps';

import Header from './Header';
import URLInput from './URLInput';
import Results from './results/Results';
import './App.css';

export default class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isFetchting: false,
      urlInput: [
        // 'http://no-such-host-example',
        // 'google.com:81', // Should time out
        // 'nmap.org',
        // 'yahoo.com:81', // Should time out
        // 'https://golang.org',
        // 'yahoo.com',
        // 'yahoo.com', // Duplicate
        // 'duckduckgo.com',
        // 'yahoo.com',
        // 'yahoo.com', // Duplicate
        // 'yahoo.com', // Duplicate
        // 'nmap.org', // Duplicate
        // 'http://bing.com',
        // 'amazon.com',
        // 'micron.com', // 10th unique
        // 'github.com', // Should be exclude from here or to end of list
        // 'bitbucket.org',
        // 'reactjs.org',
        // 'ruby-lang.org'
      ].join('\n'),
      response: {}
    };
  }

  urlInputOnChangeHandler = (event) => {
    this.setState({ urlInput: event.target.value });
  }

  urlInputOnSubmitHandler = (event) => {
    event.preventDefault();

    this.setState({ isFetchting: true });

    fetch('/api', {
      method: 'POST', // *GET, POST, PUT, DELETE, etc.
      mode: 'cors', // no-cors, cors, *same-origin
      cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
      credentials: 'same-origin', // include, *same-origin, omit
      headers: {
        'Content-Type': 'application/json',
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
      redirect: 'follow', // manual, *follow, error
      referrer: 'no-referrer', // no-referrer, *client
      body: JSON.stringify(this.state.urlInput), // body data type must match "Content-Type" header
    }).then((response) => response.json()).then((json) => {
      const response = humps.camelizeKeys(json);

      this.setState({ response, isFetchting: false });
    });
  }

  urlInputOnErrorDismiss = () => {
    this.setState((state) => ({
      response: { ...state.response, error: null }
    }));
  }

  urlInputOnWarningDismiss = () => {
    this.setState((state) => ({
      response: { ...state.response, warningMessages: [] }
    }));
  }

  urlInputOnOtherDismiss = () => {
    this.setState((state) => ({
      response: { ...state.response, statusMessages: [] }
    }));
  }

  render() {
    const { urlInput, response, isFetchting } = this.state;
    return (
      <Container className="App">
        <Divider hidden />
        <Header />
        <Divider hidden />

        <Grid columns={2}>
          <Grid.Column width={6}>
            <URLInput
              onChange={this.urlInputOnChangeHandler}
              onSumbit={this.urlInputOnSubmitHandler}
              value={urlInput}
              error={response.error}
              warningMessages={response.warningMessages}
              statusMessages={response.statusMessages}
              onErrorDismiss={this.urlInputOnErrorDismiss}
              onWarningDismiss={this.urlInputOnWarningDismiss}
              onOtherDismiss={this.urlInputOnOtherDismiss}
            />
          </Grid.Column>
          <Grid.Column width={10}>
            <Results results={response.results} isFetchting={isFetchting} />
          </Grid.Column>
        </Grid>
      </Container>
    );
  }
}
