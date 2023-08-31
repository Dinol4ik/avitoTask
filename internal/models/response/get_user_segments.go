package response

import "avitoTask/internal/models"

type GetUserSegments struct {
	Name models.UserSegments `json:"name"`
}
