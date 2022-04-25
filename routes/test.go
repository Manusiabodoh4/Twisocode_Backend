package routes

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/Manusiabodoh4/Twisocode_Backend/entity"
	"github.com/labstack/echo/v4"
)

type RoutesAccount struct {
	Router *echo.Group
}

func NewRoutesAccount(router *echo.Group) Routes {
	return &RoutesAccount{Router: router}
}

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func (st *RoutesAccount) NewCreateRoutes() {

	st.Router.GET("/decimal/:binary", func(c echo.Context) error {

		binary, _ := strconv.Atoi(c.Param("binary"))

		var (
			tmpDecimal int
			decimal    int
			counter    float64
			part       int
		)

		for binary != 0 {
			// 1001 % 10 = 1
			part = binary % 10

			if part == 0 {
				binary = binary / 10
				counter++
				continue
			}

			tmpDecimal = part * int(math.Pow(2.0, counter))
			decimal += tmpDecimal
			binary = binary / 10
			counter++
		}

		return c.JSON(http.StatusOK, decimal)

	})

	st.Router.GET("/binary/:decimal", func(c echo.Context) error {

		decimal, _ := strconv.Atoi(c.Param("decimal"))

		var (
			binary string
			satuan int
		)

		for decimal != 0 {
			satuan = decimal % 2
			binary += strconv.Itoa(satuan)
			decimal = decimal / 2
		}

		return c.JSON(http.StatusOK, reverseString(binary))

	})

	st.Router.POST("/palindrome", func(c echo.Context) error {

		body := entity.RequestJSONPalindrome{}

		err := c.Bind(&body)

		if err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		newString := strings.ReplaceAll(body.Data, " ", "")

		leng := len(newString)

		pattern := 2

		var arr []string

		for j := 0; j < 3; j++ {
			stop := leng - (pattern - 1)
			for i := 0; i < leng; i++ {
				if (stop - 1) == i {
					break
				}
				str1 := newString[i:pattern]
				str2 := newString[i+1 : pattern]
				if str1 == reverseString(str2) {
					arr = append(arr, str1)
					arr = append(arr, str2)
				}
			}
			pattern++
		}

		return c.JSON(http.StatusOK, arr)

	})

}
