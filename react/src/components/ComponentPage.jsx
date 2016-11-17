import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import NavMenu from './NavMenu';
import MaterialNav from './MaterialNav';
import Footer from './Footer';
import classnames from 'classnames';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import {deepOrange500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import Paper from 'material-ui/Paper';
import {Grid, Row, Col} from 'react-flexbox-grid/lib';

const muiTheme = getMuiTheme({
    palette: {
	accent1Color: deepOrange500,
    },
});

export default React.createClass({
    mixins: [PureRenderMixin],
    render: function() {
	return <MuiThemeProvider muiTheme={muiTheme}>
	    <div className="pageWrapper">
		<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" />
		<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet" />
		<link rel="stylesheet" href="libs/bootstrap/dist/css/bootstrap.min.css" />
		<script src="libs/jquery/dist/jquery.min.js"></script>
		<script src="libs/bootstrap/dist/js/bootstrap.min.js"></script>
		<NavMenu />
		<Grid fluid>
		    <Row>
			<Col md={12}>
			    <h1>Components</h1>
			</Col>
		    </Row>
		    <Row>
			<Col md={12}>
			    <h2>Main Menu</h2>
			</Col>
		    </Row>
		    <Row>
			<Col md={12}>
			    <MaterialNav />
			</Col>
		    </Row>
		    <Row>
			<Col md={12}>
			    <Footer />
			</Col>
		    </Row>
		</Grid>
	    </div>
	</MuiThemeProvider>;
    }
});
