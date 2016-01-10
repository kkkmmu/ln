package ln

type Hit struct {
	Shape Shape
	T     float64
}

var NoHit = Hit{nil, INF}

func (hit *Hit) Ok() bool {
	return hit.T < INF
}
