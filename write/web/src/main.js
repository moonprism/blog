import Vue from 'vue'
import App from './App.vue'

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

import store from './store'
import router from './router'

Vue.use(ElementUI)

new Vue({
  render: h => h(App),
  router,
  store,
}).$mount('#app')