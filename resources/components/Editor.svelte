<script>
export let content = "";
import MarkdownIt from 'markdown-it';

const md = new MarkdownIt();

$: contentHTML = md.render(content || '');
</script>

<style>
.editor {
  display: flex;
  flex-grow: 1;
}
.markdown, .preview {
  display: flex;
  flex: 1 1 0px;
  flex-direction: column;
}
.label {
  font-size: 1.5rem;
  margin-bottom: 8px;
}
.markdown-input, .preview-content {
  border: 1px solid #ececec;
  flex-grow: 1;
}
.markdown-input {
  font-family: 'Fira Code', monospace;
  padding: 12px;
  resize: none;
}
.preview-content {
  padding: 0 12px;
}
.spacer {
  width: 12px;
}
</style>

<div class="editor">
  <label class="markdown">
    <span class="label">Content</span>
    <textarea class="markdown-input" bind:value={content}></textarea>
  </label>
  <div class="spacer"></div>
  <label class="preview">
    <span class="label">Content Preview</span>
    <div class="preview-content">{@html contentHTML}</div>
  </label>
</div>
