package main

import (
	"fmt"
)

// | => top - bottom
// | => left - right
// L => top - left
// J => top - right
// 7 => bottom - left
// F => bottom - right
// . => empty
// S => start

type Node struct {
	coordinates    [2]int
	connectionType ConnectionType
}

func main() {
	initMap("puzzleInput.txt")
	startingNode := Map.getStartingNode()
	nodeList := []Node{startingNode}
	loop := getLoop(nodeList[0])
	for _, node := range loop {
		fmt.Println(node)
	}
	maxDistance := getMaxDistanceFromNode(loop)
	fmt.Println(maxDistance)
}

func getLoop(startingNode Node) []Node {
	loop := []Node{startingNode}

	currentIndex := 1
	for {
		parent := loop[currentIndex-1]
		parentRow := parent.coordinates[0]
		parentColumn := parent.coordinates[1]
		parentConnectionType := parent.connectionType

		isConnectedLeft, newLeftNode := getLeftConnection(parentRow, parentColumn)
		canMoveToConnection := parentConnectionType == HorizontalPipe || parentConnectionType == JPipe || parentConnectionType == SevenPipe || parentConnectionType == StartingPipe
		if isConnectedLeft && canMoveToConnection {
			isNodeAlreadyAdded := isNodeAlreadyAddedToLoop(loop, newLeftNode)
			if isNodeAlreadyAdded == false {
				loop = append(loop, newLeftNode)
				currentIndex += 1
				continue
			}
		}

		isConnectedRight, newRightNode := getRightConnection(parentRow, parentColumn)
		canMoveToConnection = parentConnectionType == HorizontalPipe || parentConnectionType == FPipe || parentConnectionType == LPipe || parentConnectionType == StartingPipe
		if isConnectedRight && canMoveToConnection {
			isNodeAlreadyAdded := isNodeAlreadyAddedToLoop(loop, newRightNode)
			if isNodeAlreadyAdded == false {
				loop = append(loop, newRightNode)
				currentIndex += 1
				continue
			}
		}

		isConnectedBottom, newBottomNode := getBottomConnection(parentRow, parentColumn)
		canMoveToConnection = parentConnectionType == FPipe || parentConnectionType == SevenPipe || parentConnectionType == VerticalPipe || parentConnectionType == StartingPipe
		if isConnectedBottom && canMoveToConnection {
			isNodeAlreadyAdded := isNodeAlreadyAddedToLoop(loop, newBottomNode)
			if isNodeAlreadyAdded == false {
				loop = append(loop, newBottomNode)
				currentIndex += 1
				continue
			}
		}

		isConnectedTop, newTopNode := getTopConnection(parentRow, parentColumn)
		canMoveToConnection = parentConnectionType == VerticalPipe || parentConnectionType == JPipe || parentConnectionType == LPipe || parentConnectionType == StartingPipe
		if isConnectedTop && canMoveToConnection {
			isNodeAlreadyAdded := isNodeAlreadyAddedToLoop(loop, newTopNode)
			if isNodeAlreadyAdded == false {
				loop = append(loop, newTopNode)
				currentIndex += 1
				continue
			}
		}

		break
	}

	return loop
}

func getMaxDistanceFromNode(loop []Node) int {
	numberOfNodesConnectedTo := len(loop)
	maxDistance := (numberOfNodesConnectedTo / 2) + numberOfNodesConnectedTo%2
	return maxDistance
}

func isNodeAlreadyAddedToLoop(loop []Node, currentNode Node) bool {
	for _, nodeInLoop := range loop {
		if nodeInLoop.coordinates[0] == currentNode.coordinates[0] && nodeInLoop.coordinates[1] == currentNode.coordinates[1] {
			return true
		}
	}

	return false
}
