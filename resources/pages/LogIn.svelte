<script>
  import { jwt } from '../stores';

  import Notification from '../components/Notification.svelte';

  let user = '';
  let password = '';

  let notificationMessage = '';
  let notificationColor = '';

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
      notificationMessage = 'Logged in';
      notificationColor = 'green';
      localStorage.setItem('jwt', json.token);
      jwt.set(json.token);
    } else {
      notificationMessage = 'Could not log in';
      notificationColor = 'red';
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

  form {
    display: flex;
    flex-direction: column;
    padding: 24px;
  }

  .field {
    display: flex;
    flex-direction: column;
    margin-bottom: 24px;
  }

  .label {
    font-size: 1.125rem;
    margin-bottom: 8px;
    text-transform: uppercase;
  }
</style>

<Notification
  bind:message="{notificationMessage}"
  color="{notificationColor}"
></Notification>

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
