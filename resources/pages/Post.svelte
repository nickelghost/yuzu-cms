<script>
  import { onMount } from 'svelte';
  import { push } from 'svelte-spa-router';

  import TopBar from '../components/TopBar.svelte';
  import Editor from '../components/Editor.svelte';
  import Modal from '../components/Modal.svelte';

  import notify from '../helpers/notify';

  import { jwt } from '../stores';

  export let params = {};

  let title = '';
  let content = '';
  let slug = '';
  let isDraft = false;

  let isTitleModalOpen = false;
  let newTitle = '';
  let isSlugModalOpen = false;
  let newSlug = '';
  let isDeleteModalOpen = false;

  onMount(async () => {
    const res = await fetch(`/api/v1/posts/${params.id}`, {
      headers: { Authorization: `Bearer ${$jwt}` },
    });
    const post = await res.json();
    title = post.title;
    content = post.content;
    slug = post.slug;
    isDraft = post.is_draft;
  });

  function updateTitle(e) {
    e.preventDefault();
    title = newTitle;
    isTitleModalOpen = false;
    newTitle = '';
  }

  function updateSlug(e) {
    e.preventDefault();
    slug = newSlug;
    isSlugModalOpen = false;
    newSlug = '';
  }

  async function deletePost() {
    const res = await fetch(`/api/v1/posts/${params.id}`, {
      method: 'DELETE',
      headers: { Authorization: `Bearer ${$jwt}` },
    });
    if (res.ok) {
      notify(`Deleted ${title}`, 'green');
      push('/posts');
    } else {
      isDeleteModalOpen = false;
      notify('Could not delete the post', 'red');
    }
  }

  async function savePost({ is_draft = false } = {}) {
    const req = {
      title,
      content,
      slug,
      is_draft,
    };
    const res = await fetch(`/api/v1/posts/${params.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${$jwt}`,
      },
      body: JSON.stringify(req),
    });
    if (res.ok) {
      notify('Post updated', 'green');
    } else {
      notify('Could not update the post', 'red');
    }
  }
</script>

<Modal bind:isOpen="{isTitleModalOpen}">
  <form class="form" on:submit="{updateTitle}">
    <label class="field">
      <span class="label">Title</span>
      <input class="input" bind:value="{newTitle}" />
    </label>
    <div class="field field-horizontal">
      <button
        class="button"
        type="button"
        on:click="{() => { isTitleModalOpen = false; }}"
      >
        Cancel
      </button>
      <div class="flex-spacer"></div>
      <button class="button button-primary" type="submit">
        Change
      </button>
    </div>
  </form>
</Modal>

<Modal bind:isOpen="{isSlugModalOpen}">
  <form class="form" on:submit="{updateSlug}">
    <label class="field">
      <span class="label">Slug</span>
      <input class="input" bind:value="{newSlug}" placeholder="{slug}" />
    </label>
    <div class="field field-horizontal">
      <button
        class="button"
        type="button"
        on:click="{() => { isSlugModalOpen = false; }}"
      >
        Cancel
      </button>
      <div class="flex-spacer"></div>
      <button class="button button-primary" type="submit">
        Change
      </button>
    </div>
  </form>
</Modal>

<Modal bind:isOpen="{isDeleteModalOpen}">
  <form class="form" on:submit="{deletePost}">
    <p>Are you sure that you want to delete {title}?</p>
    <div class="field field-horizontal">
      <button
        class="button"
        type="button"
        on:click="{() => { isDeleteModalOpen = false; }}"
      >
        Cancel
      </button>
      <div class="flex-spacer"></div>
      <button class="button button-danger" type="submit">
        Yes, delete
      </button>
    </div>
  </form>
</Modal>

<TopBar title="{title + (isDraft ? ' (draft)' : '')}">
  <button
    class="button button-danger"
    on:click="{() => { isDeleteModalOpen = true; }}"
  >
    Delete post
  </button>
  <div class="top-bar-spacer"></div>
  <button class="button" on:click="{() => { isSlugModalOpen = true; }}">
    Change slug
  </button>
  <div class="top-bar-spacer"></div>
  <button class="button" on:click="{() => { isTitleModalOpen = true; }}">
    Change title
  </button>
  <div class="top-bar-spacer"></div>
  <button class="button" on:click="{() => { savePost({ is_draft: true }); }}">
    Draft
  </button>
  <div class="top-bar-spacer"></div>
  <button class="button button-primary" on:click="{savePost}">
    Publish
  </button>
</TopBar>
<div class="content">
  <Editor bind:content="{content}"></Editor>
</div>
