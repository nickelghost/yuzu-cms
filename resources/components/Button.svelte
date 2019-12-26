<script>
export let label;
import { createEventDispatcher } from 'svelte';

const dispatch = createEventDispatcher();
function onClick() {
  dispatch('click')
}
</script>

<style>
.button {
  background-color: white;
  border: 1px solid lightgrey;
  border-radius: 4px;
  color: #222;
  cursor: pointer;
  padding: 12px 24px;
}
</style>

<button class="button" on:click={onClick}>{label}</button>
