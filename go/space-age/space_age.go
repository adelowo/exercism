package space

type Planet string

func Age(secs float64, planet Planet) float64 {

	const earthSecondsPerYear = 31557600

	var y float64

	switch planet {
	case "Mercury":
		y = 0.2408467
	case "Venus":
		y = 0.61519726
	case "Earth":
		y = 1.0
	case "Mars":
		y = 1.8808158
	case "Jupiter":
		y = 11.862615
	case "Saturn":
		y = 29.447498
	case "Uranus":
		y = 84.016846
	case "Neptune":
		y = 164.79132
	}

	return secs / (earthSecondsPerYear * y)
}
