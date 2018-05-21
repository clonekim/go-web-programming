package netserver

import (
	"github.com/labstack/echo"
)

func UserCreate(c echo.Context) (err error) {
	var user User

	if err = c.Bind(&user); err != nil {
		return
	}

	if err := c.Validate(&user); err != nil {
		return c.JSON(400, echo.Map{"validations": err})
	}

	if err := user.Save(); err != nil {
		return err
	}

	return c.JSON(200, &user)
}

/*
func CardCreate(c echo.Context) (err error) {
	var card Card

	if err = c.Bind(&card); err != nil {
		return
	}

	if err = c.Validate(&card); err != nil {
		return c.JSON(400, echo.Map{"validations": err})
	}

	if err = card.Save(); err != nil {
		return
	}

	return c.JSON(200, &card)

}
*/
