package handler

// import (
// 	"net/http"

// 	"github.com/MStation-io/api/model"
// 	"github.com/labstack/echo/v4"
// )

// func (h *Handler) GetAllCourse(c echo.Context) {
// 	var sliceUsers []model.Course
// 	result, _ := db.Query("SELECT * FROM student")
// 	for result.Next() {
// 		var s model.Course
// 		_ = result.Scan(&s.ID, &s.Fullname, &s.Age, &s.Location)
// 		sliceUsers = append(sliceUsers, s)
// 	}
// 	return c.JSON(http.StatusOK, sliceUsers)
// }
