import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';
import Paper from 'material-ui/Paper';
import Avatar from 'material-ui/Avatar';

export default React.createClass({
    mixins: [PureRenderMixin],
    render: function() {
	return <div>
	    <Card>
		<CardHeader
		    title="Song title"
		    subtitle="Song artist"
/>
		<CardText>
		    <iframe src="https://embed.spotify.com/?uri=spotify:track:2TpxZ7JUBn3uw46aR7qd6V" width="80" height="80"></iframe>
		</CardText>
	    </Card>
	</div>
    }
});
