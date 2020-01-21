<script>
import TopBar from '../components/TopBar.svelte';
import Editor from '../components/Editor.svelte';
import Input from '../components/Input.svelte';

let title = "";
let content = "";

function onClickDraft() { createPost({ is_draft: true }); }
function onClickPublish() { createPost(); }
async function createPost({
  is_draft = false,
} = {}) {
  const req = { title, content, is_draft };
  const res = await fetch(
    '/api/v1/posts',
    {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(req),
    }
  );
  if (res.ok) {
    alert('Post created');
  } else {
    alert('Error creating post')
  }
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

<TopBar title="Create a new post">
  <button class="button" on:click={onClickDraft}>Draft</button>
  <div class="top-bar-spacer"></div>
  <button class="button" on:click={onClickPublish}>Publish</button>
</TopBar>
<main class="content">
  <label class="title-section">
    <span class="label">Title</span>
    <Input name="title" bind:value={title}></Input>
  </label>
  <Editor bind:content={content}/>
</main>
