import {createApp} from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import * as ElIconModules from '@element-plus/icons-vue'

// 全局组件挂载
const app = createApp(App)
app.use(ElementPlus)

Object.keys(ElIconModules).forEach(function (key) {
    app.component(ElIconModules[key].name, ElIconModules[key])
  })
app.mount('#app')
