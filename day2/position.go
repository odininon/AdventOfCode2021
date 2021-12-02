package main

type position struct {
	horizontal, vertical, aim int
}

func newPosition(horizontal, vertical, aim int) position {
	return position{horizontal: horizontal, vertical: vertical, aim: aim}
}
