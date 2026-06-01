<template>
  <router-view />
</template>

<script setup>
import { useThemeStore } from './stores/theme.js'
// Initialize theme on app load
useThemeStore()
</script>

<style>
/* ── Theme tokens ── */
:root, [data-theme="dark"] {
  --bg-base:      #09090e;
  --bg-surface:   #0f0f18;
  --bg-elevated:  #111119;
  --bg-overlay:   #161622;
  --border:       #1c1c2a;
  --border-soft:  #141420;
  --border-mid:   #2a2a3e;
  --text-primary: #e2e2e8;
  --text-secondary: #c0c0d0;
  --text-muted:   #8888a0;
  --text-faint:   #55556a;
  --text-ghost:   #33334a;
  --text-ultra:   #2a2a3e;
  --accent:       #e8eaf0;
  --accent-hover: #ffffff;
  --sidebar-bg:   #0d0d15;
  --topbar-bg:    #0a0a10;
  --shadow:       rgba(0,0,0,0.5);
}

[data-theme="light"] {
  --bg-base:      #f4f4f8;
  --bg-surface:   #ffffff;
  --bg-elevated:  #f8f8fc;
  --bg-overlay:   #f0f0f6;
  --border:       #e0e0ec;
  --border-soft:  #ececf4;
  --border-mid:   #c8c8d8;
  --text-primary: #111118;
  --text-secondary: #2a2a3a;
  --text-muted:   #5a5a72;
  --text-faint:   #8080a0;
  --text-ghost:   #a0a0b8;
  --text-ultra:   #c0c0d0;
  --accent:       #111118;
  --accent-hover: #2a2a3a;
  --sidebar-bg:   #ffffff;
  --topbar-bg:    #ffffff;
  --shadow:       rgba(0,0,0,0.08);
}

*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

body {
  background: var(--bg-base);
  color: var(--text-primary);
  transition: background 0.25s ease, color 0.25s ease;
}
</style>