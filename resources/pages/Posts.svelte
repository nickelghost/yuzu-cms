<script>
import { onMount } from 'svelte';
import { push } from 'svelte-spa-router';
import format from 'date-fns/format';

import TopBar from '../components/TopBar.svelte';
import Content from '../components/Content.svelte';
import Button from '../components/Button.svelte';

function displayDate(dateString) {
  const date = new Date(dateString);
  return format(date, 'do MMM yyyy HH:mm')
}

let posts = [];

onMount(async () => {
  const res = await fetch('/api/v1/posts');
  posts = await res.json();
});
</script>

<style>
.table td {
  text-align: center;
}
.table tr {
  height: 30px;
}
.table tr:hover {
  background-color: lightgray;
  cursor: pointer;
}
</style>

<TopBar title="Posts">
  <Button on:click={() => push('/posts/new')} label="New"></Button>
</TopBar>
<Content>
  <table class="table">
    <tr>
      <th>ID</th>
      <th>Title</th>
      <th>Created At</th>
      <th>Updated at</th>
    </tr>
    {#each posts as post}
      <tr on:click={() => push(`/posts/${post.id}`)}>
        <td>{post.id}</td>
        <td>{post.title}</td>
        <td>{displayDate(post.created_at)}</td>
        <td>{displayDate(post.updated_at)}</td>
      </tr>
    {/each}
  </table>
</Content>
