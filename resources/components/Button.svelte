<script>
export let label;
export let color;
import { createEventDispatcher } from 'svelte';

const dispatch = createEventDispatcher();
function onClick() { dispatch('click') }
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
.button-blue {
  color: white;
  background-color: cornflowerblue;
}
.button-red {
  color: white;
  background-color: crimson;
}
.button-green {
  color: white;
  background-color: rgb(0, 209, 28);
}
.button-yellow {
  background-color: #ffd000;
}
</style>

<button
  class="button"
  class:button-blue={color === 'blue'}
  class:button-red={color === 'red'}
  class:button-green={color === 'green'}
  class:button-yellow={color === 'yellow'}
  on:click={onClick}
>
  {label}
</button>
