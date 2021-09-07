package main

import (
	"fmt"
	"test01/muke/retriever/mock"
	real2 "test01/muke/retriever/real"
	"time"
)

const url = "https://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post (url string, form map[string]string) string
}
// 组合接口
type RetrieverPoster interface {
	Retriever
	Poster
	Connect (host string)
}

func session(rp RetrieverPoster) string {
	rp.Post(url, map[string]string{
		"contents": "fake news",
	})
	return rp.Get(url)
}

func post(p Poster)  {
	p.Post(url, map[string]string{
		"name": "123",
		"age": "123",
	})
}

func download(r Retriever) string {
	return r.Get(url)
}
// 查看接口变量
func inspect(r Retriever)  {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents", v.Contents)
	case *real2.Retriever:
		fmt.Println("userAgent", v.UserAgent)
	}
}

type User struct {
	id int
	name string
}

type Manager struct {
	User
	title string
}

func (u *User)toString() {
	fmt.Printf("User->%p, %v", u, u)
}

func (m *Manager)toString()  {
	fmt.Printf("User->%p, %v", m, m)
}

func main()  {
	var r Retriever
	mockRetriever := mock.Retriever{Contents: "this is mock retriever"}
	r = &mockRetriever
	inspect(r)
	r = &real2.Retriever{
		UserAgent: "mo",
		TimeOut: time.Minute,
	}
	inspect(r)
	// type assertion 类型断言

	if realRetriever, ok := r.(*real2.Retriever); ok {
		fmt.Println("Retriever does not have property")
	} else {
		fmt.Println(realRetriever.TimeOut)
	}

	fmt.Println(session(&mockRetriever))
	//download(r)
}