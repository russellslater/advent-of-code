package device_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/russellslater/advent-of-code/2022/day-06-tuning-trouble/device"
)

func TestDevice(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		signal string
		device device.Device
		want   int
	}{
		{
			"AoC Example 1",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			device.Device{DetectionLimit: 4},
			5,
		},
		{
			"AoC Example 2",
			"nppdvjthqldpwncqszvftbrmjlhg",
			device.Device{DetectionLimit: 4},
			6,
		},
		{
			"AoC Example 3",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			device.Device{DetectionLimit: 4},
			10,
		},
		{
			"AoC Example 4",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			device.Device{DetectionLimit: 4},
			11,
		},
		{
			"AoC Example 5",
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			device.Device{DetectionLimit: 14},
			19,
		},
		{
			"AoC Example 6",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			device.Device{DetectionLimit: 14},
			23,
		},
		{
			"AoC Example 7",
			"nppdvjthqldpwncqszvftbrmjlhg",
			device.Device{DetectionLimit: 14},
			23,
		},
		{
			"AoC Example 8",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			device.Device{DetectionLimit: 14},
			29,
		},
		{
			"AoC Example 9",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			device.Device{DetectionLimit: 14},
			26,
		},
		{
			"No signal",
			"",
			device.Device{DetectionLimit: 14},
			-1,
		},
		{
			"DetectionLimit of 0",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			device.Device{DetectionLimit: 0},
			-1,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := tc.device.Detect(tc.signal)

			is.Equal(got, tc.want)
		})
	}
}
