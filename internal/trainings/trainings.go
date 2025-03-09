package trainings

import (
	"fmt"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
)

// Структура Training, представляющая тренировку.
type Training struct {
	Personal personaldata.Personal
	Type     string        // Тип тренировки (например, ходьба, бег)
	Steps    int           // Количество шагов
	Weight   float64       // Вес пользователя
	Height   float64       // Рост пользователя
	Duration time.Duration // Длительность тренировки
}

// Метод Parse, который парсит строку и инициализирует структуру Training.
func (t *Training) Parse(data string) error {
	_, err := fmt.Sscanf(data, "%s, %d, %f, %f, %v", &t.Type, &t.Steps, &t.Weight, &t.Height, &t.Duration)
	if err != nil {
		return err
	}
	return nil
}

// Метод ActionInfo возвращает строковое описание тренировки с ошибкой.
func (t *Training) ActionInfo() (string, error) {
	return fmt.Sprintf("Тип тренировки: %s\nШаги: %d\nВес: %.2f кг\nРост: %.2f см\nДлительность: %v",
		t.Type, t.Steps, t.Weight, t.Height, t.Duration), nil
}

// Метод Print выводит информацию о тренировке.
func (t *Training) Print() {
	fmt.Printf("Тип тренировки: %s\nШаги: %d\nВес: %.2f кг\nРост: %.2f см\nДлительность: %v\n",
		t.Type, t.Steps, t.Weight, t.Height, t.Duration)
}
