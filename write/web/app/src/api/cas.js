import request from '@/utils/request'

const code = {
    key() {
        return request.get('/cas/key')
    }
}

export default code
