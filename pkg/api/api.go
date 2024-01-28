package api

import (
	"context"
	"github.com/rishabh-22/blogapp-grpc/blogpost"
	"github.com/rishabh-22/blogapp-grpc/pkg/model"
)

var data = model.NewAutoIncrementMap()

type MyBlogPostServer struct {
	blogpost.UnimplementedBlogServiceServer
}

func (s MyBlogPostServer) Create(ctx context.Context, req *blogpost.CreatePost) (*blogpost.Response, error) {
	key := data.Add(req.Title, req.Content, req.Author, req.PublicationDate, req.Tags)
	return &blogpost.Response{
		PostID:          key,
		Title:           data.Data[key].Title,
		Content:         data.Data[key].Content,
		Author:          data.Data[key].Author,
		PublicationDate: data.Data[key].PublicationDate,
		Tags:            data.Data[key].Tags,
	}, nil
}

func (s MyBlogPostServer) Update(ctx context.Context, req *blogpost.UpdatePost) (*blogpost.Response, error) {
	_, exists := data.GetValueForKey(req.PostID)
	if exists {
		data.Update(req.PostID, req.Title, req.Content, req.Author, req.Tags)
		return &blogpost.Response{
			PostID:          req.PostID,
			Title:           data.Data[req.PostID].Title,
			Content:         data.Data[req.PostID].Content,
			Author:          data.Data[req.PostID].Author,
			PublicationDate: data.Data[req.PostID].PublicationDate,
			Tags:            data.Data[req.PostID].Tags,
		}, nil
	} else {
		return &blogpost.Response{
			PostID:          req.PostID,
			Title:           "No data Found to be updated",
			Content:         "No data Found to be updated",
			Author:          "No data Found to be updated",
			PublicationDate: "No data Found to be updated",
			Tags:            "No data Found to be updated",
		}, nil
	}

}

func (s MyBlogPostServer) Read(ctx context.Context, req *blogpost.ReadPost) (*blogpost.Response, error) {
	_, exists := data.GetValueForKey(req.PostID)
	if exists {
		return &blogpost.Response{
			PostID:          req.PostID,
			Title:           data.Data[req.PostID].Title,
			Content:         data.Data[req.PostID].Content,
			Author:          data.Data[req.PostID].Author,
			PublicationDate: data.Data[req.PostID].PublicationDate,
			Tags:            data.Data[req.PostID].Tags,
		}, nil
	} else {
		return &blogpost.Response{
			PostID:          req.PostID,
			Title:           "No data Found",
			Content:         "No data Found",
			Author:          "No data Found",
			PublicationDate: "No data Found",
			Tags:            "No data Found",
		}, nil
	}
}

func (s MyBlogPostServer) Delete(ctx context.Context, req *blogpost.ReadPost) (*blogpost.Message, error) {
	_, exists := data.GetValueForKey(req.PostID)
	var res string
	if exists {
		res = "success"
		data.Delete(req.PostID)
	} else {
		res = "failure"
	}
	return &blogpost.Message{
		Body: res,
	}, nil
}
