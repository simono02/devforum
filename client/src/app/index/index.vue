<template>
  <div class="landing">
    <!-- Nav -->
    <nav class="nav">
      <div class="nav-inner">
        <span class="logo">dev<span class="accent">forum</span></span>
        <div class="nav-links">
          <router-link to="/login" class="btn-ghost">Sign in</router-link>
          <router-link to="/register" class="btn-primary">Join now</router-link>
        </div>
      </div>
    </nav>

    <!-- Hero -->
    <section class="hero">
      <div class="hero-bg">
        <div class="grid-lines"></div>
        <div class="glow"></div>
      </div>
      <div class="hero-content">
        <div class="badge">// open to all developers</div>
        <h1 class="hero-title">
          Where developers<br />
          <span class="accent-text">think out loud.</span>
        </h1>
        <p class="hero-sub">
          Ask questions, share discoveries, debate ideas.<br />
          A community built by developers, for developers.
        </p>
        <div class="hero-actions">
          <router-link to="/register" class="btn-primary btn-lg">Get started free</router-link>
          <router-link to="/login" class="btn-outline btn-lg">Sign in</router-link>
        </div>
      </div>

      <!-- Floating code card -->
      <div class="code-card">
        <div class="card-header">
          <span class="dot red"></span>
          <span class="dot yellow"></span>
          <span class="dot green"></span>
          <span class="card-title">latest_post.go</span>
        </div>
        <pre class="code-block"><span class="kw">func</span> <span class="fn">GetPosts</span>(c *gin.Context) {
  <span class="kw">var</span> posts []models.Post
  db.Preload(<span class="str">"Author"</span>).
     Order(<span class="str">"created_at desc"</span>).
     Find(&amp;posts)
  c.JSON(<span class="num">200</span>, posts)
}</pre>
      </div>
    </section>

    <!-- Stats -->
    <section class="stats">
      <div class="stat">
        <span class="stat-num">12k+</span>
        <span class="stat-label">Developers</span>
      </div>
      <div class="divider"></div>
      <div class="stat">
        <span class="stat-num">48k+</span>
        <span class="stat-label">Posts</span>
      </div>
      <div class="divider"></div>
      <div class="stat">
        <span class="stat-num">200k+</span>
        <span class="stat-label">Comments</span>
      </div>
      <div class="divider"></div>
      <div class="stat">
        <span class="stat-num">99%</span>
        <span class="stat-label">Uptime</span>
      </div>
    </section>

    <!-- Features -->
    <section class="features">
      <h2 class="section-title">Everything you need to<br /><span class="accent-text">level up</span></h2>
      <div class="features-grid">
        <div class="feature-card">
          <div class="feature-icon">⚡</div>
          <h3>Fast Answers</h3>
          <p>Post a question and get answers from experienced developers within minutes.</p>
        </div>
        <div class="feature-card">
          <div class="feature-icon">🧵</div>
          <h3>Threaded Discussions</h3>
          <p>Deep-dive conversations with nested comments keep context intact.</p>
        </div>
        <div class="feature-card">
          <div class="feature-icon">🔼</div>
          <h3>Upvote System</h3>
          <p>Surface the best content. The community decides what rises to the top.</p>
        </div>
        <div class="feature-card">
          <div class="feature-icon">🛡️</div>
          <h3>Moderated Space</h3>
          <p>Admins keep the forum clean, respectful, and on-topic.</p>
        </div>
      </div>
    </section>

    <!-- CTA -->
    <section class="cta">
      <h2>Ready to join the conversation?</h2>
      <p>Create your free account in seconds.</p>
      <router-link to="/register" class="btn-primary btn-lg">Create account →</router-link>
    </section>

    <!-- Footer -->
    <footer class="footer">
      <span class="logo">dev<span class="accent">forum</span></span>
      <span class="footer-copy">© 2026 DevForum. Built for developers.</span>
    </footer>
  </div>
</template>

<script setup>
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;600&family=Syne:wght@400;600;800&display=swap');

