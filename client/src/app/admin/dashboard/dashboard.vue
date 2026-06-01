<template>
  <div class="dash">
    <div class="dash-header">
      <div>
        <div class="dash-eyebrow">
          <span class="live-badge"><span class="live-dot"></span>Live</span>
          <span class="dash-date">{{ currentDate }}</span>
        </div>
        <h1 class="dash-title">Admin Dashboard</h1>
        <p class="dash-sub">Full platform overview. All systems operational.</p>
      </div>
      <div class="header-actions">
        <router-link to="/admin/users" class="hbtn hbtn-outline">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/></svg>
          Users
        </router-link>
        <router-link to="/admin/posts" class="hbtn hbtn-primary">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
          Posts
        </router-link>
      </div>
    </div>

    <!-- Stats -->
    <div class="stats-grid">
      <div class="stat-card" v-for="(s, i) in stats" :key="s.label" :style="{ animationDelay: i * 0.07 + 's' }">
        <div class="stat-top">
          <div class="stat-icon" v-html="s.icon"></div>
          <span class="stat-trend neutral">—</span>
        </div>
        <div class="stat-value">{{ loading ? '—' : s.value }}</div>
        <div class="stat-label">{{ s.label }}</div>
        <div class="stat-bar"><div class="stat-bar-fill" :style="{ width: s.fill }"></div></div>
      </div>
    </div>

    <!-- Mid row -->
    <div class="mid-row">
      <div class="card card-feed">
        <div class="card-head">
          <h3>Recent Posts</h3>
          <span class="card-badge">Live</span>
        </div>
        <div class="mini-loading" v-if="loading"><div class="spinner-sm"></div></div>
        <div class="feed" v-else>
          <div class="feed-item" v-for="p in recentPosts" :key="p.id">
            <div class="feed-avatar">{{ initials(p.author) }}</div>
            <div class="feed-body">
              <div class="feed-text"><span class="feed-name">{{ p.author }}</span> published a post</div>
              <div class="feed-title">{{ p.title }}</div>
              <div class="feed-time">{{ formatTime(p.created_at) }}</div>
            </div>
            <div class="feed-tag post">post</div>
          </div>
          <div class="feed-empty" v-if="recentPosts.length === 0">No posts yet.</div>
        </div>
      </div>

      <div class="card card-status">
        <div class="card-head"><h3>System Status</h3></div>
        <div class="status-list">
          <div class="status-item" v-for="s in systemStatus" :key="s.label">
            <div class="status-left">
              <div class="status-dot" :class="s.status"></div>
              <span class="status-label">{{ s.label }}</span>
            </div>
            <span class="status-val" :class="s.status">{{ s.value }}</span>
          </div>
        </div>
        <div class="card-divider"></div>
        <div class="card-head" style="margin-top:0"><h3>Quick Actions</h3></div>
        <div class="quick-actions">
          <router-link to="/admin/users" class="qa-btn">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/></svg>
            <span>Manage Users</span>
            <svg class="qa-arrow" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="9 18 15 12 9 6"/></svg>
          </router-link>
          <router-link to="/admin/posts" class="qa-btn">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
            <span>Review Posts</span>
            <svg class="qa-arrow" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="9 18 15 12 9 6"/></svg>
          </router-link>
        </div>
      </div>
    </div>

    <!-- Trending topics -->
    <div class="card">
      <div class="card-head">
        <h3>Trending Topics</h3>
        <router-link to="/admin/posts" class="card-link">View all →</router-link>
      </div>
      <div class="mini-loading" v-if="loading"><div class="spinner-sm"></div></div>
      <div class="topics-grid" v-else>
        <div class="topic-chip" v-for="t in trendingTopics" :key="t.tag">
          <span class="topic-tag">{{ t.tag }}</span>
          <span class="topic-count">{{ t.count }} posts</span>
        </div>
        <div class="topic-empty" v-if="trendingTopics.length === 0">No posts yet.</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import adminService from '../../../services/admin.js'

