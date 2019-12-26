<script>
export let params = {}
import { onMount } from 'svelte';

import TopBar from '../components/TopBar.svelte';
import Button from '../components/Button.svelte';
import Content from '../components/Content.svelte';
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

<TopBar title={title}>
  <Button label="Save" on:click={savePost}></Button>
</TopBar>
<Content>
  <Editor bind:content={content}/>
</Content>
