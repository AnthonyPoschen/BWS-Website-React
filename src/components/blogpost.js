import '../App.css'

import React from 'react'
import {Link} from 'react-router'

let BlogPost = React.createClass({
    render() {
        return (
            // Blog post is setup here
            <div>
                <p>Testing</p>
                <h3>{this.props.title}</h3>
                <p>{this.props.content}</p>
            </div>
        )
    }
})

export default BlogPost