# Pingu Web

My Go approach for a frontend facing web application, instead of using Node.js

Go renders your HTML templates and serves them to the client. It also serves your static files, like CSS and JavaScript. Building assets (js and css) is done with Vite.

On top of that, we use Reflex to automatically rebuild and restart the server when changes are detected (HMR).

**This is very experimental and not recommended for production use.**

## Involved technologies
- Go Lang
- Fiber
- Vite
- Tailwind CSS
- Docker (optional)
- Reflex (optional)

## Features
Under the hood Pingu Web uses the Fiber web framework for Go. This is simply my own starterkit for a web application, 
to avoid building the same setup over and over again for small services. This boilerplat is not meant to be a full-fledged 
web application, but rather a starting point for web facing web service which needs to serve HTML, CSS and JS.

The following features have been set up:
- Routing
- HTML templating with Django-like syntax
- App config / environment variables from .env file
- Static file servings
- Vite for building assets (js and css)
- Tailwind CSS
- Docker setup
- Reflex for automatic server restarts
- Cookie handling / secure cookies
- Session handling
- Cache handling
- Logging with Slog
- Cache
- Panic recovery (return custom 500 html page and recover)
- 404 Page if route not found
