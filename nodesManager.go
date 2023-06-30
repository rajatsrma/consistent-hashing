package main

import (
	"errors"
	"fmt"
	"sort"
)

func findDataOwnerNodeForHash(hashValue int) Node {
	for _, node := range nodeDetails {
		if node.nodeHash >= hashValue {
			return node
		}
	}

	// if we couldn't return from earlier loop, means that first node is the owner
	return nodeDetails[0]
}

func AddNodeToCluster(newNode Node) {
	// add node to node cluster on correct position
	var nextNodeToNewNode = Node{}
	var foundNextNodeToNewNode = false
	for _, node := range nodeDetails {
		if node.nodeHash > newNode.nodeHash {
			nextNodeToNewNode = node
			foundNextNodeToNewNode = true
			break
		}
	}

	if !foundNextNodeToNewNode {
		nextNodeToNewNode = nodeDetails[0]
	}

	fmt.Printf("nextNode: %v \n", nextNodeToNewNode)

	// iterate over all the keys in next node to new node and insert them in new node if condition satisfied
	var usersOnNewNode []User
	var usersOnOldNode []User
	for _, user := range NodeUserCentralDB[nextNodeToNewNode.nodeIp] {
		var userIdHashValue = getHash(user.id)
		if userIdHashValue <= newNode.nodeHash {
			usersOnNewNode = append(usersOnNewNode, user)
		} else {
			usersOnOldNode = append(usersOnOldNode, user)
		}
	}

	NodeUserCentralDB[nextNodeToNewNode.nodeIp] = usersOnOldNode
	NodeUserCentralDB[newNode.nodeIp] = usersOnNewNode

	// add new node to node details list and sort basis hash score
	nodeDetails = append(nodeDetails, newNode)
	sort.Slice(nodeDetails, func(i, j int) bool {
		return nodeDetails[i].nodeHash < nodeDetails[j].nodeHash
	})
}

func RemoveNodeFromCluster(node Node) {
	// current total nodes
	var totalNodes = len(nodeDetails)

	// find node to be removed idx and new owner node
	var nodeToRemoveIdx, err = findNodeIndex(node)
	if err != nil {
		fmt.Printf("could not find node %v to remove from cluster", node)
	}

	var newOwnerNodeIdx = 0
	if (nodeToRemoveIdx + 1) < totalNodes {
		newOwnerNodeIdx = nodeToRemoveIdx + 1
	}

	// transfer the data of old node to new node
	var newOwnerNode = nodeDetails[newOwnerNodeIdx]
	NodeUserCentralDB[newOwnerNode.nodeIp] = append(NodeUserCentralDB[newOwnerNode.nodeIp], NodeUserCentralDB[node.nodeIp]...)

	// remove current Node nodeUserDB
	delete(NodeUserCentralDB, node.nodeIp)

	// remove current node from Node details
	nodeDetails[nodeToRemoveIdx] = nodeDetails[totalNodes-1]
	nodeDetails = nodeDetails[:totalNodes-1]

	// sort the node db with hash id
	sort.Slice(nodeDetails, func(i, j int) bool {
		return nodeDetails[i].nodeHash < nodeDetails[j].nodeHash
	})
}

func GetAllNodeDetails() []Node {
	return nodeDetails
}

func findNodeIndex(node Node) (int, error) {
	for idx, currNode := range nodeDetails {
		if currNode.nodeHash == node.nodeHash {
			return idx, nil
		}
	}
	return -1, errors.New("no such node found")
}
