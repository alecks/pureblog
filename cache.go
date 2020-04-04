package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

// getCSS gets the main.css file from the cache if it exists; else get from posts folder.
func getCSS() ([]byte, error) {
	if cssCache != nil {
		return cssCache, nil
	}

	file, err := ioutil.ReadFile(path.Join("styles", "main.css"))
	cssCache = file
	return file, err
}

// cssCache holds the main.css file's cache.
var cssCache []byte

// getHTML gets the main.html file from the cache if it exists; else get from posts folder.
func getHTML() (string, error) {
	if htmlCache != "" {
		return htmlCache, nil
	}

	file, err := ioutil.ReadFile(path.Join("styles", "main.html"))
	htmlCache = string(file)
	return htmlCache, err
}

// htmlCache holds the main.html file's cache.
var htmlCache string = ""

// getPost gets a post from the cache if it exists; else get from posts folder.
func getPost(post string, cache bool) (string, error) {
	// If the post id doesn't have the md extension, append to it with .md.
	if !strings.HasSuffix(post, ".md") {
		post += ".md"
	}

	// If the post is present in the cache, return it.
	if found, ok := postCache[post]; ok && cache {
		fmt.Println("cached")
		return found, nil
	}

	// Else read the file and add it to cache.
	file, err := ioutil.ReadFile(path.Join("posts", post))

	cssMain, _ := getCSS()
	htmlMain, _ := getHTML()

	res := fmt.Sprintf(
		htmlMain,
		post,
		cssMain,
		md.RenderToString(file),
	)
	postCache[post] = res
	return res, err
}

// postCache holds cached posts.
var postCache = make(map[string]string)
