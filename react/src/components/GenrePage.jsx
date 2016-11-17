import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import NavMenu from './NavMenu';
import MaterialNav from './MaterialNav';
import Footer from './Footer';
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
		<link rel="stylesheet" href="libs/bootstrap/dist/css/bootstrap.min.css" />
		<script src="libs/jquery/dist/jquery.min.js"></script>
		<script src="libs/bootstrap/dist/js/bootstrap.min.js"></script>
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
title="Track Name"
				/>
				<CardText>
				    <Row>
					<Col md={12}>
					    <Card>
						<CardHeader
title="Song title"
subtitle="Song artist"
						/>
						<CardText>
						    <iframe src="https://embed.spotify.com/?uri=spotify:track:2TpxZ7JUBn3uw46aR7qd6V" width="300" height="80" frameborder="0" allowtransparency="true"></iframe>
						</CardText>
					    </Card>
					</Col>
				    </Row>
				    <Row>
					<Col md={6}>
					    <Card>
						<CardHeader
title="Song profile"
						/>
						<CardText>
						    <Avatar size={40}>M</Avatar>
						    <Paper size={100}>Song BPM</Paper>
						</CardText>
					    </Card>
					</Col>
					<Col md={6}>
					    <Row>
						<Col md={12}>
						    <Card>
							<CardHeader
							title="Song profile"
/>
						    <CardText>
							<Avatar size={40}>M</Avatar>
							<Paper size={100}>Song BPM</Paper>
						    </CardText>
						    </Card>
						</Col>
					    </Row>
					    <Row>
						<Col md={12}>
						    <Card>
							<CardHeader
							title="Song profile"
/>
						    <CardText>
							<Avatar size={40}>M</Avatar>
							<Paper size={100}>Song BPM</Paper>
						    </CardText>
						    </Card>
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
				    <List>
					<ListItem>
					    <Card>
						<CardHeader
						    title="Track Name"
/>
						<CardText>
						    Actions
						</CardText>
					    </Card>
					</ListItem>
					<ListItem>
					    <Card>
						<CardHeader
						    title="Track Name"
						/>
						<CardText>
						    Actions
						</CardText>
					    </Card>
					</ListItem>
					<ListItem>
					    <Card>
						<CardHeader
						    title="Track Name"
						/>
						<CardText>
						    Actions
						</CardText>
					    </Card>
					</ListItem>
					<ListItem>
					    <Card>
						<CardHeader
						    title="Track Name"
						/>
						<CardText>
						    Actions
						</CardText>
					    </Card>
					</ListItem>
				    </List>
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
