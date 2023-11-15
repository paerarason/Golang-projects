package main
const m float32=0.3048

func Celsius(f float32) float32{
	return (f-32)*5/9
}

func feet(ft int) float32 {
	return float32(ft)*m
}