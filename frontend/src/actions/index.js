import axios from 'axios'

export const TEST_1 = 'TEST_1'
export const testAction = () => {
    return {
        type: TEST_1
    }
}

// Blog Actions
export const BLOG_FETCH = 'BLOG_FETCH'
export const fetchBlog = (id) => {
    return {
        type: BLOG_FETCH,
        id
    }
}

export const BLOG_PAGE_RECIEVED = 'BLOG_PAGE_RECIEVED'
export const blogPageRecieved = (data) => {
    return {
        type: BLOG_PAGE_RECIEVED,
        payload: data
    }
}

export const BLOG_PAGE_ERROR_FETCH = 'BLOG_PAGE_ERROR_FETCH'
export const blogPageErrorFetch = (error) => {
    return {
        type: 'BLOG_PAGE_ERROR_FETCH'
    }
}

export const fetchBlogPage = (page , pagesize) => {
    let url = `/api/getbloglist?page=${page}&size=${pagesize}`
    
    return (dispatch) => {
        
        axios.get(url).then(
            res => {
                if (res.status != 200)
                {
                    dispatch(blogPageErrorFetch(res.statusText))
                    return
                }
                console.log("Dispatching Recieved Blog Page")
                dispatch(blogPageRecieved(res.data)) 

            },
        ).catch( (error) => {
            dispatch(blogPageErrorFetch(error))
        })
    }
}

export const BLOG_PAGE_POST_SELECTED = 'BLOG_PAGE_POST_SELECTED'
export const blogPagePostSelected = (ID) => {
    return {
        type: 'BLOG_PAGE_POST_SELECTED',
        payload: ID
    }
}

export const BLOG_POST_FETCH = 'BLOG_POST_FETCH'
export const blogPostFetch = (Tittle , ID) => {
    var url = `/api/getblog`

    return (dispatch) => {
        axios.post(url,{"Tittle": Tittle, "ID": ID}).then(
            res => {
                if(res.status != 200)
                {
                    // Dispatch Error
                    return
                }
                // Dispatch that we got the blog post !! 
                dispatch(blogPostRecieved(res.data))
            },
        ).catch( (error) => { 
            // Dispatch Error
        })
    }
}

export const BLOG_POST_RECIEVED = 'BLOG_POST_RECIEVED'
export const blogPostRecieved = (data) => {
    return {
        type: BLOG_POST_RECIEVED,
        payload: data
    }
}