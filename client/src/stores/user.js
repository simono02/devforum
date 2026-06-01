import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import authService from '../services/auth.js'

export const useUserStore = defineStore('user', () => {
  const user = ref(authService.getUser())
  const token = ref(localStorage.getItem('token'))

  const isLoggedIn = computed(() => !!token.value)
  const isAdmin    = computed(() => user.value?.role === 'admin')
  const isMember   = computed(() => user.value?.role === 'member')

  function setAuth(data) {
    token.value = data.token
    user.value  = data.user
  }

  function clearAuth() {
    token.value = null
    user.value  = null
    authService.logout()
  }

  return { user, token, isLoggedIn, isAdmin, isMember, setAuth, clearAuth }
})