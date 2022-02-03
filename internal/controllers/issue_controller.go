package controllers

import (
	"fq_task_serivce/internal/services"
	"fq_task_serivce/internal/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func NewIssueController(issueService *services.IssueService) *IssueController {
	return &IssueController{IssueService: issueService}
}

type IssueController struct {
	IssueService *services.IssueService
}

func (ic *IssueController) NewIssue(c echo.Context) error {
	issue, err := bindIssue(c)
	if err != nil {
		return err
	}
	id, err := ic.IssueService.NewIssue(*issue)
	if err != nil {
		//todo post error type checking
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}

func (ic *IssueController) IssueById(c echo.Context) error {
	id, err := parseId(c, "id")
	if err != nil {
		return err
	}
	issue, err := ic.IssueService.IssueById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, issue)
}

func (ic *IssueController) UpdateIssue(c echo.Context) error {
	issue, err := bindIssue(c)
	if err != nil {
		return err
	}
	id0, err := ic.IssueService.UpdateIssue(*issue)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, id0)
}

func (ic *IssueController) DeleteIssue(c echo.Context) error {
	id, err := parseId(c, "id")
	if err != nil {
		return err
	}
	issue, err := ic.IssueService.DeleteIssue(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, issue)
}

func (ic *IssueController) CommentIssue(c echo.Context) error {
	id, err := parseId(c, "id")
	if err != nil {
		return err
	}
	comm, err := bindComment(c)
	if err != nil {
		return err
	}
	//todo user id
	commId, err := ic.IssueService.CommentIssue(id, *comm)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, commId)
}

func (ic *IssueController) WatchIssue(c echo.Context) error {
	id, err := parseId(c, "id")
	if err != nil {
		return err
	}
	//todo user id
	id, err = ic.IssueService.WatchIssue(id, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}

func (ic *IssueController) IssuesList(c echo.Context) error {
	issues, err := ic.IssueService.IssuesList()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, issues)
}

func (ic *IssueController) UserIssues(c echo.Context) error {
	userId, err := parseId(c, "userId")
	if err != nil {
		return err
	}
	issues, err := ic.IssueService.UserIssues(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, issues)
}

func parseId(c echo.Context, key string) (uint, error) {
	strId := c.Param(key)
	id64, err := strconv.ParseUint(strId, 0, 64)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Id must be numeric!")
	}
	return uint(id64), nil
}

func bindIssue(c echo.Context) (*types.Issue, error) {
	issue := new(types.Issue)
	if err := c.Bind(issue); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid issue strcuture!")
	}
	return issue, nil
}

func bindComment(c echo.Context) (*types.Comment, error) {
	comm := new(types.Comment)
	if err := c.Bind(comm); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid comment structure!")
	}
	return comm, nil
}
