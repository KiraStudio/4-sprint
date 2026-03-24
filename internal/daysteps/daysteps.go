package daysteps

import (
	calories "4-sprint/pkg/spentcalories"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	strSplit := strings.Split(data, ",")
	if len(strSplit) != 2 {
		return 0, time.Duration(0), fmt.Errorf("длина слайса была равна 2, так как в строке данных у нас количество шагов и продолжительность.\n")
	}

	countSteps, err := strconv.Atoi(strSplit[0])
	if err != nil {
		return 0, time.Duration(0), err
	}

	if countSteps <= 0 {
		return 0, time.Duration(0), fmt.Errorf("кол-во шагов должно быть больше нуля.\n")
	}

	t, err := time.ParseDuration(strSplit[1])
	if err != nil {
		return 0, time.Duration(0), err
	}

	return countSteps, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		fmt.Println("кол-во шагов или продолжительность или вес или высота равны или меньше нуля")
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm
	walkingSpentCalories, err := calories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if distance <= 0 || walkingSpentCalories <= 0 {
		fmt.Println("дистанция или потраченные калл. равны или меньше нуля")
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал.", steps, distance, walkingSpentCalories)
}
