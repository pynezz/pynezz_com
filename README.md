# pynezz_com

Simple website for ~~pynezz.com~~ pynezz.dev built with Go and [Go-Templ](https://templ.guide)

## Description

### Goal

The goal of this project is to create a simple website/static site generator and content management system (CMS), but from scratch. The website will be built with performance in mind and should be launchable from a single binary.

The site will utilize no external resources, meaning it will contain no tracking, no cookies, and no external assets such as fonts and icons.
Everything will be living on the server.

#### Progress

- [x] Single binary compilation
- [ ] cms cli [partially done]
    - [x] read and parse markdown files from directory
    - [ ] unpuplish posts
    - [ ] edit post (metadata, [x] content)
- [x] serve

### CMS

The `cms` module is able to read markdown files from a directory, parse them, generate a slug based on the title, insert tailwind styles, and insert them into the SQLite database. The `serve` module will then display the contents on the website.

**metadata**

The metadata fields will be used to define certain properties of the content, like the title, date, and tags/categories.

#### Planned features

- [x] Read and parse markdown files from a directory
- [x] Push the content to the frontend, with correct paths, and metadata
- [x] Display and fetch posts by tag
- [ ] Compress images for faster loading times
- [ ] Configuration file

#### Known issues

- [ ] Syntax highlighting not working properly
- [ ] Parsing code blocks still not working properly in some cases

### Frontend

The frontend is templated with [go-templ](https://github.com/a-h/templ) and styled with [Tailwind CSS](https://tailwindcss.com/).
The goal is to create a 'blazingly' fast website, where the focus will be on performance, readability, and responsiveness.

Inspiration will certainly be taken from [Hugo](https://gohugo.io/) in terms of its incredible speed and simplicity.

### Backend

The backend is written in pure Go.

#### Known issues

- [ ] Tags should contain a field containing the count of posts with the tag [need refactoring]

#### Fixes

##### [64eabb2](/commit/64eabb2)

- [x] Routing not working properly [relative path issue]

---

## Build

### Requirements

- Go
- nodejs and npm (for building tailwindcss)

### Optional

- Make  (for building the project)
- Zig compiler (for building to compatible version of glibc, as declared in the Makefile)

### Installation

**fetch the project source:**

```bash
go get github.com/pynezz/pynezz_com
```

**and then build the project:**

#### Linux

```bash
make linux
```

##### Without Make

NB: *cgo is required for sqlite driver.*

```bash
go get -u ./... # get all dependencies
npm install # install tailwindcss and dependencies

GOOS=linux GOARCH=amd64 CGO_ENABLED=1
CC="zig cc -target x86_64-linux-gnu.2.31.0" CXX="zig c++ -target x86_64-linux-gnu.2.31.0"

templ generate && npm build:css && go build -o pynezz_com_linux_amd64
```

NB: *Not tested. Just a rough idea of how to build the project without Make.*

#### Windows

```bash
make windows
```

## Usage

Run the project:

```bash
./pynezz_com_[os_arch](.exe)
```
