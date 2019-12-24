<style>
.nav {
  background-color: #3a243b;
  color: white;
  display: flex;
  flex-direction: column;
  padding: 12px;
  width: 260px;
}
.menu-label {
  color: white;
}
.menu-list a {
  color: white;
}
.menu-list a:hover {
  color: black;
}
</style>

<aside class="nav">
  <div class="menu">
    <p class="menu-label">
      Yuzu Admin
    </p>
    <ul class="menu-list">
      <li><a href="/admin/#/">Home</a></li>
      <li><a href="/admin/#/pages">Pages</a></li>
    </ul>
    <p class="menu-label">
      Posts
    </p>
    <ul class="menu-list">
      <li><a href="/admin/#/posts/new">New</a></li>
      <li><a href="/admin/#/posts">List</a></li>
    </ul>
    <p class="menu-label">
      Exit
    </p>
    <ul class="menu-list">
      <li><a href="/">Homepage</a></li>
      <li><a href="#">Log out</a></li>
    </ul>
  </div>
</aside>
