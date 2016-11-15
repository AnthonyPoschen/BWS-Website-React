import '../App.css'

import React from 'react'
import {Link} from 'react-router'

import {testAction} from '../actions'
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


let HomePage = React.createClass({
    render() {
        return <div className="PageHeader">
            <h2>Home Page</h2>
            <Link to="/Blog">Blog Page</Link>
            <p>{this.props.value}</p>
            <button onClick={this.props.myClick} > Test Button</button>
        </div>
    }
})

export default connect(mapStateToProps,mapDispatchToProps)(HomePage)