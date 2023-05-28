package mockfirebaserepository

import (
	"hackernews-service/domain"

	"github.com/stretchr/testify/mock"
)

type NewsFirebaseRepository struct {
	mock.Mock
}

func (m *NewsFirebaseRepository) GetTopStories()([]int, error){
	ret := m.Called()

	var r0 []int

	if rf, ok := ret.Get(0).(func() []int); ok {
		r0 = rf()
	}else {
		r0 = ret.Get(0).([]int)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *NewsFirebaseRepository) GetStoryById(id int)(*domain.Story, error){
	ret := m.Called(id)

	var r0 *domain.Story

	if rf, ok := ret.Get(0).(func(int) *domain.Story); ok {
		r0 = rf(id)
	}else {
		r0 = ret.Get(0).(*domain.Story)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}
