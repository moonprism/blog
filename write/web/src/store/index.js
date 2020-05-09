import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import auth from './modules/auth'
import article from './modules/article'
import setting from './modules/setting'

const store = new Vuex.Store({
    modules: {
        auth,
        article,
        setting
    }
})

export default store