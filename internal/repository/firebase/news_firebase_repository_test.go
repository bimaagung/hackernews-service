package firebaserepository_test

import (
	"bytes"
	"errors"
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
			Body: ioutil.NopCloser(bytes.NewBufferString("[36088672, 36088783, 36087442]")),
		}

		expectTopStories := []int{36088672, 36088783, 36087442}

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