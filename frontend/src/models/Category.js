import Api, {sleep} from "../core/api";

export default class Category {

  static async all() {
    try {
      const response = await Api.get("/categories")
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async create(account) {
    try {
      const response = await Api.post('/categories', account)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async findByID(id) {
    await sleep(1000)
    try {
      const response = await Api.get(`/categories/${id}`)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async update(id, account) {
    try {
      const response = await Api.put(`/categories/${id}`, account)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async delete(id) {
    try {
      const response = await Api.delete(`/categories/${id}`)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

}