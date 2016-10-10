import '../App.css'

import React from 'react'
import {Link} from 'react-router'

let BlogPage = React.createClass({
    render() {
     return <div className="PageHeader">
        <h2>Blog Page</h2>
        <Link to="/">HomePage</Link>
     </div>   
    }
})

export default BlogPage