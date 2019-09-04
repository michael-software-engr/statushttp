import React from 'react';
import PropTypes from 'prop-types';

import {
  Container
} from 'semantic-ui-react';

import Loading from './Loading';
import RTable from './RTable';

import '../App.css';

const Results = ({ results, isFetchting }) => {
  const selector = () => {
    switch (true) {
      case isFetchting: {
        return <Loading />;
      }

      case results.length > 0: {
        return <RTable results={results} />;
      }

      default: {
        return null;
      }
    }
  };

  const Content = selector();

  return (
    <Container textAlign="left">
      {Content && Content}
    </Container>
  );
};

Results.propTypes = {
  results: PropTypes.arrayOf(PropTypes.shape()),
  isFetchting: PropTypes.bool.isRequired
};

Results.defaultProps = {
  results: []
};

export default Results;
