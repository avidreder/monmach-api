import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import NavigationExpandMoreIcon from 'material-ui/svg-icons/navigation/expand-more';
import FontIcon from 'material-ui/FontIcon';
import Paper from 'material-ui/Paper';
import {Toolbar, ToolbarGroup, ToolbarSeparator, ToolbarTitle} from 'material-ui/Toolbar';
export default React.createClass({
    mixins: [PureRenderMixin],
    render: function() {
	return <div>
	    <Paper>
		<Toolbar>
		    <ToolbarGroup>
			<ToolbarTitle text={this.props.track.Name} />
		    </ToolbarGroup>
		    <ToolbarGroup>
	<FontIcon className="material-icons" onClick={() => this.props.setTrack(this.props.track)}>play_circle_outline</FontIcon>
	<FontIcon className="material-icons" onClick={() => this.props.removeFromQueue(this.props.track)}>not_interested</FontIcon>
			<FontIcon className="material-icons">playlist_add</FontIcon>
		    </ToolbarGroup>
		</Toolbar>
	    </Paper>
	</div>
    }
});