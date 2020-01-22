<script>
  import Router from 'svelte-spa-router';
  import routes from './routes';

  import Nav from './components/Nav.svelte';
</script>

<style>
.main {
  color: #474747;
  display: flex;
  font-family: Quicksand, sans-serif;
  min-height: 100vh;
}
.container {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}
</style>

<div class="main">
  <Nav/>
  <div class="container">
    <Router {routes}/>
  </div>
</div>
