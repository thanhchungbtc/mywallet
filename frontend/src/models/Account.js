import Api, {sleep} from "../core/api";

export default class Account {

  static async all() {
    try {
      const response = await Api.get("/accounts")
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async create(object) {
    try {
      const response = await Api.post('/accounts', object)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async findByID(id) {
    try {
      const response = await Api.get(`/accounts/${id}`)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async update(id, object) {
    try {
      const response = await Api.put(`/accounts/${id}`, object)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async delete(id) {
    try {
      const response = await Api.delete(`/accounts/${id}`)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

}