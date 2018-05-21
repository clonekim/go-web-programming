package netserver

import (
	"bytes"
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/neko-neko/echo-logrus/log"
	"html/template"
	"io"
	"net/http"
)

//에코 템플릿
type Template struct {
	Box *rice.Box
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	templateLayout, err := t.Box.String("layout.html")
	if err != nil {
		return err
	}

	templateString, err := t.Box.String(name)
	if err != nil {
		return err
	}

	buff := bytes.Buffer{}
	buff.WriteString(templateLayout)
	buff.WriteString(templateString)

	tmplMsg, err := template.New("Dummy").Parse(buff.String())

	if err != nil {
		Log.WithField("Template Error", err).Debug("Template error is")
		return err
	}

	return tmplMsg.Execute(w, data)
}

type Asset struct {
	Box *rice.Box
}

func NewAsset() *Asset {
	Log.Println("Asset loading")
	box, error := rice.FindBox("templates")
	if error != nil {
		panic(error)
	}

	return &Asset{
		Box: box,
	}
}

type Validations interface {
	Validate() error
}

type ValidatorCaller struct {
}

func (c *ValidatorCaller) Validate(i interface{}) (err error) {
	v := i.(Validations)
	Log.Debugf("Request is %#v", i)
	err = v.Validate()

	logEntry := Log.WithField("Validation Result", err)

	if err != nil {
		logEntry.Error("Validation Fail")
	} else {
		logEntry.Debug("Validaion Ok")
	}
	return
}

// 에코 템플릿 끝

func EchoStart() {
	asset := NewAsset()
	Log.Println("Start echo server")

	e := echo.New()
	e.Logger = log.Logger()
	e.Renderer = &Template{Box: asset.Box}
	e.Validator = &ValidatorCaller{}
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				Log.Error(err)
			}
			return err
		}
	})

	//라우터
	templateHandler := http.FileServer(asset.Box.HTTPBox())

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", echo.Map{})
	})

	v1 := e.Group("/v1")
	v1.POST("/user", UserCreate)

	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", templateHandler)))

	e.Logger.Fatal(e.Start(Conf.Port))

}