const loading = ref(false)
const allPosts = ref([])
const allUsers = ref([])

const currentDate = computed(() => new Date().toLocaleDateString('en-US', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' }))

const totalUpvotes  = computed(() => allPosts.value.reduce((s, p) => s + (p.upvotes || 0), 0))
const totalComments = computed(() => allPosts.value.reduce((s, p) => s + (p.comments || 0), 0))

const stats = computed(() => [
  {
    value: allUsers.value.length, label: 'Total Users', fill: `${Math.min(allUsers.value.length * 10, 100)}%`,
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/></svg>`
  },
  {
    value: allPosts.value.length, label: 'Total Posts', fill: `${Math.min(allPosts.value.length * 5, 100)}%`,
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>`
  },
  {
    value: totalComments.value, label: 'Total Comments', fill: `${Math.min(totalComments.value * 3, 100)}%`,
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>`
  },
  {
    value: totalUpvotes.value, label: 'Total Upvotes', fill: `${Math.min(totalUpvotes.value * 2, 100)}%`,
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="17 11 12 6 7 11"/><line x1="12" y1="18" x2="12" y2="6"/></svg>`
  },
])

const recentPosts = computed(() =>
  [...allPosts.value].sort((a, b) => new Date(b.created_at) - new Date(a.created_at)).slice(0, 5)
)

const trendingTopics = computed(() => {
  const counts = {}
  allPosts.value.forEach(p => { if (p.tag) counts[p.tag] = (counts[p.tag] || 0) + 1 })
  return Object.entries(counts).map(([tag, count]) => ({ tag, count })).sort((a, b) => b.count - a.count)
})

const systemStatus = [
  { label: 'API Server',   value: 'Online',      status: 'ok' },
  { label: 'Database',     value: 'Connected',   status: 'ok' },
  { label: 'Environment',  value: 'Development', status: 'warn' },
  { label: 'Version',      value: 'v1.0.0',      status: 'neutral' },
]

function initials(name) {
  return (name || 'U').split(' ').map(n => n[0]).join('').slice(0, 2).toUpperCase()
}
function formatTime(d) {
  if (!d) return ''
  const diff = Date.now() - new Date(d).getTime()
  const mins = Math.floor(diff / 60000)
  if (mins < 1)  return 'Just now'
  if (mins < 60) return `${mins}m ago`
  const hrs = Math.floor(mins / 60)
  if (hrs < 24)  return `${hrs}h ago`
  return `${Math.floor(hrs / 24)}d ago`
}

onMounted(async () => {
  loading.value = true
  try {
    const [posts, users] = await Promise.all([
      adminService.getPosts(),
      adminService.getUsers(),
    ])
    allPosts.value = posts
    allUsers.value = users
  } catch {
    // fail silently
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500&family=Syne:wght@400;600;700;800&display=swap');
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
.dash { font-family: 'Syne', sans-serif; display: flex; flex-direction: column; gap: 1.5rem; animation: fadeIn 0.4s ease; }

.dash-header { display: flex; align-items: flex-start; justify-content: space-between; gap: 1rem; flex-wrap: wrap; }
.dash-eyebrow { display: flex; align-items: center; gap: 0.75rem; margin-bottom: 0.6rem; }
.live-badge { display: flex; align-items: center; gap: 0.4rem; font-size: 0.72rem; font-weight: 700; color: #4caf72; background: rgba(76,175,114,0.08); border: 1px solid rgba(76,175,114,0.2); padding: 0.2rem 0.6rem; border-radius: 20px; }
.live-dot { width: 6px; height: 6px; border-radius: 50%; background: #4caf72; animation: pulse 2s infinite; }
.dash-date { font-size: 0.75rem; color: var(--text-ghost); font-family: 'JetBrains Mono', monospace; }
.dash-title { font-size: clamp(1.4rem, 3vw, 2rem); font-weight: 800; color: var(--text-primary); letter-spacing: -1px; margin-bottom: 0.3rem; }
.dash-sub { font-size: 0.85rem; color: var(--text-faint); }
.header-actions { display: flex; gap: 0.75rem; flex-shrink: 0; }
.hbtn { display: flex; align-items: center; gap: 0.5rem; padding: 0.6rem 1.1rem; border-radius: 8px; font-family: 'Syne', sans-serif; font-size: 0.82rem; font-weight: 700; text-decoration: none; transition: all 0.2s; }
.hbtn svg { width: 14px; height: 14px; }
.hbtn-outline { background: transparent; border: 1px solid var(--border-mid); color: var(--text-muted); }
.hbtn-outline:hover { color: var(--text-secondary); }
.hbtn-primary { background: var(--accent); color: var(--bg-base); border: none; }
.hbtn-primary:hover { opacity: 0.85; }

.stats-grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 1rem; }
.stat-card { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 14px; padding: 1.25rem 1.25rem 1rem; display: flex; flex-direction: column; gap: 0.3rem; animation: slideUp 0.4s ease both; transition: border-color 0.2s, transform 0.2s; }
.stat-card:hover { border-color: var(--border-mid); transform: translateY(-2px); }
.stat-top { display: flex; align-items: center; justify-content: space-between; margin-bottom: 0.5rem; }
.stat-icon { color: var(--text-ghost); }
.stat-icon svg { width: 18px; height: 18px; }
.stat-trend { font-size: 0.7rem; font-weight: 700; padding: 0.15rem 0.5rem; border-radius: 20px; }
.stat-trend.neutral { background: var(--bg-elevated); color: var(--text-ghost); }
.stat-value { font-size: 2rem; font-weight: 800; color: var(--text-primary); letter-spacing: -2px; line-height: 1; }
.stat-label { font-size: 0.75rem; color: var(--text-faint); margin-bottom: 0.75rem; }
.stat-bar { height: 3px; background: var(--bg-overlay); border-radius: 2px; overflow: hidden; }
.stat-bar-fill { height: 100%; background: linear-gradient(90deg, #3a3a5a, #8888a8); border-radius: 2px; transition: width 1s ease; }

.card { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 14px; padding: 1.25rem; }
.card-head { display: flex; align-items: center; justify-content: space-between; margin-bottom: 1.25rem; }
.card-head h3 { font-size: 0.88rem; font-weight: 700; color: var(--text-secondary); }
.card-badge { font-size: 0.65rem; font-weight: 700; background: rgba(76,175,114,0.1); color: #4caf72; border: 1px solid rgba(76,175,114,0.2); padding: 0.2rem 0.55rem; border-radius: 20px; }
.card-link { font-size: 0.78rem; color: var(--text-faint); text-decoration: none; }
.card-link:hover { color: var(--text-muted); }
.card-divider { height: 1px; background: var(--border-soft); margin: 1.25rem 0; }

.mid-row { display: grid; grid-template-columns: 1fr 340px; gap: 1rem; }

.mini-loading { display: flex; justify-content: center; padding: 2rem 0; }
.spinner-sm { width: 20px; height: 20px; border: 2px solid var(--border); border-top-color: var(--text-muted); border-radius: 50%; animation: spin 0.8s linear infinite; }

.feed { display: flex; flex-direction: column; }
.feed-item { display: flex; align-items: flex-start; gap: 0.875rem; padding: 0.75rem 0; border-bottom: 1px solid var(--border-soft); }
.feed-item:last-child { border-bottom: none; }
.feed-avatar { width: 32px; height: 32px; border-radius: 50%; background: var(--bg-overlay); border: 1px solid var(--border-mid); display: flex; align-items: center; justify-content: center; font-size: 0.68rem; font-weight: 700; color: var(--text-muted); flex-shrink: 0; }
.feed-body { flex: 1; min-width: 0; }
.feed-text { font-size: 0.78rem; color: var(--text-faint); }
.feed-name { color: var(--text-secondary); font-weight: 600; }
.feed-title { font-size: 0.82rem; color: var(--text-primary); font-weight: 600; margin-top: 0.15rem; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
.feed-time { font-size: 0.7rem; color: var(--text-ghost); margin-top: 2px; font-family: 'JetBrains Mono', monospace; }
.feed-tag { font-size: 0.65rem; font-weight: 700; padding: 0.2rem 0.5rem; border-radius: 4px; text-transform: uppercase; flex-shrink: 0; }
.feed-tag.post { background: rgba(232,234,240,0.06); color: var(--text-muted); }
.feed-empty { font-size: 0.82rem; color: var(--text-ghost); padding: 1.5rem 0; text-align: center; }

.status-list { display: flex; flex-direction: column; }
.status-item { display: flex; align-items: center; justify-content: space-between; padding: 0.6rem 0; border-bottom: 1px solid var(--border-soft); }
.status-item:last-child { border-bottom: none; }
.status-left { display: flex; align-items: center; gap: 0.6rem; }
.status-dot { width: 7px; height: 7px; border-radius: 50%; flex-shrink: 0; }
.status-dot.ok { background: #4caf72; }
.status-dot.warn { background: #e0a030; }
.status-dot.neutral { background: var(--text-ghost); }
.status-label { font-size: 0.82rem; color: var(--text-faint); }
.status-val { font-size: 0.78rem; font-family: 'JetBrains Mono', monospace; }
.status-val.ok { color: #4caf72; }
.status-val.warn { color: #e0a030; }
.status-val.neutral { color: var(--text-faint); }

.quick-actions { display: flex; flex-direction: column; gap: 0.6rem; }
.qa-btn { display: flex; align-items: center; gap: 0.75rem; padding: 0.8rem 1rem; background: var(--bg-elevated); border: 1px solid var(--border-soft); border-radius: 10px; color: var(--text-muted); font-size: 0.82rem; font-weight: 600; text-decoration: none; transition: all 0.2s; font-family: 'Syne', sans-serif; }
.qa-btn svg { width: 15px; height: 15px; flex-shrink: 0; }
.qa-btn span { flex: 1; }
.qa-arrow { color: var(--text-ultra); transition: transform 0.2s; }
.qa-btn:hover { border-color: var(--border-mid); color: var(--text-secondary); }
.qa-btn:hover .qa-arrow { transform: translateX(3px); color: var(--text-muted); }

.topics-grid { display: flex; flex-wrap: wrap; gap: 0.6rem; }
.topic-chip { display: flex; align-items: center; gap: 0.5rem; background: var(--bg-elevated); border: 1px solid var(--border-soft); border-radius: 8px; padding: 0.5rem 0.875rem; transition: border-color 0.2s; }
.topic-chip:hover { border-color: var(--border-mid); }
.topic-tag { font-size: 0.82rem; font-weight: 700; color: var(--text-secondary); }
.topic-count { font-size: 0.72rem; color: var(--text-ghost); font-family: 'JetBrains Mono', monospace; }
.topic-empty { font-size: 0.82rem; color: var(--text-ghost); }

@keyframes fadeIn { from { opacity: 0; } to { opacity: 1; } }
@keyframes slideUp { from { opacity: 0; transform: translateY(12px); } to { opacity: 1; transform: translateY(0); } }
@keyframes pulse { 0%,100% { opacity: 1; } 50% { opacity: 0.3; } }
@keyframes spin { to { transform: rotate(360deg); } }

@media (max-width: 1024px) { .mid-row { grid-template-columns: 1fr; } }
@media (max-width: 768px) {
  .stats-grid { grid-template-columns: repeat(2, 1fr); }
  .dash-header { flex-direction: column; align-items: flex-start; }
  .header-actions { width: 100%; }
  .hbtn { flex: 1; justify-content: center; }
}
@media (max-width: 480px) { .stats-grid { grid-template-columns: 1fr 1fr; } }
</style>