package main

import "testing"

func TestGetCommand(t *testing.T) {
	var tests = []struct {
		cmdStr    string
		wantCmd   Command
		wantValue int
	}{
		{"up 5", Up, 5},
		{"down 2", Down, 2},
		{"forward 9", Forward, 9},
		{"down 199", Down, 199},
		{"xxx", Up, 0}, // not recognized; default to Up, 0
	}

	for _, test := range tests {
		ansCmd, ansValue := GetCommand(test.cmdStr)
		if ansCmd != test.wantCmd || ansValue != test.wantValue {
			t.Errorf("Command should be [%d, %d], instead [%d, %d]", test.wantCmd, test.wantValue, ansCmd, ansValue)
		}
	}
}

func TestDive(t *testing.T) {
	var tests = []struct {
		commands  []string
		wantHPos  int
		wantDepth int
	}{
		{[]string{"down 2"}, 0, 2},
		{[]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}, 15, 10},
		{[]string{"up 3", "up 10"}, 0, -13},
	}

	for _, test := range tests {
		ansHPos, ansDepth := Dive(test.commands)
		if ansHPos != test.wantHPos || ansDepth != test.wantDepth {
			t.Errorf("Position should be [%d, %d], instead [%d, %d]", test.wantHPos, test.wantDepth, ansHPos, ansDepth)
		}
	}
}

func TestDiveWithAim(t *testing.T) {
	var tests = []struct {
		commands  []string
		wantHPos  int
		wantDepth int
	}{
		{[]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}, 15, 60},
		{[]string{"down 100", "forward 100"}, 100, 10000},
	}

	for _, test := range tests {
		ansHPos, ansDepth := DiveWithAim(test.commands)
		if ansHPos != test.wantHPos || ansDepth != test.wantDepth {
			t.Errorf("Position should be [%d, %d], instead [%d, %d]", test.wantHPos, test.wantDepth, ansHPos, ansDepth)
		}
	}
}
