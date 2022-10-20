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
        <template v-for="(item, i) in sidebarMenu" :key="i">
          <!-- ---------------------------------------------- -->
          <!---Single Item-->
          <!-- ---------------------------------------------- -->
          <!-- <v-subheader v-if="item.title">{{ item.title }}</v-subheader> -->
          <v-list-item :to="item.to" rounded="lg" class="mb-1">
            
              
            <v-list-item-action class="mt-1 mb-1 mr-2" v-if="item.count">
              <v-list-item-avatar start class="v-list-item-avatar--start">
                <v-icon class="feather-sm v-icon v-icon--size-default">{{
                  item.icon
                  }}</v-icon>
              
              </v-list-item-avatar>
                  <v-list-item-title v-text="item.title"></v-list-item-title>
                
              <v-chip small outlined>{{ item.count }}</v-chip>
            </v-list-item-action>
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

const sidebarMenu = [
  {
    title: "Dashboard",
    icon: "mdi-view-dashboard-outline",
    to: "/dashboard",
    count:10
  },
  {
    title: "Alert",
    icon: "mdi-alert-circle-outline",
    to: "/ui-components/alert",
    count: 10
  },
  {
    title: "Buttons",
    icon: "mdi-radiobox-blank",
    to: "/ui-components/buttons",
    count: 10
  },
  {
    title: "Cards",
    icon: "mdi-card-outline",
    to: "/ui-components/cards",
    count: 10
  },
  {
    title: "Menus",
    icon: "mdi-form-dropdown",
    to: "/ui-components/menus",
    count: 10
  },
  {
    title: "Tables",
    icon: "mdi-table",
    to: "/ui-components/tables",
    count: 10
  },
];
</script>