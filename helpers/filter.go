package helpers

import (
	"fmt"
	"strings"

	"github.com/IlhamLamp/cmty-core-service/dto"
	"gorm.io/gorm"
)

func FilterProjectsByType(q *gorm.DB, filter dto.CoreFilter) *gorm.DB {
	if filter.Types != "" {
		q = q.Where("types = ?", filter.Types)
	}
	if filter.Duration != "" {
		q = q.Where("duration = ?", filter.Duration)
	}
	if filter.Participation != "" {
		q = q.Where("participation = ?", filter.Participation)
	}
	if filter.Location != "" {
		q = q.Where("LOWER(address->>'city') LIKE ?", "%"+strings.ToLower(filter.Location)+"%")
	}
	q = applyTagsFilter(q, filter.Tags)
	q = applyMemberFilter(q, filter.Role, filter.Experience)
	return q
}

func applyTagsFilter(q *gorm.DB, tags []string) *gorm.DB {
	if len(tags) > 0 {
		for _, tag := range tags {
			q = q.Where("tags @> ?", fmt.Sprintf(`["%s"]`, tag))
		}
	}
	return q
}

func applyMemberFilter(q *gorm.DB, role string, experience string) *gorm.DB {
	if role != "" || experience != "" {
		q = q.Joins("JOIN members ON members.project_id = projects.id")
		if role != "" {
			q = q.Where("members.role_id = ?", role)
		}
		if experience != "" {
			switch experience {
			case "no_experience":
				q = q.Where("members.experience = 'no_experience'")
			case "less_than_year":
				q = q.Where("members.experience = 'less_than_year'")
			case "more_than_year":
				q = q.Where("members.experience = 'more_than_year'")
			}
		}
	}
	return q
}

// MEMBER
func FilterMembersByType(q *gorm.DB, filter dto.MemberFilter) *gorm.DB {
	if filter.Types != "" {
		q = q.Where("type = ?", filter.Types)
	}
	if filter.Role != "" {
		q = q.Where("role = ?", filter.Role)
	}
	if filter.Experience != "" {
		q = q.Where("experience = ?", filter.Experience)
	}

	return q
}
