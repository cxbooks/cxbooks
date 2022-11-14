<template>
  <v-app>
    <MessageVue />
    <!-- ---------------------------------------------- -->
    <!---Sidebar -->
    <!-- ---------------------------------------------- -->
    <v-navigation-drawer
      left
      :permanent="$vuetify.display.mdAndUp"
      elevation="10"
      app
      :temporary="$vuetify.display.mdAndDown"
      v-model="drawer"
      expand-on-hover
      v-if="store.token"
    >
      <SidebarVue />
    </v-navigation-drawer>
    <!-- ---------------------------------------------- -->
    <!---Header -->
    <!-- ---------------------------------------------- -->
    <v-app-bar elevation="0" class="v-topbar">
      <v-app-bar-nav-icon class="hidden-md-and-up" @click="drawer = !drawer" />
      <v-spacer />
      <!-- ---------------------------------------------- -->
      <!-- User Profile -->
      <!-- ---------------------------------------------- -->
      <HeaderVue />
    </v-app-bar>

    <!-- ---------------------------------------------- -->
    <!---Page Wrapper -->
    <!-- ---------------------------------------------- -->
    <v-main>
      <v-container fluid class="page-wrapper">
        <RouterView />
      </v-container>
    </v-main>
  <FooterVue />
  </v-app>
</template>

<script setup lang="ts">
import { RouterView } from "vue-router";
import { ref, onMounted } from "vue";
import SidebarVue from "@/components/Sidebar.vue";
import HeaderVue from "@/components/Header.vue";
import MessageVue from "@/components/Message.vue";
import FooterVue from "@/components/Footer.vue";

import { useUserInfo } from '@/stores';

const drawer = ref(undefined || true);
const innerW = window.innerWidth;
// const 

const store = useUserInfo();



onMounted(() => {
  if (innerW < 950) {
    drawer.value = !drawer.value;
  }

});

</script>


