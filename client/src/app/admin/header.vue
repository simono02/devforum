<template>
  <header class="topbar">
    <!-- Left -->
    <div class="topbar-left">
      <button class="menu-btn" @click="$emit('toggle-sidebar')" title="Toggle sidebar">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
          <line x1="3" y1="6" x2="21" y2="6"/>
          <line x1="3" y1="12" x2="21" y2="12"/>
          <line x1="3" y1="18" x2="21" y2="18"/>
        </svg>
      </button>
      <div class="breadcrumb">
        <span class="admin-badge">Admin</span>
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" class="bc-sep"><polyline points="9 18 15 12 9 6"/></svg>
        <span class="page-title">{{ pageTitle }}</span>
      </div>
    </div>

    <!-- Right -->
    <div class="topbar-right">

      <!-- Theme toggle -->
      <button class="icon-btn theme-toggle" @click="themeStore.toggle()" :title="isDark ? 'Switch to light mode' : 'Switch to dark mode'">
        <transition name="icon-swap" mode="out-in">
          <svg v-if="isDark" key="sun" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
            <circle cx="12" cy="12" r="5"/>
            <line x1="12" y1="1" x2="12" y2="3"/>
            <line x1="12" y1="21" x2="12" y2="23"/>
            <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"/>
            <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"/>
            <line x1="1" y1="12" x2="3" y2="12"/>
            <line x1="21" y1="12" x2="23" y2="12"/>
            <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"/>
            <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"/>
          </svg>
          <svg v-else key="moon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
            <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"/>
          </svg>
        </transition>
      </button>

      <!-- Notifications -->
      <button class="icon-btn notif-btn" title="Notifications">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
          <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"/>
          <path d="M13.73 21a2 2 0 0 1-3.46 0"/>
        </svg>
        <span class="notif-dot"></span>
      </button>

      <!-- Divider -->
      <div class="divider"></div>

      <!-- User dropdown -->
      <div class="user-menu" ref="menuRef">
        <button class="user-trigger" @click="menuOpen = !menuOpen">
          <div class="user-av">{{ initials }}</div>
          <div class="user-info" v-if="!isMobile">
            <span class="user-name">{{ user?.full_name }}</span>
            <span class="user-role">Administrator</span>
          </div>
          <svg class="chevron" :class="{ open: menuOpen }" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
            <polyline points="6 9 12 15 18 9"/>
          </svg>
        </button>

        <transition name="dropdown">
          <div class="dropdown" v-if="menuOpen">
            <div class="dropdown-header">
              <div class="dh-av">{{ initials }}</div>
              <div>
                <div class="dh-name">{{ user?.full_name }}</div>
                <div class="dh-email">{{ user?.email }}</div>
              </div>
            </div>
            <div class="dropdown-divider"></div>
            <router-link class="dropdown-item" to="/admin" @click="menuOpen = false">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><rect x="3" y="3" width="7" height="7"/><rect x="14" y="3" width="7" height="7"/><rect x="14" y="14" width="7" height="7"/><rect x="3" y="14" width="7" height="7"/></svg>
              Dashboard
            </router-link>
            <router-link class="dropdown-item" to="/admin/users" @click="menuOpen = false">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/></svg>
              Manage Users
            </router-link>
            <div class="dropdown-divider"></div>
            <button class="dropdown-item danger" @click="handleLogout">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/><polyline points="16 17 21 12 16 7"/><line x1="21" y1="12" x2="9" y2="12"/></svg>
              Sign out
            </button>
          </div>
        </transition>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '../../stores/user.js'
import { useThemeStore } from '../../stores/theme.js'

defineEmits(['toggle-sidebar'])

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const themeStore = useThemeStore()

const menuOpen = ref(false)
const menuRef = ref(null)
const isMobile = ref(window.innerWidth < 640)

const user = computed(() => userStore.user)
const isDark = computed(() => themeStore.theme === 'dark')
const initials = computed(() => {
  const name = user.value?.full_name || 'A'
  return name.split(' ').map(n => n[0]).join('').slice(0, 2).toUpperCase()
})

const titles = {
  '/admin': 'Dashboard',
  '/admin/posts': 'Posts',
  '/admin/users': 'Users',
}
const pageTitle = computed(() => titles[route.path] || 'Admin')

function handleLogout() {
  menuOpen.value = false
  userStore.clearAuth()
  router.push('/login')
}

function handleClickOutside(e) {
  if (menuRef.value && !menuRef.value.contains(e.target)) {
    menuOpen.value = false
  }
}

function handleResize() {
  isMobile.value = window.innerWidth < 640
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  window.addEventListener('resize', handleResize)
})
onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
  window.removeEventListener('resize', handleResize)
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Syne:wght@400;600;700;800&display=swap');
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

.topbar {
  height: 56px;
  background: var(--topbar-bg);
  border-bottom: 1px solid var(--border);
  display: flex; align-items: center; justify-content: space-between;
  padding: 0 1.5rem;
  position: sticky; top: 0; z-index: 100;
  transition: background 0.25s ease, border-color 0.25s ease;
  font-family: 'Syne', sans-serif;
  flex-shrink: 0;
}

