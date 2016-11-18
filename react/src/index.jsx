import React from 'react';
import ReactDOM from 'react-dom';
import thunkMiddleware from 'redux-thunk'
import createLogger from 'redux-logger'
import {Router, Route, hashHistory} from 'react-router';
import {createStore, applyMiddleware} from 'redux';
import {Provider} from 'react-redux';
import reducer from './reducer';
import App from './components/App';
import HomePage from './components/HomePage';
import ComponentPage from './components/ComponentPage';
import {GenrePageContainer} from './components/GenrePage';
import injectTapEventPlugin from 'react-tap-event-plugin';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';

injectTapEventPlugin();
// import {fromJS} from 'immutable';

const loggerMiddleware = createLogger();
var store = createStore(reducer, applyMiddleware(thunkMiddleware, loggerMiddleware));

store.dispatch({type: 'SET_STATE'});
/* if (process.env.liveData) {
 *     store.dispatch(actionCreators.retrieveStorage())
 *     store.dispatch(actionCreators.fetchShows());
 * }*/

const routes = <Route component={App}>
    <Route path="/" component={HomePage} />
    <Route path="/components" component={ComponentPage} />
    <Route path="/genre" component={GenrePageContainer} />
</Route>;

ReactDOM.render(
    <Provider store={store}>
	<Router history={hashHistory}>{routes}</Router>
    </Provider>,
    document.getElementById('app')
);
