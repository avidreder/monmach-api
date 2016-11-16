import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {Map} from 'immutable';
import classnames from 'classnames';
import * as actionCreators from '../action_creators';
import {connect} from 'react-redux';

export const ErrorMessage = React.createClass({
  mixins: [PureRenderMixin],
  render: function() {
    return <div style={{textAlign:"center"}} className={classnames('row', 'well')}>
    <div className={classnames('col-md-12', 'col-sm-12','error_message')}>
      <h3>{this.props.errorMessage}</h3>
      <button onClick={() => this.props.fetchShows()}><span className={classnames('glyphicon', 'glyphicon-refresh')}></span>Retry</button>
    </div>
  </div>;
  }
});

function mapStateToProps(state) {
  return {};
}

function mapDispatchToProps(dispatch) {
  return {
    fetchShows: function() {
      dispatch(actionCreators.fetchShows());
    }
  }
}

export const ErrorMessageContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ErrorMessage);
