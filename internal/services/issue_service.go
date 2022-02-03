package services

import (
	"fmt"
	"fq_task_serivce/internal/types"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const existsByIdQuery = "select exists(select i.id from issues i where id = ?)"

func NewIssueService(db *gorm.DB) *IssueService {
	return &IssueService{DB: db}
}

type IssueService struct {
	DB *gorm.DB
}

func (is *IssueService) NewIssue(issue types.Issue) (id uint, err error) {
	if err := (is.DB.Create(&issue)).Error; err != nil {
		return 0, err
	}
	return issue.ID, nil
}

func (is *IssueService) IssueById(id uint) (types.Issue, error) {
	if !is.existsById(id) {
		return types.Issue{}, fmt.Errorf("issue by id=%d not found", id)
	}
	iss := types.Issue{}
	res := is.DB.Where("id = ?", id).Find(&iss)
	if res.Error != nil {
		return iss, res.Error
	}
	return iss, nil
}

func (is *IssueService) UpdateIssue(issue types.Issue) (id uint, err error) {
	if err := (is.DB.Save(&issue)).Error; err != nil {
		return 0, err
	}
	return issue.ID, nil
}

func (is *IssueService) DeleteIssue(id uint) (types.Issue, error) {
	var issue types.Issue
	err := is.DB.Clauses(clause.Returning{}).
		Unscoped().
		Where("id = ?", id).
		Delete(&issue).
		Error
	if err != nil {
		return types.Issue{}, err
	}
	return issue, nil
}

func (is *IssueService) CommentIssue(id uint, comment types.Comment) (commentId uint, err error) {
	issue, err := is.IssueById(id)
	if err != nil {
		return 0, err
	}
	issue.Comments = append(issue.Comments, comment)
	is.DB.Save(&issue)
	return comment.ID, nil
}

func (is *IssueService) WatchIssue(id0 uint, userId uint) (id uint, err error) {
	return 0, nil
}

func (is *IssueService) IssuesList() ([]types.Issue, error) {
	return nil, nil
}

func (is *IssueService) UserIssues(userId uint) ([]types.Issue, error) {
	return nil, nil
}

func (is *IssueService) existsById(id uint) bool {
	var exists bool
	is.DB.Raw(existsByIdQuery, id).Scan(&exists)
	return exists
}
