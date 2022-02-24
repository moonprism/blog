import request from '@/utils/request'

const code = {
    list(page) {
        return request.get('/code', {params: {page: page}})
    },
    create(data) {
        return request.post('/code', data)
    },
    update(id, data) {
        return request.put('/code/'+id, data)
    },
    delete(id) {
        return request.delete('/code/'+id)
    },
    detail(id) {
        return request.get('/code/'+id)
    }
}

export default code
