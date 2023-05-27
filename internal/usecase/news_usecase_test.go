package usecase_test

import (
	"errors"
	"hackernews-service/domain"
	mockfirebaserepository "hackernews-service/internal/mock/repository/firebase"
	"hackernews-service/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewsUC_GetAll(t *testing.T){

	mockNewsFirebaseRepository := new(mockfirebaserepository.NewsFirebaseRepository)

	t.Run("success get all stories", func(t *testing.T) {
		// Arrange 
		story := &domain.Story{
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

		expectStories := []*domain.ResStories{
			{
				ID: 36093995,
				By: "martincmartin",
				Descendants: 24,
				TotalComment: 3,
				Score: 38,
				Time: 1685188865,
				Title: "Downtown San Francisco is at a tipping-point",
				URL: "https://www.economist.com/united-states/2023/05/25/downtown-san-francisco-is-at-a-tipping-point",
			},
		}

		mockNewsFirebaseRepository.On("GetTopStories").Return([]int{36088672}, nil).Once()
		mockNewsFirebaseRepository.On("GetStoryById", mock.AnythingOfType("int")).Return(story, nil).Once()

		uc := usecase.NewNewsUsecase(
			mockNewsFirebaseRepository,
		)

		// Action
		stories, err := uc.GetAll()

		if err != nil {
			t.Errorf("GetAll error: %s", err)
		}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expectStories, stories)

		mockNewsFirebaseRepository.AssertCalled(t, "GetTopStories")
		mockNewsFirebaseRepository.AssertCalled(t, "GetStoryById", 36088672)
	})

	t.Run("failed get all stories", func(t *testing.T) {
		// Arrange 
		mockNewsFirebaseRepository.On("GetTopStories").Return([]int{}, errors.New("error GetTopStories in news firebase repository"))
		
		uc := usecase.NewNewsUsecase(
			mockNewsFirebaseRepository,
		)

		// Action
		_, err := uc.GetAll()

		// Assert
		assert.Error(t, err)
	})

}