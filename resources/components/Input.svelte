<script>
export let value;
</script>

<style>
.input {
  border: 1px solid #ececec;
  padding: 8px;
}
</style>

<input class="input" {...$$props} bind:value={value}>
