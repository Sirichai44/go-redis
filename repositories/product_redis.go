package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryRedis{db, redisClient}
}

func (r productRepositoryRedis) GetProducts() (products []product, err error) {
	key := "repository::GetProducts"

	// Redis Get
	productJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productJson), &products)
		if err == nil {
			fmt.Println("From Redis")
			return products, nil
		}
	}

	// Database
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	if err != nil {
		fmt.Println("Error find: ", err)
		return nil, err
	}

	// Redis Set
	data, err := json.Marshal(products)
	if err != nil {
		fmt.Println("Error marshal: ", err)
		return nil, err
	}
	fmt.Println("data: ", string(data))

	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		fmt.Println("Error set: ", err)
		return nil, err
	}

	fmt.Println("From Database")
	return products, nil
}
