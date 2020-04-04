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

	file, err := ioutil.ReadFile(path.Join("styles", "post.css"))
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

	file, err := ioutil.ReadFile(path.Join("styles", "post.html"))
	htmlCache = string(file)
	return htmlCache, err
}

// htmlCache holds the main.html file's cache.
var htmlCache string = ""

// getPost gets a post from the cache if it exists; else get from posts folder.
func getPost(post string) (string, error) {
	// If the post id doesn't have the md extension, append to it with .md.
	if !strings.HasSuffix(post, ".md") {
		post += ".md"
	}

	// If the post is present in the cache, return it.
	if found, ok := postCache[post]; ok && shouldCache {
		return found, nil
	}

	// Else read the file and add it to cache.
	file, err := ioutil.ReadFile(path.Join("posts", post))

	cssMain, err := getCSS()
	htmlMain, err := getHTML()
	if err != nil {
		panic(err)
	}

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

// getPostPreview gets the template for posts being displayed in the post list from cache if present; else read file.
func getPostPreview() (string, error) {
	if shouldCache && postPreviewCache != "" {
		return postPreviewCache, nil
	}

	file, err := ioutil.ReadFile(path.Join("styles", "post_preview.html"))
	postPreviewCache = string(file)

	return postPreviewCache, err
}

// postPreviewCache holds the post list entry template.
var postPreviewCache = ""

// getPostListTemplate gets the template for posts being displayed in the post list from cache if present; else read file.
func getPostListTemplate() (string, error) {
	if shouldCache && postListTemplateCache != "" {
		return postListTemplateCache, nil
	}

	file, err := ioutil.ReadFile(path.Join("styles", "post_list.html"))
	postListTemplateCache = string(file)

	return postListTemplateCache, err
}

// postListTemplateCache holds the post list entry template.
var postListTemplateCache = ""

// getPostList gets a list of posts for the homepage from the cache if present; else read dir.
func getPostList() (res string, err error) {
	if shouldCache && postListCache != "" {
		return postListCache, nil
	}

	postDirectory, err := ioutil.ReadDir("posts")

	for _, v := range postDirectory {
		postName := v.Name()
		preview, err := getPostPreview()
		if err != nil {
			panic(err)
		}
		res += fmt.Sprintf(preview, "/"+postName, postName)
	}

	plTemplate, err := getPostListTemplate()
	res = fmt.Sprintf(plTemplate, res)

	postListCache = res
	return
}

// postListCache holds the post list.
var postListCache = ""
