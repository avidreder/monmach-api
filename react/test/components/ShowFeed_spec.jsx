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
import {ShowFeed} from '../../src/components/ShowFeed';
import {expect} from 'chai'

describe('ShowFeed', () => {

  it('renders all components', () => {
    const shows = [
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
    ];
    const deck = [];
    const component = renderIntoDocument(
      <ShowFeed showData={shows}
                deck={deck} />
    );
    const showTiles = scryRenderedDOMComponentsWithClass(component, 'show_tile');

    expect(showTiles.length).to.equal(2);
  });

  it('renders as a pure component', () => {
    const shows = [
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
    ];
    const container = document.createElement('div');
    let component = ReactDOM.render(
      <ShowFeed showData={shows} />,
      container
    );

    let showTime = findRenderedDOMComponentWithClass(component, 'show_time');
    expect(showTime.textContent).to.equal('8:30 PM');

    shows[0].time="7:30";

    component = ReactDOM.render(
      <ShowFeed showData={shows} />,
      container
    );

    showTime = findRenderedDOMComponentWithClass(component, 'show_time');
    expect(showTime.textContent).to.equal('8:30 PM');
  });

  it('does update DOM when prop changes', () => {
    let shows = [
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
    ];
    const container = document.createElement('div');
    let component = ReactDOM.render(
      <ShowFeed showData={shows} />,
      container
    );

    let showTime = findRenderedDOMComponentWithClass(component, 'show_time');
    expect(showTime.textContent).to.equal('8:30 PM');

    shows = [
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
    ];

    component = ReactDOM.render(
      <ShowFeed showData={shows} />,
      container
    );

    showTime = findRenderedDOMComponentWithClass(component, 'show_time');
    expect(showTime.textContent).to.equal('7:30 PM');
  });

});
