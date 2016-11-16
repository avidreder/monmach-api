import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {connect} from 'react-redux';
import {ShowFeedContainer} from './ShowFeed';
import {ShowVideosContainer} from './ShowVideos';
import {NavMenuContainer} from './NavMenu';
import {FilterMenuContainer} from './FilterMenu';
import {ModalWindowContainer} from './Modal';
import Footer from './Footer';
import LoadingMessage from './LoadingMessage';
import {ErrorMessageContainer} from './ErrorMessage';
import _ from 'lodash';
import {List, Map} from 'immutable';
import * as actionCreators from '../action_creators';
import classnames from 'classnames';
import * as filters from '../filters';

export const HomePage = React.createClass({
  mixins: [PureRenderMixin],
  render: function() {
    return <div className="pageWrapper">
      <link rel="stylesheet" href="libs/bootstrap/dist/css/bootstrap.min.css" />
      <script src="libs/jquery/dist/jquery.min.js"></script>
      <script src="libs/bootstrap/dist/js/bootstrap.min.js"></script>
      <NavMenuContainer />
      <ModalWindowContainer internalComponent={
        <ShowFeedContainer feedType="deck"
                           showData={this.props.deck}
                           paginated={false} />
         }
                            modalType="deck" />
      <ModalWindowContainer internalComponent={<ShowVideosContainer />}
                          modalType="shows" />
      {this.props.isLoading ?
        <LoadingMessage /> :
        this.props.errors.size > 0 ?
          <ErrorMessageContainer errorMessage="Error Fetching Shows" /> :
        this.props.shows.size == 0 ?
          <ErrorMessageContainer errorMessage="No Results" /> :
          <div className={classnames('row', 'well')}>
            <div className={classnames('col-md-12', 'col-sm-12')}>
              <FilterMenuContainer />
            </div>
            <div className={classnames('col-md-12', 'col-sm-12')}>
              <ShowFeedContainer feedType="recommended_shows"
                                 showData={filters.filterRecommendedShows(this.props.shows, this.props.recommendedVenues)}
                                 paginated={true} />
            </div>
            <div className={classnames('col-md-12', 'col-sm-12')}>
              <ShowFeedContainer feedType="shows"
                                 showData={filters.filterShows(this.props.shows, this.props.filters)}
                                 paginated={false} />
            </div>
          </div>
      }
      <Footer />
    </div>;
  }
});

function mapStateToProps(state) {
  return {
    shows: state.get('shows'),
    deck: state.get('deck'),
    filters: state.get('filters'),
    isLoading: state.get('isLoading'),
    errors: state.get('errors'),
    recommendedVenues: state.get('recommendedVenues')
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
    fetchShows: function() {
      dispatch(actionCreators.fetchShows());
    }
  }
}

export const HomePageContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(HomePage);
