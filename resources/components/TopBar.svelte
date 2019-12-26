<script>
export let title;
</script>

<style>
.top-bar {
  align-items: center;
  background: whitesmoke;
  border-bottom: 1px solid lightgrey;
  display: flex;
  height: 56px;
  padding: 0 12px;
}
.title {
  margin: 0;
}
.spacer {
  flex-grow: 1;
}
.right {
  display: flex;
}
</style>

<div class="top-bar">
  <h2 class="title">{title}</h2>
  <div class="spacer"></div>
  <div class="right">
    <slot></slot>
  </div>
</div>
