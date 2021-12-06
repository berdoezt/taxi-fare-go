package model

type DistanceMeter struct {
	ElapsedTime string
	Mileage     float64
}

func (d *DistanceMeter) GetElapsedTime() string {
	return d.ElapsedTime
}

func (d *DistanceMeter) GetMileage() float64 {
	return d.Mileage
}

func (d *DistanceMeter) WithElapsedTime(elapsedTime string) *DistanceMeter {
	d.ElapsedTime = elapsedTime
	return d
}

func (d *DistanceMeter) WithMileage(mileage float64) *DistanceMeter {
	d.Mileage = mileage
	return d
}
