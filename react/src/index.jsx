import React from 'react';
import ReactDOM from 'react-dom';
import thunkMiddleware from 'redux-thunk'
import createLogger from 'redux-logger'
import {Router, Route, hashHistory} from 'react-router';
import {createStore, applyMiddleware} from 'redux';
import {Provider} from 'react-redux';
import reducer from './reducer';
import App from './components/App';
import {HomePageContainer} from './components/HomePage';
import {fromJS} from 'immutable';
import * as actionCreators from './action_creators';
require('./style.css');

const loggerMiddleware = createLogger();
var store = createStore(reducer, applyMiddleware(thunkMiddleware, loggerMiddleware));

store.dispatch({type: 'SET_STATE'});
store.dispatch(actionCreators.retrieveStorage())
store.dispatch(actionCreators.fetchShows());

const routes = <Route component={App}>
  <Route path="/" component={HomePageContainer} />
</Route>;

ReactDOM.render(
  <Provider store={store}>
    <Router history={hashHistory}>{routes}</Router>
  </Provider>,
  document.getElementById('app')
);

