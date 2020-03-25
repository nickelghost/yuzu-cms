import Redirect from './components/Redirect.svelte';
import Posts from './pages/Posts.svelte';
import CreatePost from './pages/CreatePost.svelte';
import Post from './pages/Post.svelte';
import Pages from './pages/Pages.svelte';
import NotFound from './pages/NotFound.svelte';

const routes = {
  '/': Redirect,
  '/posts': Posts,
  '/posts/new': CreatePost,
  '/posts/:id': Post,
  '/pages': Pages,
  '*': NotFound,
};

export default routes;
