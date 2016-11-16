import {List, Map, fromJS} from 'immutable';
import _ from 'lodash'

export var mock = {
  store: {},
  getItem: function(key) {
    return this.store[key];
  },
  setItem: function(key, value) {
    this.store[key] = value.toString();
  },
  clear: function() {
    this.store = {};
  }
};

if (typeof localStorage === "undefined" || localStorage === null) {
  var localStorage = mock;
}

const initialState = fromJS({
  debug: true,
  shows: [
    {
      bands: [ 'Bluegrass Tuesdays', 'w', ' Pete Kartsounes and Friends' ],
      time: '8:30 PM',
      date: 'Tue, 5/10/16',
      price: 0,
      id: '101cf2acfdff4d8da64af777299d6f9d',
      venue: 'The Ranger Station PDX',
      address: '4260 SE Hawthorne Blvd',
      image: 'https://citysparkstorage.blob.core.windows.net/portalimages/portalimages/39ea9bf3-4639-44c2-a9be-b637647ba7c7.medium.png',
      neighborhood: 'SE',
      popularity: 4
    },
    {
      bands: [ 'Cool Band', 'Another Cool Band'],
      time: '7:30 PM',
      date: 'Mon, 5/9/16',
      price: 5,
      id: '7d4ef2acfdff4d8da64af777299d7d4e',
      venue: 'My Favorite Bar',
      address: '434 NW Hermosa',
      image: 'https://citysparkstorage.blob.core.windows.net/portalimages/portalimages/39ea9bf3-4639-44c2-a9be-b637647ba7c7.medium.png',
      neighborhood: 'NW',
      popularity: 10
    }
  ],
  deck: [
    {
      bands: [ 'Bluegrass Tuesdays', 'w', ' Pete Kartsounes and Friends' ],
      time: '8:30 PM',
      date: 'Tue, 5/10/16',
      price: 0,
      id: '101cf2acfdff4d8da64af777299d6f9d',
      venue: 'The Ranger Station PDX',
      address: '4260 SE Hawthorne Blvd',
      image: 'https://citysparkstorage.blob.core.windows.net/portalimages/portalimages/39ea9bf3-4639-44c2-a9be-b637647ba7c7.medium.png',
      neighborhood: 'SE',
      popularity: 4
    }
  ],
  dateList: ['Mon, 5/9/16','Tue, 5/10/16','Wed, 5/11/16'],
  venueList: ['The Ranger Station PDX','My Favorite Bar'],
  neighborhoodList: ['SE','SW','NE','NW','N'],
  recommendedVenues: ['My Favorite Bar'],
  filters: {},
  modalsOpen: [],
  isLoading: false,
  errors: [],
  activeShow: {},
  activeBand: {},
  videos: [],
  isLoadingVideos: false
});

function setState(state, newState) {
  return state.merge(newState);
}

function addToDeck(state, show) {
  let currentDeck = state.get('deck');
  if (currentDeck && !currentDeck.contains(fromJS(show))) {
    return setStorage(state.set('deck', currentDeck.push(fromJS(show))));
  } else {
    return state;
  }
}

function removeFromDeck(state, show) {
  let currentDeck = state.get('deck');
  if (currentDeck && currentDeck.includes(fromJS(show))) {
    return setStorage(state.set('deck', currentDeck.delete(currentDeck.indexOf(fromJS(show)))));
  } else {
    return state;
  }
}

function addFilter(state, filterKey, filterValue) {
  let currentFilters = state.get('filters');
  if (currentFilters) {
    return setStorage(state.set('filters', currentFilters.set(filterKey, filterValue)));
  } else {
    return state;
  }
}

function removeFilter(state, filterKey) {
  let currentFilters = state.get('filters');
  if (currentFilters && currentFilters.has(filterKey)) {
    return setStorage(state.set('filters', currentFilters.delete(filterKey)));
  } else {
    return state;
  }
}

function retrieveStorage(state) {
  return localStorage.getItem('state') != null ? fromJS(JSON.parse(localStorage.getItem('state'))) : state;
}

function setStorage(state) {
  localStorage.setItem('state', JSON.stringify(state.toJS()));
  return state;
}

function updateVenueList (state, venueList) {
  let newState = state;
  let filters = state.get('filters').toJS();
  newState = newState.set('venueList', fromJS(venueList));
  return _.contains(venueList,filters.venue) ? setStorage(newState) : setStorage(newState.set('filters', fromJS(_.omit(filters,'venue'))));
}

