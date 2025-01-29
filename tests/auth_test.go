package tests

// import (
// 	"context"
// 	"testing"

// 	"github.com/ivanbatutin921/Anti-bruteforce/internal/models"
// 	"github.com/ivanbatutin921/Anti-bruteforce/internal/services"
// 	pb "github.com/ivanbatutin921/Anti-bruteforce/internal/services/protobuf"
// 	db "github.com/ivanbatutin921/Anti-bruteforce/internal/database/postgresql"


// 	"github.com/stretchr/testify/assert"
// )

// func TestAuthorization(t *testing.T) {
// 	// Создаем тестовый сервер
// 	s := &Server{}

// 	// Создаем тестовый контекст
// 	ctx := context.Background()

// 	// Тестовый запрос с существующим логином
// 	req1 := &pb.AuthRequest{
// 		Login:    "test_login",
// 		Password: "test_password",
// 		Ip:       "192.168.1.1",
// 	}

// 	// Тестовый запрос с несуществующим логином
// 	req2 := &pb.AuthRequest{
// 		Login:    "new_login",
// 		Password: "new_password",
// 		Ip:       "192.168.1.2",
// 	}

// 	// Тестовый запрос с заблокированным IP
// 	req3 := &pb.AuthRequest{
// 		Login:    "test_login",
// 		Password: "test_password",
// 		Ip:       "192.168.1.3",
// 	}

// 	// Создаем тестовую запись в базе данных
// 	auth := models.Auth{
// 		Login:    "test_login",
// 		Password: "test_password",
// 		Ip:       "192.168.1.1",
// 	}
// 	db.DB.Create(&auth)

// 	// Тестируем запрос с существующим логином
// 	resp1, err1 := s.Authorization(ctx, req1)
// 	assert.Nil(t, err1)
// 	assert.False(t, resp1.Ok)

// 	// Тестируем запрос с несуществующим логином
// 	resp2, err2 := s.Authorization(ctx, req2)
// 	assert.Nil(t, err2)
// 	assert.True(t, resp2.Ok)

// 	// Тестируем запрос с заблокированным IP
// 	services.CheckIp(req3.Ip)
// 	resp3, err3 := s.Authorization(ctx, req3)
// 	assert.Nil(t, err3)
// 	assert.False(t, resp3.Ok)

// 	// Удаляем тестовую запись из базы данных
// 	db.DB.Delete(&auth)
// }

// func TestAuthorization_TokenBucket(t *testing.T) {
// 	// Создаем тестовый сервер
// 	s := &Server{}

// 	// Создаем тестовый контекст
// 	ctx := context.Background()

// 	// Тестовый запрос
// 	req := &pb.AuthRequest{
// 		Login:    "test_login",
// 		Password: "test_password",
// 		Ip:       "192.168.1.1",
// 	}

// 	// Тестируем несколько запросов подряд
// 	for i := 0; i < 5; i++ {
// 		resp, err := s.Authorization(ctx, req)
// 		if i < 3 {
// 			assert.Nil(t, err)
// 			assert.True(t, resp.Ok)
// 		} else {
// 			assert.Nil(t, err)
// 			assert.False(t, resp.Ok)
// 		}
// 	}
// }