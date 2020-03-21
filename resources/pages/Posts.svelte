<script>
  import { onMount } from 'svelte';
  import { push } from 'svelte-spa-router';
  import format from 'date-fns/format';

  import TopBar from '../components/TopBar.svelte';

  import { jwt } from '../stores';

  let posts = [];

  onMount(async () => {
    const res = await fetch('/api/v1/posts', {
      headers: { Authorization: `Bearer ${$jwt}` },
    });
    posts = await res.json();
  });

  function displayDate(dateString) {
    const date = new Date(dateString);
    return format(date, 'do MMM yyyy HH:mm');
  }

  function redirectToNew() {
    push('/posts/new');
  }
</script>

<style>
  .post {
    display: flex;
    border: 1px solid #ececec;
    margin-bottom: 24px;
    padding: 24px;
  }

  .post:hover {
    /* background-color: rgba(0, 0, 0, 0.05); */
    box-shadow: 0 0 8px #ececec;
    cursor: pointer;
  }

  .draft {
    background-color: #fff9c6;
  }

  .title-section {
    flex-shrink: 0;
    width: 20%;
  }

  .title {
    font-weight: bold;
    font-size: 1.25rem;
    margin: 0;
  }

  .preview-section {
    flex-shrink: 0;
    font-size: 1.125rem;
    padding: 0 12px;
    width: 40%;
  }

  .preview {
    margin: 0;
  }

  .created-section {
    padding: 0 12px;
  }

  .created-section,
  .updated-section {
    flex-shrink: 0;
    width: 20%;
  }

  .time-heading {
    display: block;
    font-size: 0.875rem;
    font-weight: bold;
    text-transform: uppercase;
  }
</style>

<TopBar title="Posts">
  <button class="button button-primary" on:click="{redirectToNew}">
    New
  </button>
</TopBar>
<div class="content">
  {#each posts as post}
  <div
    class="post"
    class:draft="{post.is_draft}"
    on:click="{() => push(`/posts/${post.id}`)}"
  >
    <div class="title-section">
      <h3 class="title">
        {post.title}
      </h3>
    </div>
    <div class="preview-section">
      <p class="preview">
        {post.content_preview}...
      </p>
    </div>
    <div class="created-section">
      <span class="time-heading">Created at</span>
      {displayDate(post.created_at)}
    </div>
    <div class="updated-section">
      <span class="time-heading">Updated at</span>
      {displayDate(post.updated_at)}
    </div>
  </div>
  {/each}
</div>
