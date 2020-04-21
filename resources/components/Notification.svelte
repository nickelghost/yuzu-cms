<script>
  import { fade } from 'svelte/transition';

  export let message;
  export let color;
</script>

<style>
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
