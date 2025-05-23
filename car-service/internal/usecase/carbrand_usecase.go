package usecase

import (
	"context"
	"ecommerce/car-service/internal/domain"
	"ecommerce/car-service/internal/messaging"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type CarBrandUsecase interface {
	Create(brand *domain.CarBrand) error
	GetByID(id string) (*domain.CarBrand, error)
	Update(brand *domain.CarBrand) error
	Delete(id string) error
	List() ([]domain.CarBrand, error)
}

type carBrandUsecase struct {
	repo  domain.CarBrandRepository
	cache *redis.Client
}

func NewCarBrandUsecase(repo domain.CarBrandRepository, cache *redis.Client) CarBrandUsecase {
	return &carBrandUsecase{repo: repo, cache: cache}
}

func (u *carBrandUsecase) Create(brand *domain.CarBrand) error {
	natsConn := messaging.NewNATSConnection()
	defer natsConn.Close()

	event := map[string]interface{}{
		"event": "CarBrandCreated",
		"brand": brand, // Your CarBrand struct
	}
	payload, _ := json.Marshal(event)
	messaging.Publish(natsConn, "carbrand.created", payload)

	return u.repo.Create(brand)
}
func (u *carBrandUsecase) GetByID(id string) (*domain.CarBrand, error) {
	var carBrand domain.CarBrand
	val, err := u.cache.Get(context.Background(), "carbrand:"+id).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &carBrand)
		return &carBrand, nil
	}

	cb, err := u.repo.GetByID(id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get car brand: %v", err)
	}
	if cb == nil {
		return nil, status.Errorf(codes.NotFound, "car brand not found")
	}

	data, _ := json.Marshal(cb)
	u.cache.Set(context.Background(), "carbrand:"+id, data, 0)

	return cb, nil
}
func (u *carBrandUsecase) Update(brand *domain.CarBrand) error {
	return u.repo.Update(brand)
}
func (u *carBrandUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
func (u *carBrandUsecase) List() ([]domain.CarBrand, error) {
	//natsConn := messaging.NewNATSConnection()
	//defer natsConn.Close()
	//go listenForCarBrandCreated(natsConn)
	return u.repo.List()
}

func listenForCarBrandCreated(nc *nats.Conn) {
	messaging.Subscribe(nc, "carbrand.created", func(msg *nats.Msg) {
		log.Printf("Received carbrand.created: %s", string(msg.Data))
		// Process message (e.g., update cache, notify users, etc.)
	})
}
