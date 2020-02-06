import request from '@/utils/request'

const tag = {
    list() {
        return request.get('/tag')
    },
    create(data) {
        return request.post('/tag', data)
    },
    update(id, data) {
        return request.put('/tag/'+id, data)
    },
    delete(id) {
        return request.delete('/tag/'+id)
    }
}

export default tag
