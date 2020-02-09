<script>
  export let title;
</script>

<style>
  .top-bar {
    align-items: center;
    background: #fafafa;
    display: flex;
    height: 56px;
    padding: 0 24px;
  }

  .title {
    font-size: 1.5rem;
    font-weight: normal;
    margin: 0;
  }

  .right {
    display: flex;
    margin-left: auto;
  }
</style>

<div class="top-bar">
  <h2 class="title">{title}</h2>
  <div class="right">
    <slot></slot>
  </div>
</div>
