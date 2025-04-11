package assembler

import (
	"test-sharing-vision/service/model"
	"test-sharing-vision/service/model/request"
	"time"
)

func AssemblerToPosts(request request.InsertPosts) model.Posts {
	return model.Posts{
		Title:       request.Title,
		Content:     request.Content,
		Category:    request.Category,
		Status:      request.Status,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
}

func AssemblerToUpdatePosts(request request.UpdatePosts) model.UpdatePosts {
	return model.UpdatePosts{
		Title:       request.Title,
		Content:     request.Content,
		Category:    request.Category,
		Status:      request.Status,
		UpdatedDate: time.Now(),
	}
}
