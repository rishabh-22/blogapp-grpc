package main

import (
	"context"
	"github.com/rishabh-22/blogapp-grpc/blogpost"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Blog struct {
	Title           string
	Content         string
	Author          string
	PublicationDate string
	Tags            string
}
type AutoIncrementMap struct {
	data    map[int64]Blog
	counter int64
}

func NewAutoIncrementMap() *AutoIncrementMap {
	return &AutoIncrementMap{
		data:    make(map[int64]Blog),
		counter: 0,
	}
}

func (m *AutoIncrementMap) Add(value1 string, value2 string, value3 string, value4 string, value5 string) int64 {
	obj := Blog{
		Title:           value1,
		Content:         value2,
		Author:          value3,
		PublicationDate: value4,
		Tags:            value5,
	}
	m.counter++
	m.data[m.counter] = obj
	return m.counter
}
func (m *AutoIncrementMap) Update(key int64, value1 string, value2 string, value3 string, value4 string) {
	if obj, exists := m.data[key]; exists {
		obj.Title = value1
		obj.Content = value2
		obj.Author = value3
		obj.Tags = value4
		m.data[key] = obj
	}
}

func (m *AutoIncrementMap) Delete(key int64) {
	delete(m.data, key)
}

func (m *AutoIncrementMap) GetValueForKey(key int64) (Blog, bool) {
	obj, exists := m.data[key]
	return obj, exists
}

var data = NewAutoIncrementMap()

type myBlogPostServer struct {
	blogpost.UnimplementedBlogServiceServer
}

func (s myBlogPostServer) Create(ctx context.Context, req *blogpost.CreatePost) (*blogpost.Response, error) {
	key := data.Add(req.Title, req.Content, req.Author, req.PublicationDate, req.Tags)
	return &blogpost.Response{
		PostID:          key,
		Title:           data.data[key].Title,
		Content:         data.data[key].Content,
		Author:          data.data[key].Author,
		PublicationDate: data.data[key].PublicationDate,
		Tags:            data.data[key].Tags,
	}, nil
}

func (s myBlogPostServer) Update(ctx context.Context, req *blogpost.UpdatePost) (*blogpost.Response, error) {
	_, exists := data.GetValueForKey(req.PostID)
	if exists {
		data.Update(req.PostID, req.Title, req.Content, req.Author, req.Tags)
		return &blogpost.Response{
			PostID:          req.PostID,
			Title:           data.data[req.PostID].Title,
			Content:         data.data[req.PostID].Content,
			Author:          data.data[req.PostID].Author,
			PublicationDate: data.data[req.PostID].PublicationDate,
			Tags:            data.data[req.PostID].Tags,
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

func (s myBlogPostServer) Read(ctx context.Context, req *blogpost.ReadPost) (*blogpost.Response, error) {
	_, exists := data.GetValueForKey(req.PostID)
	if exists {
		return &blogpost.Response{
			PostID:          req.PostID,
			Title:           data.data[req.PostID].Title,
			Content:         data.data[req.PostID].Content,
			Author:          data.data[req.PostID].Author,
			PublicationDate: data.data[req.PostID].PublicationDate,
			Tags:            data.data[req.PostID].Tags,
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

func (s myBlogPostServer) Delete(ctx context.Context, req *blogpost.ReadPost) (*blogpost.Message, error) {
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

func main() {

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myBlogPostServer{}
	blogpost.RegisterBlogServiceServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
