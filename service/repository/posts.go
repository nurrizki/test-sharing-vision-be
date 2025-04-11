package repository

import (
	"database/sql"
	"fmt"
	"test-sharing-vision/constant"
	"test-sharing-vision/internal/database"
	"test-sharing-vision/service/model"
	"test-sharing-vision/service/model/request"
)

func InsertBulk(data []model.Posts) error {
	// Insert data to posts
	err := database.ArticleDB.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func InsertDataPosts(data model.Posts) error {
	// Insert data to posts
	err := database.ArticleDB.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func GetDataPosts(req request.GetPosts) ([]model.Posts, int, error) {
	var posts []model.Posts
	var total int64
	query := database.ArticleDB.Model(&model.Posts{})

	if req.Title != "" {
		query = query.Where("title LIKE ?", "%"+req.Title+"%")
	}
	if req.Content != "" {
		query = query.Where("content LIKE ?", "%"+req.Content+"%")
	}
	if req.Id != "" {
		query = query.Where("id = ?", req.Id)
	}
	if req.Category != "" {
		query = query.Where("category = ?", req.Category)
	}
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	query.Count(&total)

	// Pagination
	if req.Limit > 0 {
		query = query.Limit(req.Limit)
	} else {
		query = query.Limit(10) // default limit
	}

	query = query.Offset(req.Offset)

	err := query.Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	if len(posts) == 0 {
		return nil, 0, sql.ErrNoRows
	}

	return posts, int(total), err
}

func DeleteDataPosts(id string) error {
	// Delete data posts
	result := database.ArticleDB.Delete(&model.Posts{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("data not found")
	}

	return nil
}

func UpdateDataPosts(data model.UpdatePosts, id string) error {
	// update data to posts
	result := database.ArticleDB.Table(constant.Posts).Where("id = ?", id).Updates(&data)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("data not found")
	}

	return nil
}
