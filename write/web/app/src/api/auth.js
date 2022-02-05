import request from '@/utils/request'

const auth = {
    login(data) {
        return request.post('/auth/login', data)
    }
}

export default auth