*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

.landing {
  font-family: 'Syne', sans-serif;
  background: #0a0a0f;
  color: #e2e2e8;
  min-height: 100vh;
  overflow-x: hidden;
}

/* ── Nav ── */
.nav {
  position: fixed; top: 0; left: 0; right: 0; z-index: 100;
  border-bottom: 1px solid #1e1e2e;
  background: rgba(10,10,15,0.85);
  backdrop-filter: blur(12px);
}
.nav-inner {
  max-width: 1100px; margin: 0 auto;
  padding: 1rem 2rem;
  display: flex; align-items: center; justify-content: space-between;
}
.logo { font-size: 1.4rem; font-weight: 800; letter-spacing: -0.5px; }
.accent { color: #e8eaf0; }
.nav-links { display: flex; gap: 0.75rem; align-items: center; }

/* ── Buttons ── */
.btn-ghost {
  color: #a0a0b0; text-decoration: none; padding: 0.5rem 1rem;
  border-radius: 6px; font-family: 'Syne', sans-serif; font-size: 0.9rem;
  transition: color 0.2s;
}
.btn-ghost:hover { color: #e2e2e8; }

.btn-primary {
  background: #e8eaf0; color: #0a0a0f; text-decoration: none;
  padding: 0.5rem 1.25rem; border-radius: 6px;
  font-family: 'Syne', sans-serif; font-size: 0.9rem; font-weight: 600;
  transition: background 0.2s, transform 0.15s;
  display: inline-block;
}
.btn-primary:hover { background: #ffffff; transform: translateY(-1px); }

.btn-outline {
  border: 1px solid #2e2e42; color: #e2e2e8; text-decoration: none;
  padding: 0.75rem 1.75rem; border-radius: 6px;
  font-family: 'Syne', sans-serif; font-size: 1rem; font-weight: 600;
  transition: border-color 0.2s, transform 0.15s;
  display: inline-block;
}
.btn-outline:hover { border-color: #e8eaf0; transform: translateY(-1px); }

.btn-lg { padding: 0.85rem 2rem; font-size: 1rem; }

/* ── Hero ── */
.hero {
  min-height: 100vh;
  display: flex; align-items: center;
  padding: 8rem 2rem 4rem;
  max-width: 1100px; margin: 0 auto;
  position: relative;
}
.hero-bg {
  position: fixed; inset: 0; z-index: 0; pointer-events: none;
}
.grid-lines {
  position: absolute; inset: 0;
  background-image:
    linear-gradient(rgba(232,234,240,0.03) 1px, transparent 1px),
    linear-gradient(90deg, rgba(232,234,240,0.03) 1px, transparent 1px);
  background-size: 60px 60px;
}
.glow {
  position: absolute; top: 20%; left: 50%; transform: translateX(-50%);
  width: 600px; height: 400px;
  background: radial-gradient(ellipse, rgba(232,234,240,0.06) 0%, transparent 70%);
}
.hero-content { position: relative; z-index: 1; max-width: 560px; }
.badge {
  display: inline-block;
  font-family: 'JetBrains Mono', monospace;
  font-size: 0.8rem; color: #c8cad4;
  background: rgba(232,234,240,0.06);
  border: 1px solid rgba(232,234,240,0.15);
  padding: 0.35rem 0.85rem; border-radius: 20px;
  margin-bottom: 1.5rem;
  animation: fadeUp 0.6s ease both;
}
.hero-title {
  font-size: clamp(2.8rem, 6vw, 4.5rem);
  font-weight: 800; line-height: 1.05;
  letter-spacing: -2px;
  margin-bottom: 1.25rem;
  animation: fadeUp 0.6s 0.1s ease both;
}
.accent-text { color: #e8eaf0; }
.hero-sub {
  font-size: 1.1rem; color: #7a7a90; line-height: 1.7;
  margin-bottom: 2.5rem;
  animation: fadeUp 0.6s 0.2s ease both;
}
.hero-actions {
  display: flex; gap: 1rem; flex-wrap: wrap;
  animation: fadeUp 0.6s 0.3s ease both;
}

/* ── Code card ── */
.code-card {
  position: absolute; right: 0; top: 50%; transform: translateY(-50%);
  background: #111118; border: 1px solid #1e1e2e;
  border-radius: 12px; overflow: hidden; width: 340px;
  box-shadow: 0 24px 80px rgba(0,0,0,0.6);
  animation: fadeUp 0.8s 0.4s ease both;
  z-index: 1;
}
.card-header {
  background: #161622; padding: 0.75rem 1rem;
  display: flex; align-items: center; gap: 0.4rem;
  border-bottom: 1px solid #1e1e2e;
}
.dot { width: 10px; height: 10px; border-radius: 50%; }
.dot.red { background: #ff5f57; }
.dot.yellow { background: #febc2e; }
.dot.green { background: #28c840; }
.card-title {
  font-family: 'JetBrains Mono', monospace;
  font-size: 0.75rem; color: #555568; margin-left: 0.5rem;
}
.code-block {
  font-family: 'JetBrains Mono', monospace;
  font-size: 0.82rem; line-height: 1.7;
  padding: 1.25rem 1.25rem;
  color: #c0c0d0; white-space: pre;
}
.kw { color: #c792ea; }
.fn { color: #82aaff; }
.str { color: #c3e88d; }
.num { color: #f78c6c; }

/* ── Stats ── */
.stats {
  display: flex; align-items: center; justify-content: center;
  gap: 3rem; padding: 3rem 2rem;
  border-top: 1px solid #1e1e2e;
  border-bottom: 1px solid #1e1e2e;
  background: #0d0d14;
  flex-wrap: wrap;
}
.stat { text-align: center; }
.stat-num {
  display: block; font-size: 2rem; font-weight: 800;
  color: #e8eaf0; letter-spacing: -1px;
}
.stat-label { font-size: 0.85rem; color: #555568; }
.divider { width: 1px; height: 40px; background: #1e1e2e; }

/* ── Features ── */
.features {
  max-width: 1100px; margin: 0 auto;
  padding: 6rem 2rem;
}
.section-title {
  font-size: clamp(2rem, 4vw, 3rem);
  font-weight: 800; letter-spacing: -1.5px;
  margin-bottom: 3rem; line-height: 1.1;
}
.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 1.5rem;
}
.feature-card {
  background: #0f0f18; border: 1px solid #1e1e2e;
  border-radius: 12px; padding: 2rem;
  transition: border-color 0.2s, transform 0.2s;
}
.feature-card:hover { border-color: rgba(232,234,240,0.2); transform: translateY(-3px); }
.feature-icon { font-size: 1.75rem; margin-bottom: 1rem; }
.feature-card h3 { font-size: 1.1rem; font-weight: 600; margin-bottom: 0.6rem; }
.feature-card p { font-size: 0.9rem; color: #666680; line-height: 1.6; }

/* ── CTA ── */
.cta {
  text-align: center; padding: 6rem 2rem;
  background: #0d0d14;
  border-top: 1px solid #1e1e2e;
}
.cta h2 { font-size: 2.2rem; font-weight: 800; letter-spacing: -1px; margin-bottom: 0.75rem; }
.cta p { color: #666680; margin-bottom: 2rem; }

/* ── Footer ── */
.footer {
  padding: 2rem; border-top: 1px solid #1e1e2e;
  display: flex; align-items: center; justify-content: space-between;
  max-width: 1100px; margin: 0 auto;
  flex-wrap: wrap; gap: 1rem;
}
.footer-copy { font-size: 0.85rem; color: #444458; }

/* ── Animations ── */
@keyframes fadeUp {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}

@media (max-width: 768px) {
  .code-card { display: none; }
  .hero { justify-content: center; text-align: center; }
  .hero-actions { justify-content: center; }
  .divider { display: none; }
}
</style>