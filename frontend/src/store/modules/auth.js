import Auth from "../../models/Auth";

const state = {
  user: null,
}

const mutations = {
  loginSuccess(state, user) {
    state.user = user
  },

  verifySuccess(state, user) {
    state.user = user
  },

  verifyFailed(state) {
    state.user = user
  },
}

const actions = {
  async login({commit}, credentials) {
    const res = await Auth.login(credentials)
    commit('loginSuccess', res)
  },

  async verify({commit, dispatch}) {
    try {
      const res = await Auth.verifyToken()
      commit('loginSuccess', res)
    } catch (e) {
      dispatch('logout')
    }
  },

  async logout({commit}) {
    await Auth.logout()
  },
}

export default {
  state, mutations, actions,
}