

import React from 'react'
import AppBar from 'material-ui/AppBar';
import {Tabs, Tab} from 'material-ui/Tabs';
import {browserHistory} from 'react-router'


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
                <AppBar title="My App">
                    <Tabs value={this.props.location.pathname}>
                        <Tab value="/" label="Home" data-route="/" onActive={handleActive}>
                        </Tab>
                        <Tab value="/Blog" label="Blog" data-route="/Blog" onActive={handleActive}>
                        </Tab>
                    </Tabs>
                </AppBar>
                
                <div>{this.props.children}</div>
            </div>
        
    }
})

export default PageTemplate 