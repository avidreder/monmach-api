import {List, Map, fromJS} from 'immutable';
import {expect} from 'chai';

import * as filters from '../src/filters';

describe('filters', () => {

  it('returns objects that match multiple criteria', () => {
    const shows = [
      {
        date: 'Tue, 5/10/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX',
        neighborhood: 'NW'
      },
      {
        date: 'Mon, 5/9/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX',
        neighborhood: 'SE'
      }
    ];
    const activeFilters = {
      date: 'Tue, 5/10/16',
      venue: 'The Ranger Station PDX',
      neighborhood: 'NW'
    };

    const filteredShows = filters.filterShows(shows, activeFilters);
    expect(filteredShows).to.deep.equal([
        {
          date: 'Tue, 5/10/16',
          id: '101cf2acfdff4d8da64af777299d6f9d',
          venue: 'The Ranger Station PDX',
          neighborhood: 'NW'
        }
    ]);
  })

  it('returns immutable objects that match multiple criteria', () => {
    const shows = fromJS([
      {
        date: 'Tue, 5/10/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      },
      {
        date: 'Mon, 5/9/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      }
    ]);
    const activeFilters = fromJS({
      date: 'Tue, 5/10/16',
      venue: 'The Ranger Station PDX'
    });

    const filteredShows = filters.filterShows(shows, activeFilters);
    expect(filteredShows).to.deep.equal(fromJS([
        {
          date: 'Tue, 5/10/16',
          id: '101cf2acfdff4d8da64af777299d6f9d',
          venue: 'The Ranger Station PDX'
        }
    ]));
  })

  it('returns all shows when there are no filters', () => {
    const shows = [
      {
        date: 'Tue, 5/10/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      },
      {
        date: 'Mon, 5/9/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      }
    ];
    const activeFilters = {};

    const filteredShows = filters.filterShows(shows, activeFilters);
    expect(filteredShows).to.deep.equal(shows);
  })

  it('returns all shows as immutable objects if there are no filters', () => {
    const shows = fromJS([
      {
        date: 'Tue, 5/10/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      },
      {
        date: 'Mon, 5/9/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      }
    ]);
    const activeFilters = Map();

    const filteredShows = filters.filterShows(shows, activeFilters);
    expect(filteredShows).to.deep.equal(shows);
  })

  it('does not returns those that do not match multiple criteria', () => {
    const shows = [
      {
        date: 'Tue, 5/10/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      },
      {
        date: 'Mon, 5/9/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      }
    ];
    const activeFilters = {
      date: 'Tue, 5/10/16',
      venue: 'My Favorite Bar'
    };

    const filteredShows = filters.filterShows(shows, activeFilters);
    expect(filteredShows).to.deep.equal([]);
  })

  it('does not returns immutable objects that do not match multiple criteria', () => {
    const shows = fromJS([
      {
        date: 'Tue, 5/10/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      },
      {
        date: 'Mon, 5/9/16',
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      }
    ]);
    const activeFilters = fromJS({
      date: 'Tue, 5/10/16',
      venue: 'My Favorite Bar'
    });

    const filteredShows = filters.filterShows(shows, activeFilters);
    expect(filteredShows).to.deep.equal(List());
  })

  it('returns objects that are at a recommended venue', () => {
    const shows = [
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'My Favorite Bar'
      },
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      }
    ];
    const recommendedVenues = ['One Bar','Another Bar','My Favorite Bar'];

    const filteredShows = filters.filterRecommendedShows(shows, recommendedVenues);
    expect(filteredShows).to.deep.equal([{
      id: '101cf2acfdff4d8da64af777299d6f9d',
      venue: 'My Favorite Bar'
    }]);
  })

  it('returns immutable objects that are at a recommended venue', () => {
    const shows = fromJS([
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'My Favorite Bar'
      },
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      }
    ]);
    const recommendedVenues = fromJS(['One Bar','Another Bar','My Favorite Bar']);

    const filteredShows = filters.filterRecommendedShows(shows, recommendedVenues);
    expect(filteredShows).to.deep.equal(fromJS([{
      id: '101cf2acfdff4d8da64af777299d6f9d',
      venue: 'My Favorite Bar'
    }]));
  })

  it('does not return objects that are not at a recommended venue', () => {
    const shows = [
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'My Favorite Bar'
      },
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      }
    ];
    const recommendedVenues = ['One Bar','Another Bar'];

    const filteredShows = filters.filterRecommendedShows(shows, recommendedVenues);
    expect(filteredShows).to.deep.equal([]);
  })

  it('does not return immutable objects that are not at a recommended venue', () => {
    const shows = fromJS([
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'My Favorite Bar'
      },
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      }
    ]);
    const recommendedVenues = fromJS(['One Bar','Another Bar']);

    const filteredShows = filters.filterRecommendedShows(shows, recommendedVenues);
    expect(filteredShows).to.deep.equal(List());
  })

  it('does not return any objects if there are no recommended venues', () => {
    const shows = [
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'My Favorite Bar'
      },
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      }
    ];
    const recommendedVenues = [];

    const filteredShows = filters.filterRecommendedShows(shows, recommendedVenues);
    expect(filteredShows).to.deep.equal([]);
  })

  it('does not return any immutable objects if there are no recommended venues', () => {
    const shows = fromJS([
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'My Favorite Bar'
      },
      {
        id: '101cf2acfdff4d8da64af777299d6f9d',
        venue: 'The Ranger Station PDX'
      }
    ]);
    const recommendedVenues = List();

    const filteredShows = filters.filterRecommendedShows(shows, recommendedVenues);
    expect(filteredShows).to.deep.equal(List());
  })
});
