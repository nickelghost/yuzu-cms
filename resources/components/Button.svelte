<script>
import { createEventDispatcher } from 'svelte';
export let label;
const dispatch = createEventDispatcher();
function onClick() { dispatch('click') }
</script>

<style>
.button {
  background-color: white;
  border: 1px solid #ececec;
  cursor: pointer;
  padding: 12px 24px;
}
</style>

<button class="button" on:click={onClick}>
  {label}
</button>
