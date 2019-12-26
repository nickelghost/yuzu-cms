<script>
import Nav from './components/Nav.svelte';
import Router from 'svelte-spa-router';
import routes from './routes';
</script>

<style>
.main {
  color: #333;
  display: flex;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
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
