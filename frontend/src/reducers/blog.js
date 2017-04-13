export default function blog(
    state = {
        BlogPageItems: [],
        blogPosts: [],
        selected: null} , 
    action) 
    {
    switch (action.type) {
        case 'BLOG_PAGE_RECIEVED':
            // Temp to print to console
            console.log("Blogs updated xD")
            console.log({...state , BlogPageItems: action.payload })

            return Object.assign({},...state , {BlogPageItems: action.payload });
        case 'BLOG_PAGE_POST_SELECTED':
            console.log("Changing Selected Element");
            return Object.assign({},...state,{selected: action.payload});
        case 'BLOG_FETCH':
        // fetch a blog in detail
            return {blogs:[{id:1,Title:'Test Blog',date:'somedate',content:'this is the content for a post',author:'Zanven'}],activeBlog: null};
        case 'BLOG_FETCH_PAGE':
        // fetch a page of blog titles and basic information. 
            return {blogs:[{id:1,Title:'Test Blog',content:'<p>Test <b>Content</b><p>',date:'somedate',author:'Zanven'},{id:2,Title:'Blog 2',content:'<h3>Yet another post</h3>',date:'somedate',author:'ShellSpoon'}],activeBlog: null};
        
        case 'BLOG_POST_RECIEVED':
            var bps = state.blogPosts.slice(0)
            bps[action.payload.ID] = action.payload
            console.log("Recieved Blog Post")
            console.log(bps)
            return Object.assign({},state,{ blogPosts: bps },{selected: action.payload.ID} )
        default:
            return state
        }
    
    }
