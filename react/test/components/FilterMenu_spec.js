import React from 'react';
import ReactDOM from 'react-dom';
import {
  renderIntoDocument,
  findRenderedDOMComponentWithClass,
  scryRenderedDOMComponentsWithClass,
  scryRenderedDOMComponentsWithTag,
  Simulate
} from 'react-addons-test-utils';
import {fromJS, List, Map} from 'immutable';
import {FilterMenu} from '../../src/components/FilterMenu';
import {expect} from 'chai'

describe('Filter Menu', () => {

  it('renders options for active filters', () => {
    const shows = fromJS([
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
    ]);
    const filters = fromJS({
      date: 'Tue, 5/10/16',
      venue: 'The Ranger Station PDX',
      neighborhood: 'NW'
    });
    const venueList = fromJS(
      ['The Ranger Station PDX','My Favorite Bar']
    );
    const neighborhoodList = fromJS(
      ['NW','SE']
    );
    const dateList = fromJS(
      ['Mon, 5/9/16','Tue, 5/10/16']
    );
    const component = renderIntoDocument(
      <FilterMenu shows={shows}
                  filters={filters}
                  dateList={dateList}
                  venueList={venueList}
                  neighborhoodList={neighborhoodList} />
    );
    const dateButtons = scryRenderedDOMComponentsWithClass(component, 'date_button');
    expect(dateButtons.length).to.equal(2);
    const venueOptions = scryRenderedDOMComponentsWithClass(component, 'list_item');
    expect(venueOptions.length).to.equal(2);
  });

  it('renders no options if filters are inactive', () => {
    const shows = fromJS([
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
    ]);
    const filters = fromJS({});
    const venueList = fromJS(
      ['The Ranger Station PDX','My Favorite Bar']
    );
    const neighborhoodList = fromJS(
      ['NW','SE']
    );
    const dateList = fromJS(
      ['Mon, 5/9/16','Tue, 5/10/16']
    );
    const component = renderIntoDocument(
      <FilterMenu shows={shows}
                  filters={filters}
                  dateList={dateList}
                  venueList={venueList}
                  neighborhoodList={neighborhoodList} />
    );
    const dateButtons = scryRenderedDOMComponentsWithClass(component, 'date_button');
    expect(dateButtons.length).to.equal(0);
    const venueOptions = scryRenderedDOMComponentsWithClass(component, 'list_item');
    expect(venueOptions.length).to.equal(0);
  });

});
