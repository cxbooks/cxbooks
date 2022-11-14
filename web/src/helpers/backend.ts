import axios from "axios";

import type { AxiosRequestConfig } from "axios";

import type { RespData } from "@/types";

import router from '@/router/index'

import { useUserInfo, msgStore } from '@/stores';


const instance = axios.create({
    baseURL: import.meta.env.VITE_BASE_API,
    timeout: 10000,
    withCredentials: true,
    headers:{
        "Content-type":"application/json"
    }
});


function getErrorMessage(error: unknown) {
    if (error instanceof Error) return error.message
    return String(error)
}


const backend = async <T = any>(config: AxiosRequestConfig): Promise<RespData<T>> => {

    try {
        const { data } = await instance.request<RespData<T>>(config)
        data.code === 0
            ? console.log(data.message) // 成功消息提示
            : console.error(data.message) // 失败消息提示
        return data
    } catch (error: any) {

        if (!error.response) {
            //TODO alert error
            //     ElMessageBox.confirm(`
            // <p>检测到请求错误</p>
            // <p>${error}</p>
            // `, '请求报错', {
            //         dangerouslyUseHTMLString: true,
            //         distinguishCancelAndClose: true,
            //         confirmButtonText: '稍后重试',
            //         cancelButtonText: '取消'
            //     })
            const message = '网络或接口异常，稍后重试'
            console.error(message) // 失败消息提示
            return {
                code: -1,
                message,
                data: null as any
            }
        }

        // const user = userStore()

        switch (error.response.status) {
            case 401:
                const userStore = useUserInfo()
                userStore.SetToken('')
                localStorage.clear()
                router.push({ name: 'Login', replace: true })
                break
            // case 403:

        }
        return {
            code: -1,
            message:"",
            data: null as any
        }
       
    }


}






export default backend;