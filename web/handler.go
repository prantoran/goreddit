package web

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/prantoran/goreddit"
)

func NewHandler(store goreddit.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)

	h.Get("/", h.Home())
	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadList())
		r.Get("/new", h.ThreadCreate())
		r.Post("/", h.ThreadStore())
		r.Get("/{id}", h.ThreadShow())
		r.Post("/{id}/delete", h.ThreadDelete())
		r.Get("/{id}/new", h.PostCreate())
		r.Post("/{id}", h.PostStore())
		r.Get("/{threadID}/{postID}", h.PostShow())
		r.Post("/{threadID}/{postID}", h.CommentStore())
	})

	return h
}

type Handler struct {
	*chi.Mux

	store goreddit.Store
}

func (h *Handler) Home() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/home.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) ThreadList() http.HandlerFunc {
	// Initializations here are done once
	type data struct {
		Threads []goreddit.Thread
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/threads.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data{Threads: tt})
	}
}

func (h *Handler) ThreadCreate() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/thread_create.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) ThreadShow() http.HandlerFunc {
	type data struct {
		Thread goreddit.Thread
		Posts  []goreddit.Post
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/thread.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		t, err := h.store.Thread(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pp, err := h.store.PostsByThread(t.ID)
		tmpl.Execute(w, data{Thread: t, Posts: pp})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (h *Handler) ThreadStore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		description := r.FormValue("description")

		if err := h.store.CreateThread(&goreddit.Thread{
			ID:          uuid.New(),
			Title:       title,
			Description: description,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}

func (h *Handler) ThreadDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "invalid thread id", http.StatusBadRequest)
			return
		}

		if err := h.store.DeleteThread(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}

func (h *Handler) PostCreate() http.HandlerFunc {
	type data struct {
		Thread goreddit.Thread
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/post_create.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "invalid thread id", http.StatusBadRequest)
			return
		}
		t, err := h.store.Thread(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data{Thread: t})
	}
}

func (h *Handler) PostShow() http.HandlerFunc {
	type data struct {
		Thread   goreddit.Thread
		Post     goreddit.Post
		Comments []goreddit.Comment
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/post.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := chi.URLParam(r, "postID")
		threadIDStr := chi.URLParam(r, "threadID")

		postID, err := uuid.Parse(postIDStr)
		if err != nil {
			http.Error(w, "invalid thread id", http.StatusBadRequest)
			return
		}

		threadID, err := uuid.Parse(threadIDStr)
		if err != nil {
			http.Error(w, "invalid thread id", http.StatusBadRequest)
			return
		}

		p, err := h.store.Post(postID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cc, err := h.store.CommentsByPost(p.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t, err := h.store.Thread(threadID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data{Thread: t, Post: p, Comments: cc})
	}
}

func (h *Handler) PostStore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		content := r.FormValue("content")

		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "invalid thread id", http.StatusBadRequest)
			return
		}
		t, err := h.store.Thread(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		p := &goreddit.Post{
			ID:       uuid.New(),
			ThreadID: t.ID,
			Title:    title,
			Content:  content,
		}
		if err := h.store.CreatePost(p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads/"+t.ID.String()+"/"+p.ID.String(), http.StatusFound)
	}
}

func (h *Handler) CommentStore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := r.FormValue("content")

		idStr := chi.URLParam(r, "postID")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "invalid post id", http.StatusBadRequest)
			return
		}

		if err := h.store.CreateComment(&goreddit.Comment{
			ID:      uuid.New(),
			PostID:  id,
			Content: content,
			Votes:   0,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}
