import React from 'react';
import classnames from 'classnames';

export default React.createClass({
  render: function() {
    return <div className={classnames('navbar', 'navbar-bottom', 'row')} id="home-navbar">
      <div className="container">
        <div className="navbar-left"><h3>&#169; 2016 Andrew Reder</h3></div>
        <div className="navbar-right">
          <a target="_blank" href="https://github.com/avidreder"><img src="img/github.png" width="40" /></a>
          <a target="_blank" href="https://www.linkedin.com/pub/andrew-reder/62/59b/b81/en"><img src="img/linkedin.png" width="40" /></a>
        </div>
      </div>
    </div>;
  }
});
