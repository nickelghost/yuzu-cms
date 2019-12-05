import Home from './pages/Home.svelte';
import Posts from './pages/Posts.svelte';
import Post from './pages/Post.svelte';
import NotFound from './pages/NotFound.svelte';

const routes = {
  '/': Home,
  '/posts': Posts,
  '/posts/:id': Post,
  '*': NotFound,
};

export default routes;
