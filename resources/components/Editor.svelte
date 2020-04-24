<script>
  import * as MarkdownIt from 'markdown-it';

  export let content = '';

  const md = new MarkdownIt();

  $: contentHTML = md.render(content || '');
</script>

<style>
  .editor {
    display: flex;
    flex-grow: 1;
  }

  .markdown,
  .preview {
    flex: 1 1 0px;
  }

  .markdown {
    margin-bottom: 0;
  }

  .markdown-input,
  .preview-content {
    flex-grow: 1;
  }

  .markdown-input {
    resize: none;
  }

  .preview-content {
    border: 1px solid #ececec;
    padding: 0 8px;
  }

  .spacer {
    width: 12px;
  }
</style>

<div class="editor">
  <label class="field markdown">
    <span class="label">Content</span>
    <textarea class="input markdown-input" bind:value="{content}"></textarea>
  </label>
  <div class="spacer"></div>
  <label class="field preview">
    <span class="label">Content Preview</span>
    <div class="preview-content">{@html contentHTML}</div>
  </label>
</div>
