import { createApp } from 'vue'
import './style.css'
import 'ant-design-vue/dist/reset.css'
import Antd from 'ant-design-vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
const pinia = createPinia();

const app = createApp(App);
app.use(Antd)
app.use(router)
app.use(pinia);
app.mount('#app')
