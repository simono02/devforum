<template>
  <div class="auth-wrap">
    <div class="left-panel">
      <router-link to="/" class="logo">dev<span class="accent">forum</span></router-link>
      <div class="left-content">
        <div class="terminal">
          <div class="terminal-bar">
            <span class="dot r"></span><span class="dot y"></span><span class="dot g"></span>
            <span class="terminal-title">session.go</span>
          </div>
          <pre class="terminal-code"><span class="c-kw">func</span> <span class="c-fn">Login</span>(c *gin.Context) {
  <span class="c-kw">var</span> req LoginRequest
  c.BindJSON(&amp;req)

  user, err := <span class="c-fn">FindUser</span>(req.Email)
  <span class="c-kw">if</span> err != nil {
    c.JSON(<span class="c-num">401</span>, <span class="c-str">"invalid credentials"</span>)
    <span class="c-kw">return</span>
  }

  token := <span class="c-fn">GenerateJWT</span>(user.ID)
  c.JSON(<span class="c-num">200</span>, gin.H{
    <span class="c-str">"token"</span>: token,
    <span class="c-str">"user"</span>:  user,
  })
}</pre>
        </div>
        <div class="left-stats">
          <div class="ls"><span class="ls-n">12k+</span><span class="ls-l">Developers</span></div>
          <div class="ls-sep"></div>
          <div class="ls"><span class="ls-n">48k+</span><span class="ls-l">Posts</span></div>
          <div class="ls-sep"></div>
          <div class="ls"><span class="ls-n">200k+</span><span class="ls-l">Comments</span></div>
        </div>
      </div>
      <p class="left-tagline">Where developers think out loud.</p>
    </div>

    <div class="right-panel">
      <div class="form-card">
        <div class="form-header">
          <h1>Welcome back</h1>
          <p>Sign in to continue to DevForum</p>
        </div>

        <div class="form">
          <div class="field" :class="{ focused: focused === 'id' }">
            <label>Email or username</label>
            <div class="input-wrap">
              <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
              <input
                v-model="form.identifier"
                type="text"
                placeholder="you@email.com"
                autocomplete="username"
                @focus="focused = 'id'"
                @blur="focused = ''"
              />
            </div>
          </div>

          <div class="field" :class="{ focused: focused === 'pw' }">
            <label>Password</label>
            <div class="input-wrap">
              <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><rect x="3" y="11" width="18" height="11" rx="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></svg>
              <input
                v-model="form.password"
                :type="showPw ? 'text' : 'password'"
                placeholder="••••••••"
                autocomplete="current-password"
                @focus="focused = 'pw'"
                @blur="focused = ''"
              />
              <button type="button" class="eye-btn" @click="showPw = !showPw" tabindex="-1">
                <svg v-if="!showPw" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
                <svg v-else viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"/><line x1="1" y1="1" x2="23" y2="23"/></svg>
              </button>
            </div>
          </div>

          <div v-if="error" class="alert-error">{{ error }}</div>
          <div v-if="debugMsg" class="alert-debug">{{ debugMsg }}</div>

          <button type="button" class="btn-submit" :disabled="loading" @click="handleLogin">
            <span v-if="loading" class="spinner"></span>
            <template v-else>Sign in <span class="btn-arrow">→</span></template>
          </button>
        </div>

        <div class="divider"><span>or</span></div>
        <div class="oauth-row">
          <button class="oauth-btn" type="button">
            <svg viewBox="0 0 24 24" width="18" height="18"><path fill="#4285F4" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/><path fill="#34A853" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/><path fill="#FBBC05" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/><path fill="#EA4335" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/></svg>
            Continue with Google
          </button>
          <button class="oauth-btn" type="button">
            <svg viewBox="0 0 24 24" width="18" height="18" fill="currentColor"><path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z"/></svg>
            Continue with GitHub
          </button>
        </div>
        <p class="switch-link">Don't have an account? <router-link to="/register">Create one</router-link></p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import authService from '../../services/auth.js'
import { useUserStore } from '../../stores/user.js'

