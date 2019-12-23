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

function savePost() {
  console.log(newContent);
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
