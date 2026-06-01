<template>
  <aside class="sidebar" :class="{ collapsed }">
    <div class="sidebar-top">
      <router-link to="/" class="logo">
        <span v-if="!collapsed">dev<span class="accent">forum</span></span>
        <span v-else>df</span>
      </router-link>
      <button class="collapse-btn" @click="$emit('toggle')" :title="collapsed ? 'Expand' : 'Collapse'">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
          <polyline v-if="!collapsed" points="15 18 9 12 15 6"/>
          <polyline v-else           points="9 18 15 12 9 6"/>
        </svg>
      </button>
    </div>

    <div class="user-badge" v-if="!collapsed">
      <div class="avatar">{{ initials }}</div>
      <div class="user-info">
        <div class="user-name">{{ user?.full_name }}</div>
        <div class="user-role">Member</div>
      </div>
    </div>
    <div class="avatar-sm-wrap" v-else>
      <div class="avatar">{{ initials }}</div>
    </div>

    <nav class="nav">
      <div class="nav-section" v-if="!collapsed">Menu</div>

      <router-link class="nav-item" to="/dashboard" exact-active-class="active">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
          <rect x="3" y="3" width="7" height="7"/><rect x="14" y="3" width="7" height="7"/>
          <rect x="14" y="14" width="7" height="7"/><rect x="3" y="14" width="7" height="7"/>
        </svg>
        <span v-if="!collapsed">Dashboard</span>
      </router-link>

      <router-link class="nav-item" to="/dashboard/posts" active-class="active">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
          <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
          <polyline points="14 2 14 8 20 8"/>
          <line x1="16" y1="13" x2="8" y2="13"/>
          <line x1="16" y1="17" x2="8" y2="17"/>
        </svg>
        <span v-if="!collapsed">Posts</span>
      </router-link>

      <div class="nav-section" v-if="!collapsed">Account</div>

      <router-link class="nav-item" to="/dashboard/profile" active-class="active">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
          <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
          <circle cx="12" cy="7" r="4"/>
        </svg>
        <span v-if="!collapsed">Profile</span>
      </router-link>
    </nav>

    <div class="sidebar-bottom">
      <button class="nav-item logout" @click="handleLogout">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
          <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
          <polyline points="16 17 21 12 16 7"/>
          <line x1="21" y1="12" x2="9" y2="12"/>
        </svg>
        <span v-if="!collapsed">Sign out</span>
      </button>
    </div>
  </aside>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../../stores/user.js'

defineProps({ collapsed: Boolean })
defineEmits(['toggle'])

const router    = useRouter()
const userStore = useUserStore()
const user      = computed(() => userStore.user)
const initials  = computed(() => {
  const name = user.value?.full_name || 'U'
  return name.split(' ').map(n => n[0]).join('').slice(0, 2).toUpperCase()
})

function handleLogout() {
  userStore.clearAuth()
  router.push('/login')
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Syne:wght@400;600;700;800&display=swap');
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

.sidebar {
  width: 220px; height: 100vh;
  background: var(--sidebar-bg); border-right: 1px solid var(--border);
  display: flex; flex-direction: column;
  font-family: 'Syne', sans-serif;
  transition: width 0.25s ease, background 0.25s, border-color 0.25s;
  overflow: hidden;
}
.sidebar.collapsed { width: 64px; }

.sidebar-top {
  display: flex; align-items: center; justify-content: space-between;
  padding: 1rem; border-bottom: 1px solid var(--border);
  flex-shrink: 0; min-height: 56px;
}
.logo { font-size: 1.2rem; font-weight: 800; color: var(--text-primary); text-decoration: none; white-space: nowrap; overflow: hidden; }
.accent { color: var(--accent); }
.collapse-btn {
  background: none; border: none; cursor: pointer;
  color: var(--text-ghost); padding: 4px; border-radius: 6px;
  display: flex; align-items: center; transition: color 0.2s, background 0.2s;
  flex-shrink: 0;
}
.collapse-btn svg { width: 15px; height: 15px; }
.collapse-btn:hover { color: var(--text-muted); background: var(--bg-elevated); }

.user-badge {
  display: flex; align-items: center; gap: 0.75rem;
  padding: 0.875rem 1rem; border-bottom: 1px solid var(--border); flex-shrink: 0;
}
.avatar-sm-wrap {
  display: flex; justify-content: center;
  padding: 0.875rem 0; border-bottom: 1px solid var(--border); flex-shrink: 0;
}
.avatar {
  width: 32px; height: 32px; border-radius: 50%;
  background: var(--bg-elevated); border: 1px solid var(--border-mid);
  display: flex; align-items: center; justify-content: center;
  font-size: 0.68rem; font-weight: 700; color: var(--text-muted); flex-shrink: 0;
}
.user-name { font-size: 0.8rem; font-weight: 600; color: var(--text-secondary); white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
.user-role { font-size: 0.67rem; color: var(--text-faint); margin-top: 1px; }

.nav { flex: 1; padding: 0.5rem 0; display: flex; flex-direction: column; gap: 1px; overflow-y: auto; }
.nav-section {
  font-size: 0.62rem; font-weight: 700; color: var(--text-ultra);
  letter-spacing: 0.09em; text-transform: uppercase;
  padding: 0.75rem 1rem 0.3rem;
}
.nav-item {
  display: flex; align-items: center; gap: 0.7rem;
  padding: 0.6rem 0.75rem; margin: 0 0.5rem;
  border-radius: 8px; color: var(--text-faint);
  text-decoration: none; font-size: 0.84rem; font-weight: 600;
  transition: background 0.15s, color 0.15s;
  cursor: pointer; border: none; background: none; width: calc(100% - 1rem);
  font-family: 'Syne', sans-serif; white-space: nowrap;
}
.nav-item svg { width: 16px; height: 16px; flex-shrink: 0; }
.nav-item:hover { background: var(--bg-elevated); color: var(--text-muted); }
.nav-item.active { background: var(--bg-elevated); color: var(--text-primary); }
.nav-item.create {
  color: var(--text-secondary); border: 1px dashed var(--border-mid);
  margin-top: 0.25rem;
}
.nav-item.create:hover { border-color: var(--border-mid); background: var(--bg-elevated); }
.nav-item.logout:hover { color: #e07070; background: rgba(220,60,60,0.06); }
.sidebar.collapsed .nav-item { justify-content: center; padding: 0.6rem; }
.sidebar.collapsed .nav-section { display: none; }

.sidebar-bottom { padding: 0.5rem 0 0.75rem; border-top: 1px solid var(--border); flex-shrink: 0; }
</style>