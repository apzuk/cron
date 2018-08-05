package cron

import (
	"testing"
	"gotest.tools/assert"
)

func TestParse_Wildcard(t *testing.T) {
	c := Parse("*", "*", "*", "*", "*")

	assert.Assert(t, c != nil)
	assert.DeepEqual(t, c.Minute, []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	})

	assert.DeepEqual(t, c.Hour, []int{
		0, 1, 2, 3, 4, 5, 6,
		7, 8, 9, 10, 11, 12,
		13, 14, 15, 16, 17, 18,
		19, 20, 21, 22, 23,
	})

	assert.DeepEqual(t, c.Day, []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31,
	})

	assert.DeepEqual(t, c.Month, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	assert.DeepEqual(t, c.WeekDay, []int{0, 1, 2, 3, 4, 5, 6})
}

func TestParse_Value(t *testing.T) {
	// lower edge case
	c := Parse("0", "0", "1", "1", "0")
	assert.Assert(t, c != nil)
	assert.DeepEqual(t, c.Minute, []int{0})
	assert.DeepEqual(t, c.Hour, []int{0})
	assert.DeepEqual(t, c.Day, []int{1})
	assert.DeepEqual(t, c.Month, []int{1})
	assert.DeepEqual(t, c.WeekDay, []int{0})

	// values in the middle
	c = Parse("15", "6", "10", "2", "5")
	assert.Assert(t, c != nil)
	assert.DeepEqual(t, c.Minute, []int{15})
	assert.DeepEqual(t, c.Hour, []int{6})
	assert.DeepEqual(t, c.Day, []int{10})
	assert.DeepEqual(t, c.Month, []int{2})
	assert.DeepEqual(t, c.WeekDay, []int{5})

	// upper edge case
	c = Parse("59", "23", "31", "12", "6")
	assert.Assert(t, c != nil)
	assert.DeepEqual(t, c.Minute, []int{59})
	assert.DeepEqual(t, c.Hour, []int{23})
	assert.DeepEqual(t, c.Day, []int{31})
	assert.DeepEqual(t, c.Month, []int{12})
	assert.DeepEqual(t, c.WeekDay, []int{6})
}

func TestParse_ValueInvalid(t *testing.T) {
	// minute invalid value
	assertInvalidParse(t,"60", "*", "*", "*", "*")
	assertInvalidParse(t,"-1", "*", "*", "*", "*")
	assertInvalidParse(t,"ss", "*", "*", "*", "*")

	// hour invalid value
	assertInvalidParse(t,"*", "24", "*", "*", "*")
	assertInvalidParse(t,"*", "-1", "*", "*", "*")
	assertInvalidParse(t,"*", "qq", "*", "*", "*")

	// day invalid value
	assertInvalidParse(t,"*", "*", "32", "*", "*")
	assertInvalidParse(t,"*", "*", "0", "*", "*")
	assertInvalidParse(t,"*", "*", "yy", "*", "*")

	// month invalid value
	assertInvalidParse(t,"*", "*", "*", "13", "*")
	assertInvalidParse(t,"*", "*", "*", "0", "*")
	assertInvalidParse(t,"*", "*", "*", "uu", "*")

	// weekday invalid value
	assertInvalidParse(t,"*", "*", "*", "*", "7")
	assertInvalidParse(t,"*", "*", "*", "*", "-1")
	assertInvalidParse(t,"*", "*", "*", "*", "ii")
}

func TestParse_ValueRange(t *testing.T) {
	c := Parse("2-15", "10-13", "4-9", "3-7", "4-6")
	assert.Assert(t, c != nil)

	assert.DeepEqual(t, c.Minute, []int{2, 3, 4, 5, 6, 7, 8, 9, 10,11, 12, 13, 14, 15})
	assert.DeepEqual(t, c.Hour, []int{10, 11, 12, 13})
	assert.DeepEqual(t, c.Day, []int{4, 5, 6, 7, 8, 9})
	assert.DeepEqual(t, c.Month, []int{3, 4, 5, 6, 7})
	assert.DeepEqual(t, c.WeekDay, []int{4, 5, 6})
}

func TestParse_ValueRangeEdges(t *testing.T) {
	c := Parse("0-59", "0-23", "1-31", "1-12", "0-6")
	assert.Assert(t, c != nil)

	assert.DeepEqual(t, c.Minute, []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	})
	assert.DeepEqual(t, c.Hour, []int{
		0, 1, 2, 3, 4, 5, 6,
		7, 8, 9, 10, 11, 12,
		13, 14, 15, 16, 17, 18,
		19, 20, 21, 22, 23,
	})

	assert.DeepEqual(t, c.Day, []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31,
	})

	assert.DeepEqual(t, c.Month, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	assert.DeepEqual(t, c.WeekDay, []int{0, 1, 2, 3, 4, 5, 6})
}

