<template>
  <div class="page">
    <div class="page-header">
      <div>
        <h1 class="page-title">Posts</h1>
        <p class="page-sub">Browse and manage your forum posts.</p>
      </div>
      <button class="btn-primary" @click="newPostModal = true">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/></svg>
        New Post
      </button>
    </div>

    <div class="tabs">
      <button class="tab" :class="{ active: tab === 'all' }"   @click="tab = 'all'">All Posts</button>
      <button class="tab" :class="{ active: tab === 'mine' }"  @click="tab = 'mine'">My Posts</button>
      <button class="tab" :class="{ active: tab === 'saved' }" @click="tab = 'saved'">Saved</button>
    </div>

    <div class="toolbar">
      <div class="search-wrap">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>
        <input v-model="search" type="text" placeholder="Search posts..." />
      </div>
      <div class="tag-filters">
        <button class="tag-chip" :class="{ active: tagFilter === '' }" @click="tagFilter = ''">All</button>
        <button class="tag-chip" v-for="t in allTags" :key="t" :class="{ active: tagFilter === t }" @click="tagFilter = t">{{ t }}</button>
      </div>
    </div>

    <!-- Loading -->
    <div class="loading-state" v-if="loading">
      <div class="spinner-lg"></div>
      <p>Loading posts...</p>
    </div>

    <!-- All posts -->
    <div class="posts-list" v-else-if="tab === 'all'">
      <div class="post-card" v-for="p in filteredPosts" :key="p.id">
        <div class="post-main">
          <div class="post-top">
            <div class="post-author">
              <div class="author-av">{{ initials(p.author) }}</div>
              <div>
                <div class="author-name">{{ p.author }}</div>
                <div class="author-time">{{ formatTime(p.created_at) }}</div>
              </div>
            </div>
            <span class="post-tag" v-if="p.tag">{{ p.tag }}</span>
          </div>
          <h3 class="post-title" @click="viewPost(p)">{{ p.title }}</h3>
          <p class="post-excerpt">{{ p.content }}</p>
          <div class="post-stats">
            <span class="pstat">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="17 11 12 6 7 11"/><line x1="12" y1="18" x2="12" y2="6"/></svg>
              {{ p.upvotes }}
            </span>
          </div>
        </div>
        <div class="post-actions">
          <button class="pa-btn" :class="{ active: p.upvoted }" title="Upvote" @click="handleUpvote(p)">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="17 11 12 6 7 11"/><line x1="12" y1="18" x2="12" y2="6"/></svg>
          </button>
          <button class="pa-btn" :class="{ active: p.saved }" title="Save" @click="p.saved = !p.saved">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/></svg>
          </button>
          <button class="pa-btn" title="Comment" @click="openCommentModal(p)">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
          </button>
          <button class="pa-btn" title="View" @click="viewPost(p)">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
          </button>
        </div>
      </div>
      <div class="empty-state" v-if="filteredPosts.length === 0">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
        <p>No posts yet. Be the first to write one.</p>
        <button class="empty-btn" @click="newPostModal = true">Write the first post →</button>
      </div>
    </div>

    <!-- My posts -->
    <div class="posts-list" v-else-if="tab === 'mine'">
      <div class="post-card" v-for="p in myPosts" :key="p.id">
        <div class="post-main">
          <div class="post-top">
            <span class="post-tag" v-if="p.tag">{{ p.tag }}</span>
            <span class="author-time">{{ formatTime(p.created_at) }}</span>
          </div>
          <h3 class="post-title" @click="viewPost(p)">{{ p.title }}</h3>
          <p class="post-excerpt">{{ p.content }}</p>
          <div class="post-stats">
            <span class="pstat">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="17 11 12 6 7 11"/><line x1="12" y1="18" x2="12" y2="6"/></svg>
              {{ p.upvotes }}
            </span>
          </div>
        </div>
        <div class="post-actions">
          <button class="pa-btn" title="Edit" @click="editPost(p)">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/></svg>
          </button>
          <button class="pa-btn danger" title="Delete" @click="confirmDeletePost(p)">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="3 6 5 6 21 6"/><path d="M19 6l-1 14H6L5 6"/><path d="M10 11v6"/><path d="M14 11v6"/></svg>
          </button>
        </div>
      </div>
      <div class="empty-state" v-if="myPosts.length === 0">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
        <p>You haven't written any posts yet.</p>
        <button class="empty-btn" @click="newPostModal = true">Write your first post →</button>
      </div>
    </div>

    <!-- Saved -->
    <div class="posts-list" v-else-if="tab === 'saved'">
      <div class="post-card" v-for="p in savedPosts" :key="p.id">
        <div class="post-main">
          <div class="post-top">
            <div class="post-author">
              <div class="author-av">{{ initials(p.author) }}</div>
              <div>
                <div class="author-name">{{ p.author }}</div>
                <div class="author-time">{{ formatTime(p.created_at) }}</div>
              </div>
            </div>
            <span class="post-tag" v-if="p.tag">{{ p.tag }}</span>
          </div>
          <h3 class="post-title">{{ p.title }}</h3>
          <p class="post-excerpt">{{ p.content }}</p>
        </div>
        <div class="post-actions">
          <button class="pa-btn active" title="Unsave" @click="p.saved = false">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/></svg>
          </button>
        </div>
      </div>
      <div class="empty-state" v-if="savedPosts.length === 0">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/></svg>
        <p>No saved posts yet.</p>
      </div>
    </div>

    <!-- View Post Modal with Comments -->
    <Teleport to="body">
      <div class="overlay" v-if="viewModal" @click.self="closeView">
        <div class="modal modal-lg">
          <div class="modal-head">
            <h3>Post</h3>
            <button class="modal-close" @click="closeView">✕</button>
          </div>
          <div class="modal-body" v-if="activePost">
            <!-- Post content -->
            <div class="view-meta">
              <div class="post-author">
                <div class="author-av">{{ initials(activePost.author) }}</div>
                <div>
                  <div class="author-name">{{ activePost.author }}</div>
                  <div class="author-time">{{ formatTime(activePost.created_at) }}</div>
                </div>
              </div>
              <span class="post-tag" v-if="activePost.tag">{{ activePost.tag }}</span>
            </div>
            <h2 class="view-title">{{ activePost.title }}</h2>
            <p class="view-content">{{ activePost.content }}</p>
            <div class="view-stats">
              <span class="pstat">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><polyline points="17 11 12 6 7 11"/><line x1="12" y1="18" x2="12" y2="6"/></svg>
                {{ activePost.upvotes }} upvotes
              </span>
              <span class="pstat">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
                {{ comments.length }} comments
              </span>
            </div>

            <!-- Comments section -->
            <div class="comments-section">
              <div class="comments-header">
                <h4>Comments <span class="comments-count">{{ comments.length }}</span></h4>
              </div>

              <!-- Loading -->
              <div class="comments-loading" v-if="commentsLoading">
                <div class="spinner-sm"></div>
              </div>

              <!-- Comments list -->
              <div class="comments-list" v-else>
                <div class="comment-item" v-for="c in comments" :key="c.id">
                  <div class="comment-av">{{ initials(c.author) }}</div>
                  <div class="comment-body">
                    <div class="comment-meta">
                      <span class="comment-author">{{ c.author }}</span>
                      <span class="comment-time">{{ formatTime(c.created_at) }}</span>
                      <span class="comment-edited" v-if="c.is_edited">edited</span>
                    </div>
                    <!-- Editing inline -->
                    <div v-if="editingCommentId === c.id" class="comment-edit-wrap">
                      <textarea v-model="editCommentContent" rows="2" class="comment-edit-input"></textarea>
                      <div class="comment-edit-actions">
                        <button class="cta-btn cta-cancel" @click="cancelEditComment">Cancel</button>
                        <button class="cta-btn cta-save" @click="saveEditComment(c)" :disabled="commentSaving">
                          {{ commentSaving ? 'Saving...' : 'Save' }}
                        </button>
                      </div>
                    </div>
                    <p class="comment-content" v-else>{{ c.content }}</p>
                    <!-- Comment actions (own comments) -->
                    <div class="comment-actions" v-if="c.user_id === currentUserId">
                      <button class="ca-btn" @click="startEditComment(c)">Edit</button>
                      <button class="ca-btn danger" @click="deleteComment(c)">Delete</button>
                    </div>
                  </div>
                </div>
                <div class="no-comments" v-if="comments.length === 0">
                  No comments yet. Be the first to reply.
                </div>
              </div>

              <!-- Write comment -->
              <div class="comment-form">
                <div class="comment-av">{{ initials(user?.full_name) }}</div>
                <div class="comment-input-wrap">
                  <textarea
                    v-model="newComment"
                    rows="2"
                    placeholder="Write a comment..."
                    class="comment-input"
                    @keydown.ctrl.enter="submitComment"
                  ></textarea>
                  <div class="comment-form-foot">
                    <span class="comment-hint">Ctrl+Enter to post</span>
                    <button class="cta-btn cta-post" @click="submitComment" :disabled="!newComment.trim() || commentSaving">
                      <span v-if="commentSaving" class="spinner"></span>
                      <span v-else>Post</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="modal-foot">
            <button class="mbtn mbtn-outline" @click="closeView">Close</button>
            <button class="mbtn mbtn-primary" @click="handleUpvote(activePost)" :disabled="activePost?.upvoted">
              {{ activePost?.upvoted ? 'Upvoted ✓' : 'Upvote' }}
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- New / Edit Post Modal -->
    <Teleport to="body">
      <div class="overlay" v-if="newPostModal || editModal" @click.self="closePostModal">
        <div class="modal modal-lg">
          <div class="modal-head">
            <h3>{{ editModal ? 'Edit Post' : 'New Post' }}</h3>
            <button class="modal-close" @click="closePostModal">✕</button>
          </div>
          <div class="modal-body">
            <div class="form">
              <div class="field">
                <label>Title <span class="req">*</span></label>
                <input v-model="postForm.title" type="text" placeholder="What's your question or topic?" />
              </div>
              <div class="field">
                <label>Tag <span class="req">*</span></label>
                <select v-model="postForm.tag">
                  <option value="">Select a tag</option>
                  <option v-for="t in allTags" :key="t" :value="t">{{ t }}</option>
                </select>
              </div>
              <div class="field">
                <label>Content <span class="req">*</span></label>
                <textarea v-model="postForm.content" rows="8" placeholder="Share your knowledge, ask a question, or start a discussion..."></textarea>
                <div class="char-count">{{ postForm.content.length }} characters</div>
              </div>
              <div v-if="postFormError" class="form-error">{{ postFormError }}</div>
            </div>
          </div>
          <div class="modal-foot">
            <button class="mbtn mbtn-outline" @click="closePostModal">Cancel</button>
            <button class="mbtn mbtn-primary" @click="submitPost" :disabled="postLoading">
              <span v-if="postLoading" class="spinner"></span>
              <span v-else>{{ editModal ? 'Save Changes' : 'Publish Post' }}</span>
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- ── Comment Modal ── -->
    <Teleport to="body">
      <div class="overlay" v-if="commentModal" @click.self="closeCommentModal">
        <div class="modal modal-comment">
          <div class="modal-head">
            <div class="cm-head-left">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
              <div>
                <h3>Comment on post</h3>
                <p class="cm-post-title">{{ commentTargetPost?.title }}</p>
              </div>
            </div>
            <button class="modal-close" @click="closeCommentModal">✕</button>
          </div>

          <div class="modal-body">
            <!-- Existing comments preview -->
            <div class="cm-existing" v-if="comments.length > 0">
              <div class="cm-existing-label">{{ comments.length }} comment{{ comments.length !== 1 ? 's' : '' }}</div>
              <div class="cm-preview-list">
                <div class="cm-preview-item" v-for="c in comments.slice(0, 3)" :key="c.id">
                  <div class="comment-av sm">{{ initials(c.author) }}</div>
                  <div class="cm-preview-body">
                    <span class="cm-preview-author">{{ c.author }}</span>
                    <span class="cm-preview-text">{{ c.content }}</span>
                  </div>
                </div>
                <div class="cm-more" v-if="comments.length > 3" @click="openFullView">
                  View all {{ comments.length }} comments →
                </div>
              </div>
            </div>
            <div class="cm-no-comments" v-else-if="!commentsLoading">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
              <span>No comments yet. Be the first!</span>
            </div>
            <div class="mini-loading" v-else-if="commentsLoading">
              <div class="spinner-sm"></div>
            </div>

            <!-- Write comment -->
            <div class="cm-write">
              <div class="cm-write-top">
                <div class="comment-av">{{ initials(user?.full_name) }}</div>
                <span class="cm-writing-as">{{ user?.full_name }}</span>
              </div>
              <textarea
                v-model="newComment"
                rows="4"
                placeholder="Write your comment..."
                class="cm-textarea"
                @keydown.ctrl.enter="submitCommentFromModal"
                autofocus
              ></textarea>
              <div class="cm-char-count">{{ newComment.length }} characters</div>
              <div v-if="commentError" class="cm-error">{{ commentError }}</div>
            </div>
          </div>

          <div class="modal-foot">
            <button class="mbtn mbtn-outline" @click="closeCommentModal">Cancel</button>
            <button class="mbtn mbtn-primary" @click="submitCommentFromModal" :disabled="!newComment.trim() || commentSaving">
              <span v-if="commentSaving" class="spinner"></span>
              <span v-else>
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="width:14px;height:14px"><line x1="22" y1="2" x2="11" y2="13"/><polygon points="22 2 15 22 11 13 2 9 22 2"/></svg>
                Post Comment
              </span>
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Delete Confirm Modal -->
    <Teleport to="body">
      <div class="overlay" v-if="deleteModal" @click.self="deleteModal = false">
        <div class="modal modal-sm">
          <div class="modal-head">
            <h3>Delete Post</h3>
            <button class="modal-close" @click="deleteModal = false">✕</button>
          </div>
          <div class="modal-body">
            <div class="confirm-icon">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
            </div>
            <p class="confirm-text">Delete <strong>{{ activePost?.title }}</strong>? This cannot be undone.</p>
          </div>
          <div class="modal-foot">
            <button class="mbtn mbtn-outline" @click="deleteModal = false">Cancel</button>
            <button class="mbtn mbtn-danger" @click="doDeletePost">Delete</button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Toast -->
    <Teleport to="body">
      <div class="toast" v-if="toast.show" :class="toast.type">{{ toast.msg }}</div>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useUserStore } from '../../../stores/user.js'
