import auth from '@/api/auth'
import {getToken, setToken, removeToken} from '@/utils/auth'

export default {
    namespaced: true,
    state: {
        token: getToken()
    },
    mutations: {
        setToken: (state, token) => {
            state.token = token
        },
        resetToken: (state) => {
            state.token = ''
        }
    },
    getters: {
        token: state => state.token
    },
    actions: {
        login({ commit }, userInfo) {
            const { username, password } = userInfo
            return new Promise((resolve, reject) => {
                auth.login({ username: username, password: password }).then(response => {
                    const { data } = response
                    commit('setToken', data.token)
                    setToken(data.token)
                    resolve(data)
                }).catch(error => {
                    reject(error)
                })
            })
        },
        reset({ commit }) {
            return new Promise(resolve => {
                removeToken()
                commit('resetToken')
                resolve()
            })
        },
    }
}
