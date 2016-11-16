export default function blog(state = {
blogs: [{id:0,Title:'',content:'',date:'',author:''}],
activeBlog: null,

} , action) {
    switch (action.type) {
        case 'BLOG_FETCH':
        // fetch a blog in detail
            return {blogs:[{id:1,Title:'Test Blog',date:'somedate',content:'this is the content for a post',author:'Zanven'}],activeBlog: null};
        case 'BLOG_FETCH_PAGE':
        // fetch a page of blog titles and basic information. 
            return {blogs:[{id:1,Title:'Test Blog',content:'<p>Test <b>Content</b><p>',date:'somedate',author:'Zanven'},{id:2,Title:'Blog 2',content:'<h3>Yet another post</h3>',date:'somedate',author:'ShellSpoon'}],activeBlog: null};
        default:
            return state
    }
    
}