import postService from '../../../services/posts.js'
import commentService from '../../../services/comments.js'

const userStore = useUserStore()
const user      = computed(() => userStore.user)
const currentUserId = computed(() => userStore.user?.id)

const tab           = ref('all')
const search        = ref('')
const tagFilter     = ref('')
const viewModal     = ref(false)
const newPostModal  = ref(false)
const editModal     = ref(false)
const deleteModal   = ref(false)
const activePost    = ref(null)
const postLoading   = ref(false)
const postFormError = ref('')
const loading       = ref(false)
const toast         = ref({ show: false, msg: '', type: 'success' })

// ── Comment state ─────────────────────────────────────────────────────────────
const comments           = ref([])
const commentsLoading    = ref(false)
const newComment         = ref('')
const commentSaving      = ref(false)
const editingCommentId   = ref(null)
const editCommentContent = ref('')
const commentModal       = ref(false)
const commentTargetPost  = ref(null)
const commentError       = ref('')

const allTags  = ['Go', 'Vue', 'PostgreSQL', 'Docker', 'JavaScript', 'TypeScript', 'Python', 'REST API', 'GraphQL', 'Auth']
const postForm = ref({ title: '', tag: '', content: '' })

const communityPosts = ref([])
const myPosts        = ref([])

// ── Load on mount ─────────────────────────────────────────────────────────────
onMounted(async () => {
  loading.value = true
  try {
    const [all, mine] = await Promise.all([
      postService.getAll(),
      postService.getMine(),
    ])
    communityPosts.value = all.map(p => ({ ...p, upvoted: false, saved: false }))
    myPosts.value        = mine
  } catch (e) {
    showToast('Failed to load posts', 'danger')
  } finally {
    loading.value = false
  }
})

