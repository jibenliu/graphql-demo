package model

import "time"

type Hello struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type News struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

var now = time.Now()

var NewsList = []News{
	{
		Id:       1,
		Title:    "Go lang study plan",
		Content:  "Good Good Study, Day Day Up!!",
		Author:   "Vinli Cheung",
		CreateAt: now,
		UpdateAt: now,
	},
	{
		Id:       2,
		Title:    "Gin framwork",
		Content:  "It is very easy to learn!",
		Author:   "Vinli Cheung",
		CreateAt: now,
		UpdateAt: now,
	},
}

func (hello *Hello) Query(id int, name string) (hellos []Hello, err error) {
	allHello := []Hello{
		{0, "vinli"},
		{1, "daisy"},
	}
	for _, v := range allHello {
		if id >= 0 && len(name) <= 0 {
			if v.Id == id {
				hellos = append(hellos, v)
			}
		}

		if id < 0 && len(name) > 0 {
			if v.Name == name {
				hellos = append(hellos, v)
			}
		}

		if id >= 0 && len(name) > 0 {
			if v.Name == name && v.Id == id {
				hellos = append(hellos, v)
			}
		}

		if id < 0 && len(name) <= 0 {
			hellos = append(hellos, v)
		}
	}
	return
}
