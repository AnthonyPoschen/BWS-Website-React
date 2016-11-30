import React from 'react'

let AdminPage = React.createClass({
    // force login if not logged in
    // display blog posts and the ability to edit
    // TODO: create blog posts
    // TODO: edit blog posts
    // TODO: publish blog posts
    // TODO: preview blog posts
    render() {
        return (
            <div> 
                <p>Admin Page ooooo</p>
                <p>Here is where all the functionality like writing posts and editing and publishing them should be</p>
            </div>
        )
    }
})

export default AdminPage

/* 
Notes:
    Left pane menu to display all information
    regarding functionallity of the admin Page
    like reviewing preview posts or writing new ones
    or even editing exsisting ones. 

    make page only viewable with admin privilages somehow.
    possibly with a google login edirect if not logged in.
    and auth all apis for posting checking if it comes
    from a logged in client.
*/