// ── Computed ──────────────────────────────────────────────────────────────────
const filteredPosts = computed(() => {
  let list = communityPosts.value
  if (tagFilter.value) list = list.filter(p => p.tag === tagFilter.value)
  if (search.value) {
    const q = search.value.toLowerCase()
    list = list.filter(p =>
      p.title.toLowerCase().includes(q) ||
      (p.author || '').toLowerCase().includes(q)
    )
  }
  return list
})

const savedPosts = computed(() => communityPosts.value.filter(p => p.saved))

// ── Helpers ───────────────────────────────────────────────────────────────────
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

// ── Actions ───────────────────────────────────────────────────────────────────
function viewPost(p) {
  activePost.value = p
  viewModal.value  = true
  loadComments(p.id)
}

function closeView() {
  viewModal.value        = false
  comments.value         = []
  newComment.value       = ''
  editingCommentId.value = null
}

// Open the quick comment modal from a post card
function openCommentModal(p) {
  commentTargetPost.value = p
  commentModal.value      = true
  commentError.value      = ''
  newComment.value        = ''
  comments.value          = []
  loadComments(p.id)
}

function closeCommentModal() {
  commentModal.value      = false
  commentTargetPost.value = null
  newComment.value        = ''
  commentError.value      = ''
}

// Open the full view modal from inside the comment modal
function openFullView() {
  commentModal.value = false
  viewPost(commentTargetPost.value)
}

