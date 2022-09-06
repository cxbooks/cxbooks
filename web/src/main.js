
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'

import { createStore } from 'vuex'

// Create a new store instance.
const store = createStore({
    state() {
        return {
            count: 0,
            loading: true,
            nav: true,
            role: 1000,
            nickname: "",
            is_login: false,
            sys: {
                socials: [], allow: {},
            }
        }
    },
    mutations: {
        increment(state) {
            state.count++
        },
        loaded(state) {
            state.loading = false;
        },
    }
})

loadFonts()

createApp(App)
    .use(router)
    .use(vuetify)
    .use(store)
    .mount('#app')
