import fetch from 'isomorphic-fetch';
import _ from 'lodash';

// export function openModal(modalType) {
//   return {
//     type: 'OPEN_MODAL',
//     modalType
//   };
// }

// export function closeModal(modalType) {
//   return {
//     type: 'CLOSE_MODAL',
//     modalType
//   };
// }

export function setState(state) {
  return {
    type: 'SET_STATE',
    state
  };
}

export function setTrack(track) {
  return {
    type: 'SET_CURRENT_TRACK',
    track
  };
}

export function removeFromQueue(track) {
  return {
    type: 'REMOVE_FROM_QUEUE',
    track
  };
}

export function addGenre(track) {
  return {
    type: 'ADD_GENRE',
    track
  };
}

// export function requestShows() {
//   return {
//     type: 'REQUEST_SHOWS'
//   }
// }

// export function receiveShowsSuccess(shows) {
//   return {
//     type: 'RECEIVE_SHOWS_SUCCESS',
//     shows: _.uniq(shows,'id')
//   }
// }

// export function receiveShowsError(error) {
//   return {
//     type: 'RECEIVE_SHOWS_ERROR',
//     error: error
//   }
// }

// export function sendMultiple(shows, dispatch) {
//   dispatch(receiveShowsSuccess(shows));
//   dispatch(updateVenueList(shows));
//   dispatch(updateDateList(shows));
// }

// function handleErrors(response) {
//     if (!response.ok) {
//         throw Error(response.statusText);
//     }
//     return response;
// }

// export function fetchShows() {
//   let searchParameters = JSON.stringify({
//     startDate : moment().format(),
//     results: 10,
//     longitude:-122.675628662109,
//     latitude:45.511791229248
//   });

//   return function (dispatch) {

//     dispatch(requestShows())

//     return fetch('http://localhost:3000/bitshows', {
//         method: 'get',
//         headers: {
//             "Content-type": "application/json"
//         },
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
