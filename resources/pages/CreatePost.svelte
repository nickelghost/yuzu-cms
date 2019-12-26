<script>
import TopBar from '../components/TopBar.svelte';
import Button from '../components/Button.svelte';
import Content from '../components/Content.svelte';
import Editor from '../components/Editor.svelte';

let title = "";
let content = "";

async function createPost() {
  const req = { title, content };
  const res = await fetch(
    '/api/v1/posts',
    {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(req),
    }
  );
  if (res.ok) {
    alert('Post created');
  } else {
    alert('Error creating post')
  }
}
</script>

<TopBar title="Create a new post">
  <Button label="Save" color="green" on:click={createPost}></Button>
</TopBar>
<Content>
  <h3>Title</h3>
  <input class="title-input" name="title" bind:value={title} />
  <h3>Content</h3>
  <Editor bind:content={content}/>
</Content>
