<script>
  import notify from '../helpers/notify';

  import { jwt } from '../stores';

  let user = '';
  let password = '';

  async function logIn(e) {
    e.preventDefault();
    const req = { user, password };
    const res = await fetch('/api/v1/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(req),
    });
    const json = await res.json();
    if (res.ok) {
      notify('Logged in', 'green');
      localStorage.setItem('jwt', json.token);
      jwt.set(json.token);
    } else {
      notify('Could not log in', 'red');
    }
  }
</script>

<style>
  .main {
    align-items: center;
    background: linear-gradient(90deg, #515093 0%, rgb(90, 111, 180) 100%);
    display: flex;
    height: 100vh;
    justify-content: center;
  }

  .container {
    background: white;
    width: 480px;
  }

  .title {
    font-family: Quicksand, sans-serif;
    font-size: 3rem;
    font-weight: normal;
    margin: 24px 0;
    text-align: center;
  }
</style>

<main class="main">
  <div class="container">
    <h1 class="title">Yuzu</h1>
    <form class="form" on:submit="{logIn}">
      <label class="field">
        <span class="label">User</span>
        <input class="input" bind:value="{user}" />
      </label>
      <label class="field">
        <span class="label">Password</span>
        <input type="password" class="input" bind:value="{password}" />
      </label>
      <button type="submit" class="button button-primary">Log in</button>
    </form>
  </div>
</main>
