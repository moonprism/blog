import request from '@/utils/request'

const setting = {
    detail() {
        return request.get('/system/info')
    },
    update(data) {
        return request.put('/system/info/theme', data)
    },
    updateEmail(data) {
        return request.put('/system/info/notice', data)
    }
}

export default setting
