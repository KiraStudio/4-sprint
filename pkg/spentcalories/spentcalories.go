package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	strSplit := strings.Split(data, ",")
	if len(strSplit) != 3 {
		return 0, "", time.Duration(0), fmt.Errorf(`длина слайса была равна 3, 
													так как в строке данных у нас количество шагов, 
													вид активности и продолжительность.`)
	}

	countSteps, err := strconv.Atoi(strSplit[0])
	if err != nil {
		return 0, "", time.Duration(0), err
	}

	if countSteps <= 0 {
		return 0, "", time.Duration(0), fmt.Errorf("кол-во шагов меньше или равно нулю.")
	}

	category := strSplit[1]

	t, err := time.ParseDuration(strSplit[2])
	if err != nil {
		return 0, "", time.Duration(0), err
	}

	return countSteps, category, t, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := stepLengthCoefficient * height
	destination := stepLength * float64(steps)
	destination /= mInKm

	return destination
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	destination := distance(steps, height)
	hours := float64(duration / time.Hour)
	if hours == 0 {
		return 0
	}
	averageSpeed := destination / float64(hours)

	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, category, duration, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	if steps <= 0 {
		log.Println("кол-во шагов меньше или равно нулю")
		return "", err
	}

	if duration <= 0 {
		log.Println("продолжительность тренировки меньше или равна нулю")
		return "", err
	}

	if category == "" {
		log.Println("не указан тип тренировки")
		return "", err
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInHours := float64(duration / time.Hour)
	distance := distance(steps, height)
	runningSpentCalories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return "", err
	}

	walkingSpentCalories, err := WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return "", err
	}

	switch category {
	case "Ходьба":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2fч.\nДистанция:  %.2fкм.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", category, durationInHours, distance, meanSpeed, walkingSpentCalories), nil
	case "Бег":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2fч.\nДистанция:  %.2fкм.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", category, durationInHours, distance, meanSpeed, runningSpentCalories), nil
	default:
		return "неизвестный тип тренировки", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("кол-во шагов меньше или равно нулю.")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес не может быть нулевым или отрицательным.")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост не может быть нулевым или отрицательным.")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность не может быть нулевая или отрицательная.")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := int(duration / time.Minute)

	return (weight * meanSpeed * float64(durationInMinutes)) / minInH, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("кол-во шагов меньше или равно нулю.")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес не может быть нулевым или отрицательным.")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост не может быть нулевым или отрицательным.")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность не может быть нулевая или отрицательная.")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := int(duration / time.Minute)

	walkingSpentCalories := ((weight * meanSpeed * float64(durationInMinutes)) / minInH) * walkingCaloriesCoefficient
	if walkingSpentCalories < 0 {
		return 0, fmt.Errorf("Ошибка при вычислении каллорий")
	}

	return walkingSpentCalories, nil

}
