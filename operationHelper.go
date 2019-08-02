package main

import (
	"fmt"
	"strings"
)

const replaceText string = "this name is already exists! are you want to replace?(y/n) "


/**
	first of all we scan key(name) from user and check if already exist
	if key exist showKeyExistWarning() be called
	and if key not exist so scanUsernameAndPassword() called
*/
func setOperation() {
	var name string

	fmt.Print("- enter the url, appName or anything as key: ")
	_, nameScanErr := fmt.Scanf("%s", &name)
	showErr(nameScanErr)

	if keyExists(name) {
		showKeyExistWarning(name)
		return
	}
	scanUsernameAndPassword(name)
}

/**
	show warning with title key exist and ask user to should replace
	if key is exist so ask to replace or want to choose new key for it
*/
func showKeyExistWarning(name string) {
	fmt.Print(replaceText)
	var replaceResult string
	_, replaceErr := fmt.Scanf("%s", &replaceResult)
	showErr(replaceErr)

	if replaceResult == "y" {
		scanUsernameAndPassword(name)
	}else if replaceResult == "n" {
		setOperation()
	}else {
		showKeyExistWarning(name)
	}
}

/**
	scan Username and Password as second step to put information to database
	with format (key: $key,username: $username,password: $password)
*/
func scanUsernameAndPassword(name string) {
	var username string
	var password string

	fmt.Print("- enter username: ")
	_, usernameScanErr := fmt.Scanf("%s", &username)
	showErr(usernameScanErr)

	fmt.Print("- enter password: ")
	_, passwordScanErr := fmt.Scanf("%s", &password)
	showErr(passwordScanErr)

	formattedValue := fmt.Sprintf("Key: %s,Username: %s,Password: %s", name, username, password)

	insertResult := insertKeyValue(name, formattedValue)

	if insertResult {
		fmt.Println("information inserted, well done!")
	}
	scanOperationFromUser()
}

/**
	scan key from user and call deleteKey() from redisHelper for delete key
	if deleteKey() return 1 that mean delete successful
	if deleteKey() return 0 that mean delete unsuccessful
 */
func delOperation() {
	fmt.Print("please enter key to delete: ")

	var key string
	_,keyScanErr := fmt.Scanf("%s", &key)
	showErr(keyScanErr)

	if deleteKey(key) == 1 {
		fmt.Println("delete successful.")
	}else {
		fmt.Println("there is no key found with name", key)
	}
	scanOperationFromUser()
}

/*
	get specific information with the chosen key at the first time when
	user insert the data to database
*/
func getOperation() {
	var key string

	fmt.Print("enter the key you want to get information: ")
	_, scanErr := fmt.Scanf("%s", &key)
	showErr(scanErr)

	result := getKey(key)
	if result == "" {
		scanOperationFromUser()
		return
	}

	fmt.Println("information for key", key, ":")

	resultSplit := strings.Split(result, ",")

	fmt.Println("  Username:", resultSplit[1])
	fmt.Println("  Password:", resultSplit[2])

	scanOperationFromUser()
}

/*
	get all keys inserted in redis database and print them to user
	it's for case user forgot what key choose for specific information
	so it can be found in this list
*/

func allKeysOperation() {
	allKeys := getAllKeys()
	fmt.Println("Keys List: ")
	if len(allKeys) == 0 {
		fmt.Println("  there is nothing to show!")
	}else {
		for index, value := range allKeys{
			fmt.Println(fmt.Sprintf("  %d: %s", index + 1, value))
		}
	}
	scanOperationFromUser()
}