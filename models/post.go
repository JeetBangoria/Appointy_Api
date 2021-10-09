package models

import (
	"appointy_api/models/db"
	"time"
	"gopkg.in/mgo.v2/bson"
)



type Post struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Content string `json:"content" bson:"content"`//Caption
	ImageURL  string        `json:"imageURL" bson:"imageURL"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	User_ID bson.ObjectId `json:"userid" bson:"_userid"`
}

func newPostCollection() *db.Collection {
	return db.NewCollectionSession("posts")
}

func CreatePost(post Post) (Post, error) {
	var (
	   err error
	)
	c := newPostCollection()
	defer c.Close()

	post.ID = bson.NewObjectId()
	post.CreatedAt = time.Now()
	err = c.Session.Insert(&post)
	if err != nil {
	   return post,err
	}
	 return post, err
}


func FindPosts() ([]Post, error) {
	var (
	   err error
	   posts []Post
   )
	c := newPostCollection()
	defer c.Close()

   err = c.Session.Find(nil).Sort("-published_at").All(&posts)
   if err != nil {
	   return posts,err
	}
   return posts, err
}

func FindPost(id bson.ObjectId) (Post, error) {
	var (
	   err error
	   post Post
   )

	c := newPostCollection()
	defer c.Close()

   err = c.Session.FindId(id).One(&post)
   if err != nil {
	   return post,err
	}
   return post, err
}