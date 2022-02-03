package routers

import (
	"fq_task_serivce/internal/controllers"

	"github.com/labstack/echo"
)

func RouteIssues(e *echo.Echo, c *controllers.IssueController) {
	e.POST("/issue", c.NewIssue)
	e.GET("/issue/:id", c.IssueById)
	e.PUT("/issue/", c.UpdateIssue)
	e.DELETE("/issue/:id", c.DeleteIssue)
	e.PUT("/issue/:id/comment", c.CommentIssue)
	e.PUT("/issue/:id/watch", c.WatchIssue)
	e.GET("/issues", c.IssuesList)
	e.GET("/issues/:userId", c.UserIssues)
	//todo :userId
}
