import {Map, List, fromJS, toJS} from 'immutable';
import _ from 'lodash';

export function filterShows(shows, filters) {
  // ToDo(dev) fix this ugly lil guy
  return (_.keys(filters).length == 0 || filters == Map()) ? shows : List.isList(shows) ? fromJS(_.filter(shows.toJS(), filters.toJS())) : _.filter(shows, filters);
}

export function filterRecommendedShows(shows, venues) {
  // ToDo(dev) fix this scary function
  let filteredShows = List.isList(shows) ? List() : [];
  venues.map(venue => filteredShows = List.isList(shows) ? filteredShows.concat(fromJS(_.filter(shows.toJS(), {'venue': venue}))) : filteredShows.concat(_.filter(shows, {'venue': venue})));
  return filteredShows;
}
