package cron

import (
	"regexp"
	"strconv"
	"sort"
)

type section struct {
	min int
	max int
}

type Cron struct {
	Minute  []int
	Hour    []int
	Day     []int
	Month   []int
	WeekDay []int
	Command string
}

var (
	// `3-6`
	regexpValueRange = regexp.MustCompile("^(\\d{1,2})-(\\d{1,2})$")

	// `*/5`
	regexpWildcardEveryNth = regexp.MustCompile("^\\*/(\\d{1,2})$")

	// `2/5`
	regexpValueEveryNth = regexp.MustCompile("^(\\d{1,2})/(\\d{1,2})$")

	// `2-5/4`
	regexpRangeAndEveryNth = regexp.MustCompile("^(\\d{1,2})-(\\d{1,2})/(\\d+)$")
)

var (
	minute = section{min: 0, max: 59}
	hour = section{min: 0, max: 23}
	days = section{min: 1, max: 31}
	month = section{min: 1, max: 12}
	weekday = section{min: 0, max: 6}
)

func Parse(args... string) *Cron {
	for i := range args {

		if i + 5 > len(args) {
			return nil
		}

		m := parseMinute(args[i])
		if len(m) == 0 {
			continue
		}

		h := parseHour(args[i + 1])
		if len(h) == 0 {
			continue
		}

		d := parseDay(args[i + 2])
		if len(d) == 0 {
			continue
		}

		mo := parseMonth(args[i + 3])
		if len(mo) == 0 {
			continue
		}

		w := parseWeekday(args[i + 4])
		if len(w) == 0 {
			continue
		}

		c := &Cron{
			Minute: m,
			Hour: h,
			Day: d,
			Month: mo,
			WeekDay: w,
		}

		if len(args) > i + 5 {
			c.Command = args[i + 5]
		}

		return c
	}

	return nil
}

func parseMinute(raw string) []int {
	return parseSection(minute, raw)
}

func parseHour(raw string) []int {
	return parseSection(hour, raw)
}

func parseDay(raw string) []int {
	return parseSection(days, raw)
}

func parseMonth(raw string) []int {
	return parseSection(month, raw)
}

func parseWeekday(raw string) []int {
	return parseSection(weekday, raw)
}

func parseSection(s section, raw string) []int {
	r := s.parse(raw)

	var el = make(map[int]struct{})
	// remove duplicates
	for i := range r {
		if _, ok := el[r[i]]; ok {
			r = append(r[:i], r[i+1:]...)
		}
		el[r[i]] = struct{}{}
	}

	// sort, although if should be already sorted :)
	sort.Ints(r)

	return r
}

func (s section) defaultRange() []int {
	return customRange(s.min, s.max)
}


func (s section) inRange(v int) bool {
	return s.min <= v && s.max >= v
}

func (s section) int(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return i
}

func (s section) parse(raw string) []int {
	if raw == "*" {
		return s.defaultRange()
	}

	if intval, err := strconv.Atoi(raw); err == nil {
		if s.inRange(intval) {
			return []int{intval}
		}
		return []int{}
	}

	matches := regexpValueRange.FindStringSubmatch(raw)
	if len(matches) > 0 {
		min, max := s.int(matches[1]), s.int(matches[2])
		if s.inRange(min) && s.inRange(max) && min < max {
			return customRange(min, max)
		}
		return []int{}
	}

	matches = regexpWildcardEveryNth.FindStringSubmatch(raw)
	if len(matches) > 0 {
		v := s.int(matches[1])
		if s.inRange(v) && v > 0 {
			return rangeEveryNth(s.min, s.max, v)
		}
		return []int{}
	}

	matches = regexpValueEveryNth.FindStringSubmatch(raw)
	if len(matches) > 0 {
		min, max := s.int(matches[1]), s.max
		if s.inRange(min) && s.inRange(max) && min < max && min > 0 {
			return rangeEveryNth(min, s.max, s.int(matches[2]))
		}
		return []int{}
	}

	matches = regexpRangeAndEveryNth.FindStringSubmatch(raw)
	if len(matches) > 0 {
		min, max := s.int(matches[1]), s.int(matches[2])
		if s.inRange(min) && s.inRange(max) && min < max {
			return rangeEveryNth(min, max, s.int(matches[3]))
		}
		return []int{}
	}

	return []int{}
}

func rangeEveryNth(min, max, nth int) []int {
	var (
		r = make([]int, 0)
		i = 1
	)
	for {
		cur := i * nth
		if cur > max {
			break
		}
		if cur >= min {
			r = append(r, i * nth)
		}
		i++
	}
	return r
}

func customRange(min, max int) []int {
	var r = make([]int, 0, max - min)
	for i := min; i <= max; i++ {
		r = append(r, i)
	}
	return r
}