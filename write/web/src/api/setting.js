import request from '@/utils/request'

const setting = {
    detail() {
        return request.get('/setting')
    },
    update(data) {
        return request.put('/setting', data)
    },
}

export default setting
