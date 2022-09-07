<template>
  <div class="app-container bg-light">
    <nav v-show="userStore.user" class="navbar navbar-expand navbar-dark bg-dark">
      <div class="navbar-nav">
        <RouterLink to="/" class="nav-item nav-link">Home</RouterLink>
        <a @click="authStore.logout()" class="nav-item nav-link">Logout</a>
      </div>
    </nav>
    <div class="container pt-4 pb-4">
      <RouterView />
    </div>
  </div>
</template>


<script setup>

import { RouterLink, RouterView } from 'vue-router';
import { useAuthStore } from '@/stores';

const userStore = useAuthStore();

// export default {

//     name: 'DefaultLayout',

//     data: () => ({}),
//     head() {
//       return {
//         title: this.$config.head.title,
//         titleTemplate: this.$config.head.titleTemplate,
//       }
//     }

// }
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>
