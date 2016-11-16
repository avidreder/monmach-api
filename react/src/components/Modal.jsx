import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {connect} from 'react-redux';
import * as actionCreators from '../action_creators';
import classnames from 'classnames';
import _ from 'lodash';

export const ModalWindow = React.createClass({
  mixins: [PureRenderMixin],
  render: function() {
    if(this.props.modalsOpen.contains(this.props.modalType)) {
      return <div style={{textAlign:"center"}} className={classnames('row', 'well')}>
        <div style={{display: 'block'}} className={classnames('modal_window','modal','fade','in','modal-open','col-md-12','col-sm-12')}>
          <div className="modal-dialog">
            <div className="modal-content">
              <div className="modalBody">
                <button type="button" className="modal_close_button" onClick={() => this.props.closeModal(this.props.modalType)}>&times;</button>
                {this.props.internalComponent}
              </div>
            </div>
          </div>
        </div>
      </div>;
    } else {
      return <div style={{display: 'none'}} className={classnames('modal_window')}>
      </div>;
    }
  }
});

function mapStateToProps(state) {
  return {
    modalsOpen: state.get('modalsOpen')
  };
}

function mapDispatchToProps(dispatch) {
  return {
    closeModal: function(type) {
      dispatch(actionCreators.closeModal(type));
    }
  };
}

export const ModalWindowContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ModalWindow);
