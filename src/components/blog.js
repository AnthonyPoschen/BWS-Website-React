import React from 'react'
import {Link} from 'react-router'
import {connect} from 'react-redux'

import {well , ListGroupItem } from 'react-bootstrap'


const mapStateToProps = (state , ownProps) => {
    return {
        posts: state.blog.blogs,
    }
}

const mapDispatchToProps = (dispatch, ownProps) => {
    return {
        myClick: () => {
            console.log("Dispatching")
            dispatch(testAction())    
        }
    
    }
}

let Blog = React.createClass({

    render() {
        // TODO: fetch this infromation from the server.
        var posts = this.props.posts.slice(0,8);
        console.log('Testing homie');
        //<li key={post.id}>{post.title} - {post.content}</li>
        return (
        <div>
        <well>
        {posts.map((post) => (
            <Link to={'/BlogPost/' + post.id + '/' + post.Title.replace(/ /g,"_")}>
            <ListGroupItem key={post.id}>
                <h3>{post.Title} </h3>
                <small>Author: {post.author} on {post.date}</small>
            </ListGroupItem>
            </Link>
        ))}
        </well>
        </div>
        )
    }
})

export default connect(mapStateToProps,mapDispatchToProps)(Blog)