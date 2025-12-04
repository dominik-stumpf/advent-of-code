package gridtl

type Grid[T any] [][]T
type Point struct {
	X int
	Y int
}

func checkIsInbound(predicate int, length int) bool {
	return predicate >= 0 && predicate < length
}

func (grid Grid[T]) CheckIsInbound(predicate Point) bool {
	return checkIsInbound(predicate.X, len(grid[0])) && checkIsInbound(predicate.Y, len(grid))
}

func (grid Grid[T]) GetNeighborIndices(start Point) (indices []Point) {
	if !grid.CheckIsInbound(start) {
		panic("start point is out of bounds")
	}
	for i := range 2 {
		offset := i/2 + 1
		sign := 1 - 2*(i%2)
		yNext := start.Y + (offset * sign)
		if checkIsInbound(yNext, len(grid)) {
			indices = append(indices, Point{start.X, yNext})
		}
		xNext := start.X + (offset * sign)
		if checkIsInbound(xNext, len(grid[0])) {
			indices = append(indices, Point{xNext, start.Y})
		}
	}
	return
}

func (grid Grid[T]) GetNeighborIndicesWithCorners(start Point) (indices []Point) {
	if !grid.CheckIsInbound(start) {
		panic("start point is out of bounds")
	}
	for yOffset := range 3 {
		for xOffset := range 3 {
			if xOffset == 1 && yOffset == 1 {
				continue
			}
			x := start.X - 1 + xOffset
			y := start.Y - 1 + yOffset
			if !checkIsInbound(x, len(grid[0])) || !checkIsInbound(y, len(grid)) {
				continue
			}
			indices = append(indices, Point{x, y})
		}
	}
	return
}
