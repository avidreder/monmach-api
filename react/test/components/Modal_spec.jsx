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
import {ModalWindow} from '../../src/components/Modal';
import {expect} from 'chai'

describe('Modal', () => {

  it('renders all components', () => {
    const modalsOpen = List(["deck"]);
    const component = renderIntoDocument(
      <ModalWindow internalComponent={<p className='internal_component'></p>}
                   modalType="deck"
                   modalsOpen={modalsOpen} />
    );
    const content = scryRenderedDOMComponentsWithClass(component, 'internal_component');
    expect(content.length).to.equal(1);
  });

});
