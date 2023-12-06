package main

import (
	"advent"
	"fmt"
	"os"
	"regexp"
)

type Almanac struct {
	seeds                 []int
	seedToSoil            *Catmap
	soilToFertilizer      *Catmap
	fertelizerToWater     *Catmap
	waterToLight          *Catmap
	lightToTemperature    *Catmap
	temperatureToHumidity *Catmap
	humidityToLocation    *Catmap
}

func lineToNumbers(line string) []int {
	stringed := regexp.MustCompile(`\s+`).Split(line, -1)
	nums := make([]int, len(stringed))
	for i, s := range stringed {
		nums[i] = advent.MustAtoi(s)
	}
	return nums
}

// parseCatmap parses and returns a *catmap starting at lines[start],
// and also the index where the map definition ends (empty line, or EOF)
func parseCatmap(lines []string, start int) (*Catmap, int) {
	maps := []*Mapunit{}
	var i int
	// The first line is the map name, and it's irrelevant for parsing
	// fmt.Printf("lines[%d]: %s\n", start, lines[start])
	for i = start + 1; i < len(lines) && lines[i] != ""; i++ {
		nums := lineToNumbers(lines[i])
		maps = append(maps, &Mapunit{
			destStart:   nums[0],
			sourceStart: nums[1],
			rangeLen:    nums[2]})

	}
	return NewCatmap(maps), i
}

func parseAlmanac(lines []string) Almanac {
	almanac := Almanac{}

	seedList := regexp.MustCompile(`\s*:\s*`).Split(lines[0], 2)[1]
	almanac.seeds = lineToNumbers(seedList)

	i := 1
	var incr int
	almanac.seedToSoil, i = parseCatmap(lines, i+1)
	almanac.soilToFertilizer, i = parseCatmap(lines, i+1)
	almanac.fertelizerToWater, i = parseCatmap(lines, i+1)
	almanac.waterToLight, i = parseCatmap(lines, i+1)
	almanac.lightToTemperature, i = parseCatmap(lines, i+1)
	almanac.temperatureToHumidity, i = parseCatmap(lines, i+1)
	almanac.humidityToLocation, i = parseCatmap(lines, i+1)

	if i+incr != len(lines) {
		panic(fmt.Sprintf("Error parsing Almanac - did not consume all lines (%d/%d)", i+incr, len(lines)))
	}
	return almanac
}

func (a *Almanac) locationForSeed(seed int) int {
	soil := a.seedToSoil.find(seed)
	fert := a.soilToFertilizer.find(soil)
	water := a.fertelizerToWater.find(fert)
	light := a.waterToLight.find(water)
	temp := a.lightToTemperature.find(light)
	hum := a.temperatureToHumidity.find(temp)
	location := a.humidityToLocation.find(hum)
	return location

}

func findLowestLocation(lines []string) int {
	almanac := parseAlmanac(lines)
	min := -1
	for _, seed := range almanac.seeds {
		if loc := almanac.locationForSeed(seed); min == -1 || loc < min {
			min = loc
		}
	}

	return min
}

func main() {
	lines := advent.Readlines(os.Args[1])
	location := findLowestLocation(lines)
	fmt.Printf("Part 1: %d\n", location)
}
