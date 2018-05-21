package netserver

import (
	"bytes"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/neko-neko/echo-logrus/log"
	"html/template"
	"io"
	"net/http"
	"os"
)

//에코 템플릿
type Template struct {
	Box *rice.Box
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	templateLayout, err := t.Box.String("layout.html")
	if err != nil {
		fmt.Println(err)
		return err
	}

	templateString, err := t.Box.String(name)
	if err != nil {
		fmt.Println(err)
		return err
	}

	buff := bytes.Buffer{}
	buff.WriteString(templateLayout)
	buff.WriteString(templateString)

	tmplMsg, err := template.New("Dummy").Parse(buff.String())

	if err != nil {
		fmt.Println(err)
		return err
	}

	return tmplMsg.Execute(w, data)
}

type Asset struct {
	Box *rice.Box
}

func NewAsset() *Asset {
	box, error := rice.FindBox("templates")
	if error != nil {
		panic(error)
	}

	return &Asset{
		Box: box,
	}
}

// 에코 템플릿 끝

func getLogger() *log.MyLogger {
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(echoLog.DEBUG)
	return log.Logger()
}

func EchoStart() {

	asset := NewAsset()
	e := echo.New()
	e.Logger = getLogger()
	e.Renderer = &Template{Box: asset.Box}
	e.Use(middleware.Recover())

	//라우터
	templateHandler := http.FileServer(asset.Box.HTTPBox())

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", echo.Map{})
	})

	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", templateHandler)))

	e.Logger.Fatal(e.Start(Conf.Port))

}
