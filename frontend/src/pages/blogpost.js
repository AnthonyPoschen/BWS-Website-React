import '../App.css'

import React from 'react'
import {Link} from 'react-router'

import {blogPostFetch} from '../actions'
import {connect} from 'react-redux'
import { PageHeader } from 'react-bootstrap'; 

const mapStateToProps = (state , ownProps) => {
    return {
        posts: state.blog.blogPosts,
        selected: state.blog.selected,
    }
}

const mapDispatchToProps = (dispatch, ownProps) => {
    return {
        FetchBlogPost: (Tittle,ID) => {
            console.log("Fetching Blog Post Tittle: " + Tittle + " ID: " + ID );
            dispatch(blogPostFetch(Tittle,ID));    
        }
    }
}

let BlogPost = React.createClass({
    componentDidMount() {
        // Call api to fetch the blog post in full with Content
        var url = window.location.toString().split("/");
        url = url[url.length - 1];
        var Tittle = decodeURI(url);
        var ID = this.props.selected;
        this.props.FetchBlogPost(Tittle,ID);
    },

    render() {
        if(this.props.selected != null) {
            console.log("Selected Element: "+ this.props.selected)
            post = this.props.posts.filter( obj => {return obj.ID == this.props.selected})[0]
        }
        var post = this.props.posts[this.props.selected]
        if(typeof(post) === 'undefined') 
            return <div></div>
        var date = new Date(post.PublishDate)
        return (
            // Blog post is setup here
            <div>
                <PageHeader>{post.Tittle}<small>Author: {post.AuthorName} - {date}</small></PageHeader>
                <div dangerouslySetInnerHTML={{__html:post.Content }}></div>
            </div>
        )
    }
})

export default connect(mapStateToProps,mapDispatchToProps)(BlogPost)