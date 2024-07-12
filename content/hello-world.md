---
path: content/hello-world.md
---

# Hello world

If you're reading this, the site is up and running!
This means that:

- The server is running (Go)
- The reverse proxy is working ([OpenResty](https://openresty.org/))
- The client side code is working ([Go-Templ](https://templ.guide))
- The markdown to HTML conversion is working (built from scratch)
- The database is working (SQLite)
- The authentication is working ([JWTs](https://jwt.io/) or [PASETO](https://paseto.io), implemented myself)
- The domain is working (pynezz.dev)
- SSL is working ([Let's Encrypt](https://letsencrypt.org))

## What is pynezz.dev?

This site is a personal project by me
([pynezz](https://github.com/pynezz))
where I'm building a content management
system (CMS) in [Go](https://go.dev/).

## $whoami

I've newly graduated with a bachelor's degree in cyber security,
and have a passion for programming and development.

This serves as a portfolio of my work,
as well as thoughts, and ideas.

## What's next?

Aside from searching for a job,
I'm developing this site to be a blog and portfolio.

As time passes, this site will gradually be updated
with new content in the changelog section.

### Aim

This site starts off as a simple static site, I'm writing this "by hand".

The goal is to be able to upload markdown files to the server and have
them automatically converted to HTML and displayed on the site,
with correct formatting and styling.

I'm writing the server and client side code from scratch
because it's fun to re-invent the wheel.

### Technologies used

- Reverse proxy with OpenResty (nginx) with some Lua scripting
- Go CMS server with SQLite database
- Authentication (undecided): maybe JWT (current), or maybe [PASETO](https://paseto.io),
    at least I'll use Argon2 if I need to store credentials...
- Frontend: Go-Templ
- Markdown to HTML conversion: built from scratch
- Hosting: DigitalOcean, but any VPS will do

### Read more

Read more about the project on the
[GitHub repository](https://github.com/pynezz/pynezz_com).
