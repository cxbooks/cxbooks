<template>
    <v-footer class="bg-indigo-lighten-1 text-center d-flex flex-column">

    <div>
        <v-btn v-for="link in footerLinks"  color="white" class="mx-4" :icon="link.icon" variant="text" target="_blank" :href="link.to"></v-btn>
    </div>
    <div class="pt-0">
        本站基于Vue3 Golang构建，感谢开源界的力量。所有资源搜集于互联网，如有侵权请邮件联系。
    </div>
    <v-divider></v-divider>
    </v-footer>
</template>

<script setup lang="ts">


const footerLinks = [
    {
        title: "Github",
        icon: "mdi-github",
        to: "https://github.com/cxbooks/cxbooks"
    },
    {
        title: "Docker",
        icon: "mdi-docker",
        to: "https://hub.docker.com/r/cxbooks/cxbooks"
    },
    {
        title: "Project",
        icon: "mdi-twitter",
        to: "http://cxbooks.org"
    },
    {
        title: "Github",
        icon: "mdi-github",
        to: "https://github.com/cxbooks/cxbooks"
    }
]


</script>

