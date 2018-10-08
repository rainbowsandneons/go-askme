package answer

import "github.com/bashmohandes/go-askme/model"

// Repository represents the basic answer repo functionality
type Repository interface {
	LoadAnswers(userID models.UniqueID) []*models.Answer
	AddLike(answer *models.Answer, user *models.User)
	RemoveLike(answer *models.Answer, user *models.User)
	GetLikesCount(answer *models.Answer) uint
	Add(answer *models.Answer)
}
