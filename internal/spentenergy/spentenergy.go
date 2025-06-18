package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("все параметры должныбыть больше нуля")
	}

	WalkingMeanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * WalkingMeanSpeed * durationInMinutes) / minInH
	adjustedCalories := spentCalories * walkingCaloriesCoefficient
	return adjustedCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("все параметры должны быть больше нуля")
	}

	RunningMeanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * RunningMeanSpeed * durationInMinutes) / minInH
	return spentCalories, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	speed := distance / duration.Hours()
	return speed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
