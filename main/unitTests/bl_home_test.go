package unittests

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository/mocks"
)

func (s *MyFirstSuite) TestUpdateHomeBL(t provider.T) {
	// Создаем новый контроллер для управления mock-объектами
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks.NewMockHomeRepository(ctrl)

	// Данные для обновления дома
	homeUpdate := pkg.UpdateNameHome{
		NewName: "home1",
	}

	repos := repository.NewRepository(connDB)
	services := service.NewServicesPsql(repos)

	// Создаем экземпляр сервиса с замоканным репозиторием
	homeService := &services.IHome.UpdateHome{
		Repo: mockRepo,
	}

	// Ожидаем, что метод UpdateHome репозитория будет вызван один раз с правильным параметром
	mockRepo.EXPECT().UpdateHome(homeUpdate).Return(nil)

	// Вызываем метод и проверяем результат
	err := homeService.UpdateHome(homeUpdate)
	assert.NoError(t, err)
}