async function submitCommentFromModal() {
  if (!newComment.value.trim()) return
  commentSaving.value = true
  commentError.value  = ''
  try {
    const created = await commentService.create(commentTargetPost.value.id, newComment.value.trim())
    comments.value.push(created)
    newComment.value = ''
    showToast('Comment posted!', 'success')
    closeCommentModal()
  } catch (e) {
    commentError.value = e.response?.data?.error || 'Failed to post comment.'
  } finally {
    commentSaving.value = false
  }
}

async function loadComments(postId) {
  commentsLoading.value = true
  try {
    comments.value = await commentService.getByPost(postId)
  } catch {
    showToast('Failed to load comments', 'danger')
  } finally {
    commentsLoading.value = false
  }
}

async function submitComment() {
  if (!newComment.value.trim()) return
  commentSaving.value = true
  try {
    const created = await commentService.create(activePost.value.id, newComment.value.trim())
    comments.value.push(created)
    newComment.value = ''
  } catch (e) {
    showToast(e.response?.data?.error || 'Failed to post comment', 'danger')
  } finally {
    commentSaving.value = false
  }
}

function startEditComment(c) {
  editingCommentId.value  = c.id
  editCommentContent.value = c.content
}

function cancelEditComment() {
  editingCommentId.value  = null
  editCommentContent.value = ''
}

async function saveEditComment(c) {
  if (!editCommentContent.value.trim()) return
  commentSaving.value = true
  try {
    const updated = await commentService.update(c.id, editCommentContent.value.trim())
    const idx = comments.value.findIndex(x => x.id === updated.id)
    if (idx !== -1) comments.value[idx] = updated
    cancelEditComment()
  } catch (e) {
    showToast(e.response?.data?.error || 'Failed to update comment', 'danger')
  } finally {
    commentSaving.value = false
  }
}

async function deleteComment(c) {
  try {
    await commentService.delete(c.id)
    comments.value = comments.value.filter(x => x.id !== c.id)
  } catch (e) {
    showToast(e.response?.data?.error || 'Failed to delete comment', 'danger')
  }
}

function editPost(p) {
  activePost.value = p
  postForm.value   = { title: p.title, tag: p.tag, content: p.content }
  editModal.value  = true
}

function confirmDeletePost(p) { activePost.value = p; deleteModal.value = true }

function closePostModal() {
  newPostModal.value  = false
  editModal.value     = false
  postForm.value      = { title: '', tag: '', content: '' }
  postFormError.value = ''
}

async function submitPost() {
  postFormError.value = ''
  if (!postForm.value.title.trim())   { postFormError.value = 'Title is required.'; return }
  if (!postForm.value.tag)            { postFormError.value = 'Please select a tag.'; return }
  if (!postForm.value.content.trim()) { postFormError.value = 'Content is required.'; return }

  postLoading.value = true
  try {
    if (editModal.value) {
      const updated = await postService.update(activePost.value.id, postForm.value)
      // update in both lists
      const mi = myPosts.value.findIndex(p => p.id === updated.id)
      if (mi !== -1) myPosts.value[mi] = updated
      const ci = communityPosts.value.findIndex(p => p.id === updated.id)
      if (ci !== -1) communityPosts.value[ci] = { ...updated, upvoted: communityPosts.value[ci].upvoted, saved: communityPosts.value[ci].saved }
      showToast('Post updated', 'success')
    } else {
      const created = await postService.create(postForm.value)
      // add to top of both lists
      myPosts.value.unshift(created)
      communityPosts.value.unshift({ ...created, upvoted: false, saved: false })
      showToast('Post published!', 'success')
    }
    closePostModal()
  } catch (e) {
    postFormError.value = e.response?.data?.error || 'Failed to save post. Please try again.'
  } finally {
    postLoading.value = false
  }
}

