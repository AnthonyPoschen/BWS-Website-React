import React from 'react'
import {Link} from 'react-router'
import {connect} from 'react-redux'

import {well , ListGroupItem } from 'react-bootstrap'
import {fetchBlogPage} from '../actions'

const mapStateToProps = (state , ownProps) => {
    return {
        pageitems: state.blog.BlogPageItems,
    }
}

const mapDispatchToProps = (dispatch, ownProps) => {
    return {
        onLoad: () => {
            console.log("Fetching page")
            dispatch(fetchBlogPage(0,50))
        }
    }
}

let Blog = React.createClass({
    componentDidMount(){this.props.onLoad();},

    render() {
        if(this.props.pageitems === null){
            return (<div></div>)
        }
        
        // TODO: fetch this infromation from the server.
        var posts = this.props.pageitems.slice(0,8);
        
        //<li key={post.id}>{post.title} - {post.content}</li>
        return (
        <div>
        
        {posts.map((post) => (
            <Link key={post.ID} to={'/BlogPost/' + post.ID + '/' + post.Tittle}>
            <div >
                <p>{post.Tittle} </p>
                <small>Author: {post.AuthorName} on {post.PublishDate}</small>
            </div>
            </Link>
        ))}
        
        </div>
        )
    }
})

export default connect(mapStateToProps,mapDispatchToProps)(Blog)