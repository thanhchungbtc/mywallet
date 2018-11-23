import Api from "../core/api";
import moment from "moment";

export default class Expense {

  static map(item) {
    return {
      ...item,
      use_date: moment(item.use_date).format('YYYY-MM-DD')
    }
  }

  static async all(filter) {
    try {
      const response = await Api.get("/expenses", {
        params: filter
      })
      return response.data.map(this.map)
    }catch (e) {
      throw e.response.data
    }
  }

  static async findByID(id) {
    try {
      const response = await Api.get(`/expenses/${id}`)
      return this.map(response.data)
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
      return this.map(response.data)
    } catch (e) {
      throw e.response.data
    }
  }

  static async update(id, object) {
    try {
      const payload = {
        ...object,
        use_date: new Date(object.use_date)
      }
      const response = await Api.put(`/expenses/${id}`, payload)
      return this.map(response.data)
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