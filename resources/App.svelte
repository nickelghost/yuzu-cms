<script>
  import Router from 'svelte-spa-router';

  import Nav from './components/Nav.svelte';
  import LogIn from './pages/LogIn.svelte';

  import routes from './routes';
  import { jwt } from './stores';
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
  {#if $jwt}
  <Nav></Nav>
  {/if}
  <div class="container">
    {#if $jwt}
    <Router {routes}></Router>
    {:else}
    <LogIn></LogIn>
    {/if}
  </div>
</div>
