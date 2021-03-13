<script>
  import { push } from 'svelte-spa-router';

  import TopBar from '../components/TopBar.svelte';
  import Editor from '../components/Editor.svelte';
  import Modal from '../components/Modal.svelte';

  import notify from '../helpers/notify';

  import { jwt } from '../stores';

  let title = '';
  let content = '';
  let slug = '';

  let isSlugModalOpen = false;

  function onTitleChange(e) {
    const oldTitle = title;
    title = e.target.value;
    if (oldTitle.toLowerCase().split(' ').join('-') === slug) {
      slug = title.toLowerCase().split(' ').join('-');
    }
  }

  async function createPost({ is_draft = false } = {}) {
    const req = {
      title,
      content,
      slug,
      is_draft,
    };
    const res = await fetch('/api/v1/posts', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${$jwt}`,
      },
      body: JSON.stringify(req),
    });
    if (res.ok) {
      notify('Post created', 'green');
      push('/posts');
    } else {
      notify('Could not create the post', 'red');
    }
  }
</script>

<Modal bind:isOpen="{isSlugModalOpen}">
  <form class="form" on:submit="{(e) => e.preventDefault()}">
    <label class="field">
      <span class="label">Slug</span>
      <input class="input" bind:value="{slug}" />
    </label>
    <div class="field field-horizontal">
      <div class="flex-spacer"></div>
      <button
        class="button"
        type="button"
        on:click="{() => { isSlugModalOpen = false; }}"
      >
        Close
      </button>
    </div>
  </form>
</Modal>

<TopBar title="New Post">
  <button class="button" on:click="{() => { isSlugModalOpen = true; }}">
    Custom slug
  </button>
  <div class="top-bar-spacer"></div>
  <button class="button" on:click="{() => { createPost({ is_draft: true }); }}">
    Draft
  </button>
  <div class="top-bar-spacer"></div>
  <button class="button button-primary" on:click="{createPost}">
    Publish
  </button>
</TopBar>
<main class="content">
  <label class="field">
    <span class="label">Title</span>
    <input
      class="input"
      name="title"
      value="{title}"
      on:change="{onTitleChange}"
    />
  </label>
  <Editor bind:content="{content}"></Editor>
</main>
