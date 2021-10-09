package controllers

import (
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"appointy_api/models"
	"github.com/revel/revel"
)

type PostController struct {
	*revel.Controller
}

func (c PostController) Index() revel.Result {	
	var (
		posts []models.Post
		err error
	)
	posts, err = models.FindPosts()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(posts)
}

func (c PostController) Show(id string) revel.Result {
	var (
		post models.Post
		err error
		postID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid post id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}
	postID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid post id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}
	post, err = models.FindPost(postID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(post)
}

func (c PostController) Create() revel.Result {
	var (
		post models.Post
		err error
	)
	err = json.NewDecoder(c.Request.GetBody()).Decode(&post)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}
	post, err = models.CreatePost(post)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(post)
}
