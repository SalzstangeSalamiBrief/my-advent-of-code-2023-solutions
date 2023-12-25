package main

type ConnectionType string

const (
	HorizontalPipe ConnectionType = "-"
	VerticalPipe   ConnectionType = "|"
	LPipe          ConnectionType = "L"
	JPipe          ConnectionType = "J"
	FPipe          ConnectionType = "F"
	SevenPipe      ConnectionType = "7"
	None           ConnectionType = "."
	StartingPipe   ConnectionType = "S"
)

func getTopConnection(rowNumber int, columnNumber int) (isConnected bool, newNode Node) {
	aboveRowNumber := rowNumber - 1
	if aboveRowNumber < 0 {
		return isConnected, newNode
	}

	cellAbove := Map[aboveRowNumber][columnNumber]
	newNode = Node{
		coordinates: [2]int{aboveRowNumber, columnNumber},
	}

	switch cellAbove {
	case string(VerticalPipe):
		newNode.connectionType = VerticalPipe
		isConnected = true
	case string(SevenPipe):
		newNode.connectionType = SevenPipe
		isConnected = true
	case string(FPipe):
		newNode.connectionType = FPipe
		isConnected = true
	}

	return isConnected, newNode
}

func getBottomConnection(rowNumber int, columnNumber int) (isConnected bool, newNode Node) {
	belowRowNumber := rowNumber + 1
	if belowRowNumber > len(Map)-1 {
		return isConnected, newNode
	}

	cellBelow := Map[belowRowNumber][columnNumber]
	newNode = Node{
		coordinates: [2]int{belowRowNumber, columnNumber},
	}
	switch cellBelow {
	case string(VerticalPipe):
		newNode.connectionType = VerticalPipe
		isConnected = true
	case string(LPipe):
		newNode.connectionType = LPipe
		isConnected = true
	case string(JPipe):
		newNode.connectionType = JPipe
		isConnected = true
	}

	return isConnected, newNode
}

func getLeftConnection(rowNumber int, columnNumber int) (isConnected bool, newNode Node) {
	leftColumnNumber := columnNumber - 1
	if leftColumnNumber < 0 {
		return isConnected, newNode
	}

	cellLeft := Map[rowNumber][leftColumnNumber]
	newNode = Node{
		coordinates: [2]int{rowNumber, leftColumnNumber},
	}

	switch cellLeft {
	case string(HorizontalPipe):
		newNode.connectionType = HorizontalPipe
		isConnected = true
	case string(LPipe):
		newNode.connectionType = LPipe
		isConnected = true
	case string(FPipe):
		newNode.connectionType = FPipe
		isConnected = true
	}

	return isConnected, newNode
}

func getRightConnection(rowNumber int, columnNumber int) (isConnected bool, newNode Node) {
	rightColumnNumber := columnNumber + 1
	if rightColumnNumber > len(Map[rowNumber])-1 {
		return isConnected, newNode
	}

	cellLeft := Map[rowNumber][rightColumnNumber]
	newNode = Node{
		coordinates: [2]int{rowNumber, rightColumnNumber},
	}

	switch cellLeft {
	case string(HorizontalPipe):
		newNode.connectionType = HorizontalPipe
		isConnected = true
	case string(JPipe):
		newNode.connectionType = JPipe
		isConnected = true
	case string(SevenPipe):
		newNode.connectionType = SevenPipe
		isConnected = true
	}

	return isConnected, newNode
}
