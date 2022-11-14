<template>
<v-app>
    <div class="login">
        <v-row justify="center" class="fill-center">
            <v-col xs="12" sm="8" md="4">
                <v-card v-if="show_login" class="elevation-12">
                    <v-toolbar dark color="primary">
                        <v-toolbar-title>欢迎访问</v-toolbar-title>
                    </v-toolbar>
                    <v-card-text>
                        <v-form @submit.prevent="do_login">
                            <v-text-field prepend-icon="person" v-model="user.account" label="用户名" type="text">
                            </v-text-field>
                            <v-text-field prepend-icon="lock" v-model="user.password" label="密码" type="password" id="password">
                            </v-text-field>
                            <p class="text-right">
                                <a @click="show_login = !show_login"> 忘记密码? </a>
                            </p>
                            <div align="center">
                                <v-btn type="submit" large rounded color="primary">登录</v-btn>
                            </div>
                        </v-form>
                    </v-card-text>
    
                    <v-card-text v-if="0 > 0">
                        <v-divider></v-divider>
                        <div align="center">
                            <br />
                            <small>使用社交网络账号登录</small>
                            <br />
                            <!-- <template v-for="s in socials" :key="s.text">
                                <v-btn small outlined :href="'/auth/login/' + s.value">{{ s.text }}</v-btn>
                                &nbsp;
                            </template> -->
                        </div>
                    </v-card-text>
                </v-card>
    
                <v-card v-else class="elevation-12">
                    <v-toolbar dark color="red">
                        <v-toolbar-title>重置密码</v-toolbar-title>
                    </v-toolbar>
                    <v-card-text v-if="!show_login">
                        <v-form @submit.prevent="do_reset">
                            <v-text-field prepend-icon="person" v-model="user.account" label="用户名" type="text">
                            </v-text-field>
                            <v-text-field prepend-icon="email" v-model="email" label="注册邮箱" type="text"
                                autocomplete="old-email"></v-text-field>
                        </v-form>
                        <div align="center">
                            <v-btn rounded color="" class="mr-5" @click="show_login = !show_login">返回</v-btn>
                            <v-btn rounded dark color="red" @click="do_reset">重置密码</v-btn>
                        </div>
                    </v-card-text>
    
                </v-card>
            </v-col>
        </v-row>
    </div>
</v-app>
</template>

<script setup lang='ts'>

import { defineComponent,ref, reactive  } from 'vue';
import router from '@/router';

import { useUserInfo,msgStore } from '@/stores';
// import User from '@/types/user';
import type {RespData} from '@/types';



let user = reactive({
    account: "",
    password: "",
});

const show_login = ref(true)// 提示文字

const email = ref('') // 是否显示关闭按钮
const closeBtnType = ref('icon') // icon or text 关闭按钮的类型
const timeout = ref(5000) // 自动关闭的时间
const callback = ref(null) // 关闭后的回调

const do_reset = function () {
    var data = new URLSearchParams();
    // data.append("username", this.account);
    // data.append("email", this.email);
    // this.$backend("/user/reset", {
    //     method: "POST",
    //     body: data,
    //     }).then((rsp) => {
    //         if (rsp.err == "ok") {
    //         this.alert.type = "success";
    //         this.alert.msg = "重置成功！请查阅密码通知邮件。";
    //     } else {
    //         this.alert.type = "error";
    //         this.alert.msg = rsp.msg;
    //     }
    // });
}

const do_login = async () => {

    const userInfo = useUserInfo()

    const code = await userInfo.Login(user.account, user.password) 
    
  
    if (code) {
        console.log(`跳转首页`);
        router.push('/home');
    }
    //else  弹窗错误提示

    
}




    
   




</script>

<style>
.fill-center {
    margin-top: 6%;
}
</style>
