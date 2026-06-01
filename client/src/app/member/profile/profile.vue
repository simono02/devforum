<template>
  <div class="page">
    <h1 class="page-title">Profile</h1>

    <div class="profile-grid">
      <!-- Profile card -->
      <div class="profile-card">
        <div class="profile-hero">
          <div class="profile-av">{{ initials }}</div>
          <div class="edit-av-btn" title="Change avatar">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/></svg>
          </div>
        </div>
        <div class="profile-info">
          <h2 class="profile-name">{{ user?.full_name }}</h2>
          <div class="profile-username">@{{ user?.username }}</div>
          <div class="profile-badges">
            <span class="badge member">Member</span>
            <span class="badge active">Active</span>
          </div>
        </div>
        <div class="profile-stats">
          <div class="ps"><span class="ps-num">0</span><span class="ps-label">Posts</span></div>
          <div class="ps-div"></div>
          <div class="ps"><span class="ps-num">0</span><span class="ps-label">Comments</span></div>
          <div class="ps-div"></div>
          <div class="ps"><span class="ps-num">0</span><span class="ps-label">Upvotes</span></div>
        </div>
      </div>

      <!-- Edit form -->
      <div class="edit-card">
        <div class="edit-head">
          <h3>Account Details</h3>
          <span class="edit-note">Changes are saved immediately</span>
        </div>

        <div class="form">
          <div class="form-row">
            <div class="field">
              <label>Full Name</label>
              <input v-model="form.fullName" type="text" :placeholder="user?.full_name" />
            </div>
            <div class="field">
              <label>Username</label>
              <div class="input-prefix">
                <span>@</span>
                <input v-model="form.username" type="text" :placeholder="user?.username" />
              </div>
            </div>
          </div>
          <div class="field">
            <label>Email Address</label>
            <input v-model="form.email" type="email" :placeholder="user?.email" />
          </div>
          <div class="field">
            <label>Bio <span class="label-opt">optional</span></label>
            <textarea v-model="form.bio" rows="3" placeholder="Tell the community a bit about yourself..."></textarea>
          </div>

          <div v-if="saveMsg" class="save-msg" :class="saveMsgType">{{ saveMsg }}</div>

          <div class="form-foot">
            <button class="btn-save" @click="saveProfile" :disabled="saving">
              <span v-if="saving" class="spinner"></span>
              <span v-else>Save Changes</span>
            </button>
          </div>
        </div>

        <div class="section-divider"></div>

        <!-- Change password -->
        <div class="edit-head">
          <h3>Change Password</h3>
        </div>
        <div class="form">
          <div class="field">
            <label>Current Password</label>
            <input v-model="pw.current" type="password" placeholder="••••••••" />
          </div>
          <div class="form-row">
            <div class="field">
              <label>New Password</label>
              <input v-model="pw.new1" type="password" placeholder="min. 8 characters" />
            </div>
            <div class="field">
              <label>Confirm Password</label>
              <input v-model="pw.new2" type="password" placeholder="repeat password" />
            </div>
          </div>
          <div class="form-foot">
            <button class="btn-outline" @click="changePassword">Update Password</button>
          </div>
        </div>

        <div class="section-divider"></div>

        <!-- Danger zone -->
        <div class="danger-zone">
          <div class="dz-text">
            <h4>Delete Account</h4>
            <p>Permanently delete your account and all your posts. This cannot be undone.</p>
          </div>
          <button class="btn-danger" @click="showDeleteModal = true">Delete Account</button>
        </div>
      </div>
    </div>

    <!-- Delete confirm modal -->
    <Teleport to="body">
      <div class="overlay" v-if="showDeleteModal" @click.self="showDeleteModal = false">
        <div class="modal">
          <div class="modal-head">
            <h3>Delete Account</h3>
            <button class="modal-close" @click="showDeleteModal = false">✕</button>
          </div>
          <div class="modal-body">
            <div class="confirm-icon">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
            </div>
            <p class="confirm-text">Are you absolutely sure? All your posts and data will be permanently deleted.</p>
            <div class="field" style="margin-top:1rem">
              <label>Type <strong>delete</strong> to confirm</label>
              <input v-model="deleteConfirm" type="text" placeholder="delete" />
            </div>
          </div>
          <div class="modal-foot">
            <button class="mbtn mbtn-outline" @click="showDeleteModal = false">Cancel</button>
            <button class="mbtn mbtn-danger" :disabled="deleteConfirm !== 'delete'" @click="deleteAccount">Delete Account</button>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../../../stores/user.js'

const router    = useRouter()
const userStore = useUserStore()
const user      = computed(() => userStore.user)

const initials = computed(() => {
  const name = user.value?.full_name || 'U'
  return name.split(' ').map(n => n[0]).join('').slice(0,2).toUpperCase()
})

const saving  = ref(false)
const saveMsg = ref('')
const saveMsgType = ref('success')
const showDeleteModal = ref(false)
const deleteConfirm   = ref('')

const form = ref({ fullName: '', username: '', email: '', bio: '' })
const pw   = ref({ current: '', new1: '', new2: '' })

