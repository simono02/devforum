import api from './api.js'

const adminService = {
  // ── Users ──────────────────────────────────────────────
  async getUsers() {
    const { data } = await api.get('/api/users')
    return data
  },

  async getStats() {
    const { data } = await api.get('/api/users/stats')
    return data
  },

  async updateUser(id, payload) {
    const { data } = await api.put(`/api/users/${id}`, payload)
    return data
  },

  async deleteUser(id) {
    const { data } = await api.delete(`/api/users/${id}`)
    return data
  },

  // ── Posts (reuse post endpoints) ───────────────────────
  async getPosts() {
    const { data } = await api.get('/api/posts')
    return data
  },

  async deletePost(id) {
    const { data } = await api.delete(`/api/posts/${id}`)
    return data
  },
}

export default adminService