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
    <textarea class="textarea" bind:value={content}/>
  </div>
  <div class="preview">
    {@html contentHTML}
  </div>
</div>
