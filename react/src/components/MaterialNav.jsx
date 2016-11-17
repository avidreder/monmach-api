import React from 'react';
import { Link } from 'react-router';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import IconMenu from 'material-ui/IconMenu';
import IconButton from 'material-ui/IconButton';
import FontIcon from 'material-ui/FontIcon';
import AppBar from 'material-ui/AppBar';
import FloatingActionButton from 'material-ui/FloatingActionButton';
import NavigationExpandMoreIcon from 'material-ui/svg-icons/navigation/expand-more';
import {List, ListItem} from 'material-ui/List';
import MenuItem from 'material-ui/MenuItem';
import DropDownMenu from 'material-ui/DropDownMenu';
import RaisedButton from 'material-ui/RaisedButton';
import Drawer from 'material-ui/Drawer';
import Avatar from 'material-ui/Avatar';
import {Toolbar, ToolbarGroup, ToolbarSeparator, ToolbarTitle} from 'material-ui/Toolbar';
import FlatButton from 'material-ui/FlatButton';

export default React.createClass({
    mixins: [PureRenderMixin],
    getInitialState: function() {
	return {value: 3, open: false};
    },
    handleToggle: function(){
	this.setState({open: !this.state.open});
    },
    handleChange: function(event, index, value){
	this.setState({value});
    },
    handleClose: function(){
	this.setState({open: false});
    },
    render: function() {
	return <div>
	    <Drawer
          docked={false}
          width={200}
          open={this.state.open}
          onRequestChange={(open) => this.setState({open})}
            >
	    <List>
	    <ListItem leftAvatar={<Avatar src="../img/GridLogo.png" size={40} onTouchTap={this.handleToggle} />} onTouchTap={this.handleClose}><Link to="/">Avidreder.net</Link></ListItem>
            <ListItem onTouchTap={this.handleClose} leftAvatar={<Avatar src="../img/Hawk.png" size={40} onTouchTap={this.handleToggle} />}><Link to="/components">ShowHawk</Link></ListItem>
	    <ListItem onTouchTap={this.handleClose} leftAvatar={<Avatar size={40} onTouchTap={this.handleToggle}>M</Avatar>}><Link to="/">MonsterMachine</Link></ListItem>
	    </List>
            </Drawer>
	    <AppBar
	title={"Avidreder.net"}
	onTitleTouchTap={this.handleToggle}
	onLeftIconButtonTouchTap={this.handleToggle}
    iconElementRight={<div><Avatar
          src="../img/GridLogo.png"
          size={40}
          onTouchTap={this.handleToggle}
            />
	    <Avatar
          src="../img/Hawk.png"
          size={40}
          onTouchTap={this.handleToggle}
            />
	    <Avatar
          size={40}
          onTouchTap={this.handleToggle}
            >M</Avatar></div>} />
	    </div>;
  }
});
