<script>
export let params = {}
import { onMount } from 'svelte';
import Editor from '../components/Editor.svelte';

let title = "";
let content = "";

onMount(async () => {
  const res = await fetch(`/api/v1/posts/${params.id}`);
  const post = await res.json();
  title = post.title;
  content = post.content;
});

async function savePost() {
  const req = { title, content };
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
    <h1>{title}</h1>
    <div class="top-bar-spacer"></div>
    <button on:click={savePost}>Save</button>
  </div>
  <Editor bind:content={content}/>
</div>
