import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import "@/scss/style.scss";
import VueApexCharts from "vue3-apexcharts";
import { createPinia } from 'pinia';
// @ts-ignore:next-line

const app = createApp(App);
app.use(createPinia());
app.use(VueApexCharts);
app.use(router);
app.use(vuetify).mount("#app");
