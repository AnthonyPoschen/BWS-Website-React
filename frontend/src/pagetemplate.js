import './App.css'


import React from 'react'
//import AppBar from 'material-ui/AppBar';
//import {Tabs, Tab} from 'material-ui/Tabs';
import {browserHistory, Link} from 'react-router'
//import {Toolbar, ToolbarGroup, ToolbarSeparator, ToolbarTitle} from 'material-ui/Toolbar';
import { Button, Navbar, Nav, NavItem, NavDropdown, Col, Grid } from 'react-bootstrap';

import { LinkContainer } from 'react-router-bootstrap';

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
                            <Link to="/">Brainwave Studios</Link>
                        </Navbar.Brand>
                    </Navbar.Header>
                    <Nav>
                        <LinkContainer to="/"><NavItem eventKey={1} >Home</NavItem></LinkContainer>
                        <LinkContainer to="/Blog"><NavItem eventKey={2} >Blog</NavItem></LinkContainer>
                        <LinkContainer to="/About"><NavItem eventKey={3} >About</NavItem></LinkContainer>
                    </Nav>
                </Navbar>
                <Grid>
                    {/* Main Content*/}
                    <div>{this.props.children}</div>
                    {/* Footer */}
                </Grid>
            </div>
        
    }
})

export default PageTemplate 