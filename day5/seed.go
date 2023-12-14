package main

import (
	"advent"
	"fmt"
	"os"
	"regexp"
)

type Almanac struct {
	seeds                 []int
	seedRanges            []*numRange
	seedToSoil            *categoryMap
	soilToFertilizer      *categoryMap
	fertelizerToWater     *categoryMap
	waterToLight          *categoryMap
	lightToTemperature    *categoryMap
	temperatureToHumidity *categoryMap
	humidityToLocation    *categoryMap
}

func parseAlmanac(lines []string) Almanac {
	almanac := Almanac{}

	seedList := regexp.MustCompile(`\s*:\s*`).Split(lines[0], 2)[1]
	almanac.seeds = advent.LineToNumbers(seedList)
	almanac.seedRanges = parseSeedRangesLine(seedList)

	i := 1
	var incr int
	almanac.seedToSoil, i = parseCategorymap(lines, i+1)
	almanac.soilToFertilizer, i = parseCategorymap(lines, i+1)
	almanac.fertelizerToWater, i = parseCategorymap(lines, i+1)
	almanac.waterToLight, i = parseCategorymap(lines, i+1)
	almanac.lightToTemperature, i = parseCategorymap(lines, i+1)
	almanac.temperatureToHumidity, i = parseCategorymap(lines, i+1)
	almanac.humidityToLocation, i = parseCategorymap(lines, i+1)

	if i+incr != len(lines) {
		panic(fmt.Sprintf("Error parsing Almanac - did not consume all lines (%d/%d)", i+incr, len(lines)))
	}
	return almanac
}

// parseCategorymap parses and returns a *categoryMap starting at lines[start],2
// and also the index where the map definition ends (empty line, or EOF)
func parseCategorymap(lines []string, start int) (*categoryMap, int) {
	maps := []*mapping{}
	var i int
	// The first line is the map name, and it's irrelevant for parsing
	// fmt.Printf("lines[%d]: %s\n", start, lines[start])
	for i = start + 1; i < len(lines) && lines[i] != ""; i++ {
		nums := advent.LineToNumbers(lines[i])
		maps = append(maps, &mapping{
			destStart:   nums[0],
			sourceStart: nums[1],
			rangeLen:    nums[2]})

	}
	return newCategoryMap(maps), i
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

type numRange struct {
	start  int
	_range int
}

func parseSeedRangesLine(line string) []*numRange {
	nums := advent.LineToNumbers(line)
	ranges := make([]*numRange, len(nums)/2)
	for i := 0; i < len(nums); i += 2 {
		ranges[i/2] = &numRange{start: nums[i], _range: nums[i+1]}
	}
	return ranges
}

func findLowestLocationInRange(almanac *Almanac, sr *numRange, ch chan int) {
	fmt.Printf("Finding lowest location for range %d-%d\n", sr.start, sr.start+sr._range)
	min := -1
	for i := sr.start; i < sr.start+sr._range; i++ {
		if loc := almanac.locationForSeed(i); min == -1 || loc < min {
			min = loc
		}
	}
	fmt.Printf("Found lowest location %d\n", min)
	ch <- min
}

func findLowestLocationFromRanges(lines []string) int {
	almanac := parseAlmanac(lines)
	ch := make(chan int)

	for _, sr := range almanac.seedRanges {
		go findLowestLocationInRange(&almanac, sr, ch)
	}
	min := -1
	for i := 0; i < len(almanac.seedRanges); i++ {
		loc := <-ch
		if loc < min || min == -1 {
			min = loc
		}
	}
	return min
}

func main() {
	lines := advent.Readlines(os.Args[1])
	location := findLowestLocation(lines)
	fmt.Printf("Part 1: %d\n", location)
	locationForRange := findLowestLocationFromRanges(lines)
	fmt.Printf("Part 2 %d\n", locationForRange)
}
