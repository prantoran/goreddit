package web

import "encoding/gob"

type FormErrors map[string]string

func init() { // automatically run when the package is imported
	gob.Register(CreatePostForm{}) // register the structs for session storage
	gob.Register(CreateThreadForm{})
	gob.Register(CreateCommentForm{})
	gob.Register(FormErrors{})
	gob.Register(RegisterForm{})
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

type RegisterForm struct {
	Username      string `form:"username"`
	Password      string `form:"password"`
	UsernameTaken bool   `form:"username_taken"`

	Errors FormErrors
}

func (f *RegisterForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Username == "" {
		f.Errors["Username"] = "Please enter a username."
	} else if f.UsernameTaken {
		f.Errors["Username"] = "Username is already taken."
	}

	if f.Password == "" {
		f.Errors["Password"] = "Please enter a password."
	} else if len(f.Password) < 8 {
		f.Errors["Password"] = "Password must be at least 8 characters long."
	}

	return len(f.Errors) == 0
}
