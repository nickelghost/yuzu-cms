<script>
  import Router from 'svelte-spa-router';

  import Notifications from './components/Notifications.svelte';
  import Nav from './components/Nav.svelte';
  import LogIn from './pages/LogIn.svelte';

  import routes from './routes';
  import { jwt } from './stores';
</script>

<style>
  .main {
    color: #474747;
    display: flex;
    font-family: 'Fira Sans', sans-serif;
    min-height: 100vh;
  }

  .container {
    display: flex;
    flex-direction: column;
    flex-grow: 1;
  }

  .nav-placeholder {
    flex-shrink: 0;
    width: 200px;
  }
</style>

<div class="main">
  <Notifications></Notifications>
  {#if $jwt}
  <Nav></Nav>
  <div class="nav-placeholder"></div>
  {/if}
  <div class="container">
    {#if $jwt}
    <Router {routes}></Router>
    {:else}
    <LogIn></LogIn>
    {/if}
  </div>
</div>
