package genarator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGenerateFile(t *testing.T) {
    arrivalTimeRange := []int{0, 0} 
    exeTimeRange := []int{1, 8} 
    numberOfLines := 10
    outputFile := "out_test.txt" 

    _, err := GenerateFile(arrivalTimeRange, exeTimeRange, numberOfLines, outputFile)
    if err != nil {
        t.Fatalf("failed to generate file: %v", err)
    }

    file, err := os.Open(outputFile)
    if err != nil {
        t.Fatalf("could not open generated file: %v", err)
    }
    defer file.Close()


    lineCount := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        lineCount++

        parts := make([]string, 2)
        _, err := fmt.Sscanf(line, "%s %s", &parts[0], &parts[1])
        if err != nil {
            t.Fatalf("line does not match format: %v", err)
        }

        arrivalTime, err := strconv.Atoi(parts[0])
        if err != nil || arrivalTime < arrivalTimeRange[0] || arrivalTime > arrivalTimeRange[1] {
            t.Fatalf("arrival time out of range: %d", arrivalTime)
        }

        exeTime, err := strconv.Atoi(parts[1])
        if err != nil || exeTime < exeTimeRange[0] || exeTime > exeTimeRange[1] {
            t.Fatalf("execution time out of range: %d", exeTime)
        }
    }

    if lineCount != numberOfLines {
        t.Errorf("expected %d lines, got %d", numberOfLines, lineCount)
    }

    os.Remove(outputFile)
}