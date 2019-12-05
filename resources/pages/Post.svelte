<script>
import { onMount } from 'svelte';
import MarkdownIt from 'markdown-it';
export let params = {}

const md = new MarkdownIt();

let post = {};

$: contentHTML = md.render(post.content || '');

onMount(async () => {
  const res = await fetch(`/api/v1/posts/${params.id}`);
  post = await res.json();
});
</script>

<style>
.editor {
  display: flex;
}

.markdown, .preview {
  flex-grow: 1;
  padding: 8px;
}

.markdown {
  display: flex;
}

.textarea {
  flex-grow: 1;
}
</style>

<h1>{post.title}</h1>

<div class="editor">
  <div class="markdown">
    <textarea class="textarea" bind:value={post.content}></textarea>
  </div>
  <div class="preview">
    {@html contentHTML}
  </div>
</div>
