<script>
  import { onMount } from 'svelte';
  import { push } from 'svelte-spa-router';
  import format from 'date-fns/format';

  import TopBar from '../components/TopBar.svelte';

  let posts = [];

  onMount(async () => {
    const res = await fetch('/api/v1/posts');
    posts = await res.json();
  });

  function displayDate(dateString) {
    const date = new Date(dateString);
    return format(date, 'do MMM yyyy HH:mm');
  }

  function redirectToNew() {
    push('/posts/new');
  }
</script>

<style>
  .table {
    border-spacing: 0;
    border-collapse: collapse;
    width: 100%;
  }

  .table th,
  .table td {
    text-align: center;
    padding: 16px 0;
  }

  .table tr:hover:not(:first-child) {
    background-color: whitesmoke;
    cursor: pointer;
  }
</style>

<TopBar title="Posts">
  <button class="button button-primary" on:click="{redirectToNew}">
    New
  </button>
</TopBar>
<div class="content">
  <table class="table">
    <tr>
      <th>ID</th>
      <th>Title</th>
      <th>Draft</th>
      <th>Created at</th>
      <th>Updated at</th>
    </tr>
    {#each posts as post}
    <tr on:click="{() => push(`/posts/${post.id}`)}">
      <td>{post.id}</td>
      <td>{post.title}</td>
      <td>{post.is_draft ? 'Yes' : 'No'}</td>
      <td>{displayDate(post.created_at)}</td>
      <td>{displayDate(post.updated_at)}</td>
    </tr>
    {/each}
  </table>
</div>
