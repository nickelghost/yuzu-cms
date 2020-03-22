<script>
  import { onMount } from 'svelte';

  import TopBar from '../components/TopBar.svelte';

  import { jwt } from '../stores';

  let pages = [];
  let posts = [];

  async function getPages() {
    const res = await fetch('/api/v1/pages', {
      headers: { Authorization: `Bearer ${$jwt}` },
    });
    pages = await res.json();
  }

  async function getPosts() {
    const res = await fetch('/api/v1/posts/titles', {
      headers: { Authorization: `Bearer ${$jwt}` },
    });
    posts = await res.json();
  }

  onMount(() => {
    getPages();
    getPosts();
  });
</script>

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
      <td><input /></td>
      <td><input type="checkbox" checked="{true}" /></td>
      <td>
        <select>
          {#each posts as post}
          <option>{post.title}</option>
          {/each}
        </select>
      </td>
      <td><button class="button button-primary">Add</button></td>
    </tr>
    {#each pages as page}
    <tr>
      <td>{page.index}</td>
      <td>{page.slug}</td>
      <td><input type="checkbox" disabled checked="{page.in_navigation}" /></td>
      <td>{page.post.title}</td>
      <td><button class="button">Delete</button></td>
    </tr>
    {/each}
  </table>
</div>
