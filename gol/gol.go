package main

import "fmt"

func calculateNextState(p golParams, world [][]byte) [][]byte {
	//imgWidth := p.imageWidth
	//imgHeight := p.imageHeight

	tempWorld := make([][]byte, len(world))
	copy(tempWorld, world)

	var turnOffList []cell
	var turnOnList []cell

	for rowIndex, rowList := range world {
		for cellIndex, cellValue := range rowList {

			var cellIndexPlus = (cellIndex + 1) % len(rowList)
			var cellIndexMinus = (((cellIndex - 1) % len(rowList)) + len(rowList)) % len(rowList) // SO, Liam Kelly
			var rowIndexPlus = (rowIndex + 1) % len(rowList)
			var rowIndexMinus = (((rowIndex - 1) % len(rowList)) + len(rowList)) % len(rowList)

			// 1. any live cell with fewer than two live neighbours dies
			// 2. any live cell with two or three live neighbours is unaffected
			// 3. any live cell with more than three live neighbours dies
			if cellValue == 255 {
				var liveCellNeighbours = 0

				// Right
				if world[rowIndex][cellIndexPlus] == 255 {
					liveCellNeighbours++ //fmt.Println("{", cellIndex, ",", rowIndex, "}", "right")
				}

				// Down - Right
				if world[rowIndexPlus][cellIndexPlus] == 255 {
					liveCellNeighbours++
				}

				// Down
				if world[rowIndexPlus][cellIndex] == 255 {
					liveCellNeighbours++
				}

				// Down - Left
				if world[rowIndexPlus][cellIndexMinus] == 255 {
					liveCellNeighbours++
				}

				// Left
				if world[rowIndex][cellIndexMinus] == 255 {
					liveCellNeighbours++
				}

				// Up - Left
				if world[rowIndexMinus][cellIndexMinus] == 255 {
					liveCellNeighbours++
				}

				// Up
				if world[rowIndexMinus][cellIndex] == 255 {
					liveCellNeighbours++
				}

				// Up - Right
				if world[rowIndexMinus][cellIndexPlus] == 255 {
					liveCellNeighbours++
				}

				fmt.Println("{", cellIndex, ",", rowIndex, "}", "Live Neighbours =>>>", liveCellNeighbours)

				// Evaluation
				if liveCellNeighbours < 2 || liveCellNeighbours > 3 {
					fmt.Println("Cell Killed ", liveCellNeighbours, "neighbours")
					turnOffList = append(turnOffList, cell{rowIndex, cellIndex})
				}
			} else {
				// 4. any dead cell with exactly three live neighbours becomes alive

				deadCellNeighbours := 0

				// Right
				if world[rowIndex][cellIndexPlus] == 255 {
					deadCellNeighbours++
				}

				// Up - Right
				if world[rowIndexPlus][cellIndexPlus] == 255 {
					deadCellNeighbours++
				}

				// Up
				if world[rowIndexPlus][cellIndex] == 255 {
					deadCellNeighbours++
				}

				// Up - Left
				if world[rowIndexPlus][cellIndexMinus] == 255 {
					deadCellNeighbours++
				}

				// Left
				if world[rowIndex][cellIndexMinus] == 255 {
					deadCellNeighbours++
				}

				// Down - Left
				if world[rowIndexMinus][cellIndexMinus] == 255 {
					deadCellNeighbours++
				}

				// Down
				if world[rowIndexMinus][cellIndex] == 255 {
					deadCellNeighbours++
				}

				// Down - Right
				if world[rowIndexMinus][cellIndexPlus] == 255 {
					deadCellNeighbours++
				}

				// Evaluation
				if deadCellNeighbours == 3 {
					fmt.Println("Cell Born. dead cell neighbours: ", deadCellNeighbours, "\n ")
					turnOnList = append(turnOnList, cell{rowIndex, cellIndex})
				}
			}
		}
	}

	for _, cell := range turnOffList {
		tempWorld[cell.x][cell.y] = 0
	}

	for _, cell := range turnOnList {
		tempWorld[cell.x][cell.y] = 255
	}

	return tempWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	// loop through world pixels and extract live ones
	var liveCells []cell

	for rowIndex, rowList := range world {
		for cellIndex, cellValue := range rowList {
			if cellValue == 255 {
				liveCell := cell{x: cellIndex, y: rowIndex}
				liveCells = append(liveCells, liveCell)

			}
		}
	}

	return liveCells
}