/* Left */
.topbar-left { display: flex; align-items: center; gap: 0.875rem; }
.menu-btn {
  background: none; border: none; cursor: pointer;
  color: var(--text-ghost); padding: 0.4rem; border-radius: 6px;
  display: flex; align-items: center; transition: color 0.2s, background 0.2s;
}
.menu-btn svg { width: 18px; height: 18px; }
.menu-btn:hover { color: var(--text-muted); background: var(--bg-elevated); }

.breadcrumb { display: flex; align-items: center; gap: 0.5rem; }
.admin-badge {
  font-size: 0.62rem; font-weight: 700; text-transform: uppercase; letter-spacing: 0.07em;
  background: var(--bg-elevated); color: var(--text-ghost);
  border: 1px solid var(--border-mid); padding: 0.18rem 0.55rem; border-radius: 4px;
}
.bc-sep { width: 14px; height: 14px; color: var(--text-ghost); }
.page-title { font-size: 0.92rem; font-weight: 700; color: var(--text-secondary); letter-spacing: -0.2px; }

/* Right */
.topbar-right { display: flex; align-items: center; gap: 0.5rem; }

.icon-btn {
  background: none; border: none; cursor: pointer;
  color: var(--text-ghost); padding: 0.45rem; border-radius: 8px;
  display: flex; align-items: center; justify-content: center;
  position: relative; transition: color 0.2s, background 0.2s;
}
.icon-btn svg { width: 18px; height: 18px; }
.icon-btn:hover { color: var(--text-muted); background: var(--bg-elevated); }

.theme-toggle { font-size: 0; }

.notif-dot {
  position: absolute; top: 6px; right: 6px;
  width: 7px; height: 7px; border-radius: 50%;
  background: #e05555; border: 1.5px solid var(--topbar-bg);
}

.divider { width: 1px; height: 20px; background: var(--border); margin: 0 0.25rem; }

/* User menu */
.user-menu { position: relative; }
.user-trigger {
  display: flex; align-items: center; gap: 0.6rem;
  background: none; border: none; cursor: pointer;
  padding: 0.3rem 0.5rem; border-radius: 8px;
  transition: background 0.2s;
}
.user-trigger:hover { background: var(--bg-elevated); }

.user-av {
  width: 30px; height: 30px; border-radius: 50%;
  background: var(--bg-overlay); border: 1px solid var(--border-mid);
  display: flex; align-items: center; justify-content: center;
  font-size: 0.65rem; font-weight: 700; color: var(--text-muted);
  flex-shrink: 0;
}
.user-info { display: flex; flex-direction: column; align-items: flex-start; }
.user-name { font-size: 0.8rem; font-weight: 600; color: var(--text-secondary); line-height: 1.2; }
.user-role { font-size: 0.65rem; color: var(--text-faint); }
.chevron { width: 14px; height: 14px; color: var(--text-ghost); transition: transform 0.2s; }
.chevron.open { transform: rotate(180deg); }

/* Dropdown */
.dropdown {
  position: absolute; top: calc(100% + 8px); right: 0;
  width: 220px;
  background: var(--bg-surface);
  border: 1px solid var(--border);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 8px 32px var(--shadow);
  z-index: 200;
}
.dropdown-header {
  display: flex; align-items: center; gap: 0.75rem;
  padding: 0.875rem 1rem;
}
.dh-av {
  width: 34px; height: 34px; border-radius: 50%;
  background: var(--bg-elevated); border: 1px solid var(--border-mid);
  display: flex; align-items: center; justify-content: center;
  font-size: 0.7rem; font-weight: 700; color: var(--text-muted);
  flex-shrink: 0;
}
.dh-name { font-size: 0.85rem; font-weight: 600; color: var(--text-primary); }
.dh-email { font-size: 0.72rem; color: var(--text-faint); margin-top: 1px; }

.dropdown-divider { height: 1px; background: var(--border-soft); }

.dropdown-item {
  display: flex; align-items: center; gap: 0.75rem;
  padding: 0.7rem 1rem;
  color: var(--text-muted); font-size: 0.85rem; font-weight: 600;
  text-decoration: none; cursor: pointer;
  background: none; border: none; width: 100%; text-align: left;
  font-family: 'Syne', sans-serif;
  transition: background 0.15s, color 0.15s;
}
.dropdown-item svg { width: 15px; height: 15px; flex-shrink: 0; }
.dropdown-item:hover { background: var(--bg-elevated); color: var(--text-primary); }
.dropdown-item.danger:hover { background: rgba(220,60,60,0.06); color: #e07070; }

/* Transitions */
.dropdown-enter-active, .dropdown-leave-active { transition: opacity 0.15s ease, transform 0.15s ease; }
.dropdown-enter-from, .dropdown-leave-to { opacity: 0; transform: translateY(-6px); }

.icon-swap-enter-active, .icon-swap-leave-active { transition: opacity 0.15s ease, transform 0.15s ease; }
.icon-swap-enter-from { opacity: 0; transform: rotate(-30deg) scale(0.8); }
.icon-swap-leave-to { opacity: 0; transform: rotate(30deg) scale(0.8); }

@media (max-width: 640px) {
  .topbar { padding: 0 1rem; }
  .page-title { display: none; }
}
</style>