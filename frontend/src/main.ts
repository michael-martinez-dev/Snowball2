import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import 'vuetify/dist/vuetify.min.css'
import store from './store'

loadFonts()

createApp(App)
  .use(vuetify)
  .use(store)
  .mount('#app')
