<template>
  <div class="app-layout" :data-theme="themeStore.theme">
    <Sidebar :collapsed="collapsed" @toggle="collapsed = !collapsed" />
    <div class="main-area">
      <Header @toggle-sidebar="collapsed = !collapsed" />
      <main class="content">
        <router-view />
      </main>
    </div>
    <div class="mob-overlay" v-if="!collapsed && isMobile" @click="collapsed = true"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useThemeStore } from '../../stores/theme.js'
import Sidebar from './sidebar.vue'
import Header from './header.vue'

const themeStore = useThemeStore()
const collapsed  = ref(false)
const isMobile   = ref(window.innerWidth < 768)

function handleResize() {
  isMobile.value = window.innerWidth < 768
  if (isMobile.value) collapsed.value = true
}

onMounted(() => { window.addEventListener('resize', handleResize); handleResize() })
onUnmounted(() => window.removeEventListener('resize', handleResize))
</script>

<style scoped>
.app-layout {
  display: flex; height: 100vh; overflow: hidden;
  background: var(--bg-base); transition: background 0.25s;
}
.app-layout > :first-child {
  position: sticky; top: 0; height: 100vh;
  overflow-y: auto; overflow-x: hidden; flex-shrink: 0;
}
.main-area {
  flex: 1; display: flex; flex-direction: column;
  min-width: 0; height: 100vh; overflow: hidden;
}
.main-area > :first-child { flex-shrink: 0; }
.content { flex: 1; overflow-y: auto; padding: 2rem 1.75rem; }
.mob-overlay {
  display: none; position: fixed; inset: 0;
  background: rgba(0,0,0,0.5); z-index: 99; backdrop-filter: blur(2px);
}
@media (max-width: 768px) {
  .mob-overlay { display: block; }
  .app-layout > :first-child { position: fixed; left: 0; top: 0; z-index: 100; height: 100vh; }
  .content { padding: 1.5rem 1rem; }
}
</style>