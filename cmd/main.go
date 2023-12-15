package main

import "os"

var (
	redis_addr = os.Getenv("REDIS_ADDR")
	redis_pass = os.Getenv("REDIS_PASS")
	redis_db = os.Getenv("REDIS_DB")
)

func init() {
	
}
func main() {}
