import React from 'react'
import {Link} from 'react-router'
import Blog from '../components/blog.js'
import BlogPost from './blogpost.js'
import {fetchBlogPage} from '../actions'
import {connect} from 'react-redux'
import { PageHeader , Well, ListGroup, ListGroupItem } from 'react-bootstrap'; 

const mapStateToProps = (state , ownProps) => {
    return {
        blogs: state.blog.blogs
    }
}

const mapDispatchToProps = (dispatch, ownProps) => {
    return {
        myClick: () => {
            dispatch(testAction())    
        },
        nextPage: (curpage,pagesize) => {
            dispatch(fetchBlogPage(curpage+1,pagesize))
        },

        loadPage: () => {
            dispatch(fetchBlogPage(0,20))
        }
    
    }
}

let BlogPage = React.createClass({

    componentDidMount() {
        this.props.loadPage()
    },

    render() {
    if (this.props.blogs === null) {
        return <div> </div>
    }
     return <div>
        <PageHeader>Blog</PageHeader>
        <well>
        {this.props.blogs.map((post) => (
            <Link to={'/BlogPost/' + post.ID + '/' + post.Tittle}>
            <ListGroupItem key={post.ID} header={post.Tittle}>
                <small>Author: {post.AuthorName} on {post.PublishDate}</small>
            </ListGroupItem>
            </Link>
        ))}
        </well>
     </div>   
    }
})

export default connect(mapStateToProps,mapDispatchToProps)(BlogPage)