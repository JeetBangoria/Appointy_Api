package models
import (
	"appointy_api/models/db"
	"time"
	"gopkg.in/mgo.v2/bson"
)



type Post struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Title string `json:"title" bson:"title"`
	SlugUrl string `json:"slug_url" bson:"slug_url"`
	Content string `json:"content" bson:"content"`
	PublishedAt time.Time `json:"published_at" bson:"published_at"` // check this
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
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