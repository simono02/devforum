<template>
  <div class="dash">
    <!-- Welcome -->
    <div class="welcome">
      <div class="welcome-text">
        <div class="welcome-eyebrow">
          {{ timeGreeting }},
          <span class="welcome-name">{{ user?.full_name?.split(' ')[0] }}</span> 👋
        </div>
        <h1 class="welcome-title">Your Dashboard</h1>
        <p class="welcome-sub">Track your posts, comments, and activity on DevForum.</p>
      </div>
      <button class="btn-write" @click="goToPosts">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/></svg>
        Write a post
      </button>
    </div>

    <!-- Stats -->
    <div class="stats-grid">
      <div class="stat-card" v-for="(s, i) in stats" :key="s.label" :style="{ animationDelay: i * 0.07 + 's' }">
        <div class="stat-icon" v-html="s.icon"></div>
        <div class="stat-value">{{ loading ? '—' : s.value }}</div>
        <div class="stat-label">{{ s.label }}</div>
      </div>
    </div>

    <!-- Two columns -->
    <div class="two-col">
      <!-- My recent posts -->
      <div class="card">
        <div class="card-head">
          <h3>My Recent Posts</h3>
          <router-link to="/dashboard/posts" class="card-link">View all →</router-link>
        </div>

        <!-- Loading -->
        <div class="mini-loading" v-if="loading">
          <div class="spinner-sm"></div>
        </div>

        <!-- Posts list -->
        <div class="my-posts-list" v-else-if="myPosts.length > 0">
          <div class="my-post-item" v-for="p in myPosts.slice(0, 4)" :key="p.id">
            <div class="mpi-top">
              <span class="mpi-tag" v-if="p.tag">{{ p.tag }}</span>
              <span class="mpi-time">{{ formatTime(p.created_at) }}</span>
            </div>
            <div class="mpi-title">{{ p.title }}</div>
            <div class="mpi-stats">
              <span class="mpi-stat">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="17 11 12 6 7 11"/><line x1="12" y1="18" x2="12" y2="6"/></svg>
                {{ p.upvotes }}
              </span>
            </div>
          </div>
        </div>

        <!-- Empty -->
        <div class="empty-state" v-else>
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
          <p>No posts yet</p>
          <router-link to="/dashboard/posts" class="empty-link">Write your first post →</router-link>
        </div>
      </div>

      <!-- Trending community posts -->
      <div class="card">
        <div class="card-head"><h3>Trending Now</h3></div>

        <div class="mini-loading" v-if="loading">
          <div class="spinner-sm"></div>
        </div>

        <div class="trending-list" v-else-if="trending.length > 0">
          <div class="trending-item" v-for="t in trending" :key="t.id">
            <div class="ti-left">
              <span class="ti-tag" v-if="t.tag">{{ t.tag }}</span>
              <span class="ti-title">{{ t.title }}</span>
            </div>
            <div class="ti-upvotes">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="17 11 12 6 7 11"/><line x1="12" y1="18" x2="12" y2="6"/></svg>
              {{ t.upvotes }}
            </div>
          </div>
        </div>

        <div class="empty-state" v-else>
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="17 11 12 6 7 11"/><line x1="12" y1="18" x2="12" y2="6"/></svg>
          <p>No posts yet in the community.</p>
        </div>
      </div>
    </div>

    <!-- Community stats banner -->
    <div class="community-banner">
      <div class="cb-stat" v-for="s in community" :key="s.label">
        <span class="cb-num">{{ loading ? '—' : s.value }}</span>
        <span class="cb-label">{{ s.label }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../../../stores/user.js'
import postService from '../../../services/posts.js'

const router    = useRouter()
const userStore = useUserStore()
const user      = computed(() => userStore.user)
const loading   = ref(false)

const myPosts  = ref([])
const allPosts = ref([])

const timeGreeting = computed(() => {
  const h = new Date().getHours()
  if (h < 12) return 'Good morning'
  if (h < 17) return 'Good afternoon'
  return 'Good evening'
})

// Real stats derived from API data
const totalComments = computed(() => allPosts.value.reduce((sum, p) => sum + (p.comments || 0), 0))
const myComments    = computed(() => myPosts.value.reduce((sum, p) => sum + (p.comments || 0), 0))

const stats = computed(() => [
  {
    value: myPosts.value.length,
    label: 'Posts written',
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>`
  },
  {
    value: myComments.value,
    label: 'Comments on my posts',
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>`
  },
  {
    value: myPosts.value.reduce((sum, p) => sum + (p.upvotes || 0), 0),
    label: 'Upvotes received',
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="17 11 12 6 7 11"/><line x1="12" y1="18" x2="12" y2="6"/></svg>`
  },
  {
    value: allPosts.value.length,
    label: 'Community posts',
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/></svg>`
  },
])

// Top 4 posts by upvotes
const trending = computed(() =>
  [...allPosts.value]
    .sort((a, b) => (b.upvotes || 0) - (a.upvotes || 0))
    .slice(0, 4)
)

const community = computed(() => [
  { value: allPosts.value.length, label: 'Total Posts' },
  { value: totalComments.value,   label: 'Total Comments' },
  { value: myPosts.value.reduce((s, p) => s + (p.upvotes || 0), 0), label: 'My Upvotes' },
  { value: '99%', label: 'Uptime' },
])

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

function goToPosts() {
  router.push('/dashboard/posts')
}

onMounted(async () => {
  loading.value = true
  try {
    const [all, mine] = await Promise.all([
      postService.getAll(),
      postService.getMine(),
    ])
    allPosts.value = all
    myPosts.value  = mine
  } catch {
    // fail silently — dashboard still renders with zeros
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Syne:wght@400;600;700;800&display=swap');
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
.dash { font-family: 'Syne', sans-serif; display: flex; flex-direction: column; gap: 1.5rem; animation: fadeIn 0.3s ease; }

.welcome {
  background: var(--bg-surface); border: 1px solid var(--border);
  border-radius: 16px; padding: 1.75rem;
  display: flex; align-items: center; justify-content: space-between; gap: 1rem; flex-wrap: wrap;
}
.welcome-eyebrow { font-size: 0.85rem; color: var(--text-faint); margin-bottom: 0.4rem; }
.welcome-name { color: var(--text-primary); font-weight: 700; }
.welcome-title { font-size: clamp(1.3rem, 3vw, 1.75rem); font-weight: 800; color: var(--text-primary); letter-spacing: -0.5px; margin-bottom: 0.3rem; }
.welcome-sub { font-size: 0.85rem; color: var(--text-faint); }
.btn-write {
  display: flex; align-items: center; gap: 0.5rem;
  background: var(--accent); color: var(--bg-base);
  border: none; border-radius: 10px; padding: 0.7rem 1.25rem;
  font-family: 'Syne', sans-serif; font-size: 0.88rem; font-weight: 700;
  cursor: pointer; white-space: nowrap; transition: opacity 0.2s; flex-shrink: 0;
}
.btn-write svg { width: 15px; height: 15px; }
.btn-write:hover { opacity: 0.85; }

.stats-grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 1rem; }
.stat-card {
  background: var(--bg-surface); border: 1px solid var(--border); border-radius: 12px;
  padding: 1.25rem; display: flex; flex-direction: column; gap: 0.4rem;
  animation: slideUp 0.4s ease both; transition: border-color 0.2s, transform 0.2s;
}
.stat-card:hover { border-color: var(--border-mid); transform: translateY(-2px); }
.stat-icon { color: var(--text-ghost); }
.stat-icon svg { width: 18px; height: 18px; }
.stat-value { font-size: 1.8rem; font-weight: 800; color: var(--text-primary); letter-spacing: -1px; line-height: 1; margin-top: 0.25rem; }
.stat-label { font-size: 0.75rem; color: var(--text-faint); }

.two-col { display: grid; grid-template-columns: 1fr 1fr; gap: 1rem; }
.card { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 14px; padding: 1.25rem; }
.card-head { display: flex; align-items: center; justify-content: space-between; margin-bottom: 1.25rem; }
.card-head h3 { font-size: 0.9rem; font-weight: 700; color: var(--text-secondary); }
.card-link { font-size: 0.78rem; color: var(--text-faint); text-decoration: none; }
.card-link:hover { color: var(--text-muted); }

.mini-loading { display: flex; justify-content: center; padding: 2rem 0; }
.spinner-sm { width: 20px; height: 20px; border: 2px solid var(--border); border-top-color: var(--text-muted); border-radius: 50%; animation: spin 0.8s linear infinite; }

.my-posts-list { display: flex; flex-direction: column; gap: 0; }
.my-post-item { padding: 0.75rem 0; border-bottom: 1px solid var(--border-soft); }
.my-post-item:last-child { border-bottom: none; }
.mpi-top { display: flex; align-items: center; justify-content: space-between; margin-bottom: 0.3rem; }
.mpi-tag { font-size: 0.65rem; font-weight: 700; background: var(--bg-elevated); color: var(--text-muted); padding: 0.15rem 0.45rem; border-radius: 4px; }
.mpi-time { font-size: 0.68rem; color: var(--text-ghost); font-family: monospace; }
.mpi-title { font-size: 0.85rem; font-weight: 600; color: var(--text-primary); line-height: 1.4; margin-bottom: 0.3rem; display: -webkit-box; -webkit-line-clamp: 1; -webkit-box-orient: vertical; overflow: hidden; }
.mpi-stats { display: flex; gap: 0.75rem; }
.mpi-stat { display: flex; align-items: center; gap: 0.25rem; font-size: 0.7rem; color: var(--text-ghost); font-family: monospace; }
.mpi-stat svg { width: 11px; height: 11px; }

.empty-state { display: flex; flex-direction: column; align-items: center; gap: 0.6rem; padding: 2rem 1rem; text-align: center; }
.empty-state svg { width: 28px; height: 28px; color: var(--text-ghost); }
.empty-state p { font-size: 0.85rem; color: var(--text-faint); }
.empty-link { font-size: 0.82rem; color: var(--text-muted); text-decoration: none; }
.empty-link:hover { color: var(--text-primary); }

.trending-list { display: flex; flex-direction: column; }
.trending-item { display: flex; align-items: center; justify-content: space-between; gap: 0.75rem; padding: 0.65rem 0; border-bottom: 1px solid var(--border-soft); }
.trending-item:last-child { border-bottom: none; }
.ti-left { display: flex; align-items: center; gap: 0.6rem; min-width: 0; flex: 1; }
.ti-tag { font-size: 0.68rem; font-weight: 700; background: var(--bg-elevated); color: var(--text-muted); padding: 0.18rem 0.5rem; border-radius: 4px; white-space: nowrap; flex-shrink: 0; }
.ti-title { font-size: 0.82rem; color: var(--text-muted); white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
.ti-upvotes { display: flex; align-items: center; gap: 0.3rem; font-size: 0.72rem; color: var(--text-ghost); flex-shrink: 0; font-family: monospace; }
.ti-upvotes svg { width: 12px; height: 12px; }

.community-banner {
  background: var(--bg-surface); border: 1px solid var(--border); border-radius: 14px;
  padding: 1.25rem 2rem; display: flex; align-items: center; justify-content: space-around; flex-wrap: wrap; gap: 1.5rem;
}
.cb-stat { text-align: center; }
.cb-num { display: block; font-size: 1.5rem; font-weight: 800; color: var(--text-primary); letter-spacing: -1px; }
.cb-label { font-size: 0.75rem; color: var(--text-faint); }

@keyframes fadeIn { from { opacity: 0; } to { opacity: 1; } }
@keyframes slideUp { from { opacity: 0; transform: translateY(10px); } to { opacity: 1; transform: translateY(0); } }
@keyframes spin { to { transform: rotate(360deg); } }

@media (max-width: 900px) { .two-col { grid-template-columns: 1fr; } }
@media (max-width: 640px) {
  .stats-grid { grid-template-columns: repeat(2, 1fr); }
  .welcome { flex-direction: column; align-items: flex-start; }
  .btn-write { width: 100%; justify-content: center; }
  .community-banner { padding: 1rem; gap: 1rem; }
}
</style>