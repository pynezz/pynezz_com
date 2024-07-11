package middleware

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/pynezz/pynezz_com/internal/parser"
	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
	"gorm.io/datatypes"
)

// GetPosts retrieves all posts from the database up to a specified limit
func (d *Database) GetPosts(limit int) ([]models.PostMetadata, error) {
	var posts []models.PostMetadata
	result := d.Driver.Limit(limit).Find(&posts)
	return posts, result.Error
}

func (d *Database) GetPostBySlug(slug string) (models.PostMetadata, error) {
	var post models.PostMetadata
	result := d.Driver.Where("slug = ?", slug).First(&post)
	return post, result.Error
}

func (d *Database) GetPostByID(id uint) (models.PostMetadata, error) {
	var post models.PostMetadata
	result := d.Driver.Where("id = ?", id).First(&post)
	return post, result.Error
}

// GetPostByPath retrieves a post from the database based on the path to the markdown file
func (d *Database) GetPostByPath(path string) (models.Post, error) {
	var post models.Post
	result := d.Driver.Where("path = ?", path).First(&post)
	return post, result.Error
}

// GetPostsByTag retrieves all posts from the database that contain a specific tag
func (d *Database) GetPostsByTag(tag string) ([]models.PostMetadata, error) {
	var posts []models.PostMetadata
	result := d.Driver.Where("tags LIKE ?", fmt.Sprintf("%%%s%%", tag)).Find(&posts)
	return posts, result.Error
}

// NewPost creates a new post in the database, given the metadata of the post
// This is differennt from the parser.NewPost function, which creates the necessary metadata for a post
func (d *Database) NewPost(post models.PostMetadata) error {
	result := d.Driver.Create(&post)
	return result.Error
}

// GenerateMetadata generates metadata for a post based off the contents of a markdown file
// This is the metadata that is stored in the database for each post
// The middleware will generate this metadata when a new post is created or on a schedule
// The metadata is used to generate the post's URL, tags, and other information
// Finally, the metadata is fetched from the database, and the post is generated and displayed on the website
func (d *Database) GenerateMetadata(post *models.Post) models.PostMetadata {
	parsedMetadata, err := parser.ParseMetadata([]byte(post.Content)) // Parse the metadata of the contents of the markdown file
	if err != nil {
		fmt.Println("Error parsing metadata:", err)
	}

	pPost := parser.NewPost(post.Content) // Create a new post
	parser.SetDescription(pPost)          // Set the description of the post based off the content or the description set in the metadata

	return models.PostMetadata{
		Title: post.Title,
		Slug:  d.GenerateSlug(post.Title),
		Tags:  datatypes.JSON(strings.Join(parsedMetadata.Tags, ",")), // Convert the tags to a JSON array ( // TODO: is this valid?)
	}
}

func commonStopWord(word string) bool {
	m := map[string]bool{
		"the":  true,
		"a":    true,
		"an":   true,
		"and":  true,
		"but":  true,
		"or":   true,
		"for":  true,
		"nor":  true,
		"so":   true,
		"yet":  true,
		"with": true,
		"in":   true,
		"on":   true,
		"at":   true,
		"by":   true,
		"to":   true,
		"of":   true,
		"as":   true,
		"from": true,
		"into": true,
	}
	return m[word]
}

// GenerateSlug generates a slug from the title of a post.
//
/*
	Example: "Hello World" -> "hello-world"
	Example: "Hello World, Again" -> "hello-world"
	Example: "The quick brown fox jumps over the lazy dog" -> "the-quick-brown-fox-jumps"
	Example: "The quick brown fox jumps over the lazy dog, again" -> "the-quick-brown-fox-jumps1"
/*
	fmt.Println(GenerateSlug("The quick brown fox jumps over the lazy dog"))
	// "the-quick-brown-fox-jumps-over"

	fmt.Println(GenerateSlug("Here's the story about a little guy that lives in a blue world"))
	// heres-the-story-about-a-little

	fmt.Println(GenerateSlug("Quickly some brown fox sprung into another leap accross another lazy dog"))
	> "quickly-some-brown-fox-sprung" "Error: 'into' is a common word)"
*/
func (d *Database) GenerateSlug(title string) string {
	title = strings.ToLower(title) // Lowercase the title
	words := strings.Fields(title) // Split the title into words (Fields splits the string s around each instance of one or more consecutive white space characters)
	maxWordLength := 4

	if len(words) > maxWordLength {

		// Find the number of common descriptor words in the first "maxWordLength" words (default to four)
		for i, word := range words {
			if i > maxWordLength && !commonStopWord(word) {
				maxWordLength = i - 1 // We've found the ideal number of words
				break
			}
			if i < maxWordLength && commonStopWord(word) {
				fmt.Printf("%s is a common word\n", word)
				maxWordLength++
			}
		}

		for {
			if !commonStopWord(words[maxWordLength-1]) {
				break
			}
			maxWordLength++
		}

		words = words[:maxWordLength] // Get the first four words
	}

	for i, word := range words {
		// Remove all non-alphanumeric characters
		words[i] = strings.Map(func(r rune) rune {
			if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || (r == '-') {
				return r
			}
			return -1
		}, word)
	}

	// Join the words with a hyphen
	slug := strings.Join(words, "-")

	// Check if the slug is already in use
	var post models.PostMetadata
	d.Driver.Where("slug = ?", slug).First(&post)

	// Keep adding a number to the end of the slug until it is unique
	if post.Slug == slug {
		// If the slug is already in use, add a number to the end of the slug
		var i int
		for {
			i++
			slug = fmt.Sprintf("%s-%d", slug, i)
			d.Driver.Where("slug = ?", slug).First(&post)
			if post.Slug != slug {
				break
			}
		}
	}

	return slug
}

/*
	------------ EXPERIMENTAL FUNCTION ----------------
*/
// GenerateSlug generates a slug from the title of a post.
func (d *Database) ExperimentalSlug(title string) string {
	title = strings.ToLower(title) // Lowercase the title
	words := strings.Fields(title) // Split the title into words

	// Remove non-alphanumeric characters from each word
	for i, word := range words {
		words[i] = strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsNumber(r) {
				return r
			}
			return -1
		}, word)
	}

	// Limit the number of words in the slug
	maxWords := 4
	finalWords := words[:maxWords]

	// Add additional words until no more stopwords are found or the end is reached
	for i := maxWords; i < len(words); i++ {
		finalWords = append(finalWords, words[i])
		if !commonStopWord(words[i]) {
			break
		}
	}

	// Join the words with a hyphen
	slug := strings.Join(finalWords, "-")

	// Check if the slug is already in use
	var post models.PostMetadata
	d.Driver.Where("slug = ?", slug).First(&post)

	// Keep adding a number to the end of the slug until it is unique
	if post.Slug == slug {
		// If the slug is already in use, add a number to the end of the slug
		var i int
		for {
			i++
			slug = fmt.Sprintf("%s-%d", slug, i)
			d.Driver.Where("slug = ?", slug).First(&post)
			if post.Slug != slug {
				break
			}
		}
	}

	return slug
}
