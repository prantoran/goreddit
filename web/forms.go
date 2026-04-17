package web

import "encoding/gob"

type FormErrors map[string]string

func init() { // automatically run when the package is imported
	gob.Register(CreatePostForm{}) // register the structs for session storage
	gob.Register(CreateThreadForm{})
	gob.Register(CreateCommentForm{})
	gob.Register(FormErrors{})
}

type CreatePostForm struct {
	Title   string `form:"title"`
	Content string `form:"content"`

	Errors FormErrors
}

func (f *CreatePostForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Title == "" {
		f.Errors["Title"] = "Please enter a title."
	}
	if f.Content == "" {
		f.Errors["Content"] = "Please enter a text."
	}

	return len(f.Errors) == 0
}

type CreateThreadForm struct {
	Title       string `form:"title"`
	Description string `form:"description"`

	Errors FormErrors
}

func (f *CreateThreadForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Title == "" {
		f.Errors["Title"] = "Please enter a title."
	}
	if f.Description == "" {
		f.Errors["Description"] = "Please enter a description."
	}

	return len(f.Errors) == 0
}

type CreateCommentForm struct {
	Content string `form:"content"`

	Errors FormErrors
}

func (f *CreateCommentForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Content == "" {
		f.Errors["Content"] = "Please enter a comment."
	}

	// TODO: check email format using regex

	return len(f.Errors) == 0
}
