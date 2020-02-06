import axios from 'axios'
import store from '@/store'
import { MessageBox, Message } from 'element-ui'

var instance = axios.create({
    baseURL: process.env.VUE_APP_BASE_API,
    timeout: 3000
})

instance.interceptors.request.use(
    data => {
        // 全局附加jwt校验
        if (store.getters['auth/token']) {
            data.headers['Authorization'] = 'Bearer ' + store.getters['auth/token']
        }
        return data
    },
    error => {
        return Promise.reject(error)
    }
)

instance.interceptors.response.use(
    // 延迟执行
    res => res.status === 200 ? Promise.resolve(res) : Promise.reject(res),
    error => {
        if (error.response) {
            let data = error.response.data
            // 全局错误处理
            switch (data.key) {
                // 未授权的访问
                case 'Unauthenticated':
                    MessageBox.confirm('auth expired', '', {
                        confirmButtonText: 're login',
                        cancelButtonText: 'Cancel',
                        type: 'warning'
                    }).then(() => {
                        store.dispatch('auth/reset').then(() => {
                            location.reload()
                        })
                    })
                    break;
                default:
                    Message({
                        message: data.message,
                        type: 'error',
                        duration: 3 * 1000
                    })
                    break

            }
        }
    }
)

export default instance