import Api from "../core/api";
import session from "../core/session";

export default class Auth {

  static async login(payload) {
    const response = await Api.post("/auth/login", payload)
    session.setToken(response.data.token)
    return response.data.user
  }

  static async register(payload) {
    const res = await Api.post('/auth/register', payload)
    session.setToken(res.data.token)
    return res.data.user
  }

  static async logout() {
    session.deleteToken()
  }

  static async verifyToken() {
    const token = session.getToken()

    if (!!token) {
      const res = await Api.post('/auth/verify-token', null, {
        headers: {Authorization: `Token ${token}`}
      })
      return res.data.user
    }
  }

}