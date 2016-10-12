import '../App.css'

import React from 'react'
import {Link} from 'react-router'
import Blog from '../components/blog.js'

let BlogPage = React.createClass({
    render() {
     return <div className="PageHeader">
        <h2>Blog Page</h2>
        <Link to="/">HomePage</Link>
        <Blog />
     </div>   
    }
})

export default BlogPage