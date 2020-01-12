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
  border: 1px solid lightgrey;
  border-radius: 8px;
  flex: 1 1 0px;
}
.markdown {
  display: flex;
  resize: none;
  padding: 12px;
}
.preview {
  padding: 0 12px;
}
.spacer {
  width: 12px;
}
</style>

<div class="editor">
  <textarea class="markdown" bind:value={content}></textarea>
  <div class="spacer"></div>
  <div class="preview">{@html contentHTML}</div>
</div>
