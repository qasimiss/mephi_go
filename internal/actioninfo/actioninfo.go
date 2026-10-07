package actioninfo

import (
	"fmt"
	"log"
)

// DataParser — интерфейс для разбора строки с данными
// и формирования строки с информацией об активности.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info разбирает каждую строку набора dataset с помощью переданной реализации
// DataParser и выводит информацию об активности в стандартный поток вывода.
// Ошибки парсинга и формирования информации логируются.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка разбора строки %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации для строки %q: %v", data, err)
			continue
		}

		fmt.Println(info)
	}
}
