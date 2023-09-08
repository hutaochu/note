package questionbank

import "testing"

func Test_findDelayedArrivalTime(t *testing.T) {
	type args struct {
		arrivalTime int
		delayedTime int
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		name: "findDelayedArrivalTime",
		args: args{
			arrivalTime: 1,
			delayedTime: 24,
		},
		want: 1,
	}, {
		name: "findDelayedArrivalTime",
		args: args{
			arrivalTime: 1,
			delayedTime: 20,
		},
		want: 21,
	}, {
		name: "findDelayedArrivalTime",
		args: args{
			arrivalTime: 13,
			delayedTime: 11,
		},
		want: 0,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDelayedArrivalTime(tt.args.arrivalTime, tt.args.delayedTime); got != tt.want {
				t.Errorf("findDelayedArrivalTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
