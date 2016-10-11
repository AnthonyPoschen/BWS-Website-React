import './App.css'

import React from 'react'

import {Router , Route , Link, browserHistory} from 'react-router'

import PageTemplate from './pagetemplate.js'
import BlogPage from './pages/blogpage.js'
import HomePage from './pages/homepage.js'

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';

const muiTheme = getMuiTheme({
});
let App = React.createClass({
  render() {
    /*
    return <div className="App">
      <div className="App-heading App-flex">
        <h2>Welcome to <span className="App-react">React</span></h2>
      </div>
      <div className="App-instructions App-flex">
        <img className="App-logo" src={require('./react.svg')}/>
        <p>Edit <code>src/App.js</code> and save to hot reload your changes.</p>
      </div>
    </div>
    */
    
    return (
      /*
    <MuiThemeProvider muiTheme={muiTheme}>
      <PageTemplate></PageTemplate>
    </MuiThemeProvider>
    */
    <MuiThemeProvider muiTheme={muiTheme}>
      <Router history={browserHistory}>
        <Route component={PageTemplate}>
          <Route path="/" component={HomePage} />
          <Route path="/Blog" component={BlogPage} />
        </Route>
      </Router>
    </MuiThemeProvider>
    
    );
  }
})

export default App
