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
export const blogPageRecieved = (blogpage) => {
    return {
        type: BLOG_PAGE_RECIEVED,
        payload: blogpage
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
                dispatch(blogPageRecieved(res.data)) 

            },
        ).catch( (error) => {
            dispatch(blogPageErrorFetch(error))
        })
    }
}