<script>
  export let isOpen;
</script>

<style>
  .modal {
    align-items: center;
    bottom: 0;
    display: none;
    justify-content: center;
    left: 0;
    position: absolute;
    right: 0;
    top: 0;
    z-index: 500;
  }

  .modal-open {
    display: flex;
  }

  .modal-background {
    background: rgba(0, 0, 0, 0.6);
    bottom: 0;
    left: 0;
    position: absolute;
    right: 0;
    top: 0;
    z-index: 500;
  }

  .modal-inner {
    background-color: white;
    border-radius: 8px;
    min-width: 480px;
    z-index: 1000;
  }
</style>

<div class="modal" class:modal-open="{isOpen}">
  <div class="modal-background" on:click="{() => { isOpen = false; }}"></div>
  <div class="modal-inner">
    <slot></slot>
  </div>
</div>
