import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';
import Paper from 'material-ui/Paper';
import TrackVisualization from './TrackVisualization';

export default React.createClass({
    mixins: [PureRenderMixin],
    render: function() {
	return <div>
	    <Card>
		<CardHeader
		    title="Song profile"
		/>
		<CardMedia>
		    <TrackVisualization />
		</CardMedia>
		<CardText>
		    <Paper size={100}>Song BPM: {Math.floor(this.props.track.Features[10])}</Paper>
		</CardText>
	    </Card>
	</div>
    }
});
