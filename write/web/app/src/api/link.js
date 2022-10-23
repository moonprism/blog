import request from '@/utils/request'

const link = {
    list(page) {
        return request.get('/link', {params: {page: page}})
    },
}

export default link
