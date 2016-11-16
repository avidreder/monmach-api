import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {connect} from 'react-redux';
import _ from 'lodash';
import {List, Map} from 'immutable';
import {VideoPlayerContainer} from './VideoPlayer';
import {ModalWindowContainer} from './Modal';
import * as actionCreators from '../action_creators';
import classnames from 'classnames';
import * as filters from '../filters';

export const ShowVideos = React.createClass({
  mixins: [PureRenderMixin],
  getBands: function() {
    console.log(this.props.activeShow);
    return Map.isMap(this.props.activeShow) ? this.props.activeShow.get('bands').toArray() || [] : this.props.activeShow.bands|| [];
  },
  getValue: function(key) {
    return Map.isMap(this.props.activeShow) ? this.props.activeShow.get(key) : this.props.activeShow[key];
  },
  render: function() {
    return <div className={classnames('container-fluid')}>
      <div className={classnames('row')}>
        <div className={classnames('well', 'show_tile')}>
          <img className="show_image" src={this.getValue('image')} />
          <div className="show_venue">{this.getValue('venue')}</div>
          <div className="show_date">{this.getValue('date')}</div>
          <div className="show_time">{this.getValue('time')}</div>
          <div className="show_price">{this.getValue('price')}</div>
          <div className="show_neighborhood">{this.getValue('neighborhood')}</div>
          {this.props.isPinned ?
            <button className={classnames('btn', 'btn-default', 'btn-xs', 'remove_button')} onClick={() => this.props.removeShow(this.props.activeShow)}><span className={classnames('glyphicon', 'glyphicon-remove')} aria-hidden="true"></span></button> :
            <button className={classnames('btn', 'btn-default', 'btn-xs', 'pin_button')} onClick={() => this.props.pinShow(this.props.activeShow)}><span className={classnames('glyphicon', 'glyphicon-pushpin')} aria-hidden="true"></span></button>
          }
        </div>
      </div>
      <div className={classnames('row', 'band_list')}>
        {this.getBands().map(band =>
          <div key={"preview_" + band} className={classnames('thumbnail', 'band_tile')}>
            <button className={classnames('btn', 'btn-default', 'btn-xs', 'remove_button')} onClick={() => this.props.openModal(band)}>{band}</button>
          </div>
        )}
      </div>
      <ModalWindowContainer internalComponent={<VideoPlayerContainer />}
                            modalType="video" />
    </div>;
  }
});

function mapStateToProps(state) {
  return {
    modalsOpen: state.get('modalsOpen'),
    activeShow: state.get('activeShow'),
    activeBand: state.get('activeBand')
  };
}

function mapDispatchToProps(dispatch) {
  return {
    openModal: function(band) {
      dispatch(actionCreators.setActiveBand(band));
      dispatch(actionCreators.fetchVideos(band));
      dispatch(actionCreators.openModal('video'));
    }
  }
}

export const ShowVideosContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ShowVideos);
