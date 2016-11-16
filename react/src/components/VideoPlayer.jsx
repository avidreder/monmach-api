import React from 'react';
import PureRenderMixin from 'react-addons-pure-render-mixin';
import {connect} from 'react-redux';
import _ from 'lodash';
import {List, Map} from 'immutable';
import * as actionCreators from '../action_creators';
import classnames from 'classnames';
import * as filters from '../filters';

export const VideoPlayer = React.createClass({
  mixins: [PureRenderMixin],
  getBand: function() {
    return this.props.activeBand || {};
  },
  getValue: function(key) {
    return Map.isMap(this.props.activeShow) ? this.props.activeShow.get(key) : this.props.activeShow[key];
  },
  getVideos: function() {
    return List.isList(this.props.videos) ? this.props.videos.toArray() || [] : this.props.videos || [];
  },
  render: function() {
    return <div className={classnames('container-fluid')}>
      <div className={classnames('row')}>
        <div className={classnames('well', 'band_tile')}>
          <div className="band_name">{this.getBand()}</div>
        </div>
      </div>
      {this.props.isLoadingVideos ?
        <div style={{textAlign:"center"}} className={classnames('row', 'well')}>
          <div className={classnames('col-md-12', 'col-sm-12')}>
            <h3>Loading Videos</h3>
            <h3><span className={classnames('glyphicon', 'glyphicon-refresh', 'glyphicon-spinning')}></span></h3>
          </div>
        </div> :
        <div className={classnames('row')}>
          {this.getVideos().map(video =>
            <div key={video} className={classnames('thumbnail', 'col-md-6', 'col-sm-6', 'video_tile')}>
              <div className={classnames('caption')}>
                <div>{video.get('title')}</div>
                <iframe width="100%" src={"http://www.youtube.com/embed/" + video.get('id')}  frameBorder="0" allowFullScreen></iframe>
              </div>
            </div>
          )}
        </div>
      }
    </div>;
  }
});

function mapStateToProps(state) {
  return {
    activeBand: state.get('activeBand'),
    videos: state.get('videos'),
    isLoadingVideos: state.get('isLoadingVideos')
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchVideos: function(band) {
      dispatch(actionCreators.fetchVideos(band));
    }
  }
}

export const VideoPlayerContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(VideoPlayer);
