# Yuzu

A super fast CMS built using Go and Svelte.

### Description

I created this project to be an alternative to WordPress for the "simple blog plus some pages" type of websites. It's still very much WIP, so you'll be better off using something like WP, Hugo or ButterCMS.

### Installation

The RDBMS used for the project is PostgreSQL. Pull this repository, create a `.env` file and configure it according to `.env.example`. Then you'll be able to run the application with the standard `go run .` command.

You'll need to build the assets for the admin panel with `npm run build`. If using the default theme, you also will need to build the assets for it - navigate to `themes/default` and run `npm run build`.
