import { createApp } from 'vue'
import '@/styles/index.scss'
import App from './App.vue'
import { startMicroApp4Child } from '@/utils/microApp'
import router from '@/router'
import pinia from '@/store'
import ElementPlus from 'element-plus'
import { useApp } from '@/store/app'

const app = createApp(App)
app
  .use(router)
  .use(pinia)
  .use(ElementPlus)
  .mount('#novel-app')

const appConf = useApp()
appConf.setApp(app)
startMicroApp4Child()