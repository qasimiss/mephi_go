package personaldata

import "fmt"

// Personal хранит персональные данные пользователя.
type Personal struct {
	Name   string  // имя пользователя.
	Weight float64 // вес пользователя в килограммах.
	Height float64 // рост пользователя в метрах.
}

// Print выводит персональные данные пользователя в стандартный поток вывода.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
