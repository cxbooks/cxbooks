<template>
    <v-snackbar top :color="store.color" v-model="store.show">
        {{ store.msg }}
        <template v-slot:action="{ attrs }">
            <v-btn v-if="store.closeBtnType==='text'" text v-bind="attrs" @click="hideMessage">关闭</v-btn>
            <v-btn icon v-else v-bind="attrs" @click="hideMessage">
                <v-icon>mdi-close-circle</v-icon>
            </v-btn>
        </template>
    </v-snackbar>
</template>


<script setup lang='ts'>

import { msgStore } from '@/stores';


const store = msgStore()

const hideMessage = () => {
    store.show = false
}



</script>