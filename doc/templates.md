- current value in scope: `{{.}}`
- member: `{{.Title}}`
- iterate: 
```
{{range .Lines}}
    <li>{{.}}</li>
{{end}}
```
- if statement with template directives:
```
{{if eq .Num1 .Num2}}
{{if ne .Num1 .Num2}}
{{if gt .Num1 .Num2}}
{{if lt .Num1 .Num2}}
{{if ge .Num1 .Num2}}
{{else}}
{{end}}
```

### Nested template

```
{{block "content" .}}
This is a block content.
{{end}}
```
Child templates can override block content.

childtemplate.html:
```
{{define "content"}}
123
{{end}}
```

parsing templates:
```go
t := template.Must(template.ParseFiles("templates/layout.html", "templates/childtemplate.html"))
```

### Include templates

paragraph.html
```
{{define "paragraph"}}
    <p>{{.}}</p>
{{end}}
```

```go
t := template.Must(template.New("layout.html").ParseGlob("templates/includes/*.html"))
t = template.Must(t.ParseFiles("templates/layout.html", "templates/childtemplate.html"))
```
- Must include the first template in the hierarchy


layout.html
```
{{template "paragraph" .Text}}
```


### Example templates

layout.html
```
<!DOCTYPE html>
<html>
    <body>
        {{block "header" .}}{{end}}
        <h1>{{.Title}}</h1>
        {{template "paragraph" .Text}}
        <ul>
            {{range .Lines}}
                <li>{{.}}</li>
            {{end}}
        </ul>
        {{if eq .Num1 .Num2}}
            <p>The numbers are equal!</p>
        {{else}}
            <p>The numbers are not equal.</p>
        {{end}}
        <br />
        {{block "content" .}}
        This is a block content.
        {{end}}
        {{block "footer" .}}{{end}}
    </body>
</html>
```

childtemplate.html
```
{{define "header"}}
This is a header.
{{end}}

{{define "content"}}
123
{{end}}

{{define "footer"}}
<hr />
<p>&copy; 2026 Goreddit. All rights reserved.</p>        
{{end}}
```

includes/paragraph.html
```
{{define "paragraph"}}
    <p style="text-decoration: underline;">{{.}}</p>
{{end}}
```

web/handler.go
```go
    h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}
    //...
    h.Get("/html", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("layout.html").ParseGlob("templates/includes/*.html"))
		t = template.Must(t.ParseFiles("templates/layout.html", "templates/childtemplate.html"))

		type params struct {
			Title string
			Text  string
			Lines []string
			Num1  int
			Num2  int
		}
		t.Execute(w, params{
			Title: "Goreddit",
			Text:  "Welcome to Goreddit!",
			Lines: []string{
				"Line 1",
				"Line 2",
				"Line 3",
			},
			Num1: 42,
			Num2: 42,
		})
	})
```


## Ref
[1] https://philipptanlak.com/mastering-html-templates-in-go-the-fundamentals
[2] https://github.com/gowebexamples/goreddit-templates