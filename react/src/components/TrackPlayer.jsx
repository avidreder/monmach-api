import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import TrackActions from './TrackActions';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';
import Paper from 'material-ui/Paper';
import Avatar from 'material-ui/Avatar';

export default React.createClass({
    mixins: [PureRenderMixin],
    onChange: function() {
	console.log(this.refs.externalPlayer);
	this.refs.externalPlayer.click();
    },
    render: function() {
	return <div>
	    <Card>
		<CardTitle title={this.props.track.Name} subtitle={this.props.track.Artists} />
		<CardText>
		    <Paper>
			<iframe ref="externalPlayer" id="externalPlayer" src={"https://embed.spotify.com/?uri=spotify:track:" + this.props.track.SpotifyID} width="100%" height="80" frameBorder="0" allowTransparency="true"></iframe>
		    </Paper>
		    <TrackActions />
		</CardText>
	    </Card>
	</div>
    }
});
