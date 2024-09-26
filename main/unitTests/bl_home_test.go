package unittests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func (s *MyFirstSuite) TestUpdateHome(t provider.T) {
	// Создаем новый контроллер для управления mock-объектами
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks.NewMockHomeRepository(ctrl)

	// Данные для обновления дома
	homeUpdate = pkg.UpdateNameHome{
		NewName: "home1",
	},

	// Создаем экземпляр сервиса с замоканным репозиторием
	homeService := &service.HomeService{
		Repo: mockRepo,
	}

	// Ожидаем, что метод UpdateHome репозитория будет вызван один раз с правильным параметром
	mockRepo.EXPECT().UpdateHome(homeUpdate).Return(nil)

	// Вызываем метод и проверяем результат
	err := homeService.UpdateHome(homeUpdate)
	assert.NoError(t, err)
}

