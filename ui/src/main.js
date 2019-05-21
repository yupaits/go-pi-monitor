import Vue from 'vue'
import App from './App.vue'
import Antd from 'ant-design-vue'
import utils from './utils'

import 'ant-design-vue/dist/antd.min.css'

Vue.config.productionTip = false

Vue.use(Antd)

Vue.prototype.$utils = utils

new Vue({
  render: h => h(App),
}).$mount('#app')
