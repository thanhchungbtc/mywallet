import store from '../store'

const requireGuest = (to, from, next) => {
  console.log('requireGuest')
  return !store.getters.isAuthenticated ? next() : next("/")
}
const requireAuth = (to, from, next) => {
  console.log('requireAuth')
  store.getters.isAuthenticated ? next() : next("/login")
}

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
    component: () => import(`@/views/Login.vue`),
  },
  {
    path: '/register',
    meta: {
      public: true,
    },
    name: 'register',
    component: () => import(`@/views/Register.vue`),
  },
  {
    path: '/',
    meta: {},
    name: 'root',
    redirect: {
      name: 'dashboard'
    }
  },
  {
    path: '/dashboard',
    meta: {breadcrumb: true},
    name: 'dashboard',
    component: () => import(`@/views/Dashboard.vue`),
  },
  {
    path: '/categories',
    meta: {breadcrumb: true},
    name: 'category_list',
    component: () => import(`@/views/Category/Category.vue`),
  },
  {
    path: '/categories/new',
    meta: {breadcrumb: true},
    name: 'category_create',
    component: () => import(`@/views/Category/CategoryForm.vue`),
  },
  {
    path: '/categories/:id',
    meta: {breadcrumb: true},
    name: 'category_edit',
    component: () => import(`@/views/Category/CategoryForm.vue`),
  },
  {
    path: '/expenses',
    meta: {breadcrumb: true},
    name: 'expense',
    component: () => import(`@/views/Expense/Expense.vue`),
  },
  {
    path: '/expenses/create',
    meta: {breadcrumb: true},
    name: 'expense_create',
    component: () => import(`@/views/Expense/ExpenseForm.vue`),
  },

];
