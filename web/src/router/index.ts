import { createRouter, createWebHistory } from 'vue-router'
import { useUserInfo } from '@/stores';


const router = createRouter({
  // history: createWebHistory("/"),
  history: createWebHistory(import.meta.env.BASE_URL),
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
          name: "Book",
          path: "/book/:id",
          component: () => import("@/views/BookView.vue")
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
        }
      ]
    }, {
      path: '/login',
      name: 'signin',
      component: () => import(/* webpackChunkName: "about" */ '@/views/LoginView.vue')
    }

    

  ]
})

router.beforeEach(async (to) => {
  // redirect to login page if not logged in and trying to access a restricted page
  const publicPages = ['/login', '/login2'];
  const authRequired = !publicPages.includes(to.path);
  const auth = useUserInfo();

  console.log(`is login:`,auth.token)

  if (authRequired && !auth.token) {
    // auth.returnUrl = to.fullPath;
    return '/login';
  }
});

export default router


