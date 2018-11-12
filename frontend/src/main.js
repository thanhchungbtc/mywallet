import Vue from 'vue'
import App from './App.vue'
import store from './store'
import Vuetify from 'vuetify'

import 'material-design-icons-iconfont/dist/material-design-icons.css'
import 'vuetify/dist/vuetify.min.css'
import 'font-awesome/css/font-awesome.css';
import router from './router'
import moment from "moment";


Vue.config.productionTip = false

Vue.use(Vuetify, {
  options: {
    themeVariations: ['primary', 'secondary', 'accent'],
    extra: {
      mainToolbar: {
        color: 'primary',
      },
      sideToolbar: {
      },
      sideNav: 'primary',
      mainNav: 'primary lighten-1',
      bodyBg: '',
    }
  }
});

Vue.filter('formatDate', function (value) {
  if (value) {
    return moment(value).format('YYYY-MM-DD')
  }
})

store.dispatch('auth/verify')
  .then(() => {
    new Vue({
      store,
      router,
      render: h => h(App)
    }).$mount('#app')
  })
