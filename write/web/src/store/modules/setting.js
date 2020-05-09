import setting from '@/api/setting'

export default {
    namespaced: true,
    state: {
        setting: {
            background_image: '',
        }
    },
    getters: {
        backgroundImage: state => state.setting.background_image
    },
    mutations: {
        setBackgroundImage(state, image) {
            state.setting.background_image = image
        },
        getBackgroundImage(state) {
            setting.detail().then(response => {
                const {data} = response
                state.setting.background_image = data.background_image
            })
        }
    },
    actions: {
        post({commit}, settingData) {
            commit('setBackgroundImage', settingData.background_image)
            return new Promise((resolve, reject) => {
                setting.update(settingData).then(() => {
                    resolve()
                }).catch(error => reject(error))
            })
        },
        // wryyyyy 
        async fetch({commit}) {
            const {data} = await setting.detail()
            commit('setBackgroundImage', data.background_image)
        }
    }
}