package services

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PaginationResult struct {
	Items        interface{} `json:"items"`
	TotalCount   int64       `json:"total_count"`
	CurrentPage  int         `json:"current_page"`
	LastPage     int         `json:"last_page"`
	PreviousPage *string     `json:"previous_page"`
	NextPage     *string     `json:"next_page"`
}

func CustomPaginate(
	c *fiber.Ctx,
	query *gorm.DB,
	result interface{},
	defaultLimit int,
) (*PaginationResult, error) {

	// Parse page
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	// Parse limit
	perPage, err := strconv.Atoi(c.Query("limit", strconv.Itoa(defaultLimit)))
	if err != nil || perPage < 1 {
		perPage = defaultLimit
	}

	// Clone query for count (IMPORTANT)
	var totalCount int64
	countQuery := query.Session(&gorm.Session{})
	if err := countQuery.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	// Calculate pagination
	offset := (page - 1) * perPage
	lastPage := int(math.Ceil(float64(totalCount) / float64(perPage)))

	// Fetch paginated data
	if err := query.Offset(offset).Limit(perPage).Find(result).Error; err != nil {
		return nil, err
	}

	// Build prev / next URLs
	var prevPage *string
	var nextPage *string

	if page > 1 {
		u := c.Path() + "?page=" + strconv.Itoa(page-1) + "&limit=" + strconv.Itoa(perPage)
		prevPage = &u
	}

	if page < lastPage {
		u := c.Path() + "?page=" + strconv.Itoa(page+1) + "&limit=" + strconv.Itoa(perPage)
		nextPage = &u
	}

	return &PaginationResult{
		Items:        result,
		TotalCount:   totalCount,
		CurrentPage:  page,
		LastPage:     lastPage,
		PreviousPage: prevPage,
		NextPage:     nextPage,
	}, nil
}
