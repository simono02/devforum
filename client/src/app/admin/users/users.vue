<template>
  <div class="page">
    <div class="page-header">
      <div>
        <h1 class="page-title">Users</h1>
        <p class="page-sub">Manage all registered members and admins.</p>
      </div>
    </div>

    <div class="filters">
      <div class="search-wrap">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>
        <input v-model="search" type="text" placeholder="Search users..." />
      </div>
      <div class="filter-tabs">
        <button v-for="f in filters" :key="f.value" class="filter-tab" :class="{ active: activeFilter === f.value }" @click="activeFilter = f.value">
          {{ f.label }}<span class="filter-count">{{ f.count }}</span>
        </button>
      </div>
    </div>

    <div class="mini-loading" v-if="loading"><div class="spinner-sm"></div><p>Loading users...</p></div>

    <div class="table-card" v-else>
      <div class="table-wrap">
        <table class="table">
          <thead>
            <tr>
              <th><input type="checkbox" @change="toggleAll" /></th>
              <th>User</th>
              <th>Role</th>
              <th>Status</th>
              <th>Joined</th>
              <th>Posts</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="u in filteredUsers" :key="u.id" :class="{ selected: selected.includes(u.id) }">
              <td><input type="checkbox" :checked="selected.includes(u.id)" @change="toggleSelect(u.id)" /></td>
              <td>
                <div class="user-cell">
                  <div class="user-av">{{ initials(u.full_name) }}</div>
                  <div>
                    <div class="user-name">{{ u.full_name }}</div>
                    <div class="user-email">{{ u.email }}</div>
                  </div>
                </div>
              </td>
              <td><span class="role-badge" :class="u.role">{{ u.role }}</span></td>
              <td><span class="status-badge" :class="u.is_active ? 'active' : 'inactive'">{{ u.is_active ? 'Active' : 'Inactive' }}</span></td>
              <td class="td-mono">{{ formatDate(u.created_at) }}</td>
              <td class="td-mono">{{ u.posts || 0 }}</td>
              <td>
                <div class="row-actions">
                  <button class="ra-btn" title="View" @click="viewUser(u)">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
                  </button>
                  <button class="ra-btn" title="Edit" @click="editUser(u)">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/></svg>
                  </button>
                  <button class="ra-btn danger" title="Delete" @click="confirmDelete(u)">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="3 6 5 6 21 6"/><path d="M19 6l-1 14H6L5 6"/><path d="M10 11v6"/><path d="M14 11v6"/><path d="M9 6V4h6v2"/></svg>
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="filteredUsers.length === 0">
              <td colspan="7" class="empty-row">No users found.</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="bulk-bar" v-if="selected.length > 0">
        <span class="bulk-count">{{ selected.length }} selected</span>
        <button class="bulk-btn danger" @click="bulkDelete">Delete</button>
        <button class="bulk-clear" @click="selected = []">✕ Clear</button>
      </div>
    </div>

    <!-- View Modal -->
    <Teleport to="body">
      <div class="overlay" v-if="viewModal" @click.self="viewModal = false">
        <div class="modal">
          <div class="modal-head"><h3>User Details</h3><button class="modal-close" @click="viewModal = false">✕</button></div>
          <div class="modal-body" v-if="activeUser">
            <div class="profile-hero">
              <div class="profile-av">{{ initials(activeUser.full_name) }}</div>
              <div>
                <div class="profile-name">{{ activeUser.full_name }}</div>
                <div class="profile-email">{{ activeUser.email }}</div>
                <div class="profile-badges">
                  <span class="role-badge" :class="activeUser.role">{{ activeUser.role }}</span>
                  <span class="status-badge" :class="activeUser.is_active ? 'active' : 'inactive'">{{ activeUser.is_active ? 'Active' : 'Inactive' }}</span>
                </div>
              </div>
            </div>
            <div class="detail-grid">
              <div class="detail-item"><span class="detail-label">Username</span><span class="detail-val">@{{ activeUser.username }}</span></div>
              <div class="detail-item"><span class="detail-label">User ID</span><span class="detail-val mono">#{{ activeUser.id }}</span></div>
              <div class="detail-item"><span class="detail-label">Joined</span><span class="detail-val mono">{{ formatDate(activeUser.created_at) }}</span></div>
              <div class="detail-item"><span class="detail-label">Posts</span><span class="detail-val mono">{{ activeUser.posts || 0 }}</span></div>
            </div>
          </div>
          <div class="modal-foot">
            <button class="mbtn mbtn-outline" @click="viewModal = false">Close</button>
            <button class="mbtn mbtn-primary" @click="editUser(activeUser); viewModal = false">Edit User</button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Edit Modal -->
    <Teleport to="body">
      <div class="overlay" v-if="editModal" @click.self="editModal = false">
        <div class="modal">
          <div class="modal-head"><h3>Edit User</h3><button class="modal-close" @click="editModal = false">✕</button></div>
          <div class="modal-body" v-if="editForm">
            <div class="form-grid">
              <div class="form-field"><label>Full Name</label><input v-model="editForm.full_name" type="text" /></div>
              <div class="form-field"><label>Username</label><input v-model="editForm.username" type="text" /></div>
              <div class="form-field"><label>Email</label><input v-model="editForm.email" type="email" /></div>
              <div class="form-field">
                <label>Role</label>
                <select v-model="editForm.role"><option value="member">Member</option><option value="admin">Admin</option></select>
              </div>
              <div class="form-field">
                <label>Status</label>
                <select v-model="editForm.is_active"><option :value="true">Active</option><option :value="false">Inactive</option></select>
              </div>
            </div>
          </div>
          <div class="modal-foot">
            <button class="mbtn mbtn-outline" @click="editModal = false">Cancel</button>
            <button class="mbtn mbtn-primary" @click="saveEdit" :disabled="saving">
              <span v-if="saving" class="spinner"></span>
              <span v-else>Save Changes</span>
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Delete Modal -->
    <Teleport to="body">
      <div class="overlay" v-if="deleteModal" @click.self="deleteModal = false">
        <div class="modal modal-sm">
          <div class="modal-head"><h3>Delete User</h3><button class="modal-close" @click="deleteModal = false">✕</button></div>
          <div class="modal-body">
            <div class="confirm-icon"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg></div>
            <p class="confirm-text">Delete <strong>{{ activeUser?.full_name }}</strong>? This cannot be undone.</p>
          </div>
          <div class="modal-foot">
            <button class="mbtn mbtn-outline" @click="deleteModal = false">Cancel</button>
            <button class="mbtn mbtn-danger" @click="doDelete">Delete User</button>
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
const selected = ref([])
const viewModal = ref(false)
const editModal = ref(false)
const deleteModal = ref(false)
const activeUser = ref(null)
const editForm = ref(null)
const loading = ref(false)
const saving = ref(false)
const toast = ref({ show: false, msg: '', type: 'success' })
const users = ref([])

