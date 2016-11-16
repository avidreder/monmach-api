import {List, Map, fromJS} from 'immutable';
import {expect} from 'chai';
require('isomorphic-fetch')
var fetchMock = require('fetch-mock');

fetchMock.mock('http://localhost:3000/get-ww-events',
  [{
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
  }]
);
import reducer, {mock} from '../src/reducer';

if (typeof localStorage === "undefined" || localStorage === null) {
  var localStorage = mock;
}



describe('reducer', () => {

  it('handles SET_STATE', () => {
    const initialState = Map();
    const action = {
      type:'SET_STATE',
      state: Map({
        shows: fromJS([
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
        ])
      })
    };

    const nextState = reducer(initialState, action);

    expect(nextState).to.equal(fromJS({
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
        }
      ]
    }));
  });

  it('handles SET_STATE with plain JS payload', () => {
    const initialState = Map();
    const action = {
      type: 'SET_STATE',
      state: {
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
          }
        ]
      }
    };
    const nextState = reducer(initialState, action);

    expect(nextState).to.equal(fromJS({
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
        }
      ]
    }));
  });

  // The test data in reducer.js (for development), makes this test a pain
  //
  // it('handles SET_STATE without initial state', () => {
  //   const action = {
  //     type: 'SET_STATE',
  //     state: {
  //       shows: [
  //         {
  //           bands: [ 'Bluegrass Tuesdays', 'w', ' Pete Kartsounes and Friends' ],
  //           time: '8:30 PM',
  //           date: 'Tue, 5/10/16',
  //           price: 0,
  //           id: '101cf2acfdff4d8da64af777299d6f9d',
  //           venue: 'The Ranger Station PDX',
  //           address: '4260 SE Hawthorne Blvd',
  //           image: 'https://citysparkstorage.blob.core.windows.net/portalimages/portalimages/39ea9bf3-4639-44c2-a9be-b637647ba7c7.medium.png',
  //           neighborhood: 'SE',
  //           popularity: 4
  //         },
  //         {
  //           bands: [ 'Cool Band', 'Another Cool Band'],
  //           time: '7:30 PM',
  //           date: 'Mon, 5/9/16',
  //           price: 5,
  //           id: '7d4ef2acfdff4d8da64af777299d7d4e',
  //           venue: 'My Favorite Bar',
  //           address: '434 NW Hermosa',
  //           image: 'https://citysparkstorage.blob.core.windows.net/portalimages/portalimages/39ea9bf3-4639-44c2-a9be-b637647ba7c7.medium.png',
  //           neighborhood: 'NW',
  //           popularity: 10
  //         }
  //       ],
  //       deck: [
  //         {
  //           bands: [ 'Bluegrass Tuesdays', 'w', ' Pete Kartsounes and Friends' ],
  //           time: '8:30 PM',
  //           date: 'Tue, 5/10/16',
  //           price: 0,
  //           id: '101cf2acfdff4d8da64af777299d6f9d',
  //           venue: 'The Ranger Station PDX',
  //           address: '4260 SE Hawthorne Blvd',
  //           image: 'https://citysparkstorage.blob.core.windows.net/portalimages/portalimages/39ea9bf3-4639-44c2-a9be-b637647ba7c7.medium.png',
  //           neighborhood: 'SE',
  //           popularity: 4
  //         }
  //       ]
  //     }
  //   };
  //   const nextState = reducer(undefined, action);

  //   expect(nextState).to.equal(fromJS({
  //     shows: [
  //       {
  //         bands: [ 'Bluegrass Tuesdays', 'w', ' Pete Kartsounes and Friends' ],
  //         time: '8:30 PM',
  //         date: 'Tue, 5/10/16',
  //         price: 0,
  //         id: '101cf2acfdff4d8da64af777299d6f9d',
  //         venue: 'The Ranger Station PDX',
  //         address: '4260 SE Hawthorne Blvd',
  //         image: 'https://citysparkstorage.blob.core.windows.net/portalimages/portalimages/39ea9bf3-4639-44c2-a9be-b637647ba7c7.medium.png',
  //         neighborhood: 'SE',
  //         popularity: 4
  //       },
  //       {
  //         bands: [ 'Cool Band', 'Another Cool Band'],
  //         time: '7:30 PM',
  //         date: 'Mon, 5/9/16',
  //         price: 5,
  //         id: '7d4ef2acfdff4d8da64af777299d7d4e',
  //         venue: 'My Favorite Bar',
  //         address: '434 NW Hermosa',
  //         image: 'https://citysparkstorage.blob.core.windows.net/portalimages/portalimages/39ea9bf3-4639-44c2-a9be-b637647ba7c7.medium.png',
  //         neighborhood: 'NW',
  //         popularity: 10
  //       }
  //     ],
  //     deck: [
  //       {
  //         bands: [ 'Bluegrass Tuesdays', 'w', ' Pete Kartsounes and Friends' ],
  //         time: '8:30 PM',
  //         date: 'Tue, 5/10/16',
  //         price: 0,
  //         id: '101cf2acfdff4d8da64af777299d6f9d',
  //         venue: 'The Ranger Station PDX',
  //         address: '4260 SE Hawthorne Blvd',
  //         image: 'https://citysparkstorage.blob.core.windows.net/portalimages/portalimages/39ea9bf3-4639-44c2-a9be-b637647ba7c7.medium.png',
  //         neighborhood: 'SE',
  //         popularity: 4
  //       }
  //     ],
  //     dateList: ['Mon, 5/9/16','Tue, 5/10/16','Wed, 5/11/16'],
  //     venueList: ['The Ranger Station PDX','My Favorite Bar'],
  //     neighborhoodList: ['SE','SW','NE','NW','N'],
  //     recommendedVenues: ['Lovecraft Bar','Alberta Street Pub'],
  //     filters: {},
  //     modalsOpen: [],
  //     isLoading: false,
  //     errors: [],
  //     activeShow: {},
  //     activeBand: {},
  //     videos: [],
  //     isLoadingVideos: false
  //   }));
  // });

  it('handles ADD_TO_DECK by adding show to deck', () => {
    const state = fromJS({
      deck: []
    });
    const action = {
      type: 'ADD_TO_DECK',
      show: {
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
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
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
      ]
    }));
  });

  it('handles ADD_TO_DECK by not adding a duplicate show to deck', () => {
    const state = fromJS({
      deck: [{
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
      }]
    });
    const action = {
      type: 'ADD_TO_DECK',
      show: {
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
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
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
      ]
    }));
  });

  it('handles REMOVE_FROM_DECK by removing show from deck', () => {
    const state = fromJS({
      deck: [{
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
      }]
    });
    const action = {
      type: 'REMOVE_FROM_DECK',
      show: {
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
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      deck: []
    }));
  });

  it('handles REMOVE_FROM_DECK by not deleting a non-existent show', () => {
    const state = fromJS({
      deck: []
    });
    const action = {
      type: 'REMOVE_FROM_DECK',
      show: {
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
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      deck: []
    }));
  });

  it('handles ADD_FILTER by adding a non-existent filter', () => {
    const state = fromJS({
      filters: {}
    });
    const action = {
      type: 'ADD_FILTER',
      filterKey: 'venue',
      filterValue: 'Test Venue'
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      filters: {
        venue: 'Test Venue'
      }
    }));
  });

  it('handles ADD_FILTER by updating an existing filter', () => {
    const state = fromJS({
      filters: {
        venue: 'Test Venue'
      }
    });
    const action = {
      type: 'ADD_FILTER',
      filterKey: 'venue',
      filterValue: 'New Venue'
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      filters: {
        venue: 'New Venue'
      }
    }));
  });

  it('handles REMOVE_FILTER by removing an existing filter', () => {
    const state = fromJS({
      filters: {
        venue: 'Test Venue'
      }
    });
    const action = {
      type: 'REMOVE_FILTER',
      filterKey: 'venue'
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      filters: {}
    }));
  });

  it('handles UPDATE_DATE_LIST by removing an invalid filter and setting dateList', () => {
    const state = fromJS({
      filters: {
        venue: 'Test Venue',
        date: 'Tue, Dec 31, 2016'
      }
    });
    const action = {
      type: 'UPDATE_DATE_LIST',
      dateList: ['Mon, 5/9/16','Tue, 5/10/16','Wed, 5/11/16']
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      filters: {
        venue: 'Test Venue'
      },
      dateList: ['Mon, 5/9/16','Tue, 5/10/16','Wed, 5/11/16']
    }));
  });

  it('handles UPDATE_DATE_LIST by preserving a valid filter and setting dateList', () => {
    const state = fromJS({
      filters: {
        venue: 'Test Venue',
        date: 'Mon, 5/9/16'
      }
    });
    const action = {
      type: 'UPDATE_DATE_LIST',
      dateList: ['Mon, 5/9/16','Tue, 5/10/16','Wed, 5/11/16']
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      filters: {
        venue: 'Test Venue',
        date: 'Mon, 5/9/16'
      },
      dateList: ['Mon, 5/9/16','Tue, 5/10/16','Wed, 5/11/16']
    }));
  });

  it('handles UPDATE_VENUE_LIST by removing an invalid filter and setting venueList', () => {
    const state = fromJS({
      filters: {
        venue: 'Test Venue',
        date: 'Tue, Dec 31, 2016'
      }
    });
    const action = {
      type: 'UPDATE_VENUE_LIST',
      venueList: ['The Ranger Station PDX','My Favorite Bar']
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      filters: {
        date: 'Tue, Dec 31, 2016'
      },
      venueList: ['The Ranger Station PDX','My Favorite Bar']
    }));
  });

  it('handles UPDATE_VENUE_LIST by preserving a valid filter and setting venueList', () => {
    const state = fromJS({
      filters: {
        venue: 'The Ranger Station PDX',
        date: 'Tue, Dec 31, 2016'
      }
    });
    const action = {
      type: 'UPDATE_VENUE_LIST',
      venueList: ['The Ranger Station PDX','My Favorite Bar']
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      filters: {
        venue: 'The Ranger Station PDX',
        date: 'Tue, Dec 31, 2016'
      },
      venueList: ['The Ranger Station PDX','My Favorite Bar']
    }));
  });

  it('handles REQUEST_SHOWS by setting a loading bar and calling fetch', () => {

    const state = Map();
    const action = {
      type: 'REQUEST_SHOWS'
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      isLoading: true
    }));

  });

  it('handles RECEIVE_SHOWS_SUCCESS by setting shows, removing loading bar, setting active show, and resetting errors', () => {

    const state = Map();
    const action = {
      type: 'RECEIVE_SHOWS_SUCCESS',
      shows: [{
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
      }]
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      shows: [{
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
      }],
      isLoading: false,
      errors: [],
      activeShow: {
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
    }));

  });

  it('handles RECEIVE_SHOWS_ERROR by removing shows, removing loading bar, and setting error', () => {

    const state = Map();
    const action = {
      type: 'RECEIVE_SHOWS_ERROR',
      error: 'Could not get shows'
    };
    const nextState = reducer(state, action);

    expect(nextState).to.equal(fromJS({
      shows: [],
      errors: ['Could not get shows'],
      isLoading: false,
      activeShow: {}
    }));

  });

  it('handles RETRIEVE_STORAGE by setting a valid state from stored value', () => {
    const state = Map();
    const storageValue = JSON.stringify({
      filters: {
        venue: 'The Ranger Station PDX',
        date: 'Tue, Dec 31, 2016'
      }
    });
    localStorage.setItem('state', storageValue);
    const action = {
      type: 'RETRIEVE_STORAGE',
      state: Map()
    };

    const nextState = reducer(state, action);
    expect(nextState).to.equal(fromJS({
      filters: {
        venue: 'The Ranger Station PDX',
        date: 'Tue, Dec 31, 2016'
      }
    }));
  });

  it('handles RETRIEVE_STORAGE by storing a valid state', () => {
    const state = fromJS({
      filters: {
        venue: 'The Ranger Station PDX',
        date: 'Tue, Dec 31, 2016'
      }
    });
    const storageValue = JSON.stringify({
      filters: {
        venue: 'The Ranger Station PDX',
        date: 'Tue, Dec 31, 2016'
      }
    });
    const action = {
      type: 'SET_STORAGE',
      state
    };

    const nextState = reducer(state, action);
    expect(localStorage.getItem('state')).to.equal(storageValue);
  });

});
