import (
	"github.com/redis/go-redis/v8"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func CacheProduct(product models.Product) {
	key := fmt.Sprintf("product:%d", product.ID)
	redisClient.Set(ctx, key, product, 0)
}

func GetCachedProduct(id int) *models.Product {
	key := fmt.Sprintf("product:%d", id)
	cached, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	var product models.Product
	json.Unmarshal([]byte(cached), &product)
	return &product
}