const router = useRouter()
const userStore = useUserStore()
const showPw = ref(false)
const loading = ref(false)
const error = ref('')
const debugMsg = ref('')
const focused = ref('')
const form = ref({ identifier: '', password: '' })

async function handleLogin() {
  debugMsg.value = 'Signing in...'
  error.value = ''

  if (!form.value.identifier || !form.value.password) {
    error.value = 'Please fill in both fields.'
    debugMsg.value = ''
    return
  }

  loading.value = true
  try {
    const data = await authService.login(form.value)
    userStore.setAuth(data)
    debugMsg.value = 'Success! Redirecting...'
    if (data.user.role === 'admin') {
      router.push('/admin')
    } else {
      router.push('/dashboard')
    }
  } catch (e) {
    error.value = e.response?.data?.error || e.message || 'Invalid credentials. Please try again.'
    debugMsg.value = ''
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500&family=Syne:wght@400;600;700;800&display=swap');
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

.auth-wrap {
  display: flex; min-height: 100vh;
  font-family: 'Syne', sans-serif;
  background: #09090e;
}

.left-panel {
  width: 45%; background: #0d0d15;
  border-right: 1px solid #1c1c2a;
  padding: 2.5rem;
  display: flex; flex-direction: column; justify-content: space-between;
  position: relative; overflow: hidden;
}
.left-panel::before {
  content: ''; position: absolute; inset: 0;
  background-image:
    linear-gradient(rgba(232,234,240,0.025) 1px, transparent 1px),
    linear-gradient(90deg, rgba(232,234,240,0.025) 1px, transparent 1px);
  background-size: 48px 48px; pointer-events: none;
}
.logo { font-size: 1.5rem; font-weight: 800; color: #e2e2e8; text-decoration: none; letter-spacing: -0.5px; }
.accent { color: #e8eaf0; }
.left-content { flex: 1; display: flex; flex-direction: column; justify-content: center; gap: 2.5rem; }
.terminal { background: #111119; border: 1px solid #1e1e2e; border-radius: 12px; overflow: hidden; position: relative; z-index: 1; }
.terminal-bar { background: #161622; padding: 0.65rem 1rem; display: flex; align-items: center; gap: 6px; border-bottom: 1px solid #1e1e2e; }
.dot { width: 10px; height: 10px; border-radius: 50%; }
.dot.r { background: #ff5f57; } .dot.y { background: #febc2e; } .dot.g { background: #28c840; }
.terminal-title { font-family: 'JetBrains Mono', monospace; font-size: 0.72rem; color: #444458; margin-left: 6px; }
.terminal-code { font-family: 'JetBrains Mono', monospace; font-size: 0.78rem; line-height: 1.8; padding: 1.25rem; color: #8888a0; white-space: pre; }
.c-kw { color: #c792ea; } .c-fn { color: #9ec8ff; } .c-str { color: #c3e88d; } .c-num { color: #f78c6c; }
.left-stats { display: flex; align-items: center; gap: 2rem; position: relative; z-index: 1; }
.ls { display: flex; flex-direction: column; }
.ls-n { font-size: 1.6rem; font-weight: 800; color: #e2e2e8; letter-spacing: -1px; }
.ls-l { font-size: 0.78rem; color: #44445a; margin-top: 2px; }
.ls-sep { width: 1px; height: 32px; background: #1e1e2e; }
.left-tagline { font-size: 0.85rem; color: #2e2e42; position: relative; z-index: 1; }

.right-panel { flex: 1; display: flex; align-items: center; justify-content: center; padding: 2rem; background: #09090e; }
.form-card { width: 100%; max-width: 400px; animation: fadeUp 0.5s ease both; }
.form-header { margin-bottom: 2rem; }
.form-header h1 { font-size: 2rem; font-weight: 800; color: #e2e2e8; letter-spacing: -1px; margin-bottom: 0.4rem; }
.form-header p { font-size: 0.9rem; color: #44445a; }

.form { display: flex; flex-direction: column; gap: 1rem; }
.field { display: flex; flex-direction: column; gap: 0.45rem; }
.field label { font-size: 0.8rem; font-weight: 600; color: #55556a; transition: color 0.2s; }
.field.focused label { color: #b0b0c8; }
.input-wrap { position: relative; }
.input-icon { position: absolute; left: 0.875rem; top: 50%; transform: translateY(-50%); width: 16px; height: 16px; color: #33334a; pointer-events: none; transition: color 0.2s; }
.field.focused .input-icon { color: #8888a8; }
.input-wrap input {
  width: 100%; background: #111119; border: 1px solid #1e1e2e;
  border-radius: 10px; padding: 0.8rem 1rem 0.8rem 2.75rem;
  color: #e2e2e8; font-family: 'Syne', sans-serif; font-size: 0.92rem;
  outline: none; transition: border-color 0.2s, background 0.2s;
}
.input-wrap input:focus { border-color: #3a3a54; background: #14141e; }
.input-wrap input::placeholder { color: #2a2a3e; }
.eye-btn { position: absolute; right: 0.875rem; top: 50%; transform: translateY(-50%); background: none; border: none; cursor: pointer; padding: 0; color: #33334a; transition: color 0.2s; }
.eye-btn:hover { color: #8888a8; }
.eye-btn svg { width: 16px; height: 16px; display: block; }

.alert-error { background: rgba(220,60,60,0.08); border: 1px solid rgba(220,60,60,0.18); color: #e07070; border-radius: 10px; padding: 0.75rem 1rem; font-size: 0.85rem; }
.alert-debug { background: rgba(100,180,255,0.06); border: 1px solid rgba(100,180,255,0.15); color: #88b8e0; border-radius: 10px; padding: 0.75rem 1rem; font-size: 0.85rem; }

.btn-submit {
  width: 100%; background: #e8eaf0; color: #09090e;
  border: none; border-radius: 10px; padding: 0.875rem;
  font-family: 'Syne', sans-serif; font-size: 0.95rem; font-weight: 700;
  cursor: pointer; transition: background 0.2s, transform 0.15s;
  display: flex; align-items: center; justify-content: center; gap: 0.5rem;
  margin-top: 0.5rem;
}
.btn-submit:hover:not(:disabled) { background: #ffffff; transform: translateY(-1px); }
.btn-submit:disabled { opacity: 0.5; cursor: not-allowed; }
.btn-arrow { transition: transform 0.2s; }
.btn-submit:hover .btn-arrow { transform: translateX(3px); }
.spinner { width: 18px; height: 18px; border: 2px solid rgba(9,9,14,0.2); border-top-color: #09090e; border-radius: 50%; animation: spin 0.7s linear infinite; }

.divider { display: flex; align-items: center; gap: 1rem; margin: 1.5rem 0; color: #2a2a3e; font-size: 0.8rem; }
.divider::before, .divider::after { content: ''; flex: 1; height: 1px; background: #1a1a28; }

.oauth-row { display: flex; gap: 0.75rem; }
.oauth-btn { flex: 1; display: flex; align-items: center; justify-content: center; gap: 0.6rem; background: #111119; border: 1px solid #1e1e2e; border-radius: 10px; padding: 0.75rem 0.5rem; color: #888898; font-family: 'Syne', sans-serif; font-size: 0.82rem; cursor: pointer; transition: border-color 0.2s, color 0.2s; }
.oauth-btn:hover { border-color: #2e2e42; color: #b0b0c0; }

.switch-link { text-align: center; margin-top: 1.75rem; font-size: 0.875rem; color: #44445a; }
.switch-link a { color: #c0c0d8; text-decoration: none; font-weight: 600; }
.switch-link a:hover { color: #e8eaf0; }

@keyframes fadeUp { from { opacity: 0; transform: translateY(14px); } to { opacity: 1; transform: translateY(0); } }
@keyframes spin { to { transform: rotate(360deg); } }

@media (max-width: 768px) {
  .auth-wrap { flex-direction: column; }
  .left-panel { width: 100%; min-height: auto; padding: 1.5rem; }
  .terminal { display: none; }
  .left-content { gap: 1rem; }
}
</style>