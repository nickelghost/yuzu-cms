<script>
import Nav from './components/Nav.svelte';
import Router from 'svelte-spa-router';
import routes from './routes';
</script>

<style>
:global(a) {
  color: inherit;
  text-decoration: none;
}
:global(.content) {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  padding: 24px;
}
:global(.top-bar-spacer) {
  width: 8px;
}
:global(.button) {
  background-color: white;
  border: 1px solid #ececec;
  cursor: pointer;
  padding: 12px 24px;
}
:global(.input) {
  border: 1px solid #ececec;
  padding: 8px;
}
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
