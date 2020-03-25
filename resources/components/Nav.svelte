<script>
  import { link, location, querystring } from 'svelte-spa-router';
  import { parse as parseQS } from 'qs';
  import { jwt } from '../stores';

  function logOut() {
    localStorage.removeItem('jwt');
    jwt.set('');
  }
</script>

<style>
  .nav {
    background: linear-gradient(90deg, #515093 0%, rgb(90, 111, 180) 100%);
    background-color: #515093;
    color: white;
    display: flex;
    flex-direction: column;
    flex-shrink: 0;
    font-family: Quicksand, sans-serif;
    width: 200px;
  }

  .heading {
    font-weight: normal;
    margin: 24px 0;
    text-align: center;
  }

  .links {
    display: flex;
    flex-direction: column;
    flex-grow: 1;
  }

  .link {
    cursor: pointer;
    font-size: 1.125rem;
    display: inline-block;
    padding: 11px 24px;
  }

  .link:hover {
    background: rgba(255, 255, 255, 20%);
  }

  .link-selected {
    background: rgba(255, 255, 255, 10%);
  }

  .link-sub {
    border-left: 1px solid white;
    padding-left: 8px;
  }

  .link-bottom {
    margin-top: auto;
  }
</style>

<aside class="nav">
  <h1 class="heading">Yuzu</h1>
  <nav class="links">
    <!-- <a class="link" class:link-selected="{$location === '/'}" use:link href="/">
      Home
    </a> -->
    <a
      class="link"
      class:link-selected="{$location === '/posts' && parseQS($querystring).draft !== 'true'}"
      use:link
      href="/posts"
    >
      Posts
    </a>
    <a
      class="link"
      class:link-selected="{$location === '/posts/new'}"
      use:link
      href="/posts/new"
    >
      <div class="link-sub">
        New
      </div>
    </a>
    <a
      class="link"
      class:link-selected="{$location === '/posts' && parseQS($querystring).draft === 'true'}"
      use:link
      href="/posts?draft=true"
    >
      <div class="link-sub">
        Drafts
      </div>
    </a>
    <a
      class="link"
      class:link-selected="{$location === '/pages'}"
      use:link
      href="/pages"
    >
      Pages
    </a>
    <button class="button-reset link" on:click="{logOut}">
      Log out
    </button>
    <a class="link link-bottom" href="/">
      Website
    </a>
  </nav>
</aside>
