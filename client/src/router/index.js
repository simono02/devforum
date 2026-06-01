import { createRouter, createWebHistory } from 'vue-router'
import authService from '../services/auth.js'

const routes = [
  { path: '/', component: () => import('../app/index/index.vue') },
  { path: '/login', component: () => import('../app/login/login.vue') },
  { path: '/register', component: () => import('../app/register/register.vue') },

  // ── Member (role: member) ─────────────────────────────
  {
    path: '/dashboard',
    component: () => import('../app/member/layout.vue'),
    meta: { requiresAuth: true },
    children: [
      { path: '',         component: () => import('../app/member/dashboard/dashboard.vue') },
      { path: 'posts',    component: () => import('../app/member/posts/posts.vue') },
      { path: 'profile',  component: () => import('../app/member/profile/profile.vue') },
    ]
  },

  // ── Admin (role: admin) ───────────────────────────────
  {
    path: '/admin',
    component: () => import('../app/admin/layout.vue'),
    meta: { requiresAuth: true, requiresAdmin: true },
    children: [
      { path: '',       component: () => import('../app/admin/dashboard/dashboard.vue') },
      { path: 'posts',  component: () => import('../app/admin/posts/posts.vue') },
      { path: 'users',  component: () => import('../app/admin/users/users.vue') },
    ]
  },

  { path: '/:pathMatch(.*)*', redirect: '/' }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const loggedIn = authService.isLoggedIn()
  const user     = authService.getUser()

  // Not logged in → login
  if (to.meta.requiresAuth && !loggedIn) return next('/login')

  // Non-admin trying to reach admin pages → member dashboard
  if (to.meta.requiresAdmin && user?.role !== 'admin') return next('/dashboard')

  // Admin trying to reach member dashboard → admin dashboard
  if (to.path.startsWith('/dashboard') && user?.role === 'admin') return next('/admin')

  // Logged-in users shouldn't visit login/register
  if ((to.path === '/login' || to.path === '/register') && loggedIn) {
    return next(user?.role === 'admin' ? '/admin' : '/dashboard')
  }

  next()
})

export default router