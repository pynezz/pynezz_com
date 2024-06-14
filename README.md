# website for pynezz_com

Simple website for pynezz.com built with Go and ...?

## Description

### Goal

The goal of this project is to create a simple website for pynezz.com, but from scratch. Including a light CMS for uploading new content. The website will be built with performance in mind and should be launchable from a single binary.

The site will utilize no external resources, meaning it will contain no tracking, no cookies, and no external assets such as fonts and icons.
Everything will be living on the server.

### CMS

The plan is to be able to write markdown files, which will be parsed and displayed on the website. The metadata fields will be used to define certain properties of the content, like the title, date, and tags/categories.

#### Planned features

- [ ] Read and parse markdown files from a directory
- [ ] Push the content to the frontend, with correct paths, and metadata
- [ ] Compress images for faster loading times

### Frontend

The frontend will be built with (?).
The goal is to create a 'blazingly' fast website, where the focus will be on performance, readability, and responsiveness.

Inspiration will certainly be taken from [Hugo](https://gohugo.io/) in terms of its incredible speed and simplicity.

### Backend

The backend will be built with Go.

---

## Build

### Requirements

- Go
- Make
- Probably nodejs and npm/pnpm

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