async function saveProfile() {
  saving.value = true
  await new Promise(r => setTimeout(r, 600))
  saveMsg.value = 'Profile saved successfully.'
  saveMsgType.value = 'success'
  saving.value = false
  setTimeout(() => saveMsg.value = '', 3000)
}

function changePassword() {
  if (!pw.value.current) { alert('Enter your current password.'); return }
  if (pw.value.new1 !== pw.value.new2) { alert('Passwords do not match.'); return }
  if (pw.value.new1.length < 8) { alert('Password must be at least 8 characters.'); return }
  pw.value = { current: '', new1: '', new2: '' }
  saveMsg.value = 'Password updated.'
  saveMsgType.value = 'success'
  setTimeout(() => saveMsg.value = '', 3000)
}

function deleteAccount() {
  userStore.clearAuth()
  router.push('/')
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Syne:wght@400;600;700;800&display=swap');
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
.page { font-family: 'Syne', sans-serif; display: flex; flex-direction: column; gap: 1.5rem; animation: fadeIn 0.3s ease; }

.page-title { font-size: 1.75rem; font-weight: 800; color: var(--text-primary); letter-spacing: -1px; }

.profile-grid { display: grid; grid-template-columns: 260px 1fr; gap: 1.25rem; align-items: start; }

/* Profile card */
.profile-card { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 16px; overflow: hidden; }
.profile-hero { background: var(--bg-elevated); padding: 2rem 1.5rem 1rem; display: flex; justify-content: center; position: relative; }
.profile-av { width: 72px; height: 72px; border-radius: 50%; background: var(--bg-overlay); border: 3px solid var(--bg-surface); display: flex; align-items: center; justify-content: center; font-size: 1.3rem; font-weight: 800; color: var(--text-muted); }
.edit-av-btn {
  position: absolute; bottom: 1rem; right: 1.5rem;
  width: 28px; height: 28px; border-radius: 50%;
  background: var(--bg-surface); border: 1px solid var(--border);
  display: flex; align-items: center; justify-content: center;
  cursor: pointer; color: var(--text-faint); transition: color 0.2s;
}
.edit-av-btn svg { width: 12px; height: 12px; }
.edit-av-btn:hover { color: var(--text-muted); }
.profile-info { padding: 1rem 1.5rem; text-align: center; border-bottom: 1px solid var(--border-soft); }
.profile-name { font-size: 1rem; font-weight: 700; color: var(--text-primary); margin-bottom: 0.2rem; }
.profile-username { font-size: 0.8rem; color: var(--text-faint); margin-bottom: 0.6rem; font-family: monospace; }
.profile-badges { display: flex; gap: 0.4rem; justify-content: center; }
.badge { font-size: 0.68rem; font-weight: 700; padding: 0.2rem 0.6rem; border-radius: 20px; }
.badge.member { background: var(--bg-elevated); color: var(--text-muted); border: 1px solid var(--border); }
.badge.active { background: rgba(76,175,114,0.1); color: #4caf72; }
.profile-stats { display: flex; align-items: center; justify-content: space-around; padding: 1rem 1.5rem; }
.ps { text-align: center; }
.ps-num { display: block; font-size: 1.2rem; font-weight: 800; color: var(--text-primary); }
.ps-label { font-size: 0.72rem; color: var(--text-faint); }
.ps-div { width: 1px; height: 28px; background: var(--border); }

/* Edit card */
.edit-card { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 16px; padding: 1.5rem; display: flex; flex-direction: column; gap: 1.25rem; }
.edit-head { display: flex; align-items: center; justify-content: space-between; }
.edit-head h3 { font-size: 0.9rem; font-weight: 700; color: var(--text-secondary); }
.edit-note { font-size: 0.72rem; color: var(--text-ghost); }
.section-divider { height: 1px; background: var(--border-soft); }

/* Form */
.form { display: flex; flex-direction: column; gap: 1rem; }
.form-row { display: grid; grid-template-columns: 1fr 1fr; gap: 1rem; }
.field { display: flex; flex-direction: column; gap: 0.4rem; }
.field label { font-size: 0.78rem; font-weight: 600; color: var(--text-faint); }
.label-opt { font-weight: 400; color: var(--text-ghost); }
.field input, .field textarea {
  background: var(--bg-elevated); border: 1px solid var(--border);
  border-radius: 8px; padding: 0.7rem 0.875rem;
  color: var(--text-primary); font-family: 'Syne', sans-serif; font-size: 0.88rem;
  outline: none; transition: border-color 0.2s; resize: vertical;
}
.field input:focus, .field textarea:focus { border-color: var(--border-mid); }
.field input::placeholder, .field textarea::placeholder { color: var(--text-ghost); }
.input-prefix { display: flex; align-items: center; background: var(--bg-elevated); border: 1px solid var(--border); border-radius: 8px; overflow: hidden; transition: border-color 0.2s; }
.input-prefix:focus-within { border-color: var(--border-mid); }
.input-prefix span { padding: 0 0.75rem; font-size: 0.88rem; color: var(--text-faint); background: var(--bg-overlay); border-right: 1px solid var(--border); height: 100%; display: flex; align-items: center; white-space: nowrap; }
.input-prefix input { background: none; border: none; padding: 0.7rem 0.875rem; color: var(--text-primary); font-family: 'Syne', sans-serif; font-size: 0.88rem; outline: none; flex: 1; }

.save-msg { padding: 0.65rem 1rem; border-radius: 8px; font-size: 0.82rem; }
.save-msg.success { background: rgba(76,175,114,0.08); border: 1px solid rgba(76,175,114,0.2); color: #4caf72; }
.save-msg.error   { background: rgba(220,60,60,0.08); border: 1px solid rgba(220,60,60,0.2); color: #e07070; }

.form-foot { display: flex; justify-content: flex-end; }
.btn-save {
  background: var(--accent); color: var(--bg-base);
  border: none; border-radius: 8px; padding: 0.65rem 1.5rem;
  font-family: 'Syne', sans-serif; font-size: 0.88rem; font-weight: 700;
  cursor: pointer; transition: opacity 0.2s; display: flex; align-items: center; gap: 0.5rem;
}
.btn-save:hover:not(:disabled) { opacity: 0.85; }
.btn-save:disabled { opacity: 0.5; cursor: not-allowed; }
.btn-outline {
  background: none; border: 1px solid var(--border-mid); border-radius: 8px;
  padding: 0.65rem 1.25rem; font-family: 'Syne', sans-serif; font-size: 0.88rem;
  font-weight: 600; color: var(--text-muted); cursor: pointer; transition: all 0.2s;
}
.btn-outline:hover { border-color: var(--border-mid); color: var(--text-primary); }
.spinner { width: 16px; height: 16px; border: 2px solid rgba(0,0,0,0.2); border-top-color: var(--bg-base); border-radius: 50%; animation: spin 0.7s linear infinite; }

/* Danger zone */
.danger-zone { background: rgba(220,60,60,0.04); border: 1px solid rgba(220,60,60,0.15); border-radius: 10px; padding: 1rem 1.25rem; display: flex; align-items: center; justify-content: space-between; gap: 1rem; flex-wrap: wrap; }
.dz-text h4 { font-size: 0.88rem; font-weight: 700; color: #e07070; margin-bottom: 0.2rem; }
.dz-text p { font-size: 0.78rem; color: var(--text-faint); }
.btn-danger { background: rgba(220,60,60,0.1); color: #e07070; border: 1px solid rgba(220,60,60,0.2); border-radius: 8px; padding: 0.55rem 1rem; font-family: 'Syne', sans-serif; font-size: 0.82rem; font-weight: 700; cursor: pointer; transition: background 0.2s; white-space: nowrap; }
.btn-danger:hover { background: rgba(220,60,60,0.2); }

/* Modal */
.overlay { position: fixed; inset: 0; background: rgba(0,0,0,0.6); display: flex; align-items: center; justify-content: center; z-index: 1000; padding: 1rem; backdrop-filter: blur(4px); }
.modal { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 16px; width: 100%; max-width: 420px; overflow: hidden; animation: slideUp 0.2s ease; }
.modal-head { display: flex; align-items: center; justify-content: space-between; padding: 1.25rem 1.5rem; border-bottom: 1px solid var(--border-soft); }
.modal-head h3 { font-size: 1rem; font-weight: 700; color: var(--text-primary); }
.modal-close { background: none; border: none; color: var(--text-faint); cursor: pointer; font-size: 1rem; }
.modal-body { padding: 1.5rem; display: flex; flex-direction: column; gap: 1rem; }
.modal-foot { display: flex; gap: 0.75rem; justify-content: flex-end; padding: 1.25rem 1.5rem; border-top: 1px solid var(--border-soft); }
.confirm-icon { display: flex; justify-content: center; }
.confirm-icon svg { width: 40px; height: 40px; color: #e07070; }
.confirm-text { font-size: 0.88rem; color: var(--text-muted); line-height: 1.6; text-align: center; }
.mbtn { padding: 0.65rem 1.25rem; border-radius: 8px; font-family: 'Syne', sans-serif; font-size: 0.85rem; font-weight: 700; cursor: pointer; border: none; transition: all 0.2s; }
.mbtn-outline { background: transparent; border: 1px solid var(--border-mid); color: var(--text-muted); }
.mbtn-outline:hover { color: var(--text-primary); }
.mbtn-danger { background: rgba(220,60,60,0.12); color: #e07070; border: 1px solid rgba(220,60,60,0.2); }
.mbtn-danger:hover:not(:disabled) { background: rgba(220,60,60,0.2); }
.mbtn-danger:disabled { opacity: 0.4; cursor: not-allowed; }

@keyframes fadeIn { from { opacity: 0; } to { opacity: 1; } }
@keyframes slideUp { from { opacity: 0; transform: translateY(12px); } to { opacity: 1; transform: translateY(0); } }
@keyframes spin { to { transform: rotate(360deg); } }

@media (max-width: 900px) { .profile-grid { grid-template-columns: 1fr; } }
@media (max-width: 640px) {
  .form-row { grid-template-columns: 1fr; }
  .danger-zone { flex-direction: column; align-items: flex-start; }
  .btn-danger { width: 100%; text-align: center; }
}
</style>