import React from 'react';
// import PropTypes from 'prop-types';

import logo from '../../../images/logo.svg';

const Loading = () => (
  <div className="App-loading">
    <img src={logo} className="App-logo" alt="logo" />
    <p>Checking URLs...</p>
  </div>
);

// Loading.propTypes = {
// };

export default Loading;
