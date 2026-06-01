<template>
  <div class="auth-wrap">
    <div class="left-panel">
      <router-link to="/" class="logo">dev<span class="accent">forum</span></router-link>
      <div class="left-content">
        <div class="perks">
          <div class="perk" v-for="p in perks" :key="p.title">
            <div class="perk-icon" v-html="p.icon"></div>
            <div>
              <div class="perk-title">{{ p.title }}</div>
              <div class="perk-desc">{{ p.desc }}</div>
            </div>
          </div>
        </div>
        <div class="testimonial">
          <p class="t-quote">"DevForum is the only place where I get thoughtful answers instead of copy-pasted docs."</p>
          <div class="t-author">
            <div class="t-avatar">AK</div>
            <div>
              <div class="t-name">Alex Kim</div>
              <div class="t-role">Senior Engineer @ Stripe</div>
            </div>
          </div>
        </div>
      </div>
      <p class="left-tagline">Join 12,000+ developers already on DevForum.</p>
    </div>

    <div class="right-panel">
      <div class="form-card">
        <div class="form-header">
          <h1>Create account</h1>
          <p>Free forever. No credit card required.</p>
        </div>

        <!-- Steps -->
        <div class="steps">
          <div class="step" :class="{ active: step >= 1, done: step > 1 }">
            <div class="step-circle">
              <svg v-if="step > 1" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="20 6 9 17 4 12"/></svg>
              <span v-else>1</span>
            </div>
            <span class="step-label">Account</span>
          </div>
          <div class="step-line" :class="{ done: step > 1 }"></div>
          <div class="step" :class="{ active: step >= 2 }">
            <div class="step-circle"><span>2</span></div>
            <span class="step-label">Security</span>
          </div>
        </div>

        <!-- Step 1 -->
        <div class="form" v-if="step === 1">
          <div class="field" :class="{ focused: focused === 'name' }">
            <label>Full name</label>
            <div class="input-wrap">
              <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
              <input v-model="form.fullName" type="text" placeholder="Jane Doe" @focus="focused='name'" @blur="focused=''" />
            </div>
          </div>
          <div class="field" :class="{ focused: focused === 'user' }">
            <label>Username</label>
            <div class="input-wrap">
              <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
              <input v-model="form.username" type="text" placeholder="janedoe" @focus="focused='user'" @blur="focused=''" />
              <div class="field-tag" v-if="form.username">@{{ form.username }}</div>
            </div>
          </div>
          <div class="field" :class="{ focused: focused === 'email' }">
            <label>Email address</label>
            <div class="input-wrap">
              <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/><polyline points="22,6 12,13 2,6"/></svg>
              <input v-model="form.email" type="email" placeholder="you@email.com" @focus="focused='email'" @blur="focused=''" />
            </div>
          </div>
          <div v-if="stepError" class="alert-error">{{ stepError }}</div>
          <button type="button" class="btn-submit" @click="nextStep">
            Continue <span class="btn-arrow">→</span>
          </button>
        </div>

        <!-- Step 2 -->
        <div class="form" v-if="step === 2">
          <div class="field" :class="{ focused: focused === 'pw' }">
            <label>Password</label>
            <div class="input-wrap">
              <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><rect x="3" y="11" width="18" height="11" rx="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></svg>
              <input v-model="form.password" :type="showPw ? 'text' : 'password'" placeholder="min. 8 characters" @focus="focused='pw'" @blur="focused=''" />
              <button type="button" class="eye-btn" @click="showPw = !showPw" tabindex="-1">
                <svg v-if="!showPw" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
                <svg v-else viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"/><line x1="1" y1="1" x2="23" y2="23"/></svg>
              </button>
            </div>
            <div class="strength" v-if="form.password">
              <div class="strength-bars">
                <div v-for="i in 4" :key="i" class="strength-bar" :class="{ active: passwordStrength >= i }" :style="{ background: passwordStrength >= i ? strengthColor : '' }"></div>
              </div>
              <span class="strength-text" :style="{ color: strengthColor }">{{ strengthLabel }}</span>
            </div>
            <div class="pw-rules">
              <div class="rule" :class="{ pass: form.password.length >= 8 }">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="20 6 9 17 4 12"/></svg>
                At least 8 characters
              </div>
              <div class="rule" :class="{ pass: /[A-Z]/.test(form.password) }">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="20 6 9 17 4 12"/></svg>
                One uppercase letter
              </div>
              <div class="rule" :class="{ pass: /[0-9]/.test(form.password) }">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="20 6 9 17 4 12"/></svg>
                One number
              </div>
            </div>
          </div>

          <div v-if="error" class="alert-error">{{ error }}</div>
          <div v-if="success" class="alert-success">{{ success }}</div>
          <div v-if="debugMsg" class="alert-debug">{{ debugMsg }}</div>

          <div class="step2-actions">
            <button type="button" class="btn-back" @click="step = 1">← Back</button>
            <button type="button" class="btn-submit" :disabled="loading" @click="handleSubmit" style="flex:1">
              <span v-if="loading" class="spinner"></span>
              <template v-else>Create account <span class="btn-arrow">→</span></template>
            </button>
          </div>
        </div>

        <p class="terms">By creating an account you agree to our <a href="#">Terms</a> and <a href="#">Privacy Policy</a>.</p>
        <p class="switch-link">Already have an account? <router-link to="/login">Sign in</router-link></p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import authService from '../../services/auth.js'

