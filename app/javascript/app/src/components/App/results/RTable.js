import React from 'react';
import PropTypes from 'prop-types';

import {
  Table,
  Icon,
} from 'semantic-ui-react';

const RTable = ({ results }) => (
  <Table celled>
    <Table.Header>
      <Table.Row>
        <Table.HeaderCell />
        <Table.HeaderCell>URL</Table.HeaderCell>
        <Table.HeaderCell>Status</Table.HeaderCell>
        <Table.HeaderCell>Messages</Table.HeaderCell>
      </Table.Row>
    </Table.Header>

    <Table.Body>
      {
        results.sort((a, b) => a.ix - b.ix).map(({
          ix, name, uRL, passed, statusMessage
        }) => (
          <Table.Row key={uRL}>
            <Table.Cell textAlign="right">{ix + 1}</Table.Cell>
            <Table.Cell>
              <a href={uRL} target="_blank" rel="noopener noreferrer">
                {name}
              </a>
            </Table.Cell>

            <Table.Cell
              positive={passed}
              negative={!passed}
            >
              {passed ? (
                <Icon name="checkmark" />
              ) : (
                <Icon name="x" />
              )}
              &nbsp;
              {passed ? 'Passed' : 'Failed'}
            </Table.Cell>
            <Table.Cell>{statusMessage}</Table.Cell>
          </Table.Row>
        ))
      }
    </Table.Body>
  </Table>
);

RTable.propTypes = {
  results: PropTypes.arrayOf(PropTypes.shape()),
};

RTable.defaultProps = {
  results: []
};

export default RTable;
