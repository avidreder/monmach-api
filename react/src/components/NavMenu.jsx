import React from 'react';
import { Link } from 'react-router';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import AppBar from 'material-ui/AppBar';

export default React.createClass({
  mixins: [PureRenderMixin],
  render: function() {
      return <div className="navbar navbar-default" id="home-navbar">
      <div className="container">
        <div className="navbar-header">
        <img style={{cursor: 'pointer'}} alt="Brand" src="../img/GridLogo.png" height="50" />
          <button type="button" className="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
            <span className="icon-bar"></span>
          </button>
          <Link style={{cursor: 'pointer'}} to="/" className="navbar-brand">Avidreder.net</Link>
        </div>
        <div className="collapse navbar-collapse">
          <ul className="nav navbar-nav">
            <li style={{cursor: 'pointer'}} className="navbar-link"><Link to="/">Home</Link></li>
	    <li style={{cursor: 'pointer'}} className="navbar-link"><Link to="/components">Components</Link></li>
          </ul>
        </div>
	  </div>
	  </div>;
  }
});
