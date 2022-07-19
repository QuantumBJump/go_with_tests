package blogposts_test

import (
	"errors"
	"hello/blogposts"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Oh no, I always fail")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("Successful fs", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
		)
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		})
		assertPost(t, posts[1], blogposts.Post{
			Title:       "Post 2",
			Description: "Description 2",
			Tags: []string{
				"rust",
				"borrow-checker",
			},
			Body: `B
L
M`,
		})
	})
	t.Run("Erroring fs", func(t *testing.T) {
		fs := StubFailingFS{}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err == nil {
			t.Fatal(err)
		}

		if posts != nil {
			t.Fatal("Expected no posts, got some")
		}
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
