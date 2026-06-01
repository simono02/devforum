import api from './api.js'

const authService = {
  async register(payload) {
    const { data } = await api.post('/api/auth/register', {
      full_name: payload.fullName,
      username:  payload.username,
      email:     payload.email,
      password:  payload.password,
    })
    return data
  },

  async login(payload) {
    const { data } = await api.post('/api/auth/login', {
      identifier: payload.identifier,
      password:   payload.password,
    })
    // Persist token + user
    localStorage.setItem('token', data.token)
    localStorage.setItem('user', JSON.stringify(data.user))
    return data
  },

  logout() {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  },

  getUser() {
    const raw = localStorage.getItem('user')
    return raw ? JSON.parse(raw) : null
  },

  isLoggedIn() {
    return !!localStorage.getItem('token')
  },

  isAdmin() {
    return this.getUser()?.role === 'admin'
  }
}

export default authService