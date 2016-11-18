import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import * as d3 from "d3";
import Paper from 'material-ui/Paper';
import Avatar from 'material-ui/Avatar';

export default React.createClass({
    mixins: [PureRenderMixin],
    render: function() {
	return <img src="img/chart.png" width="100%" />
    }
});
