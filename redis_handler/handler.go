package redis_handler
import "github.com/simonz05/godis/exp"
type RedisHandler struct {
  list_key string
  writer chan string
  client *redis.Client
}
func NewRedisHandler(address_arg string, list_key_arg string) (r *RedisHandler) {
  client := redis.NewClient(address_arg, 3, "")
  channel := make(chan string)
  return &RedisHandler{list_key: list_key_arg, writer: channel, client: client}
}
func (r *RedisHandler) Push(msg string) (rep *redis.Reply, err error) {
  rep, err = r.client.Call("publish", r.list_key, msg)
  return 
}






