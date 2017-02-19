import '../App.css'

import React from 'react'
import {Link} from 'react-router'

import {} from '../actions'
import {connect} from 'react-redux'
import { PageHeader } from 'react-bootstrap'; 

// TODO: LOAD BLOGPOST WHEN THIS IS MOUNTED+
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

let BlogPost = React.createClass({
    render() {
        var post = this.props.posts.filter( obj => {return obj.ID == this.props.params.ID})[0]
        console.log(this.props)
        console.log(post)
        if(typeof(post) === 'undefined') 
            return <div></div>
        return (
            // Blog post is setup here
            <div>
                <PageHeader>{post.Tittle}</PageHeader>
                <div dangerouslySetInnerHTML={{__html:post.Content }}></div>
            </div>
        )
    }
})

export default connect(mapStateToProps,mapDispatchToProps)(BlogPost)