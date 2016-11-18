import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import NavMenu from './NavMenu';
import Queue from './Queue';
import MaterialNav from './MaterialNav';
import Footer from './Footer';
import TrackProfile from './TrackProfile';
import TrackGenres from './TrackGenres';
import TrackPlaylists from './TrackPlaylists';
import TrackPlayer from './TrackPlayer';
import TrackVisualization from './TrackVisualization';
import classnames from 'classnames';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import {deepOrange500} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';
import {Grid, Row, Col} from 'react-flexbox-grid/lib';
import {List, ListItem} from 'material-ui/List';
import Avatar from 'material-ui/Avatar';
import Paper from 'material-ui/Paper';

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
		<NavMenu />
		<MaterialNav />
		<Grid fluid>
		    <Row>
			<Col md={12}>
			    <h1>Genre Title</h1>
			</Col>
		    </Row>
		    <Row>
			<Col md={6}>
			    <Card>
				<CardHeader
title="Currently Playing"
				/>
				<CardText>
				    <Row>
					<Col md={12}>
					    <TrackPlayer />
					</Col>
				    </Row>
				    <Row>
					<Col md={6}>
					    <TrackProfile />
					</Col>
					<Col md={6}>
					    <Row>
						<Col md={12}>
						    <TrackGenres />
						</Col>
					    </Row>
					    <Row>
						<Col md={12}>
						    <TrackPlaylists />
						</Col>
					    </Row>
					</Col>
				    </Row>
				</CardText>
			    </Card>
			</Col>
			<Col md={6}>
			    <Card>
				<CardHeader
				    title="Queue"
/>
				<CardText>
				    <Queue />
				</CardText>
			    </Card>
			</Col>
		    </Row>
		</Grid>
		<Footer />
	    </div>
	</MuiThemeProvider>;
    }
});
