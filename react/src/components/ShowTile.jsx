import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {connect} from 'react-redux';
import {Map} from 'immutable';
import classnames from 'classnames';

export default React.createClass({
  mixins: [PureRenderMixin],
  getValue: function(key) {
    return Map.isMap(this.props.show) ? this.props.show.get(key) : this.props.show[key];
  },
  getBands: function() {
    return Map.isMap(this.props.show) ? this.props.show.get('bands') || [] : this.props.show['bands'] || [];
  },
  render: function() {
    return <div className={classnames('thumbnail', 'show_tile')}>
      <img className="show_image" src={this.getValue('image')} />
      <div className="caption bands">
        {this.getBands().map(band =>
          <h4 key={this.props.feedType + "_" + band} className="band_name">
            {band}
          </h4>
        )}
      </div>
      <div className={classnames('caption')}>
        <div className="show_venue">{this.getValue('venue')}</div>
        <div className="show_date">{this.getValue('date')}</div>
        <div className="show_time">{this.getValue('time')}</div>
        <div className="show_price">{this.getValue('price')}</div>
        <div className="show_neighborhood">{this.getValue('neighborhood')}</div>
        {this.props.isPinned ?
          <button className={classnames('btn', 'btn-default', 'btn-xs', 'remove_button')} onClick={() => this.props.removeShow(this.props.show)}><span className={classnames('glyphicon', 'glyphicon-remove')} aria-hidden="true"></span></button> :
          <button className={classnames('btn', 'btn-default', 'btn-xs', 'pin_button')} onClick={() => this.props.pinShow(this.props.show)}><span className={classnames('glyphicon', 'glyphicon-pushpin')} aria-hidden="true"></span></button>
        }
        <button className={classnames('btn', 'btn-default', 'btn-large', 'open_modal_button')} onClick={() => this.props.openModal(this.props.show)}>Preview Show</button>
      </div>
    </div>;
  }
});
