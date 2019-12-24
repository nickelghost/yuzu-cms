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
.container {
  display: flex;
}
</style>

<div class="main">
  <Nav/>
  <div class="container is-fluid">
    <Router {routes}/>
  </div>
</div>
