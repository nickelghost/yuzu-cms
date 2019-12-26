<style>
.nav {
  background: rgb(78,0,121);
  background: linear-gradient(90deg, rgba(78,0,121,1) 0%, rgba(156,0,200,1) 100%);
  box-shadow: 0 0 10px grey;
  box-sizing: border-box;
  color: white;
  display: flex;
  flex-direction: column;
  width: 260px;
}
.heading {
  text-align: center;
}
.nav-button {
  font-size: 1.125rem;
  display: inline-block;
  padding: 12px 24px;
}
.nav-button-underline {
  background-color: white;
  height: 2px;
  margin-top: 4px;
  visibility: hidden;
  width: 100%;
}
.nav-link:hover .nav-button-underline {
  visibility: visible;
}
.spacer {
  flex-grow: 1;
}
a {
  color: inherit;
  text-decoration: none;
}
</style>

<aside class="nav">
  <h1 class="heading">Yuzu</h1>
  <a class="nav-link" href="/admin/#/">
    <div class="nav-button">
      Home
      <div class="nav-button-underline" />
    </div>
  </a>
  <a class="nav-link" href="/admin/#/posts">
    <div class="nav-button">
      Posts
      <div class="nav-button-underline" />
    </div>
  </a>
  <a class="nav-link" href="/admin/#/pages">
    <div class="nav-button">
      Pages
      <div class="nav-button-underline" />
    </div>
  </a>
  <div class="spacer"></div>
  <a class="nav-link" href="/">
    <div class="nav-button">
      Website
      <div class="nav-button-underline" />
    </div>
  </a>
</aside>
