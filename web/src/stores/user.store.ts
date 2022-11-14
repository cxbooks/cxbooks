import type { BookStats } from '@/types';
import { defineStore } from 'pinia';
import { v4 as uuid } from 'uuid';

import { ref, computed, watch } from 'vue'
import { login, getBookstats } from '@/services';

export const useUserInfo = defineStore('user_info', () => {
   
    const token = ref(window.localStorage.getItem('token') || '')

    const userInfo = ref({
        nickName: '',
        role_id: '',
        headerImg: '',
        authority: {},
        sideMode: 'dark',
    })

    const bookStats = ref<BookStats>({total:0,author:0,publisher:0,tag:0})


    const SetToken = (val: string) => {
        token.value = val
    }

    const SetUserInfo = (val: any) => {
        userInfo.value = val
    }

    const SetBookStats = (val: BookStats) => {
        bookStats.value = val
    }


    const ClearStorage = async () => {
        token.value = ''
        sessionStorage.clear()
        localStorage.clear()
    }




    const Login = async (account: string, password: string) => {

        const resp = await login(account, password)

        console.log(resp.code);
        if (resp.code === 0 ){
            
            SetToken(uuid())

            SetUserInfo(resp.data)

            const stats = await getBookstats()

            if (stats.code === 0) {
                console.log(stats.data);
                SetBookStats(stats.data)
            }

            return true
        }

        ClearStorage()

        return false

    }

    watch(() => token.value, () => {
        window.localStorage.setItem('token', token.value)
    })
    // actions: {
    //     async save(account: string, password: string) {
             
    //     }
    // }

    return {
        userInfo,
        token,
        bookStats,
        Login,
        SetToken,
    }
});

   // localStorage.setItem('user', JSON.stringify(account));

            // this.user = user;
            // this.submitted = true;
    


            // update pinia state
         

            // store user details and jwt in local storage to keep user logged in between page refreshes
            

            // redirect to previous url or default to home page

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