<script>
export let params = {}
import { onMount } from 'svelte';

import TopBar from '../components/TopBar.svelte';
import Button from '../components/Button.svelte';
import Content from '../components/Content.svelte';
import Editor from '../components/Editor.svelte';
import Modal from '../components/Modal.svelte';

let title = "";
let content = "";
let isTitleModalOpen = false;
let newTitle = "";

onMount(async () => {
  const res = await fetch(`/api/v1/posts/${params.id}`);
  const post = await res.json();
  title = post.title;
  content = post.content;
});

function openTitleModal() { isTitleModalOpen = true; }
function closeTitleModal() { isTitleModalOpen = false; }
function updateTitle() {
  title = newTitle;
  isTitleModalOpen = false;
  newTitle = "";
}

async function savePost({ is_draft = false }) {
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
    alert("Post updated");
  } else {
    alert("Could not update");
  }
}
</script>

<style>
.title-modal {
  display: flex;
  flex-direction: column;
  width: 480px;
}
.title-modal-label {
  font-size: 1.125rem;
  margin-bottom: 8px;
}
.title-modal-input {
  margin-bottom: 16px;
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
    <label class="title-modal-label">
      Enter the new title:
    </label>
    <input class="title-modal-input" bind:value={newTitle}>
    <div class="title-modal-buttons">
      <Button label="Cancel" on:click={closeTitleModal}></Button>
      <div class="spacer"></div>
      <Button label="Change" on:click={updateTitle}></Button>
    </div>
  </div>
</Modal>

<TopBar title={title}>
  <Button label="Change title" on:click={openTitleModal}></Button>
  <Button label="Draft" color="yellow" on:click={() => savePost({ is_draft: true })}></Button>
  <Button label="Publish" color="green" on:click={() => savePost({ is_draft: false })}></Button>
</TopBar>
<Content>
  <Editor bind:content={content}/>
</Content>
