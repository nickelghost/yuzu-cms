<script>
  import { onMount } from 'svelte';

  import TopBar from '../components/TopBar.svelte';

  import { jwt } from '../stores';

  let pages = [];
  let posts = [];

  const form = {
    slug: '',
    in_navigation: true,
    post_id: 0,
  };

  async function getPages() {
    const res = await fetch('/api/v1/pages', {
      headers: { Authorization: `Bearer ${$jwt}` },
    });
    if (res.ok) {
      pages = await res.json();
    }
  }

  async function getPosts() {
    const res = await fetch('/api/v1/posts/titles', {
      headers: { Authorization: `Bearer ${$jwt}` },
    });
    if (res.ok) {
      posts = await res.json();
    }
  }

  onMount(() => {
    getPages();
    getPosts();
  });

  async function addPage() {
    // eslint-disable-next-line
    const highestIndex =
      pages.length === 0
        ? 0
        : pages.reduce((p1, p2) => (p1.index > p2.index ? p1 : p2)).index;
    const res = await fetch('/api/v1/pages', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${$jwt}`,
      },
      body: JSON.stringify({
        ...form,
        index: highestIndex + 1,
      }),
    });
    if (res.ok) {
      getPages();
    }
  }

  async function deletePage(id) {
    const res = await fetch(`/api/v1/pages/${id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${$jwt}`,
      },
    });
    if (res.ok) {
      getPages();
    }
  }
</script>

<style>
  th {
    text-align: left;
  }
</style>

<TopBar title="Pages"></TopBar>

<div class="content">
  <table>
    <tr>
      <th>Position</th>
      <th>Slug</th>
      <th>In navigation</th>
      <th>Post</th>
      <th></th>
    </tr>
    <tr>
      <td></td>
      <td><input type="text" bind:value="{form.slug}" /></td>
      <td><input type="checkbox" bind:checked="{form.in_navigation}" /></td>
      <td>
        <select bind:value="{form.post_id}">
          <option value="{0}" disabled></option>
          {#each posts as post}
          <option value="{post.id}">{post.title}</option>
          {/each}
        </select>
      </td>
      <td>
        <button class="button button-primary" on:click="{addPage}">
          Add
        </button>
      </td>
    </tr>
    {#each pages as page}
    <tr>
      <td>{page.index}</td>
      <td>{page.slug}</td>
      <td><input type="checkbox" disabled checked="{page.in_navigation}" /></td>
      <td>{page.post.title}</td>
      <td>
        <button class="button" on:click="{() => deletePage(page.id)}">
          Delete
        </button>
      </td>
    </tr>
    {/each}
  </table>
</div>
