import React from 'react'
import {Link} from 'react-router'
import Blog from '../components/blog.js'
import BlogPost from '../components/blogpost.js'
import {fetchBlogPage} from '../actions'
import {connect} from 'react-redux'
import { PageHeader } from 'react-bootstrap'; 

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
     return <div>
        <PageHeader>Blog Page</PageHeader>
        {this.props.blogs.map((post) => (
            <BlogPost key={post.id} {...post} />
        ))}
     </div>   
    }
})

export default connect(mapStateToProps,mapDispatchToProps)(BlogPage)