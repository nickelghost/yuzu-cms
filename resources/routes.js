import Home from './pages/Home.svelte';
import Posts from './pages/Posts.svelte';
import CreatePost from './pages/CreatePost.svelte';
import Post from './pages/Post.svelte';
import Pages from './pages/Pages.svelte';
import NotFound from './pages/NotFound.svelte';

const routes = {
  '/': Home,
  '/posts': Posts,
  '/posts/new': CreatePost,
  '/posts/:id': Post,
  '/pages': Pages,
  '*': NotFound,
};

export default routes;
