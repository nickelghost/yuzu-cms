<script>
import Nav from './components/Nav.svelte';
import Router from 'svelte-spa-router';
import routes from './routes';
</script>

<style>
.main {
  display: flex;
  min-height: 100vh;
}
.container-wrapper {
  background-color: lightgray;
  display: flex;
  justify-content: space-around;
  flex-grow: 1;
}
.container {
  background-color: white;
  display: flex;
  flex-grow: 1;
  max-width: 970px;
  padding: 0 12px;
}
</style>

<div class="main">
  <Nav/>
  <div class="container-wrapper">
    <div class="container">
      <Router {routes}/>
    </div>
  </div>
</div>
