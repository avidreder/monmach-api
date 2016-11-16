import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {connect} from 'react-redux';
import ShowTile from './ShowTile';
import ShowPaginator from './ShowPaginator';
import _ from 'lodash';
import {List, Map} from 'immutable';
import * as actionCreators from '../action_creators';
import classnames from 'classnames';
import * as filters from '../filters';

export const ShowFeed = React.createClass({
  mixins: [PureRenderMixin],
  shouldPin: function(show) {
    return List.isList(this.props.deck) ? this.props.deck.contains(show) : _.some(this.props.deck,show);
  },
  getId: function(show) {
    return Map.isMap(show) ? show.get('id') : show.id;
  },
  render: function() {
    return <div className={classnames('container-fluid')}>
      <div className={classnames('row')}>
        <h3>{this.props.feedType}</h3>
      </div>
      <div className={classnames('row', this.props.feedType + '_feed')}>
        {this.props.paginated ?
          <ShowPaginator shows={this.props.showData}
                     deck={this.props.deck}
                     openModal={this.props.openModal}
                     dispatchAdd={this.props.dispatchAdd}
                     dispatchRemove={this.props.dispatchRemove} />
                     :
        this.props.showData.map(show =>
        <div key={this.props.feedType+'_'+this.getId(show)} className={classnames('col-md-4', 'col-sm-4')}>
          <ShowTile show={show}
                    isPinned={this.shouldPin(show)}
                    feedType={this.props.feedType}
                    removeShow={this.props.dispatchRemove}
                    pinShow={this.props.dispatchAdd}
                    openModal={this.props.openModal} />
        </div>
        )}
      </div>
    </div>;
  }
});

function mapStateToProps(state) {
  return {
    shows: state.get('shows'),
    deck: state.get('deck'),
    filters: state.get('filters')
  };
}

function mapDispatchToProps(dispatch) {
  return {
    dispatchAdd: function(show) {
      dispatch(actionCreators.addToDeck(show));
    },
    dispatchRemove: function(show) {
      dispatch(actionCreators.removeFromDeck(show));
    },
    openModal: function(show) {
      dispatch(actionCreators.setActiveShow(show));
      dispatch(actionCreators.openModal('shows'));
    }
  }
}

export const ShowFeedContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ShowFeed);
