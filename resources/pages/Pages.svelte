<script>
  import { onMount } from 'svelte';

  import TopBar from '../components/TopBar.svelte';
  import Modal from '../components/Modal.svelte';
  import PageForm from './Pages/PageForm.svelte';

  import notify from '../helpers/notify';
  import { jwt } from '../stores';

  let pages = [];
  let posts = [];

  let isNewModalOpen = false;
  let isEditModalOpen = false;

  // eslint-disable-next-line prefer-const
  let newPageForm = {};
  let editPageForm = {};

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

  async function addPage(e) {
    e.preventDefault();
    // eslint-disable-next-line operator-linebreak
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
        ...newPageForm,
        index: highestIndex + 1,
      }),
    });
    if (res.ok) {
      notify('Page added', 'green');
      getPages();
      isNewModalOpen = false;
    } else {
      notify('Could not add the page', 'red');
    }
  }

  async function updatePage(e) {
    e.preventDefault();
    const req = {
      ...editPageForm,
      position_change: 0,
    };
    const res = await fetch(`/api/v1/pages/${req.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${$jwt}`,
      },
      body: JSON.stringify(editPageForm),
    });
    if (res.ok) {
      getPages();
      notify('Page updated', 'green');
      isEditModalOpen = false;
    } else {
      notify('Could not update the page', 'red');
    }
  }

  async function deletePage(id) {
    const res = await fetch(`/api/v1/pages/${id}`, {
      method: 'DELETE',
      headers: { Authorization: `Bearer ${$jwt}` },
    });
    if (res.ok) {
      notify('Page deleted', 'green');
      getPages();
    } else {
      notify('Could not delete the page', 'red');
    }
  }

  async function changePageIndex(page, change) {
    const req = {
      ...page,
      position_change: change,
    };
    const res = await fetch(`/api/v1/pages/${page.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${$jwt}`,
      },
      body: JSON.stringify(req),
    });
    if (res.ok) {
      notify('Page position updated', 'green');
      getPages();
    } else {
      notify('Could not update page position', 'red');
    }
  }
</script>

<style>
  th {
    text-align: left;
  }

  .table {
    font-size: 1.125rem;
    width: 100%;
  }

  .table th {
    font-size: 0.875rem;
    font-weight: 500;
    text-transform: uppercase;
  }

  .slug {
    font-family: 'Fira Code', monospace;
  }
</style>

<TopBar title="Pages">
  <button
    class="button button-primary"
    on:click="{() => { isNewModalOpen = true; }}"
  >
    New
  </button>
</TopBar>

<Modal bind:isOpen="{isNewModalOpen}">
  <form class="form" on:submit="{addPage}">
    <PageForm
      posts="{posts}"
      bind:post_id="{newPageForm.post_id}"
      bind:slug="{newPageForm.slug}"
      bind:in_navigation="{newPageForm.in_navigation}"
    ></PageForm>
    <div class="field field-horizontal">
      <button
        class="button"
        type="button"
        on:click="{() => { isNewModalOpen = false; }}"
      >
        Cancel
      </button>
      <div class="flex-spacer"></div>
      <button class="button button-primary" type="submit">
        Add
      </button>
    </div>
  </form>
</Modal>

<Modal bind:isOpen="{isEditModalOpen}">
  <form class="form" on:submit="{updatePage}">
    <PageForm
      posts="{posts}"
      bind:post_id="{editPageForm.post_id}"
      bind:slug="{editPageForm.slug}"
      bind:in_navigation="{editPageForm.in_navigation}"
    ></PageForm>
    <div class="field field-horizontal">
      <button
        class="button"
        type="button"
        on:click="{() => { isEditModalOpen = false; }}"
      >
        Cancel
      </button>
      <div class="flex-spacer"></div>
      <button class="button button-primary" type="submit">
        Change
      </button>
    </div>
  </form>
</Modal>

<div class="content">
  <table class="table">
    <tr>
      <th>Position</th>
      <th>Post</th>
      <th>Slug</th>
      <th>In navigation</th>
      <th></th>
    </tr>
    {#each pages as page}
    <tr>
      <td>
        <button
          class="button button-small"
          on:click="{() => changePageIndex(page, 1)}"
        >
          +
        </button>
        {page.index}
        <button
          class="button button-small"
          on:click="{() => changePageIndex(page, -1)}"
        >
          -
        </button>
      </td>
      <td>{page.post.title}</td>
      <td class="slug">{page.slug}</td>
      <td>{page.in_navigation ? 'Yes' : 'No'}</td>
      <td>
        <button
          class="button"
          on:click="{() => { editPageForm = { ...page }; isEditModalOpen = true; }}"
        >
          Edit
        </button>
        <button
          class="button button-danger"
          on:click="{() => deletePage(page.id)}"
        >
          Delete
        </button>
      </td>
    </tr>
    {/each}
  </table>
</div>
