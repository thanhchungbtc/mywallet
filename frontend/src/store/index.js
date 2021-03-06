import Vue from 'vue'
import Vuex from 'vuex'

import auth from './modules/auth'
import app from './modules/app'
import getters from './getters'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    app: {
      namespaced: true,
      ...app,
    },
    auth:{
      namespaced: true,
      ...auth,
    },
  },
  getters,
})

export default store
