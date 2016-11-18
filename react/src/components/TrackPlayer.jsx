import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import TrackActions from './TrackActions';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';
import Paper from 'material-ui/Paper';
import Avatar from 'material-ui/Avatar';

export default React.createClass({
    mixins: [PureRenderMixin],
    render: function() {
	return <div>
	    <Card>
		<CardTitle title="Song title" subtitle="Song artist" />
		<CardText>
		    <Paper>
			<iframe src="https://embed.spotify.com/?uri=spotify:track:2TpxZ7JUBn3uw46aR7qd6V" width="100%" height="80" frameBorder="0" allowTransparency="true"></iframe>
		    </Paper>
		    <TrackActions />
		</CardText>
	    </Card>
	</div>
    }
});
