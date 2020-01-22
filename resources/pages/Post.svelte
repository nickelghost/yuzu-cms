<script>
import { onMount } from 'svelte';

import TopBar from '../components/TopBar.svelte';
import Editor from '../components/Editor.svelte';
import Modal from '../components/Modal.svelte';

export let params = {};

let title = '';
let content = '';
let isDraft = false;
let isTitleModalOpen = false;
let newTitle = '';

onMount(async () => {
  const res = await fetch(`/api/v1/posts/${params.id}`);
  const post = await res.json();
  title = post.title;
  content = post.content;
  isDraft = post.is_draft;
});

function openTitleModal() { isTitleModalOpen = true; }
function closeTitleModal() { isTitleModalOpen = false; }

function updateTitle() {
  title = newTitle;
  isTitleModalOpen = false;
  newTitle = '';
}

async function savePost({ is_draft = false } = {}) {
  const req = { title, content, is_draft };
  const res = await fetch(
    `/api/v1/posts/${params.id}`,
    {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(req),
    },
  );
  if (res.ok) {
    alert('Post updated');
  } else {
    alert('Could not update');
  }
}
function onClickDraft() { savePost({ is_draft: true }); }
function onClickPublish() { savePost(); }
</script>

<style>
.title-modal {
  display: flex;
  flex-direction: column;
  width: 480px;
}
.title-modal-field {
  display: flex;
  flex-direction: column;
  margin-bottom: 16px;
}
.title-modal-label {
  font-size: 1.125rem;
  margin-bottom: 8px;
}
.title-modal-buttons {
  display: flex;
}
.spacer {
  flex-grow: 1;
}
</style>

<Modal bind:isOpen={isTitleModalOpen}>
  <div class="title-modal">
    <label class="title-modal-field">
      <span class="title-modal-label">Enter the new title:</span>
      <input class="input" bind:value={newTitle}>
    </label>
    <div class="title-modal-buttons">
      <button class="button" on:click={closeTitleModal}>Cancel</button>
      <div class="spacer"></div>
      <button class="button" on:click={updateTitle}>Change</button>
    </div>
  </div>
</Modal>

<TopBar title={title + (isDraft ? ' (draft)' : '')}>
  <button class="button" on:click={openTitleModal}>Change title</button>
  <div class="top-bar-spacer"></div>
  <button class="button" on:click={onClickDraft}>Draft</button>
  <div class="top-bar-spacer"></div>
  <button class="button" on:click={onClickPublish}>Publish</button>
</TopBar>
<div class="content">
  <Editor bind:content={content}/>
</div>
