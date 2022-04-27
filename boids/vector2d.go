package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x + v2.x, y: v1.y + v2.y}
}

func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x - v2.x, y: v1.y - v2.y}
}

func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x * v2.x, y: v1.y * v2.y}
}

func (v Vector2D) AddV(d float64) Vector2D {
	return Vector2D{x: v.x + d, y: v.y + d}
}

func (v Vector2D) MultiplyV(d float64) Vector2D {
	return Vector2D{x: v.x * d, y: v.y * d}
}

func (v Vector2D) DivisionV(d float64) Vector2D {
	return Vector2D{x: v.x / d, y: v.y / d}
}

func (v Vector2D) Limit(lower, upper float64) Vector2D {
	return Vector2D{
		x: math.Min(math.Max(v.x, lower), upper),
		y: math.Min(math.Max(v.y, lower), upper),
	}
}

func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
