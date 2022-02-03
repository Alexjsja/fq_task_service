package types

import (
	"fq_task_serivce/internal/types/isstype"
	"fq_task_serivce/internal/types/priority"
	"fq_task_serivce/internal/types/status"
	"time"

	"github.com/StanDenisov/fq_utils/users"
	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	OwnerId      uint                   `json:"owner_id"`
	Owner        *users.User            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	ExecutorId   *uint                  `json:"executor_id,omitempty"`
	Executor     *users.User            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Watchers     []users.User           `json:"watchers,omitempty" gorm:"many2many:issue_watchers"`
	PlannedStart *time.Time             `json:"planned_start"`
	PlannedEnd   *time.Time             `json:"planned_end"`
	ActualStart  *time.Time             `json:"actual_start,omitempty"`
	ActualEnd    *time.Time             `json:"actual_end,omitempty"`
	Type         isstype.IssueType      `json:"type"`
	Status       status.IssueStatus     `json:"status"`
	Priority     priority.IssuePriority `json:"priority"`
	Comments     []Comment              `json:"comments,omitempty" gorm:"many2many:issue_comments"`
}
