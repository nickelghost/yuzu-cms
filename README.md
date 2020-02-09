# Yuzu

A super fast CMS built using Go and Svelte.

I created this project to be an alternative to WordPress for the "simple blog plus some pages" type of websites. It's still very much WIP, so you'll be better off using something like WP, Hugo or ButterCMS.

## Installation

The RDBMS used for the project is PostgreSQL. Pull this repository, create a `.env` file and configure it according to `.env.example`. Then you'll be able to run the application with the standard `go run .` command.

You'll need to build the assets for the admin panel with `npm run build`. If using the default theme, you also will need to build the assets for it - navigate to `themes/default` and run `npm run build`.

## Development

If you want to use Webpack's HMR, you need to start the server with `npm start` and set `APP_WEBPACK_FORWARD` to `true` for the backend. This will proxy the webpack dev server at `/admin`.
