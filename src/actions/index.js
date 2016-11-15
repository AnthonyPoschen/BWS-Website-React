export const testAction = () => {
    return {
        type: 'TEST_1'
    }
}

// Blog Actions
export const fetchBlog = (id) => {
    return {
        type: 'BLOG_FETCH',
        id
    }
}

export const fetchBlogPage = (page , pagesize) => {
    return {
        type: 'BLOG_FETCH_PAGE',
        page,
        pagesize
    }
}