async function doDeletePost() {
  try {
    await postService.delete(activePost.value.id)
    myPosts.value        = myPosts.value.filter(p => p.id !== activePost.value.id)
    communityPosts.value = communityPosts.value.filter(p => p.id !== activePost.value.id)
    deleteModal.value    = false
    showToast('Post deleted', 'danger')
  } catch (e) {
    showToast(e.response?.data?.error || 'Failed to delete post', 'danger')
    deleteModal.value = false
  }
}

async function handleUpvote(p) {
  if (!p || p.upvoted) return
  try {
    const updated = await postService.upvote(p.id)
    p.upvotes = updated.upvotes
    p.upvoted = true
  } catch {
    showToast('Failed to upvote', 'danger')
  }
}

function showToast(msg, type = 'success') {
  toast.value = { show: true, msg, type }
  setTimeout(() => toast.value.show = false, 3000)
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Syne:wght@400;600;700;800&display=swap');
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
.page { font-family: 'Syne', sans-serif; display: flex; flex-direction: column; gap: 1.25rem; animation: fadeIn 0.3s ease; }

.page-header { display: flex; align-items: flex-start; justify-content: space-between; gap: 1rem; flex-wrap: wrap; }
.page-title { font-size: 1.75rem; font-weight: 800; color: var(--text-primary); letter-spacing: -1px; }
.page-sub { font-size: 0.85rem; color: var(--text-faint); margin-top: 0.25rem; }
.btn-primary { display: flex; align-items: center; gap: 0.5rem; background: var(--accent); color: var(--bg-base); border: none; border-radius: 8px; padding: 0.65rem 1.1rem; font-family: 'Syne', sans-serif; font-size: 0.85rem; font-weight: 700; cursor: pointer; transition: opacity 0.2s; flex-shrink: 0; }
.btn-primary svg { width: 14px; height: 14px; }
.btn-primary:hover { opacity: 0.85; }

.tabs { display: flex; gap: 0.25rem; border-bottom: 1px solid var(--border); }
.tab { background: none; border: none; cursor: pointer; font-family: 'Syne', sans-serif; font-size: 0.85rem; font-weight: 600; color: var(--text-faint); padding: 0.6rem 1rem; border-bottom: 2px solid transparent; margin-bottom: -1px; transition: color 0.2s, border-color 0.2s; }
.tab:hover { color: var(--text-muted); }
.tab.active { color: var(--text-primary); border-bottom-color: var(--text-primary); }

.toolbar { display: flex; flex-direction: column; gap: 0.75rem; }
.search-wrap { display: flex; align-items: center; gap: 0.6rem; background: var(--bg-surface); border: 1px solid var(--border); border-radius: 8px; padding: 0.6rem 1rem; }
.search-wrap svg { width: 15px; height: 15px; color: var(--text-ghost); flex-shrink: 0; }
.search-wrap input { background: none; border: none; outline: none; color: var(--text-primary); font-family: 'Syne', sans-serif; font-size: 0.88rem; width: 100%; }
.search-wrap input::placeholder { color: var(--text-ghost); }
.tag-filters { display: flex; gap: 0.4rem; flex-wrap: wrap; }
.tag-chip { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 20px; padding: 0.3rem 0.75rem; font-family: 'Syne', sans-serif; font-size: 0.75rem; font-weight: 600; color: var(--text-faint); cursor: pointer; transition: all 0.2s; }
.tag-chip:hover { border-color: var(--border-mid); color: var(--text-muted); }
.tag-chip.active { background: var(--bg-elevated); border-color: var(--border-mid); color: var(--text-primary); }

.loading-state { display: flex; flex-direction: column; align-items: center; gap: 1rem; padding: 4rem 2rem; }
.loading-state p { font-size: 0.85rem; color: var(--text-faint); }
.spinner-lg { width: 32px; height: 32px; border: 3px solid var(--border); border-top-color: var(--text-muted); border-radius: 50%; animation: spin 0.8s linear infinite; }

.posts-list { display: flex; flex-direction: column; gap: 0.75rem; }
.post-card { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 12px; padding: 1.25rem; display: flex; gap: 1rem; align-items: flex-start; transition: border-color 0.2s, transform 0.2s; }
.post-card:hover { border-color: var(--border-mid); transform: translateY(-1px); }
.post-main { flex: 1; min-width: 0; display: flex; flex-direction: column; gap: 0.6rem; }
.post-top { display: flex; align-items: center; justify-content: space-between; flex-wrap: wrap; gap: 0.5rem; }
.post-author { display: flex; align-items: center; gap: 0.6rem; }
.author-av { width: 28px; height: 28px; border-radius: 50%; background: var(--bg-elevated); border: 1px solid var(--border-mid); display: flex; align-items: center; justify-content: center; font-size: 0.62rem; font-weight: 700; color: var(--text-muted); flex-shrink: 0; }
.author-name { font-size: 0.82rem; font-weight: 600; color: var(--text-secondary); }
.author-time { font-size: 0.7rem; color: var(--text-ghost); font-family: monospace; }
.post-tag { font-size: 0.68rem; font-weight: 700; background: var(--bg-elevated); color: var(--text-muted); padding: 0.18rem 0.55rem; border-radius: 4px; flex-shrink: 0; }
.post-title { font-size: 0.95rem; font-weight: 700; color: var(--text-primary); line-height: 1.4; cursor: pointer; transition: color 0.2s; }
.post-title:hover { color: var(--text-muted); }
.post-excerpt { font-size: 0.82rem; color: var(--text-faint); line-height: 1.6; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; }
.post-stats { display: flex; gap: 1rem; }
.pstat { display: flex; align-items: center; gap: 0.3rem; font-size: 0.75rem; color: var(--text-ghost); font-family: monospace; }
.pstat svg { width: 13px; height: 13px; }
.post-actions { display: flex; flex-direction: column; gap: 0.35rem; }
.pa-btn { background: none; border: 1px solid transparent; border-radius: 6px; padding: 0.4rem; cursor: pointer; color: var(--text-ghost); transition: all 0.15s; display: flex; align-items: center; }
.pa-btn svg { width: 15px; height: 15px; }
.pa-btn:hover { background: var(--bg-elevated); border-color: var(--border); color: var(--text-muted); }
.pa-btn.active { color: #4caf72; }
.pa-btn.danger:hover { background: rgba(220,60,60,0.06); color: #e07070; }

.empty-state { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 12px; padding: 4rem 2rem; display: flex; flex-direction: column; align-items: center; gap: 0.75rem; text-align: center; }
.empty-state svg { width: 32px; height: 32px; color: var(--text-ghost); }
.empty-state p { font-size: 0.9rem; color: var(--text-faint); }
.empty-btn { background: none; border: none; cursor: pointer; font-family: 'Syne', sans-serif; font-size: 0.85rem; color: var(--text-muted); text-decoration: underline; }
.empty-btn:hover { color: var(--text-primary); }

.overlay { position: fixed; inset: 0; background: rgba(0,0,0,0.65); display: flex; align-items: center; justify-content: center; z-index: 1000; padding: 1rem; backdrop-filter: blur(4px); animation: fadeIn 0.15s ease; }
.modal { background: var(--bg-surface); border: 1px solid var(--border); border-radius: 16px; width: 100%; max-width: 520px; overflow: hidden; animation: slideUp 0.2s ease; max-height: 90vh; display: flex; flex-direction: column; }
.modal-lg { max-width: 640px; }
.modal-sm { max-width: 380px; }
.modal-head { display: flex; align-items: center; justify-content: space-between; padding: 1.25rem 1.5rem; border-bottom: 1px solid var(--border-soft); flex-shrink: 0; }
.modal-head h3 { font-size: 1rem; font-weight: 700; color: var(--text-primary); }
.modal-close { background: none; border: none; color: var(--text-faint); cursor: pointer; font-size: 1rem; padding: 0.25rem; transition: color 0.2s; }
.modal-close:hover { color: var(--text-primary); }
.modal-body { padding: 1.5rem; overflow-y: auto; flex: 1; }
.modal-foot { display: flex; gap: 0.75rem; justify-content: flex-end; padding: 1.25rem 1.5rem; border-top: 1px solid var(--border-soft); flex-shrink: 0; flex-wrap: wrap; }

.view-meta { display: flex; align-items: center; justify-content: space-between; flex-wrap: wrap; gap: 0.5rem; margin-bottom: 1rem; }
.view-title { font-size: 1.15rem; font-weight: 800; color: var(--text-primary); line-height: 1.4; margin-bottom: 0.75rem; letter-spacing: -0.3px; }
.view-content { font-size: 0.9rem; color: var(--text-muted); line-height: 1.75; margin-bottom: 1rem; white-space: pre-wrap; }
.view-stats { display: flex; gap: 1.25rem; }

.form { display: flex; flex-direction: column; gap: 1.1rem; }
.field { display: flex; flex-direction: column; gap: 0.4rem; }
.field label { font-size: 0.78rem; font-weight: 600; color: var(--text-faint); }
.req { color: #e07070; }
.field input, .field textarea, .field select { background: var(--bg-elevated); border: 1px solid var(--border); border-radius: 8px; padding: 0.75rem 1rem; color: var(--text-primary); font-family: 'Syne', sans-serif; font-size: 0.9rem; outline: none; transition: border-color 0.2s; resize: vertical; }
.field input:focus, .field textarea:focus, .field select:focus { border-color: var(--border-mid); }
.field input::placeholder, .field textarea::placeholder { color: var(--text-ghost); }
.field select option { background: var(--bg-elevated); }
.char-count { font-size: 0.72rem; color: var(--text-ghost); text-align: right; margin-top: 0.25rem; font-family: monospace; }
.form-error { background: rgba(220,60,60,0.08); border: 1px solid rgba(220,60,60,0.2); color: #e07070; border-radius: 8px; padding: 0.7rem 1rem; font-size: 0.82rem; }

.confirm-icon { display: flex; justify-content: center; margin-bottom: 0.75rem; }
.confirm-icon svg { width: 40px; height: 40px; color: #e07070; }
.confirm-text { font-size: 0.88rem; color: var(--text-muted); line-height: 1.6; text-align: center; }
.confirm-text strong { color: var(--text-primary); }

.mbtn { padding: 0.65rem 1.25rem; border-radius: 8px; font-family: 'Syne', sans-serif; font-size: 0.85rem; font-weight: 700; cursor: pointer; border: none; transition: all 0.2s; display: flex; align-items: center; gap: 0.4rem; }
.mbtn-outline { background: transparent; border: 1px solid var(--border-mid); color: var(--text-muted); }
.mbtn-outline:hover { color: var(--text-primary); }
.mbtn-primary { background: var(--accent); color: var(--bg-base); }
.mbtn-primary:hover:not(:disabled) { opacity: 0.85; }
.mbtn-primary:disabled { opacity: 0.5; cursor: not-allowed; }
.mbtn-danger { background: rgba(220,60,60,0.1); color: #e07070; border: 1px solid rgba(220,60,60,0.2); }
.mbtn-danger:hover { background: rgba(220,60,60,0.2); }
.spinner { width: 16px; height: 16px; border: 2px solid rgba(0,0,0,0.15); border-top-color: var(--bg-base); border-radius: 50%; animation: spin 0.7s linear infinite; }

.toast { position: fixed; bottom: 1.5rem; right: 1.5rem; padding: 0.875rem 1.25rem; border-radius: 10px; font-family: 'Syne', sans-serif; font-size: 0.85rem; font-weight: 600; z-index: 2000; animation: slideUp 0.2s ease; }
.toast.success { background: var(--bg-surface); border: 1px solid rgba(76,175,114,0.3); color: #4caf72; }
.toast.danger  { background: var(--bg-surface); border: 1px solid rgba(220,60,60,0.3); color: #e07070; }

/* Comment modal */
.modal-comment { max-width: 520px; }
.cm-head-left { display: flex; align-items: flex-start; gap: 0.75rem; flex: 1; min-width: 0; }
.cm-head-left svg { width: 18px; height: 18px; color: var(--text-muted); flex-shrink: 0; margin-top: 2px; }
.cm-head-left h3 { font-size: 1rem; font-weight: 700; color: var(--text-primary); }
.cm-post-title { font-size: 0.78rem; color: var(--text-faint); margin-top: 0.15rem; display: -webkit-box; -webkit-line-clamp: 1; -webkit-box-orient: vertical; overflow: hidden; }

.cm-existing { background: var(--bg-elevated); border: 1px solid var(--border-soft); border-radius: 10px; padding: 0.875rem; margin-bottom: 1rem; }
.cm-existing-label { font-size: 0.72rem; font-weight: 700; color: var(--text-ghost); text-transform: uppercase; letter-spacing: 0.06em; margin-bottom: 0.6rem; }
.cm-preview-list { display: flex; flex-direction: column; gap: 0.5rem; }
.cm-preview-item { display: flex; align-items: flex-start; gap: 0.6rem; }
.comment-av.sm { width: 22px; height: 22px; font-size: 0.55rem; }
.cm-preview-body { flex: 1; min-width: 0; font-size: 0.8rem; line-height: 1.4; }
.cm-preview-author { font-weight: 700; color: var(--text-secondary); margin-right: 0.4rem; }
.cm-preview-text { color: var(--text-faint); display: -webkit-box; -webkit-line-clamp: 1; -webkit-box-orient: vertical; overflow: hidden; }
.cm-more { font-size: 0.75rem; color: var(--text-muted); cursor: pointer; margin-top: 0.4rem; padding-top: 0.4rem; border-top: 1px solid var(--border-soft); }
.cm-more:hover { color: var(--text-primary); }

.cm-no-comments { display: flex; align-items: center; gap: 0.6rem; font-size: 0.82rem; color: var(--text-ghost); padding: 0.875rem; background: var(--bg-elevated); border: 1px solid var(--border-soft); border-radius: 10px; margin-bottom: 1rem; }
.cm-no-comments svg { width: 16px; height: 16px; flex-shrink: 0; }

.cm-write { display: flex; flex-direction: column; gap: 0.6rem; }
.cm-write-top { display: flex; align-items: center; gap: 0.6rem; }
.cm-writing-as { font-size: 0.82rem; font-weight: 600; color: var(--text-secondary); }
.cm-textarea {
  width: 100%; background: var(--bg-elevated); border: 1px solid var(--border);
  border-radius: 10px; padding: 0.875rem 1rem;
  color: var(--text-primary); font-family: 'Syne', sans-serif; font-size: 0.9rem;
  outline: none; resize: vertical; transition: border-color 0.2s; line-height: 1.6;
}
.cm-textarea:focus { border-color: var(--border-mid); }
.cm-textarea::placeholder { color: var(--text-ghost); }
.cm-char-count { font-size: 0.7rem; color: var(--text-ghost); text-align: right; font-family: monospace; }
.cm-error { background: rgba(220,60,60,0.08); border: 1px solid rgba(220,60,60,0.2); color: #e07070; border-radius: 8px; padding: 0.65rem 1rem; font-size: 0.82rem; }

/* Comments */
.comments-section { border-top: 1px solid var(--border-soft); margin-top: 1.25rem; padding-top: 1.25rem; display: flex; flex-direction: column; gap: 1rem; }
.comments-header { display: flex; align-items: center; gap: 0.5rem; }
.comments-header h4 { font-size: 0.88rem; font-weight: 700; color: var(--text-secondary); }
.comments-count { font-size: 0.72rem; background: var(--bg-elevated); color: var(--text-faint); padding: 0.1rem 0.5rem; border-radius: 10px; font-family: monospace; }
.comments-loading { display: flex; justify-content: center; padding: 1.5rem 0; }
.comments-list { display: flex; flex-direction: column; gap: 0.875rem; }
.comment-item { display: flex; gap: 0.75rem; align-items: flex-start; }
.comment-av { width: 28px; height: 28px; border-radius: 50%; background: var(--bg-elevated); border: 1px solid var(--border-mid); display: flex; align-items: center; justify-content: center; font-size: 0.6rem; font-weight: 700; color: var(--text-muted); flex-shrink: 0; margin-top: 2px; }
.comment-body { flex: 1; min-width: 0; }
.comment-meta { display: flex; align-items: center; gap: 0.5rem; margin-bottom: 0.25rem; flex-wrap: wrap; }
.comment-author { font-size: 0.82rem; font-weight: 700; color: var(--text-secondary); }
.comment-time { font-size: 0.7rem; color: var(--text-ghost); font-family: monospace; }
.comment-edited { font-size: 0.65rem; color: var(--text-ghost); font-style: italic; }
.comment-content { font-size: 0.85rem; color: var(--text-muted); line-height: 1.6; white-space: pre-wrap; }
.comment-actions { display: flex; gap: 0.5rem; margin-top: 0.3rem; }
.ca-btn { background: none; border: none; cursor: pointer; font-family: 'Syne', sans-serif; font-size: 0.72rem; color: var(--text-ghost); padding: 0; transition: color 0.2s; }
.ca-btn:hover { color: var(--text-muted); }
.ca-btn.danger:hover { color: #e07070; }
.comment-edit-wrap { display: flex; flex-direction: column; gap: 0.4rem; }
.comment-edit-input { background: var(--bg-elevated); border: 1px solid var(--border-mid); border-radius: 8px; padding: 0.6rem 0.75rem; color: var(--text-primary); font-family: 'Syne', sans-serif; font-size: 0.85rem; outline: none; resize: vertical; width: 100%; }
.comment-edit-actions { display: flex; gap: 0.5rem; justify-content: flex-end; }
.no-comments { font-size: 0.82rem; color: var(--text-ghost); text-align: center; padding: 1.25rem 0; }

.comment-form { display: flex; gap: 0.75rem; align-items: flex-start; padding-top: 0.5rem; border-top: 1px solid var(--border-soft); }
.comment-input-wrap { flex: 1; display: flex; flex-direction: column; gap: 0.4rem; }
.comment-input { background: var(--bg-elevated); border: 1px solid var(--border); border-radius: 10px; padding: 0.7rem 0.875rem; color: var(--text-primary); font-family: 'Syne', sans-serif; font-size: 0.88rem; outline: none; resize: none; transition: border-color 0.2s; width: 100%; }
.comment-input:focus { border-color: var(--border-mid); }
.comment-input::placeholder { color: var(--text-ghost); }
.comment-form-foot { display: flex; align-items: center; justify-content: space-between; }
.comment-hint { font-size: 0.7rem; color: var(--text-ghost); font-family: monospace; }
.cta-btn { padding: 0.45rem 1rem; border-radius: 6px; font-family: 'Syne', sans-serif; font-size: 0.8rem; font-weight: 700; cursor: pointer; border: none; transition: all 0.2s; display: flex; align-items: center; gap: 0.35rem; }
.cta-post { background: var(--accent); color: var(--bg-base); }
.cta-post:hover:not(:disabled) { opacity: 0.85; }
.cta-post:disabled { opacity: 0.4; cursor: not-allowed; }
.cta-save { background: var(--accent); color: var(--bg-base); font-size: 0.75rem; padding: 0.35rem 0.75rem; }
.cta-cancel { background: var(--bg-elevated); color: var(--text-muted); border: 1px solid var(--border); font-size: 0.75rem; padding: 0.35rem 0.75rem; }

@keyframes fadeIn  { from { opacity: 0; } to { opacity: 1; } }
@keyframes slideUp { from { opacity: 0; transform: translateY(12px); } to { opacity: 1; transform: translateY(0); } }
@keyframes spin    { to { transform: rotate(360deg); } }

@media (max-width: 768px) {
  .post-card { flex-direction: column; }
  .post-actions { flex-direction: row; justify-content: flex-end; }
  .page-header { flex-direction: column; align-items: flex-start; }
  .btn-primary { width: 100%; justify-content: center; }
  .modal { max-width: 100%; border-radius: 20px 20px 0 0; position: fixed; bottom: 0; left: 0; right: 0; max-height: 92vh; }
  .overlay { align-items: flex-end; padding: 0; }
  .modal-foot { justify-content: stretch; }
  .mbtn { flex: 1; justify-content: center; }
}
</style>