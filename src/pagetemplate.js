import './App.css'

import React from 'react'
import AppBar from 'material-ui/AppBar';
import {Tabs, Tab} from 'material-ui/Tabs';
import {browserHistory} from 'react-router'
import {Toolbar, ToolbarGroup, ToolbarSeparator, ToolbarTitle} from 'material-ui/Toolbar';

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
                <Toolbar>
                    <ToolbarGroup firstChild={true}>
                        <ToolbarTitle text="Brainwave Studios" />
                        <Tabs value={this.props.location.pathname} >
                            <Tab value="/" label="Home" data-route="/" onActive={handleActive} >
                            </Tab>
                            <Tab value="/Blog" label="Blog" data-route="/Blog" onActive={handleActive} >
                            </Tab>
                        </Tabs>
                    </ToolbarGroup>
                </Toolbar>
                
                <div>{this.props.children}</div>
            </div>
        
    }
})

export default PageTemplate 