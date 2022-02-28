package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostsFromFS(f fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(f, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, file := range dir {
		post, err := getPost(f, file.Name())
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(filesystem fs.FS, f string) (Post, error) {
	postFile, err := filesystem.Open(f)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return parsePost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
	bodySeparator        = "---"
)

func parsePost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}
	readBody := func(sc *bufio.Scanner) string {
		scanner.Scan()

		buff := bytes.Buffer{}

		for scanner.Scan() {
			fmt.Fprint(&buff, scanner.Text()+"\n")
		}

		return strings.TrimSuffix(buff.String(), "\n")
	}

	post := Post{
		Title:       readLine(titleSeparator),
		Description: readLine(descriptionSeparator),
		Tags:        strings.Split(readLine(tagSeparator), ", "),
		Body:        readBody(scanner),
	}

	return post, nil
}
