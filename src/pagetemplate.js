import './App.css'


import React from 'react'
//import AppBar from 'material-ui/AppBar';
//import {Tabs, Tab} from 'material-ui/Tabs';
import {browserHistory} from 'react-router'
//import {Toolbar, ToolbarGroup, ToolbarSeparator, ToolbarTitle} from 'material-ui/Toolbar';
import { Button, Navbar, Nav, NavItem, NavDropdown } from 'react-bootstrap';

let PageTemplate = React.createClass({

    componentWillMount() {
        let hash = window.location.hash;
        hash = hash.substr(1,hash.indexOf("/",0)-1);
        //let urlPath = window.location.urlPath;
        //let currentTab = urlPath.split('/').pop();
        // you can add more validations here
        this.setState({ activeTab: hash || '/' });
    },
    render() {
        function handleActive(tab) {
            browserHistory.push(tab.props['data-route']);  
        }
        return <div>
                <Navbar>
                    <Navbar.Header>
                        <Navbar.Brand>
                            <a href="/">My App</a>
                        </Navbar.Brand>
                    </Navbar.Header>
                    <Nav>
                        <NavItem eventKey={1} href="#">Link</NavItem>
                        <NavItem eventKey={2} href="#">Link</NavItem>
                    </Nav>
                </Navbar>
                
                <div>{this.props.children}</div>
            </div>
        
    }
})

export default PageTemplate 