<script>
  export let posts;
  export let post_id;
  export let slug;
  export let in_navigation;
</script>

<label class="field">
  <span class="label">Post</span>
  <select class="input" bind:value="{post_id}">
    <option value="{0}" disabled></option>
    {#each posts as post}
    <option value="{post.id}">{post.title}</option>
    {/each}
  </select>
</label>
<label class="field">
  <span class="label">Slug</span>
  <input class="input" bind:value="{slug}" />
</label>
<label class="field">
  <span class="label">In navigation</span>
  <input class="input" type="checkbox" bind:checked="{in_navigation}" />
</label>
