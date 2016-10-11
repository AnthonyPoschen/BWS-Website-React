import React from 'react'
import {Link} from 'react-router'

// contains a refrence to all blog posts curently available
let Blog = React.createClass({

    getDefaultProps() {
    }
    render() {
        return (
        <div className="BlogContainer">
        <ol>
            {this.props.posts.map((post) => (
                <li key={post.id}>{post.tittle}</li>
            ))}
        </ol>
        </div>
        )
    }
})

export default Blog

// Blog post fetches BlogPosts