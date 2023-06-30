package main

import (
	"errors"
	"fmt"
)

func GetUserDetails(userId string) (User, error) {
	nodeForUser := FindNodeForUser(userId)
	allUserData := NodeUserCentralDB[nodeForUser.nodeIp]

	for _, user := range allUserData {
		if user.id == userId {
			return user, nil
		}
	}
	return User{}, errors.New("no user found")
}

func InsertUser(user User) {
	// find the owner node for user
	nodeForUser := FindNodeForUser(user.id)
	// fmt.Printf("Assigning user with id %v to node %v\n", user.id, nodeForUser)

	// insert new user in owner db
	NodeUserCentralDB[nodeForUser.nodeIp] = append(NodeUserCentralDB[nodeForUser.nodeIp], user)
}

func FindNodeForUser(userId string) Node {
	var userIdHashValue = getHash(userId)
	fmt.Printf("userid: %v hash is %v. \n", userId, userIdHashValue)
	return findDataOwnerNodeForHash(userIdHashValue)
}
