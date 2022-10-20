import { defineStore } from 'pinia';
import { ref, computed, watch } from 'vue'

export const msgStore = defineStore('message', () => {

    const msg = ref('')// 提示文字
    const show = ref( false) // 是否显示
    const color = ref( '' )// 颜色
    const closeBtn = ref( true) // 是否显示关闭按钮
    const closeBtnType = ref( 'icon') // icon or text 关闭按钮的类型
    const timeout = ref( 5000) // 自动关闭的时间
    const callback = ref( null) // 关闭后的回调


   
    const setHide = () => {
        show.value = false
    }


    let messageTimer:any = null;

    // const mutations = {
    //     SET_MESSAGE: (state, message: string) => {
    //         if (messageTimer) {
    //             clearTimeout(messageTimer);
    //         }
    //         if (show && timeout > 0) {
    //             // 自动关闭
    //             messageTimer = setTimeout(() => {
                    
    //                 if (callback && typeof callback === 'function') {
    //                     callback(state);
    //                 }
    //             }, timeout);
    //         }
    //     },
    // }

    const setSuccessMsg = (text: string) => {

        msg.value = text 
        color.value = 'success'
        show.value = true

    }

    const setErrorMsg = (text: string) => {

        msg.value = text
        color.value = 'error'
        show.value = true

    }

    const setWarningMsg = (text: string) => {

        msg.value = text
        color.value = 'warning'
        show.value = true

    }

    const setInfoMsg = (text: string) => {

        msg.value = text
        color.value = 'info'
        show.value = true

    }



    return {
        msg,
        show,
        color,
        closeBtn,
        closeBtnType,
        timeout,
        callback,
        setSuccessMsg,
        setErrorMsg,
        setWarningMsg,
        setInfoMsg

    }

})