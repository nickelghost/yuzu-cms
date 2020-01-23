<script>
  import { fade } from 'svelte/transition';

  export let message;
  export let color;
  export let timeout = 2000;

  let visible = false;

  $: if (message) {
    visible = true;
    const timeoutID = setTimeout(() => {
      message = '';
      visible = false;
      clearTimeout(timeoutID);
    }, timeout);
  }
</script>

<style>
  .notification-container {
    display: flex;
    left: 0;
    justify-content: center;
    position: fixed;
    right: 0;
    top: 8px;
    z-index: 1000;
  }
  .notification {
    align-items: center;
    background: rgba(102, 102, 102, 0.9);
    color: white;
    display: flex;
    justify-content: center;
    padding: 16px;
    width: 480px;
  }
  .notification-green {
    background: rgba(57, 168, 57, 0.9);
  }
  .notification-red {
    background: rgba(197, 39, 39, 0.9);
  }
</style>

{#if visible}
<div class="notification-container">
  <div
    class="notification"
    class:notification-green="{color === 'green'}"
    class:notification-red="{color === 'red'}"
    transition:fade
  >
    {message}
  </div>
</div>
{/if}
