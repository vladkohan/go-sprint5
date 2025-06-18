package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	splitComma := strings.Split(datastring, ",")
	if len(splitComma) != 3 {
		return errors.New("неверный формат данных")
	}

	t.Steps, err = strconv.Atoi(splitComma[0])
	if err != nil {
		return fmt.Errorf("неверный формат шагов: %w", err)
	}
	if t.Steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}

	t.TrainingType = splitComma[1]

	t.Duration, err = time.ParseDuration(splitComma[2])
	if err != nil {
		return fmt.Errorf("неверный формат продолжительности: %w", err)
	}
	if t.Duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return "", errors.New("неизвестный тип тренировки")
	}

	calcDistance := spentenergy.Distance(t.Steps, t.Height)
	midSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var burncalories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		burncalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		burncalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		calcDistance,
		midSpeed,
		burncalories,
	)

	return result, nil
}
