const state = {
  snackbar: {
    show: false,
    text: '',
    color: '',
  }
}

const mutations = {

}

const actions = {
  showSnackbar({state}, val) {
    state.snackbar.show = val
  },

  error({state}, message) {
    state.snackbar = {
      show: true,
      text: message,
      color: 'error',
    }
  },
  success({state}, message) {
    state.snackbar = {
      show: true,
      text: message,
      color: 'success',
    }
  },
  info({state}, message) {
    state.snackbar = {
      show: true,
      text: message,
      color: 'info',
    }
  },


}

export default {
  state, mutations, actions,
}