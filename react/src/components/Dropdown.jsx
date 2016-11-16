import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {Map} from 'immutable';
import classnames from 'classnames';

export default React.createClass({
  mixins: [PureRenderMixin],
  getInitialState: function() {
    return {value: this.props.itemList[0]};
  },
  selectItem: function(event) {
    this.setState({value: event.target.value});
    this.props.addFilter('venue', event.target.value);
  },
  getState: function() {
    return this.state.value;
  },
  render: function() {
    return <select onChange={this.selectItem} value={this.state.value} className={classnames('dropdown-list', 'form-control')}>
      {this.props.itemList.map(item =>
        <option className="list_item" key={item} value={item}>{item}</option>
      )}
    </select>;
  }
});
