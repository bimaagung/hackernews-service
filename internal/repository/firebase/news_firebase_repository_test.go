package firebaserepository_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hackernews-service/domain"
	mockhttpclient "hackernews-service/internal/mock/http_client"
	firebaserepository "hackernews-service/internal/repository/firebase"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
);

func TestNewsFirebaseRepository_GetTopStories(t *testing.T) {
	mockClient := &mockhttpclient.MockHTTPClient{}
	
	t.Run("success get top stories", func(t *testing.T) {
		// Arrange
		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Body: ioutil.NopCloser(bytes.NewBufferString("[36093995, 36088783, 36087442]")),
		}

		expectTopStories := []int{36093995, 36088783, 36087442}

		mockClient.On("Get", "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty").Return(mockResponse, nil)

		repo := firebaserepository.NewNewsFirebaseRepository(mockClient)

		// Action
		topStories, err := repo.GetTopStories()
		if err != nil {
			t.Errorf("GetTopStories error: %s", err)
		}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expectTopStories, topStories)
		mockClient.AssertCalled(t, "Get", "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	})

	t.Run("failed can't get stories from firebase", func(t *testing.T) {
		// Arrange
		mockClient.On("Get", "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty").Return(nil, errors.New("failed fetch firebase stories"))

		repo := firebaserepository.NewNewsFirebaseRepository(mockClient)

		// Action
		_, err := repo.GetTopStories()

		// Assert
		require.Error(t, err)
		mockClient.AssertCalled(t, "Get", "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	})
}

func TestNewsFirebaseRepository_GetStoryById(t *testing.T) {
	mockClient := &mockhttpclient.MockHTTPClient{}

	expectTopStories := &domain.Story{
		ID: 36093995,
		By: "martincmartin",
		Descendants: 24,
		Kids: []int{36094267, 36094186, 36094203},
		Score: 38,
		Time: 1685188865,
		Title: "Downtown San Francisco is at a tipping-point",
		Type: "story",
		URL: "https://www.economist.com/united-states/2023/05/25/downtown-san-francisco-is-at-a-tipping-point",
	}
	
	t.Run("success getStoryById", func(t *testing.T) {
		// Arrange
		buffer := new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(expectTopStories)

		if err != nil {
			t.Errorf("Error encoding story: %s", err)
			
		}

		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Body: ioutil.NopCloser(bytes.NewBuffer(buffer.Bytes())),
		}

		url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", expectTopStories.ID)
		mockClient.On("Get", url).Return(mockResponse, nil)

		repo := firebaserepository.NewNewsFirebaseRepository(mockClient)

		// Action
		story, err := repo.GetStoryById(expectTopStories.ID)
		
		if err != nil {
			t.Errorf("GetTopStories error: %s", err)
		}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expectTopStories, story)

		mockClient.AssertCalled(t, "Get", url)
	})

	t.Run("failed can't getStoryById from firebase", func(t *testing.T) {
		// Arrange
		url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", expectTopStories.ID)
		mockClient.On("Get", url).Return(nil, errors.New("failed fetch firebase story"))
	
		repo := firebaserepository.NewNewsFirebaseRepository(mockClient)

		// Action
		_, err := repo.GetStoryById(expectTopStories.ID)

		// Assert
		require.Error(t, err)

		mockClient.AssertCalled(t, "Get", url)
	})
}

func TestNewsFirebaseRepository_GetCommentById(t *testing.T) {
	mockClient := &mockhttpclient.MockHTTPClient{}

	expectTopStories := &domain.Comment{
		ID: 36094267,
		By: "spaniard89277",
		Kids: []int{36094432, 36094350, 36094381},
		Parent: 36093995,
		Text: "Make housing affordable with denser mixed-use zoning",
		Time: 1685191208,
		Type: "comment",
	}
	
	t.Run("success getCommentById", func(t *testing.T) {
		// Arrange
		buffer := new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(expectTopStories)

		if err != nil {
			t.Errorf("Error encoding story: %s", err)
			
		}

		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Body: ioutil.NopCloser(bytes.NewBuffer(buffer.Bytes())),
		}

		url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", expectTopStories.ID)
		mockClient.On("Get", url).Return(mockResponse, nil)

		repo := firebaserepository.NewNewsFirebaseRepository(mockClient)

		// Action
		story, err := repo.GetStoryById(expectTopStories.ID)

		if err != nil {
			t.Errorf("GetTopStories error: %s", err)
		}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expectTopStories, story)

		mockClient.AssertCalled(t, "Get", url)
	})

	t.Run("failed can't getStoryById from firebase", func(t *testing.T) {
		// Arrange
		url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", expectTopStories.ID)
		mockClient.On("Get", url).Return(nil, errors.New("failed fetch firebase story"))
	
		repo := firebaserepository.NewNewsFirebaseRepository(mockClient)

		// Action
		_, err := repo.GetStoryById(expectTopStories.ID)

		// Assert
		require.Error(t, err)

		mockClient.AssertCalled(t, "Get", url)
	})
}