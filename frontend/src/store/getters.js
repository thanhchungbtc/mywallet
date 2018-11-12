const getters = {
  isAuthenticated: state => !!state.auth.user,
}

export default getters