import api from './api.js'

const postService = {
  // Get all posts (public)
  async getAll() {
    const { data } = await api.get('/api/posts')
    return data
  },

  // Get my posts (authenticated)
  async getMine() {
    const { data } = await api.get('/api/posts/mine')
    return data
  },

  // Get single post
  async getById(id) {
    const { data } = await api.get(`/api/posts/${id}`)
    return data
  },

  // Create post
  async create(payload) {
    const { data } = await api.post('/api/posts', {
      title:   payload.title,
      content: payload.content,
      tag:     payload.tag,
    })
    return data
  },

  // Update post
  async update(id, payload) {
    const { data } = await api.put(`/api/posts/${id}`, {
      title:   payload.title,
      content: payload.content,
      tag:     payload.tag,
    })
    return data
  },

  // Delete post
  async delete(id) {
    const { data } = await api.delete(`/api/posts/${id}`)
    return data
  },

  // Upvote post
  async upvote(id) {
    const { data } = await api.post(`/api/posts/${id}/upvote`)
    return data
  },
}

export default postService