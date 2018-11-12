import axios from 'axios'
import Event from 'pubsub-js'
import Session from './session'

const session = new Session(window.localStorage)
const Api = axios.create({
  baseURL: "http://localhost:3000/api",
  headers: {
    common: {
      "Authorization": session.getToken()
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

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms))
}

const Auth = {
  login: async (payload) => {
    const response = await Api.post("/auth/login", payload)
    await sleep(1000)

    session.setToken(response.data.token)
    return response.data
  },

  register: async (payload) => {
    const response = await Api.post('/auth/register', payload)
    session.setToken(response.data.user.token)
    return response.data
  },

  logout: async () => {
    session.deleteToken()
  },

  verifyToken: async () => {
    const token = session.getToken()
    console.log('verify token', token)

    if (!!token) {
      const response = await Api.post('/auth/verify-token', null, {
        headers: {Authorization: `Token ${token}`}
      })
      return response.data
    }
    return null
    // const response = await Api.post('/auth/verify-token', null, {
    //   headers: {Authorization: `Token ${session.getToken()}`}
    // })
    // const response = await Api.post('/auth/verify-token')
    // session.setToken(response.data.token)
    // return response.data
  }
}

const Account = {
  async all() {
    try {
      const url = urlForPath("/api/v1/accounts")
      const response = await Api.get(url)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },
  async create(account) {
    try {
      const url = urlForPath("/api/v1/accounts")
      const response = await Api.post(url, account)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },
  async findByID(id) {
    try {
      const url = urlForPath((`/api/v1/accounts/${id}`))
      const response = await Api.get(url)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },

  async update(id, account) {
    try {
      const url = urlForPath((`/api/v1/accounts/${id}`))
      const response = await Api.put(url, account)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },
  async delete(id) {
    try {
      const url = urlForPath((`/api/v1/accounts/${id}`))
      const response = await Api.delete(url)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }
}

const Category = {
  async all() {
    try {
      const url = urlForPath("/api/v1/categories")
      const response = await Api.get(url)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },
  async create(account) {
    try {
      const url = urlForPath("/api/v1/categories")
      const response = await Api.post(url, account)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },
  async findByID(id) {
    try {
      const url = urlForPath((`/api/v1/categories/${id}`))
      const response = await Api.get(url)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },

  async update(id, account) {
    try {
      const url = urlForPath((`/api/v1/categories/${id}`))
      const response = await Api.put(url, account)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },
  async delete(id) {
    try {
      const url = urlForPath((`/api/v1/categories/${id}`))
      const response = await Api.delete(url)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }
}

const Expense = {
  async fetchAll() {
    try {
      const url = urlForPath("/api/v1/expenses")
      const response = await Api.get(url)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },
  async create(account) {
    try {
      const url = urlForPath("/api/v1/expenses")
      const response = await Api.post(url, account)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },
  async fetchById(id) {
    try {
      const url = urlForPath((`/api/v1/expenses/${id}`))
      const response = await Api.get(url)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },

  async update(id, account) {
    try {
      const url = urlForPath((`/api/v1/expenses/${id}`))
      const response = await Api.put(url, account)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  },
  async delete(id) {
    try {
      const url = urlForPath((`/api/v1/expenses/${id}`))
      const response = await Api.delete(url)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }
}

export default Api
export {
  Auth,
}
// export default {
//   Auth,
//   Account,
//   Category,
//   Expense,
//   post: (url, data) => Api.post(url, data),
//   get: (url) => {
//     if (!/^https?/.test(url)) {
//       url = `http://${url}`
//     }
//     return Api.get(url)
//   },
//   sleep,
// }