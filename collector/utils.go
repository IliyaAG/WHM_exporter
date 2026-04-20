package collector
import "strconv"

func parseFloat(s string) float64 {
    v, err := strconv.ParseFloat(s, 64)
    if err != nil {
        return 0
    }
    return v
}

func boolToFloat(i int) float64 {
    if i == 1 {
        return 1
    }
    return 0
}
