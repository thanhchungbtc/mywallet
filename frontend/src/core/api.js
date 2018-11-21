import axios from 'axios'
import Event from 'pubsub-js'
import session from './session'

const Api = axios.create({
  baseURL: "http://localhost:3000/api",
  headers: {
    common: {
      "Authorization": `Token ${session.getToken()}`
    }
  }
})

Api.interceptors.request.use(function (config) {
  Event.publish('ajax.start', config)
  return config
}, function (error) {
  return Promise.reject(error)
})

Api.interceptors.response.use(function (response) {
  Event.publish('ajax.end', response)
  return response
}, function (error) {
  if (console && console.log) {
    console.log(error)
  }
  if (error.response) {
    return Promise.reject(error.response.data)

  }
  return Promise.reject(error)
  /*  let errorMessage = 'An error occurred'
    let code = error.code
    if (error.response && error.response.data) {
      let data = error.response.data
      code = data.code
      errorMessage = data.message ? data.message : data.error
    }
    Event.publish('ajax.end')
    Event.publish('alert.error', errorMessage)
    if (code === 401) {
      window.location = '/'
    }
    return Promise.reject(error)*/
})

export const sleep = (ms) => {
  return new Promise(resolve => setTimeout(resolve, ms))
}

export default Api
