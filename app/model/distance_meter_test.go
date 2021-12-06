package model

import (
	"reflect"
	"testing"
)

func TestDistanceMeter_GetElapsedTime(t *testing.T) {
	type fields struct {
		ElapsedTime string
		Mileage     float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Case #1. when elapsed time set, should return elapsed time",
			fields: fields{
				ElapsedTime: "00:02:00.125",
			},
			want: "00:02:00.125",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DistanceMeter{
				ElapsedTime: tt.fields.ElapsedTime,
				Mileage:     tt.fields.Mileage,
			}
			if got := d.GetElapsedTime(); got != tt.want {
				t.Errorf("DistanceMeter.GetElapsedTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistanceMeter_GetMileage(t *testing.T) {
	type fields struct {
		ElapsedTime string
		Mileage     float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Case #1. When mileage set, should return mileage",
			fields: fields{
				Mileage: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DistanceMeter{
				ElapsedTime: tt.fields.ElapsedTime,
				Mileage:     tt.fields.Mileage,
			}
			if got := d.GetMileage(); got != tt.want {
				t.Errorf("DistanceMeter.GetMileage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistanceMeter_WithElapsedTime(t *testing.T) {
	type fields struct {
		ElapsedTime string
		Mileage     float64
	}
	type args struct {
		elapsedTime string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DistanceMeter
	}{
		{
			name: "Case #1. When elapsed time pass, should return DistanceMeter with elapsed time",
			args: args{
				elapsedTime: "00:02:00.125",
			},
			want: &DistanceMeter{
				ElapsedTime: "00:02:00.125",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DistanceMeter{
				ElapsedTime: tt.fields.ElapsedTime,
				Mileage:     tt.fields.Mileage,
			}
			if got := d.WithElapsedTime(tt.args.elapsedTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistanceMeter.WithElapsedTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistanceMeter_WithMileage(t *testing.T) {
	type fields struct {
		ElapsedTime string
		Mileage     float64
	}
	type args struct {
		mileage float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DistanceMeter
	}{
		{
			name: "Case #1. When mileage is pass, should return DistanceMeter with mileage",
			args: args{
				mileage: 100,
			},
			want: &DistanceMeter{
				Mileage: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DistanceMeter{
				ElapsedTime: tt.fields.ElapsedTime,
				Mileage:     tt.fields.Mileage,
			}
			if got := d.WithMileage(tt.args.mileage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistanceMeter.WithMileage() = %v, want %v", got, tt.want)
			}
		})
	}
}
