import { createRouter, createWebHistory } from 'vue-router'



const router = createRouter({
  history: createWebHistory("/"),
  // history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'index',
      redirect: "/home",
      component: () => import('@/layouts/DefaultLayout.vue'),
      children: [
        {
          name: "Home",
          path: "/home",
          component: () => import("@/views/IndexView.vue")
        },
        {
          name: "Publisher",
          path: "/publisher",
          component: () => import("@/views/IndexView.vue")
        },
        {
          name: "Author",
          path: "/author",
          component: () => import("@/views/IndexView.vue")
        },
        {
          name: "/admin/books",
          path: "/admin/books",
          component: () => import("@/views/IndexView.vue")
        },
        {
          name: "用户管理",
          path: "/admin/users",
          component: () => import("@/views/IndexView.vue")
        },
        {
          name: "系统设置",
          path: "/admin/settings",
          component: () => import("@/views/IndexView.vue")
        }, {
          path: '/login',
          name: 'signin',
          component: () => import(/* webpackChunkName: "about" */ '@/views/LoginView.vue')
        }
      ]
    },

    

  ]
})

export default router


