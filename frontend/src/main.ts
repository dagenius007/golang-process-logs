import './styles/main.css'
//@ts-ignore
import { createApp } from 'vue/dist/vue.esm-bundler'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import Popper from 'vue3-popper'
//@ts-ignore
import VueAwesomePaginate from 'vue-awesome-paginate'
import 'vue-awesome-paginate/dist/style.css'
import VueApexCharts from 'vue3-apexcharts'

const pinia = createPinia()
const app = createApp(App)

app.use(router)

app.use(pinia)
app.component('Popper', Popper)
app.use(VueAwesomePaginate)
app.use(VueApexCharts)
app.mount('#app')
