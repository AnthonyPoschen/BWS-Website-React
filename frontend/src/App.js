import './App.css'

import React from 'react'

import {Router , Route , Link, browserHistory} from 'react-router'
import { Provider } from 'react-redux'
import { createStore , applyMiddleware } from 'redux'
import thunk from 'redux-thunk'

import reducer from './reducers/index.js'

import PageTemplate from './pagetemplate.js'
import BlogPage from './pages/blogpage.js'
import HomePage from './pages/homepage.js'
import AdminPage from './pages/adminpage.js'
import AboutPage from './pages/aboutpage.js'
import BlogPost from './pages/blogpost.js'
import LoginPage from './pages/loginpage.js'

//import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
//import getMuiTheme from 'material-ui/styles/getMuiTheme';
//<!-- Latest compiled and minified CSS -->

const store = createStore(reducer,applyMiddleware(thunk))

let App = React.createClass({
  render() {
    return (
    <Provider store={store}>
      <Router history={browserHistory}>
        <Route component={PageTemplate}>
          <Route path="/" component={HomePage} />
          <Route path="/Blog/Post/:title" component={BlogPost} />
          <Route path="/Blog" component={BlogPage} />
          <Route path="/Admin" component={AdminPage} />
          <Route path="/About" component={AboutPage} />
          <Route path="/login" component={LoginPage} />
        </Route>
      </Router> 
    </Provider>
    );
  }
})

export default App
