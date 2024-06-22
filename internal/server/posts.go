package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/internal/parser"
	"github.com/pynezz/pynezz_com/templates"
)

type Metadata parser.Metadata

type PostsHandler struct{}

func (p PostsHandler) handleShowPosts(c echo.Context) error {
	id := c.Param("id")
	jwt := c.Get("jwt")
	response := fmt.Sprintln("here's post ", id)
	if jwt == nil {
		response += " (you're not logged in)"
	} else {
		// validate JWT
		response += " (you're logged in)"
	}

	return Render(c, http.StatusOK, templates.Show("posts", "here's post "+id))
}

// Fetch the last 5 posts or so
func (p PostsHandler) handleShowLastPosts(c echo.Context) error {
	limit := 5
	if c.QueryParam("limit") != "" {
		limit, _ = strconv.Atoi(c.QueryParam("limit"))
	}

	if limit < 25 {
		limit = 25
	}

	// posts, err := parser.GetLastPosts(limit)
	body := fmt.Sprintln("here's the last", limit, " posts. You can tweak the limit by adding a 'limit' query parameter to the URL.")

	return Render(c, http.StatusOK, templates.Show("post", body))
}

// Post is a struct that represents a post.
type Post struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Metadata Metadata `json:"metadata"`
	Content  string   `json:"content"`
	Path     string   `json:"path"` // Path to the markdown file - e.g. "/posts/2021-01-01-post.md"
}

// NewPostsHandler creates a new PostsHandler.
func newPostsHandler() *PostsHandler {
	return &PostsHandler{}
}

// // AddPost adds a post to the PostsHandler.
// func (ph *PostsHandler) AddPost(p Post) {
// 	ph.Posts = append(ph.Posts, p)
// }

// // GetPosts returns all the posts in the PostsHandler.
// func (ph *PostsHandler) GetPosts() []Post {
// 	return ph.Posts
// }

// // GetPostByID returns a post by its ID.
// func (ph *PostsHandler) GetPostByID(id int) *Post {
// 	for _, p := range ph.Posts {
// 		if p.ID == id {
// 			return &p
// 		}
// 	}
// 	return nil
// }

// // UpdatePost updates a post by its ID.
// func (ph *PostsHandler) UpdatePost(id int, p Post) {
// 	for i, post := range ph.Posts {
// 		if post.ID == id {
// 			ph.Posts[i] = p
// 		}
// 	}
// }

// // DeletePost deletes a post by its ID.
// func (ph *PostsHandler) DeletePost(id int) {
// 	for i, post := range ph.Posts {
// 		if post.ID == id {
// 			ph.Posts = append(ph.Posts[:i], ph.Posts[i+1:]...)
// 		}
// 	}
// }
