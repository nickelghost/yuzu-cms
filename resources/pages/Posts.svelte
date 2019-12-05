<script>
import { onMount } from 'svelte';

let posts = [];

onMount(async () => {
  const res = await fetch('/api/v1/posts');
  posts = await res.json();
});
</script>

<h1>Posts</h1>

<ul>
  {#each posts as post}
    <li>{post.title}</li>
  {/each}
</ul>
