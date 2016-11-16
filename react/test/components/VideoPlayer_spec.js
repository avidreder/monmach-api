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
import {VideoPlayer} from '../../src/components/VideoPlayer';
import {expect} from 'chai'

describe('Video Player', () => {

  it('renders all components', () => {
    const activeBand = "The Kills";
    const videos = fromJS([
      {
        id: 'MgaHy7DYs3g',
        title: 'The Kills - The Last Goodbye',
        href: '//www.youtube.com/embed/MgaHy7DYs3g',
        thumbnail: 'https://i.ytimg.com/vi/MgaHy7DYs3g/mqdefault.jpg'
      },
      {
        id: '_R83aq4Aqpk',
        title: 'The Kills - No Wow (Live in Sydney) | Moshcam',
        href: '//www.youtube.com/embed/_R83aq4Aqpk',
        thumbnail: 'https://i.ytimg.com/vi/_R83aq4Aqpk/mqdefault.jpg'
      }
    ]);
    const isLoadingVideos = false;

    const component = renderIntoDocument(
      <VideoPlayer activeBand={activeBand}
                   videos={videos}
                   isLoadingVideos={isLoadingVideos} />
    );
    const bandTile = scryRenderedDOMComponentsWithClass(component, 'band_name');
    expect(bandTile.length).to.equal(1);
    const videoTiles = scryRenderedDOMComponentsWithClass(component, 'video_tile');
    expect(videoTiles.length).to.equal(2);
  });

});
