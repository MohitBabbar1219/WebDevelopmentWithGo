# Templates

## text/template
### Steps:
#### 1. Parse - 
1. `template.ParseFile("<filename1>, <filename2>, ...")`<br/>
Loads a template. This will return a pointer to template and error. 
2. `template.ParseGlob("templates/*.gohtml")`<br/>
Parses all the template files in `template` directory with `gohtml` extension.
3. `template.Must(template.ParseGlob("templates/*.gohtml"))`<br/>
`Must` takes pointer to template and an error, same things which `ParseGlob` returns, and does the error checking for us. It returns a pointer to template which can then be used without error checking.
#### 2. Execute - 
1. Execute the template to inject data and output it to desired location. Can be done in two ways:
    1. `tpl.Execute(<object implementing writer interface>, data)`<br/>
    If only one template file was parsed.
    2. `tpl.ExecuteTemplate(<object implementing writer interface>, "template name", data)`<br/>
    If more than one template files were parsed.

### Passing data
- Template file can receive data with `{{.}}` syntax.
- `.` represents the data. Example: `<h1>{{.}}</h2>`.
- Only one piece of data can be passed on to the template by the `Execute` method but it can be a complex data type like `struct`, `slice` or whatnot.

### Variables in templates
- Variables can be declared with `{{$varName := .}}` syntax.
- These variables can be used with `{{$varName}}` syntax in the whole template.

### Composite data
- Slices
```gotemplate
{{range .}}
<h1>{{.}}</h1>
{{end}}

{{range $index, $val := .}}
<h1>{{$index}} - {{$val}}</h1>
{{end}}
```
- Maps
```gotemplate
{{range $key, $val := .}}
<h1>{{$key}} - {{$val}}</h1>
{{end}}
```
- Structs
```gotemplate
<h1>{{.Name}} - {{.Motto}}</h1>
```

### Functions in templates
- `template` provides a type `FuncMap` which can be used to aggregate all the functions which are to be passed in the template, together.
- `template.FuncMap` is `type FuncMap map[string]interface{}`
- `template.FuncMap`