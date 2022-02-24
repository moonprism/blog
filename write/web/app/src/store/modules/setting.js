import setting from '@/api/setting'

export default {
    namespaced: true,
    state: {
        setting: {
            background_image: '',
            global_css: '',
            global_js: '',
            captcha: false,
            notice: false,
        }
    },
    getters: {
        backgroundImage: state => state.setting.background_image,
        globalJS: state => state.setting.global_js,
        globalCSS: state => state.setting.global_css,
        captcha: state => state.setting.captcha,
        notice: state => state.setting.notice
    },
    mutations: {
        setBackgroundImage(state, image) {
            state.setting.background_image = image
        },
        setGlobalJS(state, js) {
            state.setting.global_js = js
        },
        setGlobalCSS(state, css) {
            state.setting.global_css = css
        },
        set(state, setting) {
            state.setting = setting
        },
        init(state) {
            setting.detail().then(response => {
                if (response.data != null) {
                    state.setting = response.data.theme
                    state.setting.captcha = response.data.captcha
                    state.setting.notice = response.data.notice
                }
            })
        }
    },
    actions: {
        post({commit}, settingData) {
            commit('set', settingData)
            return new Promise((resolve, reject) => {
                setting.update(settingData).then(() => {
                    resolve()
                }).catch(error => reject(error))
            })
        },
        // wryyyyy 
        // async fetch({commit}) {
        //     const {data} = await setting.detail()
        //     commit('setBackgroundImage', data.background_image)
        // }
    }
}
