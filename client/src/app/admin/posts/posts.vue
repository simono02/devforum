<template>
  <div class="page">
    <div class="page-header">
      <div>
        <h1 class="page-title">Posts</h1>
        <p class="page-sub">Review, moderate, and manage all forum posts.</p>
      </div>
      <div class="header-actions">
        <div class="sort-wrap">
          <select v-model="sortBy">
            <option value="newest">Newest first</option>
            <option value="oldest">Oldest first</option>
            <option value="upvotes">Most upvotes</option>
          </select>
        </div>
      </div>
    </div>

    <div class="filters">
      <div class="search-wrap">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>
        <input v-model="search" type="text" placeholder="Search posts..." />
      </div>
      <div class="filter-tabs">
        <button v-for="f in filters" :key="f.value" class="filter-tab" :class="{ active: activeFilter === f.value }" @click="activeFilter = f.value">
          {{ f.label }}<span class="filter-count">{{ f.count }}</span>
        </button>
      </div>
    </div>

    <div class="mini-loading" v-if="loading"><div class="spinner-sm"></div><p>Loading posts...</p></div>

    <div class="posts-list" v-else>
      <div class="post-card" v-for="p in filteredPosts" :key="p.id" :class="{ flagged: p.flagged }">
        <div class="post-main">
          <div class="post-meta">
            <div class="post-author">
              <div class="author-av">{{ initials(p.author) }}</div>
              <div>
                <div class="author-name">{{ p.author }}</div>
                <div class="author-time">{{ formatDate(p.created_at) }}</div>
              </div>
            </div>
            <div class="post-flags">
              <span class="post-tag" v-if="p.tag">{{ p.tag }}</span>
              <span class="flag-badge" v-if="p.flagged">⚑ Flagged</span>
              <span class="status-dot published">published</span>
            </div>
          </div>
          <h3 class="post-title" @click="viewPost(p)">{{ p.title }}</h3>
          <p class="post-excerpt">{{ p.content }}</p>
          <div class="post-stats">
            <span class="pstat">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="17 11 12 6 7 11"/><line x1="12" y1="18" x2="12" y2="6"/></svg>
              {{ p.upvotes }} upvotes
            </span>
            <span class="pstat mono">ID #{{ p.id }}</span>
          </div>
        </div>
        <div class="post-actions">
          <button class="pa-btn" title="View" @click="viewPost(p)">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
          </button>
          <button class="pa-btn" :class="{ active: p.flagged }" title="Flag" @click="toggleFlag(p)">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M4 15s1-1 4-1 5 2 8 2 4-1 4-1V3s-1 1-4 1-5-2-8-2-4 1-4 1z"/><line x1="4" y1="22" x2="4" y2="15"/></svg>
          </button>
          <button class="pa-btn danger" title="Delete" @click="confirmDelete(p)">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="3 6 5 6 21 6"/><path d="M19 6l-1 14H6L5 6"/><path d="M10 11v6"/><path d="M14 11v6"/></svg>
          </button>
        </div>
      </div>
      <div class="empty-state" v-if="filteredPosts.length === 0">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
        <p>No posts found.</p>
        <span>Posts will appear here once members start writing.</span>
      </div>
    </div>

    <!-- View Modal -->
    <Teleport to="body">
      <div class="overlay" v-if="viewModal" @click.self="viewModal = false">
        <div class="modal modal-lg">
          <div class="modal-head">
            <h3>Post Details</h3>
            <button class="modal-close" @click="viewModal = false">✕</button>
          </div>
          <div class="modal-body" v-if="activePost">
            <div class="post-detail-meta">
              <div class="post-author">
                <div class="author-av">{{ initials(activePost.author) }}</div>
                <div>
                  <div class="author-name">{{ activePost.author }}</div>
                  <div class="author-time">{{ formatDate(activePost.created_at) }}</div>
                </div>
              </div>
              <div class="post-flags">
                <span class="post-tag" v-if="activePost.tag">{{ activePost.tag }}</span>
                <span class="flag-badge" v-if="activePost.flagged">⚑ Flagged</span>
              </div>
            </div>
            <h2 class="detail-title">{{ activePost.title }}</h2>
            <p class="detail-content">{{ activePost.content }}</p>
            <div class="detail-stats">
              <span class="pstat">{{ activePost.upvotes }} upvotes</span>
              <span class="pstat mono">ID #{{ activePost.id }}</span>
            </div>
          </div>
          <div class="modal-foot">
            <button class="mbtn mbtn-outline" @click="viewModal = false">Close</button>
            <button class="mbtn mbtn-flag" @click="toggleFlag(activePost); viewModal = false">{{ activePost?.flagged ? 'Unflag' : 'Flag Post' }}</button>
            <button class="mbtn mbtn-danger" @click="confirmDelete(activePost); viewModal = false">Delete</button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Delete Confirm -->
    <Teleport to="body">
      <div class="overlay" v-if="deleteModal" @click.self="deleteModal = false">
        <div class="modal modal-sm">
          <div class="modal-head">
            <h3>Delete Post</h3>
            <button class="modal-close" @click="deleteModal = false">✕</button>
          </div>
          <div class="modal-body">
            <div class="confirm-icon"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg></div>
            <p class="confirm-text">Delete <strong>{{ activePost?.title }}</strong>? This cannot be undone.</p>
          </div>
          <div class="modal-foot">
            <button class="mbtn mbtn-outline" @click="deleteModal = false">Cancel</button>
            <button class="mbtn mbtn-danger" @click="doDelete">Delete Post</button>
          </div>
        </div>
      </div>
    </Teleport>

    <Teleport to="body">
      <div class="toast" v-if="toast.show" :class="toast.type">{{ toast.msg }}</div>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import adminService from '../../../services/admin.js'

