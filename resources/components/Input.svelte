<script>
export let value;
</script>

<style>
.input {
  border: 2px solid lightgrey;
  border-radius: 8px;
  padding: 8px;
}
</style>

<input class="input" {...$$props} bind:value={value}>
