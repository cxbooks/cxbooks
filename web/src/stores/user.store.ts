import { defineStore } from 'pinia';


export const userStore = defineStore('user_info', {
    state: () => {
        return {
            nickname: "",
            returnUrl: "",
            count: 0,
            loading: true,
            role: 1000,
            isLogin: false,
        }
    },
    actions: {
        async save(account: string, password: string) {


            localStorage.setItem('user', JSON.stringify(account));

            // this.user = user;
            // this.submitted = true;
    


            // update pinia state
         

            // store user details and jwt in local storage to keep user logged in between page refreshes
            

            // redirect to previous url or default to home page
            
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