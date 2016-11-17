import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import NavMenu from './NavMenu';
import Footer from './Footer';
import classnames from 'classnames';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import {deepOrange500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import Paper from 'material-ui/Paper';

const muiTheme = getMuiTheme({
    palette: {
	accent1Color: deepOrange500,
    },
});

const style = {
    height: 100,
    width: 100,
    margin: 20,
    textAlign: 'center',
    display: 'inline-block',
};

export default React.createClass({
    mixins: [PureRenderMixin],
    render: function() {
	return <MuiThemeProvider muiTheme={muiTheme}>
	    <div className="pageWrapper">
		<link rel="stylesheet" href="libs/bootstrap/dist/css/bootstrap.min.css" />
		<script src="libs/jquery/dist/jquery.min.js"></script>
		<script src="libs/bootstrap/dist/js/bootstrap.min.js"></script>
		<NavMenu />
		<h1>Home Page</h1>
		<Footer />
	    </div>
	</MuiThemeProvider>;
    }
});

