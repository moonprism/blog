import request from '@/utils/request'

const code = {
    list(page) {
        return request.get('/comment', {params: {page: page}})
    },
    delete(id) {
        return request.delete('/comment/'+id)
    },
}

export default code
