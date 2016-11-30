import './index.css'

import React from 'react'
import {render} from 'react-dom'

import injectTapEventPlugin from 'react-tap-event-plugin';
 
// Needed for onTouchTap 
// http://stackoverflow.com/a/34015469/988941 
injectTapEventPlugin();
 

import App from './App'

render(<App/>, document.querySelector('#app'))
