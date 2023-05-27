package usecase_test

import (
	"hackernews-service/domain"
	"hackernews-service/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewsUC_GetAll(t *testing.T){
	payload := &domain.Story{} 
	
	t.Run("success", func(t *testing.T) {
		uc := usecase.NewNewsUsecase()

		story, err := uc.GetAll(payload)

		assert.NoError(t, err)
		assert.NotNil(t, story)

	})

	t.Run("failed", func(t *testing.T) {
		uc := usecase.NewNewsUsecase()

		story, err := uc.GetAll(payload)

		assert.NoError(t, err)
		assert.NotNil(t, story)

	})
}