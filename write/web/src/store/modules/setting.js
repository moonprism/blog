import setting from '@/api/setting'

export default {
    namespaced: true,
    state: {
        setting: {
            background_image: '',
            global_css: '',
            global_js: '',
        }
    },
    getters: {
        backgroundImage: state => state.setting.background_image,
        globalJS: state => state.setting.global_js,
        globalCSS: state => state.setting.global_css
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
                state.setting = response.data
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