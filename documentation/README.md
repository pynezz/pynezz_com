---
Title: "How to write a CMS and site generator in Go from scratch"
Tags: ["go", "projects", "portfolio", "pynezz.dev"]
Created: 20.06.2024
---


**🏗️ Current status:**  Building the project

Still manually writing the html for the [changelog]. It's a hassle.

[changelog]: https://pynezz.dev/changelog.html

---

# How to write a CMS and site generator in Go from scratch

- [How to write a CMS and site generator in Go from scratch](#how-to-write-a-cms-and-site-generator-in-go-from-scratch)
  - [Planning and introduction](#planning-and-introduction)
    - [Introduction](#introduction)
      - [Dual purpose project](#dual-purpose-project)
    - [Requirements](#requirements)
      - [Markdown](#markdown)
      - [Authentication](#authentication)
      - [Theming and styles](#theming-and-styles)
      - [CLI build tool](#cli-build-tool)
      - [Web server](#web-server)
  - [The implementation, thoughts, and ideas](#the-implementation-thoughts-and-ideas)
    - [Markdown to HTML](#markdown-to-html)
      - [Content](#content)
      - [Metadata](#metadata)
    - [Implementing authentication](#implementing-authentication)

## Planning and introduction

### Introduction

This is the beginning of my journey in building a content management system (CMS), dynamic site generator, from scratch, in Go.

I recently graduated with bachelors of science in Cyber Security, so why build a CMS in Go from scratch? Well, our final semester project (bachelor thesis) was a project where we built a Security Information and Event Management (SIEM) system for log analysis, also mainly built in Go. Over the course of the project, I got more fond of the language, and want to keep honing my skills in it. Now however, instead of building more security tools, I want to do something with web tech instead.

Although my current site is up and running using Hugo, I have decided upon building my own framework.

#### Dual purpose project

The project will serve a dual purpose in relation to my recently acquired bachelors degree in cyber security.

First, it will be a fun side project that will teach me more about web, server management, development, Go, and so on.

Second, it will be a place for me to post content such as my projects, guides, and notes.

My GitHub username `pynezz` have been used as the domain name, and both `pynezz.com` and `pynezz.dev` serves the page where the result of this project will be served.

### Requirements

The project have a set of requirements, which is listed below. All the required features will be compiled into a single Go binary with a command-line interface that makes them all accessible.

#### Markdown

Progress on this part can be found [here](./parsing)

I'm writing my all my main digital notes in Obsidian due to the simplicity of writing markdown, and have done for several years now. This makes the first requirement easy, the framework would need to work with markdown.

It should come to no surprise that my framework will be inspired by [Hugo](https://gohugo.io).

More about this [here](#markdown-to-html)

#### Authentication

The next requirement is sort of obvious, but we'll need to define it anyways. The CMS should not allow other people on the internet to upload posts. Therefore, an authentication and authorization scheme is needed. This is some of the more fun parts, so I'll do this one from scratch as well.

I'm not going to roll my own crypto, just implement the authentication myself.

There's however an obvious second solution: Just have a directory with markdown files and sub directories, and routes based on the file structure .

More about this [here](#implementing-authentication)

#### Theming and styles

I don't want the page to look like it were abandoned back in 2008, so some separate front-end or integrated styling will need to be implemented. This will have to live as a separate component, implementation is undecided.

#### CLI build tool

A command-line interface (CLI) is needed to do build the site. It should feature help commands to make the different functionalities readily available.

Nice to have:  It would be desirable if no intermediate build step would be needed when adding a new markdown page, as long as it's in an existing path.

#### Web server

The web server is essential. It will live in a cloud VPS for now. I'll probably proxy pass it via the Openresty reverse proxy already set up for the domains. The web server will handle the essentials like routing, authentication and authorization, logging, and server side rendering(SSR).

I used [Go Fiber](https://gofiber.io/) in our bachelors project, but I'm considering [Echo](https://github.com/labstack/echo) for this one, not only because I've seen it be recommended by multiple people, but also because it's fun to explore different ways of doing things.

---

## The implementation, thoughts, and ideas

### Markdown to HTML

#### Content

Markdown comes in several flavors, which means I'll be able to create my own flavor for this project. However, doing so complicates the process of parsing the markdown and converting it to HTML that the browser can understand. I'm considered using [pandoc](https://pandoc.org) due to its flexibility and my own familiarity with it, which means an additional step will be needed to insert my own styles if I want to do it at this level.

I'm probably not going to write the whole parsing functionality, and rather use some open source tooling as a foundation.

#### Metadata

Markdown supports yaml formatted metadata. Parsing yaml is quite straight forward with [gopkg yaml v3](https://pkg.go.dev/gopkg.in/yaml.v3). The metadata will contain data such as tags, category, title, and id.

### Implementing authentication

> ⚠️ Undecided

- public key cryptography is probably the safest
- password based auth with argon2 hashing and sqlite and JWTs is probably easiest due to familiarity with those technologies in particular
- protocols such as [SRP](https://en.wikipedia.org/wiki/Secure_Remote_Password_protocol) would be the most interesting to implement, but sadly I have not much experience with [PAKE](https://en.wikipedia.org/wiki/Password-authenticated_key_agreement) protocols - something worth looking into in the future of the project

If I go for a directory with markdown files located on the server, authentication may be unnecessary.

---

## Sources

-  [writing github docs using markdown](https://docs.github.com/en/contributing/writing-for-github-docs/using-markdown-and-liquid-in-github-docs#callout-tags)
- [markdownlint rules](https://xiangxing98.github.io/Markdownlint_Rules.html#md001)
