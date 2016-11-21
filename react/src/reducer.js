import {List, Map, fromJS, toJS} from 'immutable';
import _ from 'lodash'

const initialState = fromJS(
    {
	currentTrack: {
	    ID: 1,
	    Name: "Wish You Were Here",
	    Artists: ["Pink Floyd"],
	    ImageURL: "https://i.scdn.co/image/aab31a87e274822dd11c1de4b6e851aa3a471500",
	    SpotifyID: "6mFkJmJqdDVQ1REhVfGgd1",
	    Genres: [],
	    Playlists: [1],
	    Rating: 0,
	    Created: 0,
	    Updated: 0,
	    Features: [0.481, 0.262, 7, -15.730, 1, 0.0414, 0.735, 0.0114, 0.832, 0.358, 122.883, 334744, 4] 
	},
	queue: {
	    ID: 1,
	    UserID: 1,
	    Name: "TestQueue",
	    MaxSize: 10,
	    TrackQueue:[
	    {
		ID: 1,
		Name: "Wish You Were Here",
		Artists: ["Pink Floyd"],
		ImageURL: "https://i.scdn.co/image/aab31a87e274822dd11c1de4b6e851aa3a471500",
		SpotifyID: "6mFkJmJqdDVQ1REhVfGgd1",
		Genres: [],
		Playlists: [1],
		Rating: 0,
		Created: 0,
		Updated: 0,
		Features: [0.481, 0.262, 7, -15.730, 1, 0.0414, 0.735, 0.0114, 0.832, 0.358, 122.883, 334744, 4] 
	    },
	    {
		ID: 2,
		Name: "The Trooper - 1998 Remastered Version",
		Artists: ["Iron Maiden"],
		ImageURL: "https://i.scdn.co/image/61db744786108eb24f8d54f6e9eaf3eb3331feea",
		SpotifyID: "3MXA2BkBk0lSuMpRoM7SK2",
		Genres: [1],
		Playlists: [1],
		Rating: 0,
		Created: 0,
		Updated: 0,
		Features: [0.288, 0.864, 4, -6.762, 0, 0.0673, 0.0172, 0.0567, 0.541, 0.640, 162.233, 250947, 4]
	    },
	    {
		ID: 3,
		Name: "Landslide",
		Artists: ["The Smashing Pumpkins"],
		ImageURL: "https://i.scdn.co/image/0fc777e6f5f813d3de6dc1cbeef322f0edd09619",
		SpotifyID: "2dc1BYopTHgviDXawShfME",
		Genres: [1],
		Playlists: [1],
		Rating: 0,
		Created: 0,
		Updated: 0,
		Features: [0.422, 0.208, 10, -15.886, 1, 0.0318, 0.896, 0.0265, 0.192, 0.377, 111.546, 190587, 3] 
	    }
	    ],
	    SeedArtists: [],
	    SeedTracks:[],
	    ListenedTracks:[],
	    Created:0,
	    Updated:0
	},
	playlist: {
	    ID: 1,
	    Name: "TestQueue",
	    UserID: 1,
	    SpotifyID: "",
	    Tracks:[],
	    ListenedTracks:[],
	    Created:0,
	    Updated:0
	},
	genre: {
	    ID: 1,
	    UserID: 1,
	    QueueID: 1,
	    Name: "Great Music",
	    Description: "My favorite songs",
	    SeedArtists: [],
	    SeedTracks:[],
	    SeedPlaylists:[],
	    AvatarURL: "",
	    Created:0,
	    Updated:0,
	    TrackBlacklist: [],
	    TrackWhitelist: []
	}
    }
);


/*{
	  "danceability" : 0.481,
	  "energy" : 0.262,
	  "key" : 7,
	  "loudness" : -15.730,
	  "mode" : 1,
	  "speechiness" : 0.0414,
	  "acousticness" : 0.735,
	  "instrumentalness" : 0.0114,
	  "liveness" : 0.832,
	  "valence" : 0.358,
	  "tempo" : 122.883,
	  "type" : "audio_features",
	  "id" : "6mFkJmJqdDVQ1REhVfGgd1",
	  "uri" : "spotify:track:6mFkJmJqdDVQ1REhVfGgd1",
	  "track_href" : "https://api.spotify.com/v1/tracks/6mFkJmJqdDVQ1REhVfGgd1",
	  "analysis_url" : "https://api.spotify.com/v1/audio-analysis/6mFkJmJqdDVQ1REhVfGgd1",
	  "duration_ms" : 334744,
	  "time_signature" : 4
}*/

function setState(state, newState) {
  return state.merge(newState);
}

function setTrack(state, track) {
    state = removeFromQueue(addToListened(state, track), track)
    return state.set('currentTrack', fromJS(track));
}

function removeFromQueue(state, track) {
    var queue = state.get('queue')
    var newQueue = queue.set('TrackQueue', fromJS(_.reject(queue.get('TrackQueue').toJS(), {ID: track.ID})))
    return state.set('queue', newQueue);
}

function addToListened(state, track) {
    var queue = state.get('queue');
    queue.set('ListenedTracks', queue.get('ListenedTracks').push(track.ID));
    return state.set('queue', queue);
}

function addGenre(state, track) {
    track.Genres.push(state.get('genre').get('ID'))
    return state.set('currentTrack', fromJS(track));
}

// function receiveShowsSuccess(state, shows) {
//   state = state.set('shows', fromJS(shows));
//   state = state.set('activeShow', fromJS(shows[0]));
//   state = state.set('errors', List());
//   return setStorage(state.set('isLoading', false));
// }

// function receiveShowsError(state, error) {
//   state = state.set('shows', List());
//   state = state.set('activeShow', Map());
//   let errors = state.get('errors');
//   if (errors) {
//     state = state.set('errors', errors.push(error));
//   } else {
//     state = state.set('errors', List.of(error));
//   }
//   return state.set('isLoading', false);
// }

// function requestShows(state) {
//   return state.set('isLoading',true);
// }

// function openModal(state, modalType) {
//   let currentModalsOpen = state.get('modalsOpen');
//   if (currentModalsOpen && !currentModalsOpen.contains(fromJS(modalType))) {
//     return setStorage(state.set('modalsOpen', currentModalsOpen.push(fromJS(modalType))));
//   } else {
//     return state;
//   }
// }

// function closeModal(state, modalType) {
//   let currentModalsOpen = state.get('modalsOpen');
//   if (currentModalsOpen && currentModalsOpen.includes(fromJS(modalType))) {
//     return setStorage(state.set('modalsOpen', currentModalsOpen.delete(currentModalsOpen.indexOf(fromJS(modalType)))));
//   } else {
//     return state;
//   }
// }

export default function(state = initialState, action) {
    switch (action.type) {
    case 'SET_CURRENT_TRACK':
	return setTrack(state, action.track);
    case 'REMOVE_FROM_QUEUE':
	return removeFromQueue(state, action.track);
    case 'ADD_TO_LISTENED':
	return addToListened(state, action.track);
    case 'ADD_GENRE':
	return addGenre(state, action.track);
	// case 'REQUEST_SHOWS':
	//   return requestShows(state);
	// case 'RECEIVE_SHOWS_SUCCESS':
	//   return receiveShowsSuccess(state, action.shows);
	// case 'RECEIVE_SHOWS_ERROR':
	//   return receiveShowsError(state, action.error
	// );
    case 'SET_STATE':
	return state.merge(action.state);
    }
    return state;
}
