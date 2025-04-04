package usecase

import (
	"context"
	"forgetful-guard/common/rdb"
	"forgetful-guard/internal/domain"
	"forgetful-guard/internal/domain/models"
	"forgetful-guard/internal/interface/oapi"
	rMock "forgetful-guard/internal/interface/repository/mock"
	uMock "forgetful-guard/internal/usecase/mock"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
)

func TestCreateGeofence_Normal(t *testing.T) {
	mockDB, smock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()

	rdb.DB = mockDB
	smock.ExpectBegin()
	smock.ExpectCommit()

	ctx := context.Background()
	title := "TestTask"
	userID := uint64(100)
	deviceToken := "dummy-token"
	polygon := [][][]float64{{{1.0, 2.0}, {2.0, 3.0}, {3.0, 1.0}, {1.0, 2.0}}}
	geofence := &domain.Geofence{Polygon: polygon}

	task := &models.Task{
		Title:  title,
		UserID: userID,
	}

	req := &oapi.Geofence{
		Title:       title,
		UserID:      userID,
		DeviceToken: deviceToken,
		Polygon:     polygon,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := uMock.NewMockUsecaseService(ctrl)
	mockService.EXPECT().CreateTask(ctx, gomock.Any(), task).Return(nil)

	mockRepository := rMock.NewMockGeofenceRepository(ctrl)
	mockRepository.EXPECT().PutGeofence(ctx, geofence).Return(nil)
	mockRepository.EXPECT().PutDeviceToken(userID, deviceToken).Return(nil)

	u := NewUsecase(mockService, mockRepository)
	if err := u.CreateGeofence(ctx, req); err != nil {
		t.Errorf("Unexpected results: %v", err)
	}
}
