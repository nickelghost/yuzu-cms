<script>
  import TopBar from '../components/TopBar.svelte';
  import Editor from '../components/Editor.svelte';
  import Notification from '../components/Notification.svelte';

  import { jwt } from '../stores';

  let title = '';
  let content = '';

  let notificationMessage = '';
  let notificationColor = '';

  async function createPost({ is_draft = false } = {}) {
    const req = { title, content, is_draft };
    const res = await fetch('/api/v1/posts', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${$jwt}`,
      },
      body: JSON.stringify(req),
    });
    if (res.ok) {
      notificationMessage = 'Post created';
      notificationColor = 'green';
    } else {
      notificationMessage = 'Could not create post';
      notificationColor = 'red';
    }
  }

  function onClickDraft() {
    createPost({ is_draft: true });
  }
  function onClickPublish() {
    createPost();
  }
</script>

<style>
  .title-section {
    display: flex;
    flex-direction: column;
    margin-bottom: 24px;
  }

  .label {
    font-size: 1.5rem;
    margin-bottom: 8px;
  }
</style>

<Notification
  bind:message="{notificationMessage}"
  color="{notificationColor}"
></Notification>

<TopBar title="New Post">
  <button class="button" on:click="{onClickDraft}">Draft</button>
  <div class="top-bar-spacer"></div>
  <button class="button button-primary" on:click="{onClickPublish}">
    Publish
  </button>
</TopBar>
<main class="content">
  <label class="title-section">
    <span class="label">Title</span>
    <input class="input" name="title" bind:value="{title}" />
  </label>
  <Editor bind:content="{content}"></Editor>
</main>
