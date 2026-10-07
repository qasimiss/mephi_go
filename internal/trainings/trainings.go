package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training хранит данные о тренировке и персональные данные пользователя.
type Training struct {
	Steps                 int           // количество шагов за тренировку.
	TrainingType          string        // тип тренировки.
	Duration              time.Duration // длительность тренировки.
	personaldata.Personal               // встроенные персональные данные пользователя.
}

// Parse разбирает строку формата "3456,Ходьба,3h00m" и заполняет поля структуры.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("неверное количество элементов: получено %d, ожидается 3", len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга количества шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля: %d", steps)
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности тренировки: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность тренировки должна быть больше нуля: %v", duration)
	}

	t.Steps = steps
	t.TrainingType = parts[1] // тип тренировки валидируется в ActionInfo
	t.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о проведённой тренировке.
func (t Training) ActionInfo() (string, error) {
	var spentCalories func(steps int, weight, height float64, duration time.Duration) (float64, error)

	switch t.TrainingType {
	case "Ходьба":
		spentCalories = spentenergy.WalkingSpentCalories
	case "Бег":
		spentCalories = spentenergy.RunningSpentCalories
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	calories, err := spentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories,
	), nil
}
