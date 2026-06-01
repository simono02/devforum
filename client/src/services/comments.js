import api from './api.js'

const commentService = {
  async getByPost(postId) {
    const { data } = await api.get(`/api/posts/${postId}/comments`)
    return data
  },

  async create(postId, content) {
    const { data } = await api.post(`/api/posts/${postId}/comments`, { content })
    return data
  },

  async update(commentId, content) {
    const { data } = await api.put(`/api/comments/${commentId}`, { content })
    return data
  },

  async delete(commentId) {
    const { data } = await api.delete(`/api/comments/${commentId}`)
    return data
  },
}

export default commentService