<script>
import Editor from '../components/Editor.svelte';

let title = "";
let content = "";

function titleUpdate(e) {
  title = e.target.value;
}

function contentUpdate(e) {
  content = e.detail.content;
}

async function createPost() {
  const req = { title, content };
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
.create-post-page {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}
.top-bar {
  display: flex;
}
.title-input {
  flex-grow: 1;
}
</style>

<div class="create-post-page">
  <div class="top-bar">
    <input class="title-input" name="title" on:change={titleUpdate} />
    <button on:click={createPost}>Create</button>
  </div>
  <Editor on:change={contentUpdate}/>
</div>
