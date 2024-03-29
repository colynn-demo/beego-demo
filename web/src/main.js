import Vue from 'vue'
import App from './App.vue'

import Element from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css';


// Vue.config.productionTip = false

Vue.use(Element)

new Vue({
  el: '#app',
  render: h => h(App),
})
