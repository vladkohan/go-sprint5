package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		// Распарсим данные
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Ошибка парсинга: %v", err)
			continue
		}

		// Получим информацию о действии
		info, err2 := dp.ActionInfo()
		if err2 != nil {
			log.Printf("Ошибка получения информации: %v", err2)
			continue
		}

		// Выведем информацию
		fmt.Println(info)
	}

}
