package genarator

import (
	"fmt"
	"math/rand"
	"os"
)

func GenerateFile(arrivalTimeRange []int, exeTimeRange []int, numberOfLines int, outputFile string) error {
    file, err := os.Create(outputFile)
    if err != nil {
        return fmt.Errorf("could not create file: %v", err)
    }
    defer file.Close() 

    for i := 0; i < numberOfLines; i++ {
        arrivalTime := rand.Intn(arrivalTimeRange[1]-arrivalTimeRange[0]+1) + arrivalTimeRange[0]
        exeTime     := rand.Intn(exeTimeRange[1]-exeTimeRange[0]+1) + exeTimeRange[0]
        line := fmt.Sprintf("%d %d\n", arrivalTime, exeTime)

        _, err := file.WriteString(line)
        if err != nil {
            return fmt.Errorf("could not write to file: %v", err)
        }
    }

    return nil
}
