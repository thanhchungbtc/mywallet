import Api from './api'

export default class Session {
  constructor(storage) {
    this.storage = storage
  }

  setToken(token) {
    console.log('setToken')
    this.storage.setItem('token', token)
    Api.defaults.headers.common['Authorization'] = `Token ${token}`
  }

  getToken() {
    return this.storage.getItem('token')
  }

  deleteToken() {
    this.storage.removeItem('token')
    Api.defaults.headers.common['Authorization'] = ``
  }
}