func TestParse_ValueRangeInvalid(t *testing.T) {
	// Minute invalid range
	assertInvalidParse(t, "1-60", "*", "*", "*", "*")
	assertInvalidParse(t, "1-aa", "*", "*", "*", "*")
	assertInvalidParse(t, "aa-6", "*", "*", "*", "*")
	assertInvalidParse(t, "aa-bb", "*", "*", "*", "*")
	assertInvalidParse(t, "13-4", "*", "*", "*", "*")
	assertInvalidParse(t, "78-98", "*", "*", "*", "*")

	// Hour invalid range
	assertInvalidParse(t, "*", "4-24", "*", "*", "*")
	assertInvalidParse(t, "*", "3-aa", "*", "*", "*")
	assertInvalidParse(t, "*", "aa-8", "*", "*", "*")
	assertInvalidParse(t, "*", "aa-bb", "*", "*", "*")
	assertInvalidParse(t, "*", "7-2", "*", "*", "*")
	assertInvalidParse(t, "*", "45-65", "*", "*", "*")

	// Day invalid range
	assertInvalidParse(t, "*", "*", "0-7", "*", "*")
	assertInvalidParse(t, "*", "*", "5-32", "*", "*")
	assertInvalidParse(t, "*", "*", "4-aa", "*", "*")
	assertInvalidParse(t, "*", "*", "aa-7", "*", "*")
	assertInvalidParse(t, "*", "*", "aa-bb*", "*", "*")
	assertInvalidParse(t, "*", "*", "8-4", "*", "*")
	assertInvalidParse(t, "*", "*", "56-76", "*", "*")

	// Month invalid range
	assertInvalidParse(t, "*", "*", "*", "0-4", "*")
	assertInvalidParse(t, "*", "*", "*", "6-14", "*")
	assertInvalidParse(t, "*", "*", "*", "3-aa", "*")
	assertInvalidParse(t, "*", "*", "*", "aa-5", "*")
	assertInvalidParse(t, "*", "*", "*", "aa-bb", "*")
	assertInvalidParse(t, "*", "*", "*", "5-2", "*")
	assertInvalidParse(t, "*", "*", "*", "20-24", "*")

	// Weekday invalid range
	assertInvalidParse(t, "*", "*", "*", "*", "3-9")
	assertInvalidParse(t, "*", "*", "*", "*", "2-aa")
	assertInvalidParse(t, "*", "*", "*", "*", "aa-4")
	assertInvalidParse(t, "*", "*", "*", "*", "aa-bb")
	assertInvalidParse(t, "*", "*", "*", "*", "5-3")
	assertInvalidParse(t, "*", "*", "*", "*", "34-38")
}

func TestParse_WildcardEveryNth(t *testing.T) {
	c := Parse("*/13", "*/6", "*/10", "*/3", "*/2")
	assert.Assert(t, c != nil)
	assert.DeepEqual(t, c.Minute, []int{13, 26, 39, 52})
	assert.DeepEqual(t, c.Hour, []int{6, 12, 18})
	assert.DeepEqual(t, c.Day, []int{10, 20, 30})
	assert.DeepEqual(t, c.Month, []int{3, 6, 9, 12})
	assert.DeepEqual(t, c.WeekDay, []int{2, 4, 6})
}

func TestParse_WildcardEveryNthEdges(t *testing.T) {
	c := Parse("*/1", "*/1", "*/1", "*/1", "*/1")
	assert.Assert(t, c != nil)
	assert.DeepEqual(t, c.Minute, []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	})

	assert.DeepEqual(t, c.Hour, []int{
		1, 2, 3, 4, 5, 6,
		7, 8, 9, 10, 11, 12,
		13, 14, 15, 16, 17, 18,
		19, 20, 21, 22, 23,
	})

	assert.DeepEqual(t, c.Day, []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31,
	})

	assert.DeepEqual(t, c.Month, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	assert.DeepEqual(t, c.WeekDay, []int{1, 2, 3, 4, 5, 6})
}

func TestParse_WildcardEveryNthInvalid(t *testing.T) {
	// Minute invalid every nth
	assertInvalidParse(t, "*/60", "*", "*", "*", "*")
	assertInvalidParse(t, "*/0", "*", "*", "*", "*")
	assertInvalidParse(t, "*/aa", "*", "*", "*", "*")
	assertInvalidParse(t, "*/89", "*", "*", "*", "*")
	assertInvalidParse(t, "*/-18", "*", "*", "*", "*")

	// Hour invalid every nth
	assertInvalidParse(t, "*", "*/24", "*", "*", "*")
	assertInvalidParse(t, "*", "*/0", "*", "*", "*")
	assertInvalidParse(t, "*", "*/aa", "*", "*", "*")
	assertInvalidParse(t, "*", "*/45", "*", "*", "*")
	assertInvalidParse(t, "*", "*/-10", "*", "*", "*")

	// Day invalid every nth
	assertInvalidParse(t, "*", "*", "*/32", "*", "*")
	assertInvalidParse(t, "*", "*", "*/0", "*", "*")
	assertInvalidParse(t, "*", "*", "*/aa", "*", "*")
	assertInvalidParse(t, "*", "*", "*/45", "*", "*")
	assertInvalidParse(t, "*", "*", "*/-10", "*", "*")

	// Month invalid every nth
	assertInvalidParse(t, "*", "*", "*", "*/13", "*")
	assertInvalidParse(t, "*", "*", "*", "*/0", "*")
	assertInvalidParse(t, "*", "*", "*", "*/aa", "*")
	assertInvalidParse(t, "*", "*", "*", "*/20", "*")
	assertInvalidParse(t, "*", "*", "*", "*/-3", "*")

	// Weekday invalid every nth
	assertInvalidParse(t, "*", "*", "*", "*", "*/9")
	assertInvalidParse(t, "*", "*", "*", "*", "*/0")
	assertInvalidParse(t, "*", "*", "*", "*", "*/aa")
	assertInvalidParse(t, "*", "*", "*", "*", "*/55")
	assertInvalidParse(t, "*", "*", "*", "*", "*/-4")
}

