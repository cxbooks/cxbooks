import { defineStore } from 'pinia';

import { fetchWrapper } from '@/helpers';

import router  from '@/router';

const baseUrl = `${process.env.VUE_APP_API_PATH}`;


export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        // initialize state from local storage to enable user to stay logged in
        user: JSON.parse(localStorage.getItem('user')),
        returnUrl: null,
        count: 0,
        loading: true,
        nav: true,
        role: 1000,
        nickname: "",
        is_login: false,
        sys: {
            socials: [], allow: {},
        }
    }),
    actions: {
        async login(username, password) {

            console.log(process.env.VUE_APP_SECRET)
            const user = await fetchWrapper.post(`${baseUrl}/user/login`, { username, password });

            // update pinia state
            this.user = user;

            // store user details and jwt in local storage to keep user logged in between page refreshes
            localStorage.setItem('user', JSON.stringify(user));

            // redirect to previous url or default to home page
            router.push(this.returnUrl || '/');
        },
        logout() {
            this.user = null;
            localStorage.removeItem('user');
            router.push('/login');
        }
    }
});


// Create a new store instance.
// const store = createStore({
//     state() {
//         return {
//             count: 0,
//             loading: true,
//             nav: true,
//             role: 1000,
//             nickname: "",
//             is_login: false,
//             sys: {
//                 socials: [], allow: {},
//             }
//         }
//     },
//     mutations: {
//         increment(state) {
//             state.count++
//         },
//         loaded(state) {
//             state.loading = false;
//         },
//     }
// })