import Vue from 'vue'
import paths from './paths'
import Router from 'vue-router'
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import store from "../store";

Vue.use(Router);

const router = new Router({
  base: '/',
  mode: 'history',
  linkActiveClass: 'active',
  routes: paths
});
// router gards
router.beforeEach((to, from, next) => {
  NProgress.start();

  if (to.meta.public === true) {
    if (!store.getters.isAuthenticated ){
      return next()
    }
    return next({name: 'dashboard'})
  }

  // non public router
  if (store.getters.isAuthenticated ){
    return next()
  }
  return next({name: 'login'})

});

router.afterEach((to, from) => {
  NProgress.done();
});

export default router;
