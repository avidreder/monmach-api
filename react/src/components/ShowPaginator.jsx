import React from 'react';
import ShowTile from './ShowTile';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {Map, List} from 'immutable';
import classnames from 'classnames';
import _ from 'lodash';

export default React.createClass({
  mixins: [PureRenderMixin],
  getInitialState: function() {
    return {
      pageIndex: 0,
      showsPerPage: 3,
      maxPage: Math.floor(this.props.shows.size / 3)
    }
  },
  shouldPin: function(show) {
    return List.isList(this.props.deck) ? this.props.deck.contains(show) : _.some(this.props.deck,show);
  },
  getId: function(show) {
    return Map.isMap(show) ? show.get('id') : show.id || "";
  },
  pageUp: function() {
    if (this.state.pageIndex < this.state.maxPage) {
      this.setState({pageIndex: this.state.pageIndex + 1});
    }
  },
  pageDown: function() {
    if (this.state.pageIndex > 0) {
      this.setState({pageIndex: this.state.pageIndex - 1});
    }
  },
  getShows: function() {
    return List.isList(this.props.shows) ? this.props.shows.toArray() || [] : this.props.shows || [];
  },
  getVisibleShows: function() {
    let currentIndex = this.state.pageIndex;
    return _.filter(this.getShows(), function(show, index) {
      return Math.floor(index / 3) == currentIndex;
    });
  },
  render: function() {
    return <div className={classnames('recommended_show_feed')}>
      <div className={classnames('row', 'show_paginator', 'paginator')}>
        {this.getVisibleShows().map(show =>
          <div key={"recommended_shows_"+this.getId(show)} className={classnames('col-md-4', 'col-sm-4')}>
            <ShowTile show={show}
                      isPinned={this.shouldPin(show)}
                      feedType="show"
                      openModal={this.props.openModal}
                      removeShow={this.props.dispatchRemove}
                      pinShow={this.props.dispatchAdd} />
          </div>
        )}
      </div>
      <div style={{textAlign: 'center'}} className={classnames('row')}>
          <nav>
            <ul className="pagination">
              <li><a className="page_down" onClick={this.pageDown}><span className={classnames('glyphicon', 'glyphicon-chevron-left')}></span></a></li>
              {_.range(this.state.maxPage+1).map(item =>
                this.state.pageIndex == item ? <li key={item} className="active"><a></a></li> : <li key={item} className="inactive"><a></a></li>
              )}
              <li><a className="page_up" onClick={this.pageUp}><span className={classnames('glyphicon', 'glyphicon-chevron-right')}></span></a></li>
            </ul>
          </nav>
      </div>
    </div>;
  }
});
