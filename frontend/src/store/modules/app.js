const state = {
  drawer: true,
}

const mutations = {

}

const actions = {
  APP_DRAWER_TOGGLE({state}) {
    state.drawer = !state.drawer
  }
}

export default {
  state, mutations, actions,
}