import request from '@/utils/request'

const article = {
    list(page) {
        return request.get('/article', {params: {page: page}})
    },
    create(data) {
        return request.post('/article', data)
    },
    update(id, data) {
        return request.put('/article/'+id, data)
    },
    delete(id) {
        return request.delete('/article/'+id)
    },
    detail(id) {
        return request.get('/article/'+id)
    }
}

export default article