const search = ref('')
const activeFilter = ref('all')
const sortBy = ref('newest')
const viewModal = ref(false)
const deleteModal = ref(false)
const activePost = ref(null)
const loading = ref(false)
const toast = ref({ show: false, msg: '', type: 'success' })
const posts = ref([])

const filters = computed(() => [
  { label: 'All',     value: 'all',     count: posts.value.length },
  { label: 'Flagged', value: 'flagged', count: posts.value.filter(p => p.flagged).length },
])

const filteredPosts = computed(() => {
  let list = [...posts.value]
  if (activeFilter.value === 'flagged') list = list.filter(p => p.flagged)
  if (search.value) {
    const q = search.value.toLowerCase()
    list = list.filter(p => p.title.toLowerCase().includes(q) || (p.author || '').toLowerCase().includes(q))
  }
  if (sortBy.value === 'newest') list.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
  else if (sortBy.value === 'oldest') list.sort((a, b) => new Date(a.created_at) - new Date(b.created_at))
  else if (sortBy.value === 'upvotes') list.sort((a, b) => b.upvotes - a.upvotes)
  return list
})

function initials(name) { return (name || 'U').split(' ').map(n => n[0]).join('').slice(0,2).toUpperCase() }
function formatDate(d) { return d ? new Date(d).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' }) : '—' }
function viewPost(p) { activePost.value = p; viewModal.value = true }
function confirmDelete(p) { activePost.value = p; deleteModal.value = true }
function toggleFlag(p) { p.flagged = !p.flagged; showToast(p.flagged ? 'Post flagged' : 'Post unflagged', p.flagged ? 'warn' : 'success') }

async function doDelete() {
  try {
    await adminService.deletePost(activePost.value.id)
    posts.value = posts.value.filter(p => p.id !== activePost.value.id)
    deleteModal.value = false
    showToast('Post deleted', 'danger')
  } catch { showToast('Failed to delete post', 'danger') }
}

function showToast(msg, type = 'success') { toast.value = { show: true, msg, type }; setTimeout(() => toast.value.show = false, 3000) }

onMounted(async () => {
  loading.value = true
  try {
    posts.value = (await adminService.getPosts()).map(p => ({ ...p, flagged: p.is_flagged || false, pinned: p.is_pinned || false }))
  } catch { showToast('Failed to load posts', 'danger') }
  finally { loading.value = false }
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500&family=Syne:wght@400;600;700;800&display=swap');
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
.page { font-family: 'Syne', sans-serif; display: flex; flex-direction: column; gap: 1.5rem; animation: fadeIn 0.3s ease; }
.page-header { display: flex; align-items: flex-start; justify-content: space-between; gap: 1rem; flex-wrap: wrap; }
.page-title { font-size: 1.75rem; font-weight: 800; color: var(--text-primary); letter-spacing: -1px; }
.page-sub { font-size: 0.85rem; color: var(--text-faint); margin-top: 0.25rem; }
.header-actions { display: flex; gap: 0.75rem; align-items: center; }
.sort-wrap select { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 8px; padding: 0.6rem 0.875rem; color: var(--text-muted); font-family: 'Syne', sans-serif; font-size: 0.82rem; outline: none; cursor: pointer; }
.sort-wrap select option { background: var(--bg-surface); }
.filters { display: flex; align-items: center; gap: 0.75rem; flex-wrap: wrap; }
.search-wrap { display: flex; align-items: center; gap: 0.6rem; background: var(--bg-surface); border: 1px solid var(--border); border-radius: 8px; padding: 0.55rem 0.875rem; flex: 1; min-width: 200px; max-width: 320px; }
.search-wrap svg { width: 15px; height: 15px; color: var(--text-ghost); flex-shrink: 0; }
.search-wrap input { background: none; border: none; outline: none; color: var(--text-primary); font-family: 'Syne', sans-serif; font-size: 0.85rem; width: 100%; }
.search-wrap input::placeholder { color: var(--text-ultra); }
.filter-tabs { display: flex; gap: 0.35rem; flex-wrap: wrap; }
.filter-tab { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 6px; padding: 0.45rem 0.75rem; font-family: 'Syne', sans-serif; font-size: 0.78rem; font-weight: 600; color: var(--text-faint); cursor: pointer; transition: all 0.2s; display: flex; align-items: center; gap: 0.4rem; }
.filter-tab:hover { color: var(--text-muted); }
.filter-tab.active { background: var(--bg-overlay); border-color: var(--border-mid); color: var(--text-secondary); }
.filter-count { font-size: 0.68rem; background: var(--bg-elevated); padding: 0.1rem 0.4rem; border-radius: 10px; color: var(--text-faint); }
.mini-loading { display: flex; flex-direction: column; align-items: center; gap: 0.75rem; padding: 3rem; }
.mini-loading p { font-size: 0.85rem; color: var(--text-faint); }
.spinner-sm { width: 24px; height: 24px; border: 2px solid var(--border); border-top-color: var(--text-muted); border-radius: 50%; animation: spin 0.8s linear infinite; }
.posts-list { display: flex; flex-direction: column; gap: 0.75rem; }
.post-card { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 12px; padding: 1.25rem; display: flex; align-items: flex-start; gap: 1rem; transition: border-color 0.2s, transform 0.2s; }
.post-card:hover { border-color: var(--border-mid); transform: translateY(-1px); }
.post-card.flagged { border-color: rgba(220,140,30,0.3); background: rgba(220,140,30,0.03); }
.post-main { flex: 1; min-width: 0; display: flex; flex-direction: column; gap: 0.6rem; }
.post-meta { display: flex; align-items: center; justify-content: space-between; flex-wrap: wrap; gap: 0.5rem; }
.post-author { display: flex; align-items: center; gap: 0.6rem; }
.author-av { width: 28px; height: 28px; border-radius: 50%; background: var(--bg-overlay); border: 1px solid var(--border-mid); display: flex; align-items: center; justify-content: center; font-size: 0.62rem; font-weight: 700; color: var(--text-muted); flex-shrink: 0; }
.author-name { font-size: 0.82rem; font-weight: 600; color: var(--text-secondary); }
.author-time { font-size: 0.7rem; color: var(--text-ghost); font-family: 'JetBrains Mono', monospace; }
.post-flags { display: flex; align-items: center; gap: 0.5rem; flex-wrap: wrap; }
.post-tag { font-size: 0.68rem; font-weight: 700; background: var(--bg-overlay); color: var(--text-muted); padding: 0.18rem 0.55rem; border-radius: 4px; }
.flag-badge { font-size: 0.68rem; font-weight: 700; background: rgba(220,140,30,0.1); color: #d4a030; border: 1px solid rgba(220,140,30,0.2); padding: 0.18rem 0.55rem; border-radius: 4px; }
.status-dot { font-size: 0.68rem; font-weight: 700; padding: 0.18rem 0.55rem; border-radius: 4px; text-transform: capitalize; }
.status-dot.published { background: rgba(76,175,114,0.08); color: #4caf72; }
.post-title { font-size: 0.95rem; font-weight: 700; color: var(--text-primary); cursor: pointer; transition: color 0.2s; line-height: 1.4; }
.post-title:hover { color: var(--text-muted); }
.post-excerpt { font-size: 0.82rem; color: var(--text-faint); line-height: 1.6; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; }
.post-stats { display: flex; gap: 1rem; }
.pstat { display: flex; align-items: center; gap: 0.35rem; font-size: 0.75rem; color: var(--text-ghost); font-family: 'JetBrains Mono', monospace; }
.pstat svg { width: 13px; height: 13px; }
.post-actions { display: flex; flex-direction: column; gap: 0.35rem; flex-shrink: 0; }
.pa-btn { background: none; border: 1px solid transparent; border-radius: 6px; padding: 0.4rem; cursor: pointer; color: var(--text-ghost); transition: all 0.15s; display: flex; align-items: center; }
.pa-btn svg { width: 14px; height: 14px; }
.pa-btn:hover { background: var(--bg-overlay); border-color: var(--border-mid); color: var(--text-muted); }
.pa-btn.active { color: #d4a030; }
.pa-btn.danger:hover { background: rgba(220,60,60,0.08); border-color: rgba(220,60,60,0.2); color: #cc6060; }
.empty-state { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 12px; padding: 4rem 2rem; display: flex; flex-direction: column; align-items: center; gap: 0.75rem; text-align: center; }
.empty-state svg { width: 36px; height: 36px; color: var(--text-ultra); }
.empty-state p { font-size: 1rem; font-weight: 600; color: var(--text-faint); }
.empty-state span { font-size: 0.82rem; color: var(--text-ghost); }
.overlay { position: fixed; inset: 0; background: rgba(0,0,0,0.7); display: flex; align-items: center; justify-content: center; z-index: 1000; padding: 1rem; backdrop-filter: blur(4px); animation: fadeIn 0.15s ease; }
.modal { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 16px; width: 100%; max-width: 480px; overflow: hidden; animation: slideUp 0.2s ease; max-height: 90vh; overflow-y: auto; }
.modal-lg { max-width: 600px; }
.modal-sm { max-width: 380px; }
.modal-head { display: flex; align-items: center; justify-content: space-between; padding: 1.25rem 1.5rem; border-bottom: 1px solid var(--border-soft); position: sticky; top: 0; background: var(--bg-surface); z-index: 1; }
.modal-head h3 { font-size: 1rem; font-weight: 700; color: var(--text-primary); }
.modal-close { background: none; border: none; color: var(--text-faint); cursor: pointer; font-size: 1rem; padding: 0.25rem; transition: color 0.2s; }
.modal-close:hover { color: var(--text-primary); }
.modal-body { padding: 1.5rem; display: flex; flex-direction: column; gap: 1.25rem; }
.modal-foot { display: flex; gap: 0.75rem; justify-content: flex-end; padding: 1.25rem 1.5rem; border-top: 1px solid var(--border-soft); flex-wrap: wrap; }
.post-detail-meta { display: flex; align-items: center; justify-content: space-between; flex-wrap: wrap; gap: 0.5rem; }
.detail-title { font-size: 1.1rem; font-weight: 700; color: var(--text-primary); line-height: 1.4; }
.detail-content { font-size: 0.88rem; color: var(--text-muted); line-height: 1.7; white-space: pre-wrap; }
.detail-stats { display: flex; gap: 1rem; flex-wrap: wrap; }
.confirm-icon { display: flex; justify-content: center; }
.confirm-icon svg { width: 40px; height: 40px; color: #cc6060; }
.confirm-text { font-size: 0.88rem; color: var(--text-muted); line-height: 1.6; text-align: center; }
.confirm-text strong { color: var(--text-primary); }
.mbtn { padding: 0.65rem 1.25rem; border-radius: 8px; font-family: 'Syne', sans-serif; font-size: 0.85rem; font-weight: 700; cursor: pointer; transition: all 0.2s; border: none; }
.mbtn-outline { background: transparent; border: 1px solid var(--border-mid); color: var(--text-muted); }
.mbtn-outline:hover { color: var(--text-secondary); }
.mbtn-danger { background: rgba(220,60,60,0.12); color: #e07070; border: 1px solid rgba(220,60,60,0.2); }
.mbtn-danger:hover { background: rgba(220,60,60,0.2); }
.mbtn-flag { background: rgba(220,140,30,0.1); color: #d4a030; border: 1px solid rgba(220,140,30,0.2); }
.mbtn-flag:hover { background: rgba(220,140,30,0.2); }
.toast { position: fixed; bottom: 1.5rem; right: 1.5rem; padding: 0.875rem 1.25rem; border-radius: 10px; font-family: 'Syne', sans-serif; font-size: 0.85rem; font-weight: 600; z-index: 2000; animation: slideUp 0.2s ease; }
.toast.success { background: var(--bg-surface); border: 1px solid rgba(76,175,114,0.3); color: #4caf72; }
.toast.danger  { background: var(--bg-surface); border: 1px solid rgba(220,60,60,0.3); color: #e07070; }
.toast.warn    { background: var(--bg-surface); border: 1px solid rgba(220,140,30,0.3); color: #d4a030; }
@keyframes fadeIn { from { opacity: 0; } to { opacity: 1; } }
@keyframes slideUp { from { opacity: 0; transform: translateY(12px); } to { opacity: 1; transform: translateY(0); } }
@keyframes spin { to { transform: rotate(360deg); } }
@media (max-width: 768px) {
  .post-card { flex-direction: column; }
  .post-actions { flex-direction: row; justify-content: flex-end; }
  .filters { flex-direction: column; align-items: stretch; }
  .search-wrap { max-width: 100%; }
  .page-header { flex-direction: column; align-items: flex-start; }
  .modal { max-width: 100%; border-radius: 16px 16px 0 0; position: fixed; bottom: 0; left: 0; right: 0; }
  .overlay { align-items: flex-end; padding: 0; }
  .modal-foot { flex-wrap: wrap; }
  .mbtn { flex: 1; text-align: center; }
}
</style>