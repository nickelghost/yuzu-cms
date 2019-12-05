import Home from './pages/Home.svelte';
import Posts from './pages/Posts.svelte';
import NotFound from './pages/NotFound.svelte';

const routes = {
  '/': Home,
  '/posts': Posts,
  '*': NotFound,
};

export default routes;
