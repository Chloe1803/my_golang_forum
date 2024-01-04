package forum

import "time"

type ErrorInput struct {
	ErrorMessage bool
}

type Index_data struct {
	Connected   bool
	Index_posts []Post
	Categories  []Category
}

type Post struct {
	ID             int
	Title          string
	Author         string
	Timestamp      time.Time
	Content        string
	ImageAvailable bool
	ImageURL       string
	ImageType      string
	Categories     []Category
	Nb_like        int
	Nb_dislike     int
	Nb_comments    int
}

type Category struct {
	ID   int
	Name string
}

type VP_data struct {
	Connected bool
	Post      Post
	Comments  []Comment
	Like_status    int
}


type Comment struct {
	Connected   bool
	ID          int
	Author      string
	User_id     string
	Time_stamp  time.Time
	Content     string
	Nb_like     int
	Nb_dislike  int
	Like_status int
}
