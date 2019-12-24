<script>
export let content = "";
import MarkdownIt from 'markdown-it';
import { createEventDispatcher } from 'svelte';

const md = new MarkdownIt();
const dispatch = createEventDispatcher();

$: contentHTML = md.render(content || '');

function dispatchChange() {
  dispatch('change', { content });
}
</script>

<style>
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

<div class="editor">
  <div class="markdown">
    <textarea
      class="textarea"
      bind:value={content}
      on:change={dispatchChange}
    ></textarea>
  </div>
  <div class="preview">
    {@html contentHTML}
  </div>
</div>
