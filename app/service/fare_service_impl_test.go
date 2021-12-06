package service

import (
	"reflect"
	"testing"

	"github.com/berdoezt/taxi-fare-go/app/model"
	"github.com/berdoezt/taxi-fare-go/config"
)

func TestFareServiceImpl_GetFareMeter(t *testing.T) {
	type fields struct {
		fareRules map[string]config.FareRule
	}
	type args struct {
		distanceMeters []model.DistanceMeter
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFunc func(dms []model.DistanceMeter)
		want     float64
		wantErr  bool
	}{
		{
			name: "Case #1. When mileage not pass base distance, should calculate until base distance",
			fields: fields{
				fareRules: map[string]config.FareRule{
					"base": {
						Price:    100,
						Distance: 1000,
					},
				},
			},
			args: args{
				distanceMeters: []model.DistanceMeter{
					{
						ElapsedTime: "00:00:00.000",
						Mileage:     0.0,
					},
					{
						ElapsedTime: "00:02:00.125",
						Mileage:     500.0,
					},
				},
			},
			want: 100,
		},
		{
			name: "Case #2. When mileage pass base distance but not over up_to distance, should calculate until upto distance",
			fields: fields{
				fareRules: map[string]config.FareRule{
					"base": {
						Price:    100,
						Distance: 1000,
					},
					"up_to": {
						Price:             50,
						Distance:          200,
						DistanceThreshold: 5000,
					},
				},
			},
			args: args{
				distanceMeters: []model.DistanceMeter{
					{
						ElapsedTime: "00:00:00.000",
						Mileage:     0.0,
					},
					{
						ElapsedTime: "00:02:00.125",
						Mileage:     4000.0,
					},
				},
			},
			want: 850,
		},
		{
			name: "Case #3. When mileage pass over distance, should calculate until over distance",
			fields: fields{
				fareRules: map[string]config.FareRule{
					"base": {
						Price:    100,
						Distance: 1000,
					},
					"up_to": {
						Price:             50,
						Distance:          200,
						DistanceThreshold: 5000,
					},
					"over": {
						Price:             50,
						Distance:          100,
						DistanceThreshold: 5000,
					},
				},
			},
			args: args{
				distanceMeters: []model.DistanceMeter{
					{
						ElapsedTime: "00:00:00.000",
						Mileage:     0.0,
					},
					{
						ElapsedTime: "00:02:00.125",
						Mileage:     9000.0,
					},
				},
			},
			want: 3100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FareServiceImpl{
				fareRules: tt.fields.fareRules,
			}

			got, err := f.GetFareMeter(tt.args.distanceMeters)
			if (err != nil) != tt.wantErr {
				t.Errorf("FareServiceImpl.GetFareMeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FareServiceImpl.GetFareMeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFareServiceImpl(t *testing.T) {
	type args struct {
		opts []FareServiceImplOpt
	}
	tests := []struct {
		name string
		args args
		want *FareServiceImpl
	}{
		{
			name: "Case #1. When option not passed, should create without option",
			args: args{},
			want: &FareServiceImpl{},
		},
		{
			name: "Case #2. When option passed, should create with option",
			args: args{
				opts: []FareServiceImplOpt{
					func(fsi *FareServiceImpl) {
						fsi.fareRules = map[string]config.FareRule{}
					},
				},
			},
			want: &FareServiceImpl{
				fareRules: map[string]config.FareRule{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFareServiceImpl(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFareServiceImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
