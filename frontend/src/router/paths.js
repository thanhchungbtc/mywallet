export default [

  {
    path: '*',
    meta: {
      public: true,
    },
    redirect: {
      path: '/404'
    }
  },
  {
    path: '/404',
    meta: {
      public: true,
    },
    name: 'notFound',
    component: () => import(
      /* webpackChunkName: "routes" */
      /* webpackMode: "lazy-once" */
      `@/views/NotFound.vue`
      )
  },
  {
    path: '/403',
    meta: {
      public: true,
    },
    name: 'accessDenied',
    component: () => import(
      /* webpackChunkName: "routes" */
      /* webpackMode: "lazy-once" */
      `@/views/Deny.vue`
      )
  },
  {
    path: '/500',
    meta: {
      public: true,
    },
    name: 'serverError',
    component: () => import(
      /* webpackChunkName: "routes" */
      /* webpackMode: "lazy-once" */
      `@/views/Error.vue`
      )
  },
  {
    path: '/login',
    meta: {
      public: true,
    },
    name: 'login',
    component: () => import(
      /* webpackChunkName: "routes" */
      /* webpackMode: "lazy-once" */
      `@/views/Login.vue`
      )
  },
  {
    path: '/',
    meta: { },
    name: 'root',
    redirect: {
      name: 'dashboard'
    }
  },
  {
    path: '/dashboard',
    meta: { breadcrumb: true },
    name: 'dashboard',
    component: () => import(
      `@/views/Dashboard.vue`
      )
  },
  {
    path: '/categories',
    meta: { breadcrumb: true },
    name: 'category',
    component: () => import(
      `@/views/Category.vue`
      )
  },

];
