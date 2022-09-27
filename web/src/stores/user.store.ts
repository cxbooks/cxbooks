import { defineStore } from 'pinia';


import router  from '@/router';



const baseUrl = `${import.meta.env.VUE_APP_API_PATH}`;


export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        // initialize state from local storage to enable user to stay logged in
        // user: JSON.parse(localStorage.getItem('user')),
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
        async save(account: string, password: string) {


            localStorage.setItem('user', JSON.stringify(account));

            // this.user = user;
            // this.submitted = true;
            router.push(this.returnUrl || '/');


            // update pinia state
         

            // store user details and jwt in local storage to keep user logged in between page refreshes
            

            // redirect to previous url or default to home page
            
        },
        logout() {
            // this.user = null;
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