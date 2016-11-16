import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {Map} from 'immutable';
import classnames from 'classnames';
import * as actionCreators from '../action_creators';

export default React.createClass({
  render: function() {
    return <div style={{textAlign:"center"}} className={classnames('row', 'well')}>
      <div className={classnames('col-md-12', 'col-sm-12')}>
        <h3>Loading Shows</h3>
        <h3><span className={classnames('glyphicon', 'glyphicon-refresh', 'glyphicon-spinning')}></span></h3>
      </div>
    </div>;
  }
});
