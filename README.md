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
    - [ ] edit post (metadata, content)
- [x] serve 

### CMS

The `cms` module is able to read markdown files from a directory, parse them, generate a slug based on the title, insert tailwind styles, and insert them into the SQLite database. The `serve` module will then display the contents on the website. 

**metadata**

The metadata fields will be used to define certain properties of the content, like the title, date, and tags/categories.

#### Planned features

- [x] Read and parse markdown files from a directory
- [x] Push the content to the frontend, with correct paths, and metadata
- [ ] Display and fetch posts by tag
- [ ] Compress images for faster loading times
- [ ] Configuration file

#### Known issues

- [ ] Filter by tags not working [error 404 due to incorrect database query, wrong model]

### Frontend

The frontend will be built with (?).
The goal is to create a 'blazingly' fast website, where the focus will be on performance, readability, and responsiveness.

Inspiration will certainly be taken from [Hugo](https://gohugo.io/) in terms of its incredible speed and simplicity.

### Backend

The backend is written in pure Go.

#### Known issues

- [ ] Routing not working properly [relative path issue]
- [ ] Tags should contain a field containing the count of posts with the tag [need refactoring] 

---

## Build

### Requirements

- Go
- Make
- nodejs and npm

### 

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

#### Windows

```bash
make windows
```

## Usage

Run the project:

```bash
./pynezz_com_[os_arch](.exe)
```
