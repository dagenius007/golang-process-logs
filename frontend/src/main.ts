import './styles/main.css'

import { createApp } from 'vue/dist/vue.esm-bundler'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import Popper from 'vue3-popper'

const pinia = createPinia()
const app = createApp(App)

app.use(router)

app.use(pinia)
app.component('Popper', Popper)
app.mount('#app')
