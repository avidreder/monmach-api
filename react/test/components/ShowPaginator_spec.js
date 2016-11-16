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
import ShowPaginator from '../../src/components/ShowPaginator';
import {expect} from 'chai';
import _ from 'lodash';

describe('Show Paginator', () => {

  it('renders all components in first page', () => {
    const showsPerPage = 3;
    const shows = _.reduce(_.range(7), function(result, value, key) {
      var obj = {
        bands: [ 'Band ' + value],
        id: value
      };
      result[value] = obj;
      return result;
    }, []);

    const deck = [];
    const component = renderIntoDocument(
      <ShowPaginator shows={shows}
                     deck={deck} />
    );
    const paginationWrapper = scryRenderedDOMComponentsWithClass(component, 'paginator');
    const bandNames = scryRenderedDOMComponentsWithClass(component, 'band_name');
    expect(bandNames.length).to.equal(3);
    expect(paginationWrapper.length).to.equal(1);
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 0")).to.be.true;
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 3")).to.be.false;
  });

  it('pages up when up button is pressed', () => {
    const showsPerPage = 3;
    const shows = fromJS(_.reduce(_.range(7), function(result, value, key) {
      var obj = {
        bands: [ 'Band ' + value],
        id: value
      };
      result[value] = obj;
      return result;
    }, []));
    const deck = [];
    let component = renderIntoDocument(
      <ShowPaginator shows={shows}
                     deck={deck} />
    );
    const paginationWrapper = scryRenderedDOMComponentsWithClass(component, 'paginator');
    expect(paginationWrapper.length).to.equal(1);

    // Page 1 has 3 shows
    let bandNames = scryRenderedDOMComponentsWithClass(component, 'band_name');
    expect(bandNames.length).to.equal(3);
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 0")).to.be.true;
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 3")).to.be.false;

    const upButton = findRenderedDOMComponentWithClass(component, 'page_up');
    Simulate.click(upButton);

    // Page 2 has 3 different shows
    bandNames = scryRenderedDOMComponentsWithClass(component, 'band_name');
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 0")).to.be.false;
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 3")).to.be.true;
    expect(bandNames.length).to.equal(3);

    Simulate.click(upButton);

    // Page 3 has 1 show
    bandNames = scryRenderedDOMComponentsWithClass(component, 'band_name');
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 3")).to.be.false;
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 6")).to.be.true;
    expect(bandNames.length).to.equal(1);

    // Pressing up on last page shouldn't do anything
    Simulate.click(upButton);

    bandNames = scryRenderedDOMComponentsWithClass(component, 'band_name');
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 6")).to.be.true;
    expect(bandNames.length).to.equal(1);
  });

  it('pages down when down button is pressed', () => {
    const showsPerPage = 3;
    const shows = fromJS(_.reduce(_.range(7), function(result, value, key) {
      var obj = {
        bands: [ 'Band ' + value],
        id: value
      };
      result[value] = obj;
      return result;
    }, []));
    const deck = [];
    let component = renderIntoDocument(
      <ShowPaginator shows={shows}
                     deck={deck} />
    );
    const paginationWrapper = scryRenderedDOMComponentsWithClass(component, 'paginator');
    expect(paginationWrapper.length).to.equal(1);

    // Go to last page
    const upButton = findRenderedDOMComponentWithClass(component, 'page_up');
    Simulate.click(upButton);
    Simulate.click(upButton);

    // Page 3 has 1 show
    let bandNames = scryRenderedDOMComponentsWithClass(component, 'band_name');
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 3")).to.be.false;
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 6")).to.be.true;
    expect(bandNames.length).to.equal(1);

    const downButton = findRenderedDOMComponentWithClass(component, 'page_down');
    Simulate.click(downButton);

    // Page 2 has 3 different shows
    bandNames = scryRenderedDOMComponentsWithClass(component, 'band_name');
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 0")).to.be.false;
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 3")).to.be.true;
    expect(bandNames.length).to.equal(3);

    Simulate.click(downButton);

    // Page 1 has 3 shows
    bandNames = scryRenderedDOMComponentsWithClass(component, 'band_name');
    expect(bandNames.length).to.equal(3);
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 0")).to.be.true;
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 3")).to.be.false;

    // Pressing up on last page shouldn't do anything
    Simulate.click(downButton);
    bandNames = scryRenderedDOMComponentsWithClass(component, 'band_name');
    expect(bandNames.length).to.equal(3);
    expect(_.contains(_.pluck(bandNames,'textContent'),"Band 0")).to.be.true;
  });

});