func TestParse_ValueEveryNth(t *testing.T) {
	c := Parse("31/13", "15/6", "19/10", "7/3", "3/2")
	assert.Assert(t, c != nil)
	assert.DeepEqual(t, c.Minute, []int{39, 52})
	assert.DeepEqual(t, c.Hour, []int{18})
	assert.DeepEqual(t, c.Day, []int{20, 30})
	assert.DeepEqual(t, c.Month, []int{9, 12})
	assert.DeepEqual(t, c.WeekDay, []int{4, 6})
}

func TestParse_ValueEveryNthInvalid(t *testing.T) {
	// Minute invalid value and every nth
	assertInvalidParse(t, "0/13", "*", "*", "*", "*")
	assertInvalidParse(t, "67/12", "*", "*", "*", "*")
	assertInvalidParse(t, "13/60", "*", "*", "*", "*")

	// Hour invalid value and every nth
	assertInvalidParse(t, "*", "0/24", "*", "*", "*")
	assertInvalidParse(t, "*", "34/2", "*", "*", "*")
	assertInvalidParse(t, "*", "2/56", "*", "*", "*")

	// Day invalid value and every nth
	assertInvalidParse(t, "*", "*", "0/32", "*", "*")
	assertInvalidParse(t, "*", "*", "36/5", "*", "*")
	assertInvalidParse(t, "*", "*", "5/60", "*", "*")

	// Month invalid value and every nth
	assertInvalidParse(t, "*", "*", "*", "0/15", "*")
	assertInvalidParse(t, "*", "*", "*", "16/4", "*")
	assertInvalidParse(t, "*", "*", "*", "7/28", "*")

	// Weekday invalid value and every nth
	assertInvalidParse(t, "*", "*", "*", "*", "0/15")
	assertInvalidParse(t, "*", "*", "*", "*", "16/4")
	assertInvalidParse(t, "*", "*", "*", "*", "4/23")
}

func TestParse_ValueRangeEveryNth(t *testing.T) {
	c := Parse("15-50/13", "8-20/6", "2-28/10", "5-11/3", "3-6/2")
	assert.Assert(t, c != nil)
	assert.DeepEqual(t, c.Minute, []int{26, 39})
	assert.DeepEqual(t, c.Hour, []int{12, 18})
	assert.DeepEqual(t, c.Day, []int{10, 20})
	assert.DeepEqual(t, c.Month, []int{6, 9})
	assert.DeepEqual(t, c.WeekDay, []int{4, 6})
}

func TestParse_ValueRangeEveryNthInvalid(t *testing.T) {
	// Minute
	assertInvalidParse(t,"23-5/13", "*", "*", "*", "*")
	assertInvalidParse(t,"4-65/13", "*", "*", "*", "*")
	// Hour
	assertInvalidParse(t,"*", "18-5/13", "*", "*", "*")
	assertInvalidParse(t,"*", "4-9/12", "*", "*", "*")
	// Day
	assertInvalidParse(t,"*", "*", "26-56/2", "*", "*")
	assertInvalidParse(t,"*", "*", "11-6/3", "*", "*")
	// Month
	assertInvalidParse(t,"*", "*", "*", "6-20/4", "*")
	assertInvalidParse(t,"*", "*", "*", "7-4/4", "*")
	// Weekday
	assertInvalidParse(t,"*", "*", "*", "*", "3-9/4")
	assertInvalidParse(t,"*", "*", "*", "*", "6-1/3")
}

func TestParse_WithCommandAndMultipleArgs(t *testing.T) {
	c := Parse("-d", "--output=html", "*", "*", "*", "*", "*", "/usr/bin/bash", "-m", "--debug")
	assert.Assert(t, c != nil)
	assert.Equal(t, c.Command, "/usr/bin/bash")
}

func TestParse_InvalidArguments(t *testing.T) {
	assertInvalidParse(t, "-d", "--output=html", "*", "*", "*", "*", "/usr.bin/bash")
	assertInvalidParse(t, "-d", "--output=html", "12", "-f", "*", "*", "*", "*", "/usr.bin/bash")
	assertInvalidParse(t, "*", "*", "*", "*", "45", "/usr.bin/bash")
	assertInvalidParse(t, "*", "*", "89" ,"*", "*", "/usr.bin/bash")
}

func assertInvalidParse(t *testing.T, args... string) {
	c := Parse(args...)
	assert.Assert(t, c == nil)
}