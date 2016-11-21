import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';
import DropDownMenu from 'material-ui/DropDownMenu';
import MenuItem from 'material-ui/MenuItem';
import Paper from 'material-ui/Paper';
import Avatar from 'material-ui/Avatar';

export default React.createClass({
    mixins: [PureRenderMixin],
    render: function() {
	return <div>
	    <Card>
		<CardHeader title="Track Genre" />
		<CardText>
		    <DropDownMenu value={this.props.track.Genres[0]} >
			{this.props.track.Genres.map(genre => <MenuItem key={genre} value={genre} primaryText={genre}/>)}
		    </DropDownMenu>
		</CardText>
	    </Card>
	</div>
    }
});
