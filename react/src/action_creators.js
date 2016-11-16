import fetch from 'isomorphic-fetch';
import _ from 'lodash';
var moment = require('moment');
var querystring = require('querystring');

export function openModal(modalType) {
  return {
    type: 'OPEN_MODAL',
    modalType
  };
}

export function closeModal(modalType) {
  return {
    type: 'CLOSE_MODAL',
    modalType
  };
}

export function setActiveShow(show) {
  return {
    type: 'SET_ACTIVE_SHOW',
    show
  };
}

export function setActiveBand(band) {
  return {
    type: 'SET_ACTIVE_BAND',
    band
  };
}

export function setState(state) {
  return {
    type: 'SET_STATE',
    state
  };
}

export function addToDeck(show) {
  return {
    type: 'ADD_TO_DECK',
    show
  }
}

export function removeFromDeck(show) {
  return {
    type: 'REMOVE_FROM_DECK',
    show
  }
}

export function addFilter(filterKey, filterValue) {
  return {
    type: 'ADD_FILTER',
    filterKey,
    filterValue
  }
}

export function removeFilter(filterKey) {
  return {
    type: 'REMOVE_FILTER',
    filterKey
  }
}

export function resetFilters() {
  return {
    type: 'RESET_FILTERS'
  }
}

export function updateVenueList(shows) {
  let venueList = _.pluck(_.uniq(shows, 'venue'),'venue');
  return {
    type: 'UPDATE_VENUE_LIST',
    venueList: venueList
  }
}

export function updateDateList(shows) {
  let dateList = _.pluck(_.uniq(shows, 'date'),'date');
  return {
    type: 'UPDATE_DATE_LIST',
    dateList: dateList
  }
}

export function requestShows() {
  return {
    type: 'REQUEST_SHOWS'
  }
}

export function receiveShowsSuccess(shows) {
  return {
    type: 'RECEIVE_SHOWS_SUCCESS',
    shows: _.uniq(shows,'id')
  }
}

export function receiveShowsError(error) {
  return {
    type: 'RECEIVE_SHOWS_ERROR',
    error: error
  }
}

export function requestVideos() {
  return {
    type: 'REQUEST_VIDEOS'
  }
}

export function receiveVideosSuccess(videos) {
  return {
    type: 'RECEIVE_VIDEOS_SUCCESS',
    videos: _.uniq(videos,'id')
  }
}

export function receiveVideosError(error) {
  return {
    type: 'RECEIVE_VIDEOS_ERROR',
    error: error
  }
}

export function retrieveStorage() {
  return {
    type: 'RETRIEVE_STORAGE'
  }
}

export function setStorage() {
  return {
    type: 'SET_STORAGE'
  }
}

export function sendMultiple(shows, dispatch) {
  dispatch(receiveShowsSuccess(shows));
  dispatch(updateVenueList(shows));
  dispatch(updateDateList(shows));
}

function handleErrors(response) {
    if (!response.ok) {
        throw Error(response.statusText);
    }
    return response;
}

// export function fetchShows() {
//   let searchParameters = JSON.stringify({
//     startDate : moment().format(),
//     results: 10,
//     longitude:-122.675628662109,
//     latitude:45.511791229248
//   });

//   return function (dispatch) {

//     dispatch(requestShows())

//     return fetch('http://localhost:3000/get-bit-events', {
//         method: 'post',
//         headers: {
//             "Content-type": "application/json"
//         },
//         body: searchParameters
//       })
//       .then(handleErrors)
//       .then(function(response){
//         return response.json();
//       })
//       .then(function(json){
//         sendMultiple(json,dispatch);
//       }).catch(function(error){
//         dispatch(receiveShowsError(error));
//       })
//   }
// }

export function fetchShows() {
  let searchParameters = JSON.stringify({
    startDate : moment().format(),
    results: 10,
    longitude:-122.675628662109,
    latitude:45.511791229248
  });

  return function (dispatch) {

    dispatch(requestShows())

    return fetch('http://localhost:3000/bitshows', {
        method: 'get',
        headers: {
            "Content-type": "application/json"
        },
      })
      .then(handleErrors)
      .then(function(response){
        return response.json();
      })
      .then(function(json){
        sendMultiple(json,dispatch);
      }).catch(function(error){
        dispatch(receiveShowsError(error));
      })
  }
}

export function fetchVideos(searchTerm) {
  let query = JSON.stringify({ query: searchTerm });
  console.log(query);
  console.log(searchTerm);
  return function (dispatch) {

    dispatch(requestVideos())

    return fetch('http://localhost:3000/search-videos', {
        method: 'post',
        headers: {
            "Content-type": "application/json"
        },
        body: query
      })
      .then(handleErrors)
      .then(function(response){
        console.log(response);
        return response.json();
      })
      .then(function(json){
        dispatch(receiveVideosSuccess(json));
      }).catch(function(error){
        dispatch(receiveVideosError(error));
      })
  }
}
