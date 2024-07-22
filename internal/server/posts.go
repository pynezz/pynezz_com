package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/internal/parser"
	"github.com/pynezz/pynezz_com/internal/server/middleware"
	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
	"github.com/pynezz/pynezz_com/templates"
	"github.com/pynezz/pynezzentials/ansi"
)

type Metadata parser.Metadata

type Stats struct {
	Views    int
	Posts    int
	Tags     int
	Visitors int
}

func getStats() Stats {
	views := 0
	posts, _ := middleware.DBInstance.GetPosts(-1)
	ansi.PrintBold(fmt.Sprintf("posts: %v", posts))
	visitors := 0
	tags := len(middleware.GetTags())

	return Stats{views, len(posts), tags, visitors}
}

type PostsHandler struct{}

func (p PostsHandler) handleShowPosts(c echo.Context) error {
	id := c.Param("id")
	limit := c.Param("limit")
	fmt.Println("id", id)
	fmt.Println("limit", limit)
	token := jwt.Token{}

	cookie, err := c.Request().Cookie("Authorization")
	if err != nil {
		fmt.Println("error getting cookie", err)
	}

	if cookie != nil {
		fmt.Println("cookie value", cookie.Value)
		ret, err := middleware.VerifyJWTToken(cookie.Value)
		if err != nil {
			fmt.Println("error verifying JWT token", err)
		} else {
			// TODO: Refresh token as well
			token = *ret
			fmt.Println("token verified", ret)
		}
	}

	response := fmt.Sprintln("here's post ", id)
	if token.Valid {
		response += "(you're logged in)"
	} else {
		response += "(you're not logged in)"
	}

	posts := middleware.GetPostsMetadata(0)

	return Render(c, http.StatusOK, templates.PostsList(posts))
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

	posts := middleware.GetPostsMetadata(limit)

	// posts := contentsMiddleware.GetPosts(limit)
	// posts, err := parser.GetLastPosts(limit)
	// body := fmt.Sprintln("here's the last", limit, " posts. You can tweak the limit by adding a 'limit' query parameter to the URL.")
	return Render(c, http.StatusOK, templates.PostsList(posts))
}

// NewPostsHandler creates a new PostsHandler.
func newPostsHandler() *PostsHandler {
	return &PostsHandler{}
}

func (ph *PostsHandler) GetPostBySlug(c echo.Context) error {
	slug := c.Param("slug")
	ansi.PrintInfo("[POSTSHANDLER] Post with slug " + slug + " requested")

	if slug == "" {
		return Render(c, http.StatusNotFound, templates.Show("404", "Post not found", models.Post{}))
	}

	p, err := middleware.GetPostBySlug(slug)
	if err != nil {
		ansi.PrintError("[POSTSHANDLER] Error getting post by slug: " + err.Error())
		return Render(c, http.StatusNotFound, templates.Show("404", "Post not found", models.Post{}))
	}

	return Render(c, http.StatusOK, templates.Show(p.Metadata.Title, p.Content.String(), p))
}

func (ph *PostsHandler) GetPostsByTag(c echo.Context) error {
	tag := c.Param("tag")
	ansi.PrintInfo("[POSTSHANDLER] Posts with tag " + tag + " requested")

	if tag == "" {
		return Render(c, http.StatusNotFound, templates.Show("404", "Tag not found", models.Post{}))
	}

	posts, err := middleware.GetPostsByTag(tag)
	if err != nil {
		ansi.PrintError("[POSTSHANDLER] Error getting posts by tag: " + err.Error())
		return Render(c, http.StatusNotFound, templates.Show("404", "Tag not found", models.Post{}))
	}
	return Render(c, http.StatusOK, templates.PostsList(posts))
}

func (ph *PostsHandler) GetTags(c echo.Context) error {
	tags := middleware.GetTags()

	return Render(c, http.StatusOK, templates.Tags(tags))
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
