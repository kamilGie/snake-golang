package point

type Point struct{
  X int
  Y int 
}

func NewPointAtDir(head Point, direction [4]int) Point {
	switch direction {
	case [4]int{0, 0, 0, 1}:
		return Point{head.X + 1, head.Y}
	case [4]int{0, 0, 1, 0}:
		return Point{head.X, head.Y + 1}
	case [4]int{0, 1, 0, 0}:
		return Point{head.X - 1, head.Y}
	case [4]int{1, 0, 0, 0}:
		return Point{head.X, head.Y - 1}
	default:
		return head
	}
}
