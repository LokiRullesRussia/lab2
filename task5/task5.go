package task5

//Структура — это способ объединить
// несколько связанных значений в одну «коробку» и дать ей имя.
type Rectangle struct {
	Width  float64
	Height float32
}

func SquareRectangle(r Rectangle) float64 {
	return r.Width * float64(r.Height)
}
