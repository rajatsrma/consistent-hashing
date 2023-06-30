package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"

	"github.com/google/uuid"
)

var nodeDetails = []Node{
	{
		nodeIp:   "10.131.0.1",
		nodeHash: 15,
	},
	{
		nodeIp:   "10.131.0.2",
		nodeHash: 120,
	},
	{
		nodeIp:   "10.131.0.3",
		nodeHash: 300,
	},
}

var NodeUserCentralDB = make(map[string][]User)

func main() {
	// add some users
	for i := 0; i < 5235; i++ {
		InsertUser(
			User{
				id:    uuid.NewString(),
				name:  strconv.Itoa(i),
				email: strconv.Itoa(i) + "@gmail.com",
			},
		)
	}

	initialCentralNodeUserMap := make(map[string][]User)
	for k, v := range NodeUserCentralDB {
		initialCentralNodeUserMap[k] = v
	}

	// print current node to user db
	fmt.Printf("\nnode user db = %v\n", NodeUserCentralDB)

	// remove node with hash 300 and ip
	AddNodeToCluster(
		Node{
			nodeIp:   "10.131.0.4",
			nodeHash: 230,
		},
	)

	// after adding new node with hash 240 all data between 120-240 should point to new node
	fmt.Println("===========================After Adding node 4====================")
	fmt.Printf("node user db = %v\n", NodeUserCentralDB)
	fmt.Printf("node list= %v\n", nodeDetails)

	inprocessCentralNodeUserMap := make(map[string][]User)
	for k, v := range NodeUserCentralDB {
		inprocessCentralNodeUserMap[k] = v
	}

	// remove node with hash 300 and ip
	RemoveNodeFromCluster(
		Node{
			nodeIp:   "10.131.0.4",
			nodeHash: 230,
		},
	)

	// after removing all data pointed to node 4 will point to node 1
	fmt.Println("===========================After removing node 4====================")
	fmt.Printf("node user db = %v\n", NodeUserCentralDB)
	fmt.Printf("node list= %v\n", nodeDetails)

	finalCentralNodeUserMap := make(map[string][]User)
	for k, v := range NodeUserCentralDB {
		finalCentralNodeUserMap[k] = v
	}

	for k := range initialCentralNodeUserMap {
		sort.Slice(initialCentralNodeUserMap[k], func(i, j int) bool {
			return initialCentralNodeUserMap[k][i].id < initialCentralNodeUserMap[k][j].id
		})
		sort.Slice(finalCentralNodeUserMap[k], func(i, j int) bool {
			return finalCentralNodeUserMap[k][i].id < finalCentralNodeUserMap[k][j].id
		})
	}

	var initialVsFinalData = reflect.DeepEqual(initialCentralNodeUserMap, finalCentralNodeUserMap)

	fmt.Println("+++++++++++++++comparison resuults start+++++++++++++++++++++")
	fmt.Printf("initial == final: %v\n", initialVsFinalData)
	fmt.Println("+++++++++++++++comparison resuults end+++++++++++++++++++++")
}
