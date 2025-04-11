package model

import (
	"test-sharing-vision/constant"
	"time"
)

type Posts struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"column:title;size:200" json:"title"`
	Content     string    `gorm:"column:content;type:text" json:"content"`
	Category    string    `gorm:"column:category;size:100" json:"category"`
	CreatedDate time.Time `gorm:"column:created_date;type:timestamp" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date;type:timestamp" json:"updated_date"`
	Status      string    `gorm:"column:status;size:100;comment:Publish|Draft|Thrash" json:"status"`
}

func (Posts) TableName() string {
	return constant.Posts
}

type UpdatePosts struct {
	Title       string    `gorm:"column:title"`
	Content     string    `gorm:"column:content"`
	Category    string    `gorm:"column:category"`
	UpdatedDate time.Time `gorm:"column:updated_date"`
	Status      string    `gorm:"column:status"`
}

var DummyPosts = []Posts{
	{
		Title:       "Belajar Golang untuk Pemula",
		Content:     "Artikel ini membahas dasar-dasar pemrograman menggunakan bahasa Go (Golang)...",
		Category:    "Programming",
		CreatedDate: time.Date(2024, 2, 12, 10, 30, 0, 0, time.UTC),
		UpdatedDate: time.Date(2024, 3, 1, 9, 15, 0, 0, time.UTC),
		Status:      "Publish",
	},
	{
		Title:       "Panduan Membuat REST API dengan Gin",
		Content:     "Framework Gin memudahkan pengembangan RESTful API di Golang. Dalam artikel ini...",
		Category:    "Web Development",
		CreatedDate: time.Date(2024, 6, 5, 14, 45, 0, 0, time.UTC),
		UpdatedDate: time.Date(2024, 6, 10, 11, 20, 0, 0, time.UTC),
		Status:      "Publish",
	},
	{
		Title:       "Tips Menjaga Kesehatan Mental di Era Digital",
		Content:     "Kesehatan mental menjadi isu penting di era modern, apalagi dengan tekanan media sosial...",
		Category:    "Lifestyle",
		CreatedDate: time.Date(2024, 9, 18, 8, 0, 0, 0, time.UTC),
		UpdatedDate: time.Date(2024, 9, 20, 13, 0, 0, 0, time.UTC),
		Status:      "Draft",
	},
	{
		Title:       "Rekomendasi Film Sci-Fi yang Wajib Ditonton",
		Content:     "Dari Interstellar hingga Blade Runner 2049, ini dia daftar film sci-fi terbaik menurut versi kami...",
		Category:    "Entertainment",
		CreatedDate: time.Date(2024, 11, 3, 19, 30, 0, 0, time.UTC),
		UpdatedDate: time.Date(2024, 11, 5, 21, 0, 0, 0, time.UTC),
		Status:      "Publish",
	},
	{
		Title:       "Apa itu Zero Trust Security?",
		Content:     "Zero Trust adalah pendekatan keamanan TI yang mengharuskan semua pengguna, baik di dalam maupun di luar jaringan...",
		Category:    "Cybersecurity",
		CreatedDate: time.Date(2025, 1, 10, 7, 25, 0, 0, time.UTC),
		UpdatedDate: time.Date(2025, 1, 11, 8, 50, 0, 0, time.UTC),
		Status:      "Thrash",
	},
}
