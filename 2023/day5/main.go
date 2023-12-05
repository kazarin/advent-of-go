package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Record struct {
	Source      int
	Destination int
	Range       int
}

func buildRecord(values []string) *Record {
	source, _ := strconv.Atoi(values[1])
	dst, _ := strconv.Atoi(values[0])
	r, _ := strconv.Atoi(values[2])
	return &Record{Source: source, Destination: dst, Range: r}
}

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	seeds := map[int]int{}
	soils := []*Record{}
	waters := []*Record{}
	fertilizers := []*Record{}
	lights := []*Record{}
	temperatures := []*Record{}
	humidities := []*Record{}
	locations := []*Record{}
	seedToSoil := false
	soilToFertilizer := false
	fertilizerToWater := false
	waterToLight := false
	temperatureToHumidity := false
	lightToTemperature := false
	humidityToLocation := false
	for _, row := range rows[:len(rows)-1] {
		re := regexp.MustCompile(`(\d+)`)

		if strings.HasPrefix(row, "seeds:") {
			matches := re.FindAllString(row, -1)
			for _, match := range matches {

				n, _ := strconv.Atoi(match)
				seeds[n] = n
			}
			continue
		}

		if strings.HasPrefix(row, "seed-to-soil") {
			seedToSoil = true
			continue
		}
		if strings.HasPrefix(row, "soil-to-fertilizer") {
			seedToSoil = false
			soilToFertilizer = true
			continue
		}
		if strings.HasPrefix(row, "fertilizer-to-water") {
			soilToFertilizer = false
			fertilizerToWater = true
			continue
		}
		if strings.HasPrefix(row, "water-to-light") {
			waterToLight = true
			fertilizerToWater = false
			continue
		}
		if strings.HasPrefix(row, "light-to-temperature") {
			lightToTemperature = true
			waterToLight = false
			continue
		}
		if strings.HasPrefix(row, "temperature-to-humidity") {
			temperatureToHumidity = true
			lightToTemperature = false
			continue
		}
		if strings.HasPrefix(row, "humidity-to-location") {
			humidityToLocation = true
			temperatureToHumidity = false
			continue
		}
		if seedToSoil {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				soils = append(soils, buildRecord(matches))
			}
		}
		if soilToFertilizer {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				fertilizers = append(fertilizers, buildRecord(matches))
			}
		}
		if fertilizerToWater {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				waters = append(waters, buildRecord(matches))
			}
		}
		if waterToLight {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				lights = append(lights, buildRecord(matches))
			}
		}
		if lightToTemperature {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				temperatures = append(temperatures, buildRecord(matches))
			}
		}
		if temperatureToHumidity {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				humidities = append(humidities, buildRecord(matches))
			}
		}
		if humidityToLocation {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				locations = append(locations, buildRecord(matches))
			}
		}
	}
	results := []int{}
	for _, seed := range seeds {
		result := fetchFromDict(locations, fetchFromDict(humidities, fetchFromDict(temperatures, fetchFromDict(lights, fetchFromDict(waters, fetchFromDict(fertilizers, fetchFromDict(soils, seed)))))))
		results = append(results, result)

	}

	return slices.Min(results)
}

func fetchFromDict(list []*Record, target int) int {
	for _, r := range list {
		if target >= r.Source && target < r.Source+r.Range {
			return target - r.Source + r.Destination
		}
	}
	return target
}

func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	seeds := []int{}
	soils := []*Record{}
	waters := []*Record{}
	fertilizers := []*Record{}
	lights := []*Record{}
	temperatures := []*Record{}
	humidities := []*Record{}
	locations := []*Record{}
	seedToSoil := false
	soilToFertilizer := false
	fertilizerToWater := false
	waterToLight := false
	temperatureToHumidity := false
	lightToTemperature := false
	humidityToLocation := false
	for _, row := range rows[:len(rows)-1] {
		re := regexp.MustCompile(`(\d+)`)

		if strings.HasPrefix(row, "seeds:") {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				for _, match := range matches {
					n, _ := strconv.Atoi(match)
					seeds = append(seeds, n)
				}
			}
			continue
		}

		if strings.HasPrefix(row, "seed-to-soil") {
			seedToSoil = true
			continue
		}
		if strings.HasPrefix(row, "soil-to-fertilizer") {
			seedToSoil = false
			soilToFertilizer = true
			continue
		}
		if strings.HasPrefix(row, "fertilizer-to-water") {
			soilToFertilizer = false
			fertilizerToWater = true
			continue
		}
		if strings.HasPrefix(row, "water-to-light") {
			waterToLight = true
			fertilizerToWater = false
			continue
		}
		if strings.HasPrefix(row, "light-to-temperature") {
			lightToTemperature = true
			waterToLight = false
			continue
		}
		if strings.HasPrefix(row, "temperature-to-humidity") {
			temperatureToHumidity = true
			lightToTemperature = false
			continue
		}
		if strings.HasPrefix(row, "humidity-to-location") {
			humidityToLocation = true
			temperatureToHumidity = false
			continue
		}
		if seedToSoil {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				soils = append(soils, buildRecord(matches))
			}
		}
		if soilToFertilizer {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				fertilizers = append(fertilizers, buildRecord(matches))
			}
		}
		if fertilizerToWater {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				waters = append(waters, buildRecord(matches))
			}
		}
		if waterToLight {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				lights = append(lights, buildRecord(matches))
			}
		}
		if lightToTemperature {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				temperatures = append(temperatures, buildRecord(matches))
			}
		}
		if temperatureToHumidity {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				humidities = append(humidities, buildRecord(matches))
			}
		}
		if humidityToLocation {
			matches := re.FindAllString(row, -1)
			if len(matches) > 1 {
				locations = append(locations, buildRecord(matches))
			}
		}
	}

	m := math.MaxInt
	i := 0
	for i < len(seeds) {
		start, length := seeds[i], seeds[i+1]
		for j := start; j < (start + length); j++ {
			result := fetchFromDict(locations, fetchFromDict(humidities, fetchFromDict(temperatures, fetchFromDict(lights, fetchFromDict(waters, fetchFromDict(fertilizers, fetchFromDict(soils, j)))))))
			if result < m {
				m = result

			}
		}

		i += 2
	}

	return m
}
func main() {
	fmt.Println(Part1("day5.txt"))
	fmt.Println(Part2("day5.txt"))
}
