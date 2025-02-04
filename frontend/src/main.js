import './assets/main.css'

import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import axios from "axios";

const app = createApp(App)

app.use(router)

app.config.globalProperties.$axios= axios.create({
baseURL: "http://localhost:8090",
timeout: 1000 * 5
});

app.mount('#app')
