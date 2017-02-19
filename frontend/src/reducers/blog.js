export default function blog(
    state = {
        blogs: [{ID:'',Tittle:'',Content:'',PublishDate:'',Author:''}],
        activeBlog: null} , 
    action) 
    {
    switch (action.type) {
        case 'BLOG_PAGE_RECIEVED':
            // Temp to print to console
            console.log("Blogs updated xD")
            console.log({...state , blogs: action.payload })

            return Object.assign({},...state , {blogs: action.payload });
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
