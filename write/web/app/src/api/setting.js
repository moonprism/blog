import request from '@/utils/request'

const setting = {
    detail() {
        return request.get('/system/info')
    },
    update(data) {
        return request.put('/system/info', data)
    },
}

export default setting
