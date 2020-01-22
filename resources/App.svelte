<script>
  import Router from 'svelte-spa-router';

  import Nav from './components/Nav.svelte';
  import routes from './routes';
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
  <Nav></Nav>
  <div class="container">
    <Router {routes}></Router>
  </div>
</div>
