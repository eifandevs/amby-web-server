package main

import (
    "html/template"
    "net/http"
    "io"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "./handler"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

type ServiceInfo struct {
    Title string
}

var serviceInfo = ServiceInfo {
    "サイトのタイトル",
}

  
func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // routing
    e.GET("/hello", handler.Hello())

    // html routing
    t := &Template{
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
    e.Renderer = t
    e.GET("/chat", func(c echo.Context) error {
        // テンプレートに渡す値
    
    data := struct {
        ServiceInfo
        Content_a string
        Content_b string
        Content_c string
        Content_d string
    } {
        ServiceInfo: serviceInfo,
        Content_a: "雨が降っています。",
        Content_b: "明日も雨でしょうか。",
        Content_c: "台風が近づいています。",
        Content_d: "Jun/11/2018",
    }
    return c.Render(http.StatusOK, "chat", data)
    })

    // port
    e.Start(":8080")
}