<template>
  <div>
    <!-- ---------------------------------------------- -->
    <!---Logo part -->
    <!-- ---------------------------------------------- -->
    <div class="pa-4">
      <LogoDark />
    </div>
    <!-- ---------------------------------------------- -->
    <!---Navigation -->
    <!-- ---------------------------------------------- -->
    <div class="scrollnavbar">
      <v-list class="pa-4" color="transparent">
        <!-- ---------------------------------------------- -->
        <!---Menu Loop -->
        <!-- ---------------------------------------------- -->
        <v-list-item-group  active-class="border" color="indigo">导航</v-list-item-group>
        <template v-for="(item, i) in sidebarMenu" :key="i">
          <!-- ---------------------------------------------- -->
          <!---Single Item-->
          <!-- ---------------------------------------------- -->
          <v-list-item :to="item.to" :prepend-icon="item.icon" :title="item.title" rounded="lg" class="mb-1">
            <template v-slot:append>
              <v-badge color="info" :content="item.count" inline></v-badge>
            </template>
          
          </v-list-item>
        </template>
      </v-list>
    </div>
    <div class="pa-4 ma-4 bg-light-primary rounded-lg text-center">
      <img src="@/assets/images/sidebar-buynow-bg.svg" />
      <h4 class="font-weight-regular mb-3">Get Template for Free</h4>
      <v-btn class="mb-2" href="https://www.wrappixel.com/templates/flexy-vuejs-admin-free/" block>Download Free</v-btn>
      <v-btn
        color="info"
        href="https://www.wrappixel.com/templates/flexy-vuetify-dashboard/"
        block
        >Check Pro</v-btn
      >
    </div>
  </div>
</template>
<script setup lang="ts">
  
import LogoDark from "@/components/LogoDark.vue";

import { useUserInfo } from '@/stores';


const userInfo = useUserInfo();

const stats = userInfo.bookStats;


const sidebarMenu = [
  {
    title: "分类导航",
    icon: "mdi-view-dashboard-outline",
    to: "/nav",
    count: stats.total,
  },
  {
    title: "出版社",
    icon: "mdi-alert-circle-outline",
    to: "/publishers",
    count: stats.publisher,
  },
  {
    title: "作者",
    icon: "mdi-radiobox-blank",
    to: "/authors",
    count: stats.author,
  },
  {
    title: "标签",
    icon: "mdi-card-outline",
    to: "/tags",
    count: stats.author,
  }
];
</script>