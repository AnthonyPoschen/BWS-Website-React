import '../App.css'

import React from 'react'
import {Link} from 'react-router'

import {} from '../actions'
import {connect} from 'react-redux'

const mapStateToProps = (state , ownProps) => {
    return {
        value: state.test
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
        return (
            // Blog post is setup here
            <div>
                <h3>{this.props.Title}</h3>
                <div dangerouslySetInnerHTML={{__html:this.props.content }}></div>
            </div>
        )
    }
})

export default connect(mapStateToProps,mapDispatchToProps)(BlogPost)