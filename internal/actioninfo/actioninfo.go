package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			fmt.Println("Ошибка парсинга:", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("Ошибка формирования информации:", err)
			continue
		}
		fmt.Println(info)
	}
}
