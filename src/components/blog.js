import React from 'react'
import {Link} from 'react-router'
import BlogPost from './blogpost.js'
// contains a refrence to all blog posts curently available
let Blog = React.createClass({

    render() {
        // TODO: fetch this infromation from the server.
        var posts = [
            {id: 1, title: "Blog Post 1",content: "<p>This is the text for the <b>blog post</b></p>",url:"/Blog/post/1"},
            {id: 2, title: "A seocnd post",content: "<p>look im a <b>blogger</b></p>",url:"/Blog/post/2"},
        ];
        //<li key={post.id}>{post.title} - {post.content}</li>
        return (
        <div className="BlogContainer">
            <h2>Blogs go under here</h2>
               
            {posts.map((post) => (
                <BlogPost key={post.id} {...post}></BlogPost>
            ))}
        </div>
        )
    }
})

export default Blog

// Blog post fetches BlogPosts