package validation

import "testing"

func TestFareValidation_Validate(t *testing.T) {
	type args struct {
		inputs interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Case #1. When all validation pass, should return no error",
			args: args{
				inputs: []string{
					"00:00:00.000 0.0",
					"00:01:00.123 480.9",
					"00:02:00.125 1141.2",
					"00:03:00.100 6000.8",
				},
			},
		},
		{
			name: "Case #2. When input regex is invalid, should return error",
			args: args{
				inputs: []string{
					"asdf",
					"asdf",
				},
			},
			wantErr: true,
		},
		{
			name: "Case #3. When len data < 2, should return error",
			args: args{
				inputs: []string{
					"00:01:00.123 480.9",
				},
			},
			wantErr: true,
		},
		{
			name: "Case #4. When data range > 5 minutes, should return error",
			args: args{
				inputs: []string{
					"00:00:00.000 0.0",
					"00:01:00.123 480.9",
					"00:02:00.125 1141.2",
					"00:13:00.100 6000.8",
				},
			},
			wantErr: true,
		},
		{
			name: "Case #5. When mileage is 0.0, should return error",
			args: args{
				inputs: []string{
					"00:00:00.000 0.0",
					"00:01:00.123 480.9",
					"00:02:00.125 0.0",
					"00:03:00.100 6000.8",
				},
			},
			wantErr: true,
		},
		{
			name: "Case #6. When timestamp is backdate, should return error",
			args: args{
				inputs: []string{
					"00:00:00.000 0.0",
					"00:01:00.123 480.9",
					"00:02:00.125 1142.0",
					"00:01:00.100 6000.8",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FareValidation{}
			if err := f.Validate(tt.args.inputs); (err != nil) != tt.wantErr {
				t.Errorf("FareValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
