// build f0112df575
// (go) f9759d
package main

import (
    "html/template"
    "os"
)

type Ctx struct{ Title, Content string }

func main() {
    t := template.Must(template.ParseFiles("templates/index.html"))
    b, _ := os.ReadFile("content/index.html")
    os.MkdirAll("_site", 0755)
    out, _ := os.Create("_site/index.html")
    defer out.Close()
    t.Execute(out, Ctx{Title: "girlsbaito-server", Content: string(b)})
}
