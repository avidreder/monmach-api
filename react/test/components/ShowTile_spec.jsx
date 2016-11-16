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
import ShowTile from '../../src/components/ShowTile';
import {expect} from 'chai'

describe('ShowTile', () => {

  it('renders all components', () => {
    const isPinned = false;
    const show = {
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
    };
    const component = renderIntoDocument(
      <ShowTile show={show}
                isPinned={isPinned}
                feedType="shows" />
    );
    const image = findRenderedDOMComponentWithClass(component, 'show_image');
    const bands = scryRenderedDOMComponentsWithClass(component, 'band_name');
    const venue = findRenderedDOMComponentWithClass(component, 'show_venue');
    const date = findRenderedDOMComponentWithClass(component, 'show_date');
    const time = findRenderedDOMComponentWithClass(component, 'show_time');
    const price = findRenderedDOMComponentWithClass(component, 'show_price');
    const pinButton = scryRenderedDOMComponentsWithClass(component, 'pin_button');
    const removeButton = scryRenderedDOMComponentsWithClass(component, 'remove_button');

    expect(image.src).to.equal(show.image);
    for (var i = 0; i < show.bands.length; i++) {
      expect(bands[0].textContent).to.equal(show.bands[0]);
    }
    expect(venue.textContent).to.equal(show.venue);
    expect(date.textContent).to.equal(show.date);
    expect(time.textContent).to.equal(show.time);
    expect(price.textContent).to.equal(show.price.toString());
    expect(pinButton.length).to.equal(1);
    expect(removeButton.length).to.equal(0);
  });

  it('invokes callback when pin button is clicked', () => {
    let pinnedShow;
    const pinShow = (show) => pinnedShow = show;
    const isPinned = false;
    const show = {
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
    };

    const component = renderIntoDocument(
      <ShowTile show={show}
                pinShow={pinShow}
                isPinned={isPinned} />
    );

    const button = findRenderedDOMComponentWithClass(component, 'pin_button');
    Simulate.click(button);

    expect(pinnedShow).to.equal(show);
  });

  it('invokes callback when pin button is clicked', () => {
    const isPinned = true;
    const show = {
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
    };
    let pinnedShow = show;
    const removeShow = (show) => pinnedShow = {};
    const component = renderIntoDocument(
      <ShowTile show={show}
                removeShow={removeShow}
                isPinned={isPinned} />
    );

    const removeButton = findRenderedDOMComponentWithClass(component, 'remove_button');
    Simulate.click(removeButton);

    expect(pinnedShow).to.be.empty;
  });

  it('hides pin button and shows remove button when show is pinned', () => {
    const isPinned = true;
    const show = {
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
    };

    const component = renderIntoDocument(
      <ShowTile show={show}
                isPinned={isPinned} />
    );

    const pinButton = scryRenderedDOMComponentsWithClass(component, 'pin_button');
    const removeButton = scryRenderedDOMComponentsWithClass(component, 'remove_button');

    expect(removeButton.length).to.equal(1);
    expect(pinButton.length).to.equal(0);
  });

  it('renders as a pure component', () => {
    const isPinned = false;
    const show = {
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
    };
    const container = document.createElement('div');
    let component = ReactDOM.render(
      <ShowTile show={show}
                isPinned={isPinned} />,
      container
    );

    let showTime = findRenderedDOMComponentWithClass(component, 'show_time');
    expect(showTime.textContent).to.equal('8:30 PM');

    show.time = '9:30 PM';

    component = ReactDOM.render(
      <ShowTile show={show}
                isPinned={isPinned} />,
      container
    );

    showTime = findRenderedDOMComponentWithClass(component, 'show_time');
    expect(showTime.textContent).to.equal('8:30 PM');
  });

  it('does update DOM when prop changes', () => {
    let isPinned = false;
    const show = fromJS({
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
    });
    const container = document.createElement('div');
    let component = ReactDOM.render(
      <ShowTile show={show}
                isPinned={isPinned} />,
      container
    );

    let showTime = findRenderedDOMComponentWithClass(component, 'show_time');
    expect(showTime.textContent).to.equal('8:30 PM');

    const newShow = show.set('time','9:30 PM');

    component = ReactDOM.render(
      <ShowTile show={newShow}
                isPinned={isPinned} />,
      container
    );

    showTime = findRenderedDOMComponentWithClass(component, 'show_time');
    expect(showTime.textContent).to.equal('9:30 PM');
  });

});
