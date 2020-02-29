package main

import "fmt"
import "time"
import "learngo/05ducktype/mock"
import "learngo/05ducktype/real"

const URL = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

// func retrive(r Retriever) string {
// 	return r.Get(URL)
// }

type Poster interface {
	Post(url string, form map[string]string) string
}

// func post(poster Poster) {
// 	poster.Post(URL, map[string]string{"name": "ccmouse", "course": "golang"})
// }

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(URL, map[string]string{"contents": "another"})
	return s.Get(URL)
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main() {
	var r Retriever

	r = &mock.Retriever{"test again"}
	inspect(r)
	mockRetriever := r.(*mock.Retriever)
	fmt.Println(mockRetriever.Contents)
	mrp := mockRetriever

	r = &real.Retriever{UserAgent: "Mozilla/5.0", Timeout: time.Minute}
	inspect(r)
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.Timeout)
	rrp := realRetriever

	fmt.Println()
	fmt.Println(session(mrp))
	fmt.Println(session(rrp))
}
