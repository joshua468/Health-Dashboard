
package main

import (
    "fmt"
    "os"
    "time"

    "github.com/plotly/plotly/graph"
    "github.com/plotly/plotly/graph/options"
    "github.com/plotly/plotly/graph/traces"
)

func main() {
    // Sample data
    dates := make([]time.Time, 30)
    weights := make([]float64, 30)
    steps := make([]int, 30)
    sleeps := make([]float64, 30)
    calories := make([]int, 30)
    water := make([]float64, 30)

    startDate := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
    for i := 0; i < 30; i++ {
        dates[i] = startDate.AddDate(0, 0, i)
        weights[i] = 70.0 - 0.1*float64(i)
        steps[i] = 8000 + 500*i
        sleeps[i] = 7.5 + 0.1*float64(i)
        calories[i] = 2000 - 50*i
        water[i] = 8.0 + 0.5*float64(i)
    }

    // Weight Tracker
    weightTrace := traces.Scatter{
        X: dates,
        Y: weights,
        Mode: traces.ModeLinesMarkers,
        Name: "Weight",
    }
    weightGraph := graph.Graph2d{Data: []interface{}{weightTrace}}
    weightPlot, err := weightGraph.Plot()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error plotting weight tracker:", err)
        return
    }
    fmt.Println("Weight Tracker:", weightPlot.URL)

    // Fitness Tracker
    stepsTrace := traces.Scatter{
        X: dates,
        Y: steps,
        Mode: traces.ModeLinesMarkers,
        Name: "Steps",
    }
    stepsGraph := graph.Graph2d{Data: []interface{}{stepsTrace}}
    stepsPlot, err := stepsGraph.Plot()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error plotting fitness tracker:", err)
        return
    }
    fmt.Println("Fitness Tracker:", stepsPlot.URL)

    // Sleep Tracker
    sleepsTrace := traces.Scatter{
        X: dates,
        Y: sleeps,
        Mode: traces.ModeBar,
        Name: "Sleep",
    }
    sleepsGraph := graph.Graph2d{Data: []interface{}{sleepsTrace}}
    sleepsPlot, err := sleepsGraph.Plot()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error plotting sleep tracker:", err)
        return
    }
    fmt.Println("Sleep Tracker:", sleepsPlot.URL)

    // Caloric Intake Tracker
    caloriesTrace := traces.Scatter{
        X: dates,
        Y: calories,
        Mode: traces.ModeLinesMarkers,
        Name: "Calories",
    }
    caloriesGraph := graph.Graph2d{Data: []interface{}{caloriesTrace}}
    caloriesPlot, err := caloriesGraph.Plot()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error plotting caloric intake tracker:", err)
        return
    }
    fmt.Println("Caloric Intake Tracker:", caloriesPlot.URL)

    // Hydration Tracker
    waterTrace := traces.Scatter{
        X: dates,
        Y: water,
        Mode: traces.ModeLinesMarkers,
        Name: "Water",
    }
    waterGraph := graph.Graph2d{Data: []interface{}{waterTrace}}
    waterPlot, err := waterGraph.Plot()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error plotting hydration tracker:", err)
        return
    }
    fmt.Println("Hydration Tracker:", waterPlot.URL)
}
