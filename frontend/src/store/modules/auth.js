import {Auth} from "../../core/api";

const state = {
  user: null,
}

const mutations = {
  loginSuccess(state, user) {
    state.user = user
  },
  verifySuccess(state, user) {
    state.user = user
  }
}

const actions = {
  login({commit}, credentials) {
    return Auth.login(credentials).then(user => {
      commit('loginSuccess', user)
    })
  },

  async verify({commit}) {
    const res = await Auth.verifyToken()
    commit('verifySuccess', res)
  },
  async logout({commit}) {
    return await Auth.logout()
  },
}

export default {
  state, mutations, actions,
}