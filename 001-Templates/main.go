package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var funcMap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func main() {
	//tpl := template.Must(template.ParseFiles("01-basic.gohtml", "02-data.gohtml", "03-variables.gohtml", "04-composite-data.gohtml"))
	//tpl, err := template.ParseFiles("01-basic.gohtml", "02-data.gohtml")
	//err := tpl.ExecuteTemplate(os.Stdout, "01-basic.gohtml", nil)  // simple data
	//err = tpl.ExecuteTemplate(os.Stdout, "02-data.gohtml", 42)  // simple data
	//err = tpl.ExecuteTemplate(os.Stdout, "03-variables.gohtml", "Redrum")  // variables
	//sages := []string{"Gandhi", "MLK", "Buddha", "Jesus"}
	//err = tpl.ExecuteTemplate(os.Stdout, "04-composite-data.gohtml", sages)  // composite data - slice
	//sages := map[string]string{
	//	"India": "Gandhi",
	//	"America": "MLK",
	//	"Love": "Jesus",
	//	"meditate": "Buddha",
	//}
	//err = tpl.ExecuteTemplate(os.Stdout, "04-composite-data.gohtml", sages)  // composite data - map
	//err = tpl.ExecuteTemplate(os.Stdout, "04-composite-data.gohtml", sage{Name: "Buddha", Motto: "Meditate"})  // composite data - struct
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//tpl := template.Must(template.New("05-functions.gohtml").Funcs(funcMap).ParseFiles("05-functions.gohtml"))
	//tpl := template.Must(template.New("06-pipeline.gohtml").Funcs(funcMap).ParseFiles("06-pipeline.gohtml"))
	//data := []sage{{Name: "Buddha", Motto: "Meditate"}, {Name: "Gandhi", Motto: "Non violence"}}
	//err := tpl.Execute(os.Stdout, data)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	tpl := template.Must(template.ParseGlob("08-nested-templates/*.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "parent.gohtml", nil)
	if err != nil {
		log.Panic(err)
	}
}
