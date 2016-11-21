import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import QueueItem from './QueueItem';
import {List, ListItem} from 'material-ui/List';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';

export default React.createClass({
    mixins: [PureRenderMixin],
    render: function() {
	return <div>
	{this.props.queue.TrackQueue.map(track => <QueueItem key={track.ID} removeFromQueue={this.props.removeFromQueue} setTrack={this.props.setTrack} track={track}/>)}
	</div>
    }
});
