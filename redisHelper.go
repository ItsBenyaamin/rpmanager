/**
	redis client helper for operation like:
	delete,
	get all available keys,
	get information with specific key,
	check if key exists in db,
	insert (key, value) in db
 */

package main

import "fmt"

func deleteKey(key string) int64 {
	result, _ := client.Del(key).Result()
	return result
}

func getAllKeys() []string {
	allKeys, _ := client.Keys("*").Result()
	return allKeys
}

func getKey(key string) string {
	result, getErr := client.Get(key).Result()
	if getErr != nil{
		fmt.Println("  there is no information with key", key)
		return ""
	}
	return result
}

func keyExists(k string) bool {
	result, _ := client.Exists(k).Result()
	if result == 1 {
		return true
	}
	return false
}

func insertKeyValue(key string, value string) bool {
	_, setErr := client.Set(key, value, 0).Result()
	if setErr != nil {
		showErr(setErr)
		return false
	}
	return true
}