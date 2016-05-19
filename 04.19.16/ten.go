func Mul(m [][]float64, val interface{}) [][]float64 {
	switch v := val.(type) {
	case float64:
		for i := range n {
			for j := range n[i] {
				n[i][j] *= v
			}
		}
	case []float64:
		for i := range n {
			for j := range v {
				n[i][j] *= v[j]
			}
		}
	case [][]float64:
		for i := range n {
			for j := range n[i] {
				n[i][j] *= v[i][j]
			}
		}
	default:
		panic(fmt.Println(("wrong type received", w))
	}
	return n
}


