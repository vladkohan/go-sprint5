package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	splitComma := strings.Split(datastring, ",")
	if len(splitComma) != 2 {
		return errors.New("неверный формат данных")
	}

	ds.Steps, err = strconv.Atoi(splitComma[0])
	if err != nil {
		return fmt.Errorf("неверный формат шагов: %w", err)
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}

	ds.Duration, err = time.ParseDuration(splitComma[1])
	if err != nil {
		return fmt.Errorf("неверный формат продолжительности: %w", err)
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(int(ds.Steps), float64(ds.Personal.Height))

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
