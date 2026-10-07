package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories рассчитывает количество калорий, потраченных при ходьбе:
// базовый расход калорий, умноженный на коэффициент ходьбы.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := spentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories рассчитывает количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	return spentCalories(steps, weight, height, duration)
}

// spentCalories рассчитывает базовое количество потраченных калорий:
// (вес × средняя скорость × длительность в минутах) / минут в часе.
func spentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case steps <= 0:
		return 0, fmt.Errorf("некорректное количество шагов: %d (ожидается больше нуля)", steps)
	case weight <= 0:
		return 0, fmt.Errorf("некорректный вес: %.2f (ожидается больше нуля)", weight)
	case height <= 0:
		return 0, fmt.Errorf("некорректный рост: %.2f (ожидается больше нуля)", height)
	case duration <= 0:
		return 0, fmt.Errorf("некорректная длительность: %v (ожидается больше нуля)", duration)
	}

	speed := MeanSpeed(steps, height, duration)
	return (weight * speed * duration.Minutes()) / minInH, nil
}

// MeanSpeed рассчитывает среднюю скорость в км/ч:
// дистанция, делённая на длительность в часах.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// Distance рассчитывает пройденную дистанцию в километрах:
// (количество шагов × длина шага) / метров в километре.
// Длина шага вычисляется как рост, умноженный на коэффициент длины шага.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength
	return distance / mInKm
}
