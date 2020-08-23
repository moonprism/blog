import axios from 'axios'
import store from '@/store'
import { MessageBox, Message, Loading } from 'element-ui'

var instance = axios.create({
    baseURL: process.env.VUE_APP_BASE_API,
    timeout: 3000
})

let loading
let requestNum = 0

instance.interceptors.request.use(
    data => {
        // login page not loading
        if (data.url != '/auth/login') {
            if (requestNum == 0) {
                let url = data.url
                if (data.params) {
                    url += '?'
                    Object.keys(data.params).forEach(key => {
                        url += (key+'='+data.params[key]+'&')
                    })
                    url = url.substr(0, url.length - 1)
                }
                loading = Loading.service({
                    lock: true,
                    text: '[ '+url+' ]',
                })
            }
            requestNum++
        }
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
    res => {
        requestNum--
        if (requestNum == 0) {
            loading.close()
        } else if (requestNum < 0) {
            requestNum = 0
        }
        return res.status === 200 ? Promise.resolve(res) : Promise.reject(res)
    },
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
                    break
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
