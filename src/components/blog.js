import React from 'react'
import {Link} from 'react-router'
import {connect} from 'react-redux'

import {well , ListGroupItem } from 'react-bootstrap'
import {fetchBlogPage} from '../actions'

const mapStateToProps = (state , ownProps) => {
    return {
        posts: state.blog.blogs,
    }
}

const mapDispatchToProps = (dispatch, ownProps) => {
    return {
        myClick: () => {
            console.log("Dispatching")   
        },
        onLoad: () => {
            console.log("Fetching page")
            dispatch(fetchBlogPage(0,50))
        }
    }
}

let Blog = React.createClass({
    componentDidMount(){this.props.onLoad();},
    render() {
        // TODO: fetch this infromation from the server.
        var posts = this.props.posts.slice(0,8);
        //<li key={post.id}>{post.title} - {post.content}</li>
        return (
        <div>
        
        {posts.map((post) => (
            <Link key={post.id} to={'/BlogPost/' + post.id + '/' + post.Title.replace(/ /g,"_")}>
            <div >
                <p>{post.Title} </p>
                <small>Author: {post.author} on {post.date}</small>
            </div>
            </Link>
        ))}
        
        </div>
        )
    }
})

export default connect(mapStateToProps,mapDispatchToProps)(Blog)