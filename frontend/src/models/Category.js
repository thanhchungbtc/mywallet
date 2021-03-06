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

  static async create(object) {
    try {
      const response = await Api.post('/categories', object)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async findByID(id) {
    try {
      const response = await Api.get(`/categories/${id}`)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async update(id, object) {
    try {
      const response = await Api.put(`/categories/${id}`, object)
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