import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {connect} from 'react-redux';
import {ShowFeedContainer} from './ShowFeed';
import _ from 'lodash';
import {List, Map} from 'immutable';
import * as actionCreators from '../action_creators';

export const NavMenu = React.createClass({
  mixins: [PureRenderMixin],
  render: function() {
    return <div className="navbar navbar-default" id="home-navbar">
      <div className="container">
        <div className="navbar-header">
        <img style={{cursor: 'pointer'}} onClick={() => this.props.fetchShows()} alt="Brand" src="../img/Hawk.png" height="50" />
          <button type="button" className="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
            <span className="icon-bar"></span>
          </button>
          <a style={{cursor: 'pointer'}} className="navbar-brand" onClick={() => this.props.fetchShows()}>ShowHawk</a>
        </div>
        <div className="collapse navbar-collapse">
          <ul className="nav navbar-nav">
            <li style={{cursor: 'pointer'}} className="navbar-link"><a onClick={() => this.props.fetchShows()}>Home</a></li>
            <li style={{cursor: 'pointer'}} className="navbar-link"><a onClick={() => this.props.openModal('deck')}>Deck <span className="badge">{this.props.deck.size}</span></a></li>
          </ul>
        </div>
      </div>
    </div>;
  }
});

function mapDispatchToProps(dispatch) {
  return {
    openModal: function(type) {
      dispatch(actionCreators.openModal(type));
    },
    fetchShows: function() {
      dispatch(actionCreators.fetchShows());
    }
  }
}

function mapStateToProps(state) {
  return {
    shows: state.get('shows'),
    deck: state.get('deck')
  };
}
export const NavMenuContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(NavMenu);
