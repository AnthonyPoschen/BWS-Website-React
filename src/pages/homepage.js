import '../App.css'

import React from 'react'
import {Link} from 'react-router'

let HomePage = React.createClass({
    render() {
        return <div className="PageHeader">
            <h2>Home Page</h2>
            <Link to="/Blog">Blog Page</Link>
        </div>
    }
})

export default HomePage