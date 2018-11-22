import Api, {sleep} from "../core/api";

export default class Expense {

  static async all() {
    try {
      const response = await Api.get("/expenses")
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async create(object) {
    try {
      const payload = {
        ...object,
        use_date: new Date(object.use_date)
      }
      const response = await Api.post('/expenses', payload)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async findByID(id) {
    try {
      const response = await Api.get(`/expenses/${id}`)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async update(id, object) {
    try {
      const response = await Api.put(`/expenses/${id}`, object)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

  static async delete(id) {
    try {
      const response = await Api.delete(`/expenses/${id}`)
      return response.data
    } catch (e) {
      throw e.response.data
    }
  }

}