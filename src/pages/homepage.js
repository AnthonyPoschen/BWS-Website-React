import '../App.css'

import React from 'react'
import {Link} from 'react-router'

import {testAction} from '../actions'
import Blog from '../components/blog.js'
import {connect} from 'react-redux'

import {Carousel, PageHeader , Col , Row, Grid , Button} from 'react-bootstrap'

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
        <Grid fluid>
            <Row>
                <Carousel>
                    <Carousel.Item>
                        <img width={900} height={500} alt="900x500" src="https://upload.wikimedia.org/wikipedia/commons/9/90/Dscn3308-rg-e-from-w_900x500.jpg"/>
                        <Carousel.Caption>
                            <h3>First slide label</h3>
                            <p>Nulla vitae elit libero, a pharetra augue mollis interdum.</p>
                            </Carousel.Caption>
                    </Carousel.Item>
                    <Carousel.Item>
                        <img width={900} height={500} alt="900x500" src="http://photography.cdeiwiks.com/wp-content/uploads/2013/02/NZ034-069-900x500-900x500.jpg"/>
                        <Carousel.Caption>
                            <h3>Second slide label</h3>
                            <p>Nulla vitae elit libero, a pharetra augue mollis interdum.</p>
                            </Carousel.Caption>
                    </Carousel.Item>
                </Carousel>
            </Row>
                <Row>
                    <Col sm={7} md={8}>
                        <PageHeader>
                            About Danger Zone
                        </PageHeader>
                        <p>
                            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sed feugiat augue. Mauris non ligula eget augue vehicula ultrices ac elementum metus. Proin malesuada congue augue sed faucibus. Maecenas in ex in leo sagittis ullamcorper a ut diam. Proin enim metus, faucibus quis sapien eget, gravida rhoncus arcu. Integer faucibus elit tellus, et scelerisque purus pellentesque id. Duis tristique porta dui.

                            Nunc quis erat mauris. Aliquam ultrices vulputate dui sit amet auctor. Nunc cursus, turpis et gravida tempus, lectus massa vulputate felis, in ultrices ante sapien sit amet augue. Nam lobortis dui eget tempus malesuada. Quisque ipsum sem, pulvinar quis aliquam at, tempus accumsan ex. Phasellus auctor ante non turpis vestibulum, lobortis convallis lorem fringilla. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut molestie mauris eget elit faucibus, nec egestas tellus laoreet. Proin eget ultricies velit, eget ultrices libero. Suspendisse fringilla sapien a tortor eleifend, sagittis consectetur ex tristique. Curabitur ullamcorper blandit nulla a ultrices. Cras a erat lorem. Morbi at convallis orci. Integer ac turpis nulla. Cras et tincidunt justo. Sed mattis a mauris et gravida.
                        </p>
                    </Col>
                    <Col sm={5} md={4}>
                        <div className="SidePanelContainer">
                            <h3>
                                Recent Blog Posts
                            </h3>
                            <Blog> Test</Blog>
                            <h3>
                                Get In Touch
                            </h3>
                            <div>
                                Facebook Icon | Twitter Icon | Email Icon | RSS Icon
                            </div>
                        </div>
                    </Col>
                </Row>
            </Grid>

            
            <p>{this.props.value}</p>
            <Button onClick={this.props.myClick} > Test Button</Button>
        </div>
    }
})

export default connect(mapStateToProps,mapDispatchToProps)(HomePage)