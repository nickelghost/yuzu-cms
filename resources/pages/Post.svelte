<script>
export let params = {}
import { onMount } from 'svelte';
import Editor from '../components/Editor.svelte';

let post = {};
let newContent = "";

onMount(async () => {
  const res = await fetch(`/api/v1/posts/${params.id}`);
  post = await res.json();
  newContent = post.content;
});

function newContentUpdate(e) {
  newContent = e.detail.content;
}

async function savePost() {
  const req = {
    title: post.title,
    content: newContent,
  };
  const res = await fetch(
    `/api/v1/posts/${post.id}`,
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
.post-page {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}
.top-bar {
  display: flex;
}

.top-bar-spacer {
  flex-grow: 1;
}
</style>

<div class="post-page">
  <div class="top-bar">
    <h1>{post.title}</h1>
    <div class="top-bar-spacer"></div>
    <button on:click={savePost}>Save</button>
  </div>
  <Editor
    content={post.content}
    on:change={newContentUpdate}
  />
</div>
