package domain

import "net/http"

type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

type ResStory struct {
	By          	string 		 `json:"by"`
	Descendants 	int    		 `json:"descendants"`
	ID          	int    		 `json:"id"`
	Comments        []*ResComment `json:"comments"`
	Score       	int    		 `json:"score"`
	Time        	int    		 `json:"time"`
	Title       	string 		 `json:"title"`
	Type        	string 		 `json:"type"`
	URL         	string 		 `json:"url"`
}

type Comment struct {
	By          string 	`json:"by"`
	ID          int    	`json:"id"`
	Kids        []int  	`json:"kids"`
	Parent      int 	`json:"parent"`
	Text        string 	`json:"text"`
	Time        int    	`json:"time"`
	Type        string 	`json:"type"`
}

type ResComment struct {
	By          	string 			`json:"by"`
	ID          	int    			`json:"id"`
	Comments 		[]*ResComment 	`json:"comments"`
	Parent      	int 			`json:"parent"`
	Text        	string 			`json:"text"`
	Time        	int    			`json:"time"`
	Type        	string 			`json:"type"`
}

type ResStories struct {
	By          		string 	`json:"by"`
	Descendants 		int    	`json:"descendants"`
	ID          		int    	`json:"id"`
	TotalComment        int  	`json:"total_comment"`
	Score       		int    	`json:"score"`
	Time        		int    	`json:"time"`
	Title       		string 	`json:"title"`
	URL         		string 	`json:"url"`
}

type NewsUsecase interface {
	GetAll() ([]*ResStories, error)
	GetStoryById(id int)(*ResStory, error)
	GetCommentById(id int)(*ResComment, error)
}

type NewsFirebaseRepository interface {
	GetTopStories() ([]int, error)
	GetStoryById(id int) (*Story, error)
	GetCommentById(id int) (*Comment, error)
}

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}