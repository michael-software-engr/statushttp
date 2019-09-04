import React from 'react';
import PropTypes from 'prop-types';

import {
  Container, Form, TextArea, Button,
  Message,
  Segment,
  Divider,
  List
} from 'semantic-ui-react';

const URLInput = ({
  onChange, onSumbit, value,
  error,
  warningMessages,
  statusMessages,
  onErrorDismiss,
  onWarningDismiss,
  onOtherDismiss
}) => {
  const warning = warningMessages.length > 0;

  const isWarning = error ? false : warning;

  return (
    <Container textAlign="left">
      <Form error={!!error} warning={isWarning}>
        <Segment inverted>
          Enter URLs
          <List bulleted>
            <List.Item>separated by spaces or new lines</List.Item>
            <List.Item>maximum of 10</List.Item>
            <List.Item>must be unique</List.Item>
            <List.Item>if URL has no scheme, &quot;https&quot; will be used</List.Item>
          </List>
        </Segment>
        <TextArea
          onChange={onChange}
          value={value}
          rows={10}
          placeholder={[
            'example.org',
            'httpstat.us/404'
          ].join('\n')}
        />

        <Message
          error
          header="Errors"
          list={[error]}
          onDismiss={onErrorDismiss}
        />

        <Message
          warning
          header="Warnings"
          list={warningMessages}
          onDismiss={onWarningDismiss}
        />

        {statusMessages && statusMessages.length > 0 && (
          <Message
            header="Other"
            list={statusMessages}
            onDismiss={onOtherDismiss}
          />
        )}

        <Divider hidden />

        <Button type="submit" onClick={onSumbit} disabled={!value.trim()}>Submit</Button>
      </Form>
    </Container>
  );
};

URLInput.propTypes = {
  onChange: PropTypes.func.isRequired,
  onSumbit: PropTypes.func.isRequired,
  onErrorDismiss: PropTypes.func.isRequired,
  onWarningDismiss: PropTypes.func.isRequired,
  onOtherDismiss: PropTypes.func.isRequired,
  error: PropTypes.string,
  value: PropTypes.string,
  warningMessages: PropTypes.arrayOf(PropTypes.string),
  statusMessages: PropTypes.arrayOf(PropTypes.string)
};

URLInput.defaultProps = {
  error: '',
  value: '',
  warningMessages: [],
  statusMessages: []
};

export default URLInput;
