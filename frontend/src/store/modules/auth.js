import {Auth} from "../../core/api";

const state = {
  user: null,
}

const mutations = {
  loginSuccess(state, user) {
    state.user = user
  },
  tokenValid(state, user) {
    state.user = user
  },
}

const actions = {
  login({commit}, credentials) {
    return Auth.login(credentials).then(user => {
      commit('loginSuccess', user)
    })
  },

  async verify({commit, dispatch}) {
    try {
      const res = await Auth.verifyToken()
      commit('tokenValid', res)
    } catch (e) {
      console.log(e)
      dispatch('auth/logout')
    }
  },
  async logout({commit}) {
    return Auth.logout()
  },
}

export default {
  state, mutations, actions,
}