const filters = computed(() => [
  { label: 'All',      value: 'all',      count: users.value.length },
  { label: 'Admins',   value: 'admin',    count: users.value.filter(u => u.role === 'admin').length },
  { label: 'Members',  value: 'member',   count: users.value.filter(u => u.role === 'member').length },
  { label: 'Inactive', value: 'inactive', count: users.value.filter(u => !u.is_active).length },
])

const filteredUsers = computed(() => {
  let list = users.value
  if (activeFilter.value === 'admin')    list = list.filter(u => u.role === 'admin')
  else if (activeFilter.value === 'member')   list = list.filter(u => u.role === 'member')
  else if (activeFilter.value === 'inactive') list = list.filter(u => !u.is_active)
  if (search.value) {
    const q = search.value.toLowerCase()
    list = list.filter(u => u.full_name.toLowerCase().includes(q) || u.email.toLowerCase().includes(q) || u.username.toLowerCase().includes(q))
  }
  return list
})

function initials(name) { return (name || 'U').split(' ').map(n => n[0]).join('').slice(0,2).toUpperCase() }
function formatDate(d) { return d ? new Date(d).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' }) : '—' }
function toggleAll(e) { selected.value = e.target.checked ? users.value.map(u => u.id) : [] }
function toggleSelect(id) { selected.value.includes(id) ? selected.value = selected.value.filter(i => i !== id) : selected.value.push(id) }
function viewUser(u) { activeUser.value = u; viewModal.value = true }
function editUser(u) { activeUser.value = u; editForm.value = { ...u }; editModal.value = true }
function confirmDelete(u) { activeUser.value = u; deleteModal.value = true }

async function saveEdit() {
  saving.value = true
  try {
    const updated = await adminService.updateUser(editForm.value.id, editForm.value)
    const idx = users.value.findIndex(u => u.id === updated.id)
    if (idx !== -1) users.value[idx] = updated
    editModal.value = false
    showToast('User updated', 'success')
  } catch { showToast('Failed to update user', 'danger') }
  finally { saving.value = false }
}

async function doDelete() {
  try {
    await adminService.deleteUser(activeUser.value.id)
    users.value = users.value.filter(u => u.id !== activeUser.value.id)
    deleteModal.value = false
    showToast('User deleted', 'danger')
  } catch { showToast('Failed to delete user', 'danger') }
}

async function bulkDelete() {
  try {
    await Promise.all(selected.value.map(id => adminService.deleteUser(id)))
    users.value = users.value.filter(u => !selected.value.includes(u.id))
    selected.value = []
    showToast('Users deleted', 'danger')
  } catch { showToast('Failed to delete some users', 'danger') }
}

function showToast(msg, type = 'success') { toast.value = { show: true, msg, type }; setTimeout(() => toast.value.show = false, 3000) }

onMounted(async () => {
  loading.value = true
  try { users.value = await adminService.getUsers() }
  catch { showToast('Failed to load users', 'danger') }
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
.table-card { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 14px; overflow: hidden; }
.table-wrap { overflow-x: auto; }
.table { width: 100%; border-collapse: collapse; }
.table thead tr { border-bottom: 1px solid var(--border-soft); }
.table th { padding: 0.75rem 1rem; text-align: left; font-size: 0.72rem; font-weight: 700; color: var(--text-ghost); text-transform: uppercase; letter-spacing: 0.06em; white-space: nowrap; }
.table th:first-child, .table td:first-child { padding-left: 1.25rem; width: 40px; }
.table td { padding: 0.875rem 1rem; border-bottom: 1px solid var(--border-soft); vertical-align: middle; }
.table tbody tr:last-child td { border-bottom: none; }
.table tbody tr { transition: background 0.15s; }
.table tbody tr:hover { background: var(--bg-elevated); }
.table tbody tr.selected { background: rgba(232,234,240,0.03); }
.user-cell { display: flex; align-items: center; gap: 0.75rem; }
.user-av { width: 32px; height: 32px; border-radius: 50%; background: var(--bg-overlay); border: 1px solid var(--border-mid); display: flex; align-items: center; justify-content: center; font-size: 0.68rem; font-weight: 700; color: var(--text-muted); flex-shrink: 0; }
.user-name { font-size: 0.85rem; font-weight: 600; color: var(--text-secondary); }
.user-email { font-size: 0.75rem; color: var(--text-faint); font-family: 'JetBrains Mono', monospace; }
.role-badge { font-size: 0.7rem; font-weight: 700; padding: 0.2rem 0.6rem; border-radius: 20px; text-transform: capitalize; }
.role-badge.admin { background: rgba(232,234,240,0.08); color: var(--text-secondary); border: 1px solid var(--border-mid); }
.role-badge.member { background: var(--bg-elevated); color: var(--text-faint); border: 1px solid var(--border-soft); }
.status-badge { font-size: 0.7rem; font-weight: 700; padding: 0.2rem 0.6rem; border-radius: 20px; }
.status-badge.active { background: rgba(76,175,114,0.1); color: #4caf72; }
.status-badge.inactive { background: rgba(220,60,60,0.08); color: #cc6060; }
.td-mono { font-size: 0.78rem; color: var(--text-faint); font-family: 'JetBrains Mono', monospace; }
.row-actions { display: flex; gap: 0.35rem; justify-content: flex-end; }
.ra-btn { background: none; border: 1px solid transparent; border-radius: 6px; padding: 0.35rem; cursor: pointer; color: var(--text-ghost); transition: all 0.15s; display: flex; align-items: center; }
.ra-btn svg { width: 14px; height: 14px; }
.ra-btn:hover { background: var(--bg-overlay); border-color: var(--border-mid); color: var(--text-muted); }
.ra-btn.danger:hover { background: rgba(220,60,60,0.08); border-color: rgba(220,60,60,0.2); color: #cc6060; }
.empty-row { text-align: center; color: var(--text-ghost); font-size: 0.85rem; padding: 3rem !important; }
.bulk-bar { display: flex; align-items: center; gap: 0.75rem; padding: 0.875rem 1.25rem; background: var(--bg-elevated); border-top: 1px solid var(--border-soft); flex-wrap: wrap; }
.bulk-count { font-size: 0.82rem; font-weight: 600; color: var(--text-muted); flex: 1; }
.bulk-btn { background: var(--bg-overlay); border: 1px solid var(--border-mid); border-radius: 6px; padding: 0.4rem 0.875rem; font-family: 'Syne', sans-serif; font-size: 0.78rem; font-weight: 600; color: var(--text-muted); cursor: pointer; transition: all 0.2s; }
.bulk-btn.danger:hover { background: rgba(220,60,60,0.08); border-color: rgba(220,60,60,0.2); color: #cc6060; }
.bulk-clear { background: none; border: none; color: var(--text-faint); font-size: 0.78rem; cursor: pointer; font-family: 'Syne', sans-serif; }
.overlay { position: fixed; inset: 0; background: rgba(0,0,0,0.7); display: flex; align-items: center; justify-content: center; z-index: 1000; padding: 1rem; backdrop-filter: blur(4px); animation: fadeIn 0.15s ease; }
.modal { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 16px; width: 100%; max-width: 480px; overflow: hidden; animation: slideUp 0.2s ease; }
.modal-sm { max-width: 380px; }
.modal-head { display: flex; align-items: center; justify-content: space-between; padding: 1.25rem 1.5rem; border-bottom: 1px solid var(--border-soft); }
.modal-head h3 { font-size: 1rem; font-weight: 700; color: var(--text-primary); }
.modal-close { background: none; border: none; color: var(--text-faint); cursor: pointer; font-size: 1rem; padding: 0.25rem; transition: color 0.2s; }
.modal-close:hover { color: var(--text-primary); }
.modal-body { padding: 1.5rem; display: flex; flex-direction: column; gap: 1.25rem; }
.modal-foot { display: flex; gap: 0.75rem; justify-content: flex-end; padding: 1.25rem 1.5rem; border-top: 1px solid var(--border-soft); }
.profile-hero { display: flex; align-items: center; gap: 1rem; }
.profile-av { width: 52px; height: 52px; border-radius: 50%; background: var(--bg-overlay); border: 1px solid var(--border-mid); display: flex; align-items: center; justify-content: center; font-size: 1rem; font-weight: 700; color: var(--text-muted); flex-shrink: 0; }
.profile-name { font-size: 1rem; font-weight: 700; color: var(--text-primary); margin-bottom: 0.2rem; }
.profile-email { font-size: 0.78rem; color: var(--text-faint); font-family: 'JetBrains Mono', monospace; margin-bottom: 0.5rem; }
.profile-badges { display: flex; gap: 0.5rem; flex-wrap: wrap; }
.detail-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 0.75rem; }
.detail-item { background: var(--bg-elevated); border: 1px solid var(--border-soft); border-radius: 8px; padding: 0.75rem; }
.detail-label { display: block; font-size: 0.68rem; font-weight: 700; color: var(--text-ghost); text-transform: uppercase; letter-spacing: 0.06em; margin-bottom: 0.3rem; }
.detail-val { font-size: 0.85rem; color: var(--text-secondary); font-weight: 600; }
.detail-val.mono { font-family: 'JetBrains Mono', monospace; font-size: 0.78rem; }
.form-grid { display: flex; flex-direction: column; gap: 1rem; }
.form-field { display: flex; flex-direction: column; gap: 0.4rem; }
.form-field label { font-size: 0.78rem; font-weight: 600; color: var(--text-faint); }
.form-field input, .form-field select { background: var(--bg-elevated); border: 1px solid var(--border); border-radius: 8px; padding: 0.7rem 0.875rem; color: var(--text-primary); font-family: 'Syne', sans-serif; font-size: 0.88rem; outline: none; transition: border-color 0.2s; }
.form-field input:focus, .form-field select:focus { border-color: var(--border-mid); }
.form-field select option { background: var(--bg-elevated); }
.confirm-icon { display: flex; justify-content: center; }
.confirm-icon svg { width: 40px; height: 40px; color: #cc6060; }
.confirm-text { font-size: 0.88rem; color: var(--text-muted); line-height: 1.6; text-align: center; }
.confirm-text strong { color: var(--text-primary); }
.mbtn { padding: 0.65rem 1.25rem; border-radius: 8px; font-family: 'Syne', sans-serif; font-size: 0.85rem; font-weight: 700; cursor: pointer; transition: all 0.2s; border: none; display: flex; align-items: center; gap: 0.4rem; }
.mbtn-outline { background: transparent; border: 1px solid var(--border-mid); color: var(--text-muted); }
.mbtn-outline:hover { color: var(--text-secondary); }
.mbtn-primary { background: var(--accent); color: var(--bg-base); }
.mbtn-primary:hover:not(:disabled) { opacity: 0.85; }
.mbtn-primary:disabled { opacity: 0.5; cursor: not-allowed; }
.mbtn-danger { background: rgba(220,60,60,0.12); color: #e07070; border: 1px solid rgba(220,60,60,0.2); }
.mbtn-danger:hover { background: rgba(220,60,60,0.2); }
.spinner { width: 14px; height: 14px; border: 2px solid rgba(0,0,0,0.2); border-top-color: var(--bg-base); border-radius: 50%; animation: spin 0.7s linear infinite; }
.toast { position: fixed; bottom: 1.5rem; right: 1.5rem; padding: 0.875rem 1.25rem; border-radius: 10px; font-family: 'Syne', sans-serif; font-size: 0.85rem; font-weight: 600; z-index: 2000; animation: slideUp 0.2s ease; }
.toast.success { background: var(--bg-surface); border: 1px solid rgba(76,175,114,0.3); color: #4caf72; }
.toast.danger  { background: var(--bg-surface); border: 1px solid rgba(220,60,60,0.3); color: #e07070; }
@keyframes fadeIn { from { opacity: 0; } to { opacity: 1; } }
@keyframes slideUp { from { opacity: 0; transform: translateY(12px); } to { opacity: 1; transform: translateY(0); } }
@keyframes spin { to { transform: rotate(360deg); } }
@media (max-width: 768px) {
  .detail-grid { grid-template-columns: 1fr; }
  .filters { flex-direction: column; align-items: stretch; }
  .search-wrap { max-width: 100%; }
  .page-header { flex-direction: column; align-items: flex-start; }
  .table th:nth-child(5), .table td:nth-child(5),
  .table th:nth-child(6), .table td:nth-child(6) { display: none; }
  .modal { max-width: 100%; border-radius: 16px 16px 0 0; position: fixed; bottom: 0; left: 0; right: 0; }
  .overlay { align-items: flex-end; padding: 0; }
  .modal-foot { flex-wrap: wrap; }
  .mbtn { flex: 1; justify-content: center; }
}
</style>