const router = useRouter()
const step = ref(1)
const showPw = ref(false)
const loading = ref(false)
const error = ref('')
const stepError = ref('')
const success = ref('')
const debugMsg = ref('')
const focused = ref('')
const form = ref({ fullName: '', username: '', email: '', password: '' })

const perks = [
  { icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" width="18" height="18"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>`, title: 'Ask anything', desc: 'Post questions and get answers from real developers.' },
  { icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" width="18" height="18"><polyline points="17 1 21 5 17 9"/><path d="M3 11V9a4 4 0 0 1 4-4h14"/><polyline points="7 23 3 19 7 15"/><path d="M21 13v2a4 4 0 0 1-4 4H3"/></svg>`, title: 'Share knowledge', desc: 'Write posts and help others level up.' },
  { icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" width="18" height="18"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/></svg>`, title: 'Build reputation', desc: 'Earn upvotes and grow your profile.' },
]

const passwordStrength = computed(() => {
  const p = form.value.password
  if (!p) return 0
  let s = 0
  if (p.length >= 8) s++
  if (/[A-Z]/.test(p)) s++
  if (/[0-9]/.test(p)) s++
  if (/[^A-Za-z0-9]/.test(p)) s++
  return s
})
const strengthColor = computed(() => ['#e05555','#e08c30','#d4c040','#7cfc9f'][passwordStrength.value - 1] || '#2a2a3e')
const strengthLabel = computed(() => ['Weak','Fair','Good','Strong'][passwordStrength.value - 1] || '')

function nextStep() {
  stepError.value = ''
  if (!form.value.fullName || !form.value.username || !form.value.email) {
    stepError.value = 'Please fill in all fields.'
    return
  }
  step.value = 2
}

async function handleSubmit() {
  error.value = ''
  debugMsg.value = ''
  if (!form.value.password || form.value.password.length < 8) {
    error.value = 'Password must be at least 8 characters.'
    return
  }
  loading.value = true
  debugMsg.value = 'Creating account...'
  try {
    await authService.register(form.value)
    success.value = 'Account created! Redirecting to login...'
    debugMsg.value = ''
    setTimeout(() => router.push('/login'), 1500)
  } catch (e) {
    error.value = e.response?.data?.error || e.message || 'Registration failed. Please try again.'
    debugMsg.value = ''
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500&family=Syne:wght@400;600;700;800&display=swap');
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

.auth-wrap { display: flex; min-height: 100vh; font-family: 'Syne', sans-serif; background: #09090e; }

.left-panel {
  width: 45%; background: #0d0d15; border-right: 1px solid #1c1c2a;
  padding: 2.5rem; display: flex; flex-direction: column; justify-content: space-between;
  position: relative; overflow: hidden;
}
.left-panel::before {
  content: ''; position: absolute; inset: 0; pointer-events: none;
  background-image: linear-gradient(rgba(232,234,240,0.025) 1px, transparent 1px), linear-gradient(90deg, rgba(232,234,240,0.025) 1px, transparent 1px);
  background-size: 48px 48px;
}
.logo { font-size: 1.5rem; font-weight: 800; color: #e2e2e8; text-decoration: none; letter-spacing: -0.5px; }
.accent { color: #e8eaf0; }
.left-content { flex: 1; display: flex; flex-direction: column; justify-content: center; gap: 2.5rem; }
.perks { display: flex; flex-direction: column; gap: 1.5rem; position: relative; z-index: 1; }
.perk { display: flex; align-items: flex-start; gap: 1rem; }
.perk-icon { width: 36px; height: 36px; border-radius: 8px; background: #161622; border: 1px solid #1e1e2e; display: flex; align-items: center; justify-content: center; flex-shrink: 0; color: #8888a0; }
.perk-title { font-size: 0.9rem; font-weight: 600; color: #c0c0d0; margin-bottom: 0.2rem; }
.perk-desc { font-size: 0.8rem; color: #44445a; line-height: 1.5; }
.testimonial { background: #111119; border: 1px solid #1e1e2e; border-radius: 12px; padding: 1.25rem; position: relative; z-index: 1; }
.t-quote { font-size: 0.85rem; color: #8888a0; line-height: 1.7; margin-bottom: 1rem; font-style: italic; }
.t-author { display: flex; align-items: center; gap: 0.75rem; }
.t-avatar { width: 32px; height: 32px; border-radius: 50%; background: #1e1e30; border: 1px solid #2e2e42; display: flex; align-items: center; justify-content: center; font-size: 0.7rem; font-weight: 700; color: #8888a8; }
.t-name { font-size: 0.82rem; font-weight: 600; color: #c0c0d0; }
.t-role { font-size: 0.75rem; color: #44445a; }
.left-tagline { font-size: 0.82rem; color: #2e2e42; position: relative; z-index: 1; }

.right-panel { flex: 1; display: flex; align-items: center; justify-content: center; padding: 2rem; background: #09090e; overflow-y: auto; }
.form-card { width: 100%; max-width: 420px; animation: fadeUp 0.5s ease both; padding: 0.5rem 0; }
.form-header { margin-bottom: 1.75rem; }
.form-header h1 { font-size: 2rem; font-weight: 800; color: #e2e2e8; letter-spacing: -1px; margin-bottom: 0.35rem; }
.form-header p { font-size: 0.875rem; color: #44445a; }

.steps { display: flex; align-items: center; margin-bottom: 2rem; }
.step { display: flex; align-items: center; gap: 0.5rem; }
.step-circle { width: 28px; height: 28px; border-radius: 50%; background: #111119; border: 1px solid #2a2a3e; display: flex; align-items: center; justify-content: center; font-size: 0.75rem; font-weight: 700; color: #44445a; transition: all 0.3s; }
.step-circle svg { width: 12px; height: 12px; }
.step.active .step-circle { border-color: #e8eaf0; color: #e8eaf0; }
.step.done .step-circle { background: #e8eaf0; border-color: #e8eaf0; color: #09090e; }
.step-label { font-size: 0.78rem; color: #44445a; }
.step.active .step-label { color: #c0c0d0; }
.step-line { flex: 1; height: 1px; background: #1a1a28; margin: 0 0.75rem; transition: background 0.3s; }
.step-line.done { background: #e8eaf0; }

.form { display: flex; flex-direction: column; gap: 1rem; }
.field { display: flex; flex-direction: column; gap: 0.45rem; }
.field label { font-size: 0.8rem; font-weight: 600; color: #55556a; transition: color 0.2s; }
.field.focused label { color: #b0b0c8; }
.input-wrap { position: relative; }
.input-icon { position: absolute; left: 0.875rem; top: 50%; transform: translateY(-50%); width: 16px; height: 16px; color: #33334a; pointer-events: none; transition: color 0.2s; }
.field.focused .input-icon { color: #8888a8; }
.input-wrap input { width: 100%; background: #111119; border: 1px solid #1e1e2e; border-radius: 10px; padding: 0.8rem 1rem 0.8rem 2.75rem; color: #e2e2e8; font-family: 'Syne', sans-serif; font-size: 0.92rem; outline: none; transition: border-color 0.2s; }
.input-wrap input:focus { border-color: #3a3a54; background: #14141e; }
.input-wrap input::placeholder { color: #2a2a3e; }
.field-tag { position: absolute; right: 0.875rem; top: 50%; transform: translateY(-50%); font-family: 'JetBrains Mono', monospace; font-size: 0.72rem; color: #44445a; }
.eye-btn { position: absolute; right: 0.875rem; top: 50%; transform: translateY(-50%); background: none; border: none; cursor: pointer; padding: 0; color: #33334a; transition: color 0.2s; }
.eye-btn:hover { color: #8888a8; }
.eye-btn svg { width: 16px; height: 16px; display: block; }

.strength { display: flex; align-items: center; gap: 0.75rem; margin-top: 0.5rem; }
.strength-bars { display: flex; gap: 4px; flex: 1; }
.strength-bar { flex: 1; height: 3px; border-radius: 2px; background: #1e1e2e; transition: background 0.3s; }
.strength-text { font-size: 0.75rem; font-weight: 600; min-width: 36px; }

.pw-rules { display: flex; flex-direction: column; gap: 0.4rem; margin-top: 0.6rem; }
.rule { display: flex; align-items: center; gap: 0.5rem; font-size: 0.78rem; color: #2e2e42; transition: color 0.3s; }
.rule svg { width: 12px; height: 12px; opacity: 0.3; transition: opacity 0.3s; flex-shrink: 0; }
.rule.pass { color: #6888a0; }
.rule.pass svg { opacity: 1; color: #7cfc9f; }

.alert-error { background: rgba(220,60,60,0.08); border: 1px solid rgba(220,60,60,0.18); color: #e07070; border-radius: 10px; padding: 0.75rem 1rem; font-size: 0.85rem; }
.alert-success { background: rgba(124,252,159,0.06); border: 1px solid rgba(124,252,159,0.15); color: #7cfc9f; border-radius: 10px; padding: 0.75rem 1rem; font-size: 0.85rem; }
.alert-debug { background: rgba(100,180,255,0.06); border: 1px solid rgba(100,180,255,0.15); color: #88b8e0; border-radius: 10px; padding: 0.75rem 1rem; font-size: 0.85rem; }

.btn-submit { background: #e8eaf0; color: #09090e; border: none; border-radius: 10px; padding: 0.875rem; font-family: 'Syne', sans-serif; font-size: 0.95rem; font-weight: 700; cursor: pointer; transition: background 0.2s, transform 0.15s; display: flex; align-items: center; justify-content: center; gap: 0.5rem; margin-top: 0.5rem; width: 100%; }
.btn-submit:hover:not(:disabled) { background: #ffffff; transform: translateY(-1px); }
.btn-submit:disabled { opacity: 0.5; cursor: not-allowed; }
.btn-arrow { transition: transform 0.2s; }
.btn-submit:hover .btn-arrow { transform: translateX(3px); }
.spinner { width: 18px; height: 18px; border: 2px solid rgba(9,9,14,0.2); border-top-color: #09090e; border-radius: 50%; animation: spin 0.7s linear infinite; }

.step2-actions { display: flex; gap: 0.75rem; align-items: center; margin-top: 0.5rem; }
.btn-back { background: none; border: 1px solid #1e1e2e; border-radius: 10px; padding: 0.875rem 1rem; color: #55556a; font-family: 'Syne', sans-serif; font-size: 0.9rem; cursor: pointer; white-space: nowrap; transition: border-color 0.2s, color 0.2s; }
.btn-back:hover { border-color: #2e2e42; color: #8888a0; }

.terms { text-align: center; margin-top: 1.25rem; font-size: 0.78rem; color: #2e2e42; line-height: 1.6; }
.terms a { color: #44445a; text-decoration: underline; }
.switch-link { text-align: center; margin-top: 1rem; font-size: 0.875rem; color: #44445a; }
.switch-link a { color: #c0c0d8; text-decoration: none; font-weight: 600; }
.switch-link a:hover { color: #e8eaf0; }

@keyframes fadeUp { from { opacity: 0; transform: translateY(14px); } to { opacity: 1; transform: translateY(0); } }
@keyframes spin { to { transform: rotate(360deg); } }

@media (max-width: 768px) {
  .auth-wrap { flex-direction: column; }
  .left-panel { width: 100%; min-height: auto; padding: 1.5rem; }
  .testimonial { display: none; }
  .right-panel { padding: 1.5rem 1rem; align-items: flex-start; }
}
</style>