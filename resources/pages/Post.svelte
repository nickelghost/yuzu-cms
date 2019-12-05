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
.post-page {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}
.editor {
  display: flex;
  flex-grow: 1;
}
.markdown, .preview {
  flex: 1 1 0px;
  padding: 8px;
}
.markdown {
  display: flex;
}
.textarea {
  flex-grow: 1;
  min-height: 100%;
  resize: vertical;
}
</style>

<div class="post-page">
  <h1>{post.title}</h1>
  <div class="editor">
    <div class="markdown">
      <textarea class="textarea" bind:value={post.content}></textarea>
    </div>
    <div class="preview">
      {@html contentHTML}
    </div>
  </div>
</div>
