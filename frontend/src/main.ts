import "./assets/main.css";

import { createApp } from "vue";
import { createPinia } from "pinia";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";
import Lara from "./style-preset/index";

import App from "./App.vue";
import router from "./router";

import PrimeVue from "primevue/config";

const app = createApp(App);

const ptOptions = {
    table: "p-4",
    bodyRow: "bg-slate-200 odd:bg-slate-50 border-t  first:border-none",
    headerRow: "text-blue-900 text-left border-b border-b-blue-900 rounded-xl",
    bodyCell: "bg-red-100",
};

app.use(PrimeVue, {
    pt: Lara,
});

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);

app.use(pinia);
app.use(router);

app.mount("#app");
