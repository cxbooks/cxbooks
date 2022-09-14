import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'index',
    component: () => import(/* webpackChunkName: "about" */ '../views/IndexView.vue')
  },
  {
    path: '/login',
    name: 'signin',
    component: () => import(/* webpackChunkName: "about" */ '../views/LoginView.vue')
  }

]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  base: '/ui',
  routes
})


// router.beforeEach(async (to) => {
//   // redirect to login page if not logged in and trying to access a restricted page
//   const publicPages = ['/login'];
//   const authRequired = !publicPages.includes(to.path);
//   const auth = useAuthStore();

//   if (authRequired && !auth.user) {
//     auth.returnUrl = to.fullPath;
//     return '/login';
//   }
// });

export default router
