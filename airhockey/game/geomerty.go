package game

import "math"

func normalizeMallet1Pos(pos Vec2d) Vec2d {
	return Vec2d{
		X:clamp(pos.X, -0.5 + malletRadius, 0.5 + malletRadius),
		Y:clamp(pos.Y, 0, 0.8 - malletRadius),

	}
}

func normalizeMallet2Pos(pos Vec2d) Vec2d {
	return Vec2d{
		X:clamp(pos.X, -0.5 + malletRadius, 0.5 + malletRadius),
		Y:clamp(pos.Y, -0.8 + malletRadius, 0),
	}
}

func distance(vec1, vec2 Vec2d) float64  {
	return math.Sqrt(math.Pow(vec2.X - vec1.X, 2) + math.Pow(vec2.Y - vec1.Y, 2))
}

func vecBetween(from, to Vec2d) Vec2d {
	return Vec2d{
		X: to.X - from.X,
		Y: to.Y - from.Y,
	}
}

func clamp(x float64, min float64, max float64) float64 {
	return math.Min(max, math.Max(x, min))
}
