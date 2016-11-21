import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import Paper from 'material-ui/Paper';
import Avatar from 'material-ui/Avatar';
import {Toolbar, ToolbarGroup, ToolbarSeparator, ToolbarTitle} from 'material-ui/Toolbar';
import FloatingActionButton from 'material-ui/FloatingActionButton';
import FontIcon from 'material-ui/FontIcon';

export default React.createClass({
    mixins: [PureRenderMixin],
    render: function() {
	return <div>
	    <Paper>
		<Toolbar>
		    <ToolbarGroup>
			<FontIcon className="material-icons">star_border</FontIcon>
			<FontIcon className="material-icons">star_border</FontIcon>
			<FontIcon className="material-icons">star_border</FontIcon>
			<FontIcon className="material-icons">star_border</FontIcon>
			<FontIcon className="material-icons">star_border</FontIcon>
		    </ToolbarGroup>
		    <ToolbarGroup>
			<FloatingActionButton mini={true}>
			    <FontIcon className="material-icons" onClick={() => this.props.addGenre(this.props.track)}>playlist_add</FontIcon>
			</FloatingActionButton>
			<FloatingActionButton mini={true}>
			    <FontIcon className="material-icons">not_interested</FontIcon>
			</FloatingActionButton>
			<FloatingActionButton mini={true}>
			    <FontIcon className="material-icons">menu</FontIcon>
			</FloatingActionButton>
		    </ToolbarGroup>
		</Toolbar>
	    </Paper>
	</div>
    }
});