function updateDateList (state, dateList) {
  let newState = state;
  let filters = state.get('filters').toJS();
  newState = newState.set('dateList', fromJS(dateList));
  return _.contains(dateList,filters.date) ? setStorage(newState) : setStorage(newState.set('filters', fromJS(_.omit(filters,'date'))));
}

function receiveShowsSuccess(state, shows) {
  state = state.set('shows', fromJS(shows));
  state = state.set('activeShow', fromJS(shows[0]));
  state = state.set('errors', List());
  return setStorage(state.set('isLoading', false));
}

function receiveShowsError(state, error) {
  state = state.set('shows', List());
  state = state.set('activeShow', Map());
  let errors = state.get('errors');
  if (errors) {
    state = state.set('errors', errors.push(error));
  } else {
    state = state.set('errors', List.of(error));
  }
  return state.set('isLoading', false);
}

function requestShows(state) {
  return state.set('isLoading',true);
}

function receiveVideosSuccess(state, videos) {
  state = state.set('videos', fromJS(videos));
  state = state.set('errors', List());
  return setStorage(state.set('isLoadingVideos', false));
}

function receiveVideosError(state, error) {
  state = state.set('videos', List());
  let errors = state.get('errors');
  if (errors) {
    state = state.set('errors', errors.push(error));
  } else {
    state = state.set('errors', List.of(error));
  }
  return state.set('isLoadingVideos', false);
}

function requestVideos(state) {
  return state.set('isLoadingVideos',true);
}

function resetFilters(state) {
  return setStorage(state.set('filters',Map()));
}

function openModal(state, modalType) {
  let currentModalsOpen = state.get('modalsOpen');
  if (currentModalsOpen && !currentModalsOpen.contains(fromJS(modalType))) {
    return setStorage(state.set('modalsOpen', currentModalsOpen.push(fromJS(modalType))));
  } else {
    return state;
  }
}

function closeModal(state, modalType) {
  let currentModalsOpen = state.get('modalsOpen');
  if (currentModalsOpen && currentModalsOpen.includes(fromJS(modalType))) {
    return setStorage(state.set('modalsOpen', currentModalsOpen.delete(currentModalsOpen.indexOf(fromJS(modalType)))));
  } else {
    return state;
  }
}

function setActiveShow(state, show) {
  return setStorage(state.set('activeShow',fromJS(show)));
}

function setActiveBand(state, band) {
  return setStorage(state.set('activeBand',fromJS(band)));
}

export default function(state = initialState, action) {
    switch (action.type) {
    case 'OPEN_MODAL':
      return openModal(state, action.modalType);
    case 'CLOSE_MODAL':
      return closeModal(state, action.modalType);
    case 'SET_ACTIVE_SHOW':
      return setActiveShow(state, action.show);
    case 'SET_ACTIVE_BAND':
      return setActiveBand(state, action.band);
    case 'SET_STORAGE':
      return setStorage(state);
    case 'RETRIEVE_STORAGE':
      return retrieveStorage(state);
    case 'REQUEST_SHOWS':
      return requestShows(state);
    case 'REQUEST_VIDEOS':
      return requestVideos(state);
    case 'UPDATE_VENUE_LIST':
      return updateVenueList(state, action.venueList);
    case 'UPDATE_DATE_LIST':
      return updateDateList(state, action.dateList);
    case 'RECEIVE_SHOWS_SUCCESS':
      return receiveShowsSuccess(state, action.shows);
    case 'RECEIVE_SHOWS_ERROR':
      return receiveShowsError(state, action.error);
    case 'RECEIVE_VIDEOS_SUCCESS':
      return receiveVideosSuccess(state, action.videos);
    case 'RECEIVE_VIDEOS_ERROR':
      return receiveVideosError(state, action.error);
    case 'RESET_FILTERS':
      return resetFilters(state);
    case 'ADD_FILTER':
      return addFilter(state, action.filterKey, action.filterValue);
    case 'REMOVE_FILTER':
      return removeFilter(state, action.filterKey);
    case 'SET_STATE':
      return state.merge(action.state);
    case 'ADD_TO_DECK':
      return addToDeck(state, action.show);
    case 'REMOVE_FROM_DECK':
      return removeFromDeck(state, action.show);
  }
  return state;
}
