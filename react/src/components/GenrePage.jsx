import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import NavMenu from './NavMenu';
import {QueueContainer} from './Queue';
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
import {connect} from 'react-redux';
import * as actionCreators from '../action_creators';

const muiTheme = getMuiTheme({
    palette: {
	accent1Color: deepOrange500,
    },
});

export const GenrePage = React.createClass({
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
			    <h1>{this.props.genre.get("Name")}</h1>
			</Col>
		    </Row>
		    <Row>
			<Col md={6}>
			    <Card>
				<CardTitle title="Currently Playing" />
				<CardText>
				    <Row>
					<Col md={12}>
					    <TrackPlayer track={this.props.currentTrack.toJS()} addGenre={this.props.dispatchAddGenre}/>
					</Col>
				    </Row>
				    <Row>
					<Col md={6}>
					    <TrackProfile track={this.props.currentTrack.toJS()} />
					</Col>
					<Col md={6}>
					    <Row>
						<Col md={12}>
						    <TrackGenres track={this.props.currentTrack.toJS()} />
						</Col>
					    </Row>
					    <Row>
						<Col md={12}>
						    <TrackPlaylists track={this.props.currentTrack.toJS()} />
						</Col>
					    </Row>
					</Col>
				    </Row>
				</CardText>
			    </Card>
			</Col>
			<Col md={6}>
			    <Card>
				<CardTitle title="Queue" />
				<CardText>
				    <QueueContainer addGenre={this.props.dispatchAddGenre} removeFromQueue={this.props.dispatchRemoveFromQueue} setTrack={this.props.dispatchSetTrack} />
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

function mapStateToProps(state) {
    return {
	currentTrack: state.get('currentTrack'),
	queue: state.get('queue'),
	genre: state.get('genre'),
	playlist: state.get('playlist')
    };
}

function mapDispatchToProps(dispatch) {
    return {
 	dispatchSetTrack: function(track) {
 	    dispatch(actionCreators.setTrack(track));
 	},
	dispatchRemoveFromQueue: function(track) {
 	    dispatch(actionCreators.removeFromQueue(track));
 	},
	dispatchAddGenre: function(track) {
 	    dispatch(actionCreators.addGenre(track));
 	}
    }
}

export const GenrePageContainer = connect(mapStateToProps, mapDispatchToProps)(GenrePage);
