import { createApp } from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify";
import { loadFonts } from "./plugins/webfontloader";
import "vuetify/dist/vuetify.min.css";
import { createPinia } from "pinia";

loadFonts();
const pinia = createPinia();

const app = createApp(App);
app.use(pinia);
app.use(vuetify);
app.mount("#app");
