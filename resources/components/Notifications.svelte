<script>
  import Notification from './Notification.svelte';
  import { notifications } from '../stores';
</script>

<style>
  .notifications {
    display: flex;
    flex-direction: column;
    left: 50%; /* For horizontal centering */
    justify-content: center;
    position: fixed;
    top: 8px;
    transform: translateX(-50%); /* For horizontal centering */
  }
</style>

<div class="notifications">
  {#each $notifications as n}
  <Notification message="{n.message}" color="{n.color}"></Notification>
  {/each}
</div>
