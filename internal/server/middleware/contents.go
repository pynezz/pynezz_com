package middleware

import (
	"fmt"
	"strings"

	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
	"github.com/pynezz/pynezzentials/ansi"
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
	result := d.Driver.Where(&models.PostMetadata{Slug: post.Slug}).FirstOrCreate(&post)
	ansi.PrintInfo(fmt.Sprintf("new post created: %s\n%d affected rows", post.Title, result.RowsAffected))
	return result.Error
}

func (d *Database) UpdatePostMetadata(post models.PostMetadata) error {
	result := d.Driver.Save(&post)
	return result.Error
}

func (d *Database) DeletePost(post models.PostMetadata) error {
	result := d.Driver.Delete(&post)
	return result.Error
}

func (d *Database) WriteContentsToDatabase(slug string, content []byte, pMetadata models.PostMetadata) error {
	ansi.PrintColor(ansi.Yellow, string(content))
	post := models.Post{
		Content: datatypes.JSON(content),
		Metadata: models.Metadata{
			Title:       pMetadata.Title,
			Description: pMetadata.Summary,
			Date:        pMetadata.CreatedAt,
			Tags:        datatypes.JSON(pMetadata.Tags),
		},
		Slug: slug,
	}

	result := d.Driver.Create(&post)
	ansi.PrintInfo(fmt.Sprintf("new post created: %s\n%d affected rows", post.Metadata.Title, result.RowsAffected))
	return result.Error

	// return d.Driver.Model(&models.Post{}).Where("Slug = ?", slug).Update("Content", post.Content).Error
}

// GenerateMetadata generates metadata for a post based off the contents of a markdown file.
// It writes the metadata and post to the database.
// The middleware will generate this metadata when a new post is created or on a schedule.
// The metadata is used to generate the post's URL, tags, and other information.
// Later, the metadata will be used to fetch from the database,
// and the post is generated and displayed on the website
// func (d *Database) GenerateMetadata(content []byte) models.PostMetadata {
// 	parsedMetadata, err := parser.ParseMetadata(content) // Parse the metadata of the contents of the markdown file
// 	if err != nil {
// 		fmt.Println("Error parsing metadata:", err)
// 	}

// 	pPost := parser.NewPost(string(content)) // Create a new post
// 	parser.SetDescription(pPost)             // Set the description of the post based off the content or the description set in the metadata

// 	ansi.PrintColor(ansi.Purple, "Title: "+parsedMetadata.Title)
// 	md := models.PostMetadata{
// 		Title: parsedMetadata.Title,
// 		Slug:  d.GenerateSlug(parsedMetadata.Title),
// 		Tags:  datatypes.JSON(strings.Join(parsedMetadata.Tags, ",")), // Convert the tags to a JSON array ( // TODO: is this valid?)
// 	}

// 	// Write post to database
// 	d.NewPost(md)

// 	return md
// }

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

	fmt.Println("Words: ", words)
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
			if (r >= 'a' && r <= 'z') || (r == '-') {
				return r
			}
			return -1
		}, word)
	}

	// Join the words with a hyphen
	slug := strings.Join(words, "-")

	// Check if the slug is already in use
	var post []models.PostMetadata
	d.Driver.Where("slug LIKE ?", slug+"%").Find(&post)
	if len(post) == 0 {
		ansi.PrintInfo("No posts found with slug: " + slug + " returning slug for use")
		return slug
	}

	ansi.PrintWarning("Posts found with slug: " + slug + " generating new slug")

	// Find the highest numbered slug
	maxNum := 0
	for _, p := range post {
		var num int
		fmt.Sscanf(p.Slug, slug+"-%d", &num)
		if num > maxNum {
			maxNum = num
		}
	}

	return fmt.Sprintf("%s-%d", slug, maxNum+1)
}
