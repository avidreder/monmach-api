import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {connect} from 'react-redux';
import _ from 'lodash';
import {List, Map, toJS} from 'immutable';
import * as actionCreators from '../action_creators';
import classnames from 'classnames';
import Dropdown from './Dropdown';
import moment from 'moment';

export const FilterMenu = React.createClass({
  mixins: [PureRenderMixin],
  getDateList: function () {
    return this.props.dateList.toArray();
  },
  getVenueList: function () {
    return this.props.venueList.toArray();
  },
  getNeighborhoodList: function() {
    return this.props.neighborhoodList.toArray();
  },
  getItemValue: function(item, key) {
    return Map.isMap(item) ? item.get(key) : item[key];
  },
  isActive: function(filter) {
    return Map.isMap(this.props.filters) ? this.props.filters.has(filter) : _.has(this.props.filters,filter);
  },
  render: function() {
    return <div className="container-fluid">
      <div className={classnames('well', 'row')}>
        <div className={classnames('col-md-4','col-sm-4')}>
          {this.isActive('date') ? <h5>Date</h5> : null}
          {this.isActive('date') ?
            this.getDateList().map(date =>
              <button key={date} className={classnames('date_button','btn', 'btn-default')} onClick={() => this.props.addFilter('date', this.getItemValue({date},'date'))}>{date}</button>
            )
            : <button className="add_date_button" onClick={() => this.props.addFilter('date',moment().format('ddd; M/D/YY'))}>Date</button>
          }
        </div>
        <div className={classnames('col-md-4','col-sm-4')}>
          {this.isActive('venue') ? <h5>Venue</h5> : null}
          {this.isActive('venue') ?
            <Dropdown itemList={this.getVenueList()}
                      addFilter={this.props.addFilter} /> :
            <button className="add_venue_button" onClick={() => this.props.addFilter('venue',this.getVenueList()[0])}>Venue</button>
          }
        </div>
        <div className={classnames('col-md-4','col-sm-4')}>
          <button type="button" className={classnames('reset_button','btn', 'btn-default')} onClick={() => this.props.resetFilters()}>Clear</button>
        </div>
      </div>
    </div>;
  }
});

function mapStateToProps(state) {
  return {
    shows: state.get('shows'),
    filters: state.get('filters'),
    dateList: state.get('dateList'),
    venueList: state.get('venueList'),
    neighborhoodList: state.get('neighborhoodList')
  };
}

function mapDispatchToProps(dispatch) {
  return {
    resetFilters: function() {
      dispatch(actionCreators.resetFilters());
    },
    addFilter: function(filterKey, filterValue) {
      dispatch(actionCreators.addFilter(filterKey, filterValue));

    },
    removeFilter: function(filterKey) {
      dispatch(actionCreators.removeFilter(filterKey));
    }
  }
}

export const FilterMenuContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(FilterMenu);
