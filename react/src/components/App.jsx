import React from 'react';
import {List, Map, fromJS} from 'immutable';

export default React.createClass({
  render: function() {
    return React.cloneElement(this.props.children, {
    });
  }
});
