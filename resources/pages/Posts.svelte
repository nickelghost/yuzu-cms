<script>
import { onMount } from 'svelte';
import { push } from 'svelte-spa-router';

let posts = [];

onMount(async () => {
  const res = await fetch('/api/v1/posts');
  posts = await res.json();
});
</script>

<style>
.posts-page {
  display: flex;
  flex-grow: 1;
  flex-direction: column;
}
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

<div class="posts-page">
  <h1>Posts</h1>
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
        <td>{post.created_at}</td>
        <td>{post.updated_at}</td>
      </tr>
    {/each}
  </table>
</div>
