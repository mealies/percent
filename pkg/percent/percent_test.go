package percent_test

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/mealies/percent/internal/resources"
	"github.com/mealies/percent/pkg/percent"
	"testing"
)

func TestPercent(t *testing.T) {
	t.Parallel()

	type in struct {
		percent float64
		value   float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid positive input",
			in: in{
				percent: 25,
				value:   100,
			},
			want: want{
				value: 25,
				err:   nil,
			},
		},
		{
			name: "valid negative input",
			in: in{
				percent: 50,
				value:   -200,
			},
			want: want{
				value: -100,
				err:   nil,
			},
		},
		{
			name: "invalid percentage",
			in: in{
				percent: 150,
				value:   100,
			},
			want: want{
				value: 0.0,
				err:   resources.ErrOutOfRange,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := percent.Percent(tt.in.percent, tt.in.value)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Percent() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("Percent(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}

func TestOf(t *testing.T) {
	t.Parallel()

	type in struct {
		part  float64
		total float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid positive input",
			in: in{
				part:  25,
				total: 100,
			},
			want: want{
				value: 25,
				err:   nil,
			},
		},
		{
			name: "valid negative input",
			in: in{
				part:  -200,
				total: -50,
			},
			want: want{
				value: 400,
				err:   nil,
			},
		},
		{
			name: "invalid part",
			in: in{
				part:  150,
				total: 100,
			},
			want: want{
				value: 0.0,
				err:   resources.ErrPartGreaterThanTotal,
			},
		},
		{
			name: "invalid total",
			in: in{
				part:  150,
				total: 0,
			},
			want: want{
				value: 0.0,
				err:   resources.ErrDivideByZero,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := percent.Of(tt.in.part, tt.in.total)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Of() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("Of(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}

func TestChange(t *testing.T) {
	t.Parallel()

	type in struct {
		oldValue float64
		newValue float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid increase input",
			in: in{
				oldValue: 25,
				newValue: 100,
			},
			want: want{
				value: 300,
				err:   nil,
			},
		},
		{
			name: "valid decrease input",
			in: in{
				oldValue: -50,
				newValue: -200,
			},
			want: want{
				value: -300,
				err:   nil,
			},
		},
		{
			name: "invalid zero value",
			in: in{
				oldValue: 0,
				newValue: 100,
			},
			want: want{
				value: 0,
				err:   resources.ErrDivideByZero,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := percent.Change(tt.in.oldValue, tt.in.newValue)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Change() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("Change(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}

func TestRemain(t *testing.T) {
	t.Parallel()

	type in struct {
		percent float64
		value   float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid positive input",
			in: in{
				percent: 25,
				value:   100,
			},
			want: want{
				value: 75,
				err:   nil,
			},
		},
		{
			name: "valid negative input",
			in: in{
				percent: 50,
				value:   -200,
			},
			want: want{
				value: -100,
				err:   nil,
			},
		},
		{
			name: "invalid percentage input",
			in: in{
				percent: 150,
				value:   100,
			},
			want: want{
				value: 0,
				err:   resources.ErrOutOfRange,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := percent.Remain(tt.in.percent, tt.in.value)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Remain() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("Remain(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}

func TestFromRatio(t *testing.T) {
	t.Parallel()

	type in struct {
		ratio float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid input",
			in: in{
				ratio: 0.25,
			},
			want: want{
				value: 25,
				err:   nil,
			},
		},
		{
			name: "invalid input",
			in: in{
				ratio: 2,
			},
			want: want{
				value: 0,
				err:   resources.ErrOutOfRange,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := percent.FromRatio(tt.in.ratio)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("FromRatio() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("FromRatio(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}

func TestToRatio(t *testing.T) {
	t.Parallel()

	type in struct {
		percent float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid positive input",
			in: in{
				percent: 50,
			},
			want: want{
				value: 0.5,
				err:   nil,
			},
		},
		{
			name: "invalid input",
			in: in{
				percent: 150,
			},
			want: want{
				value: 0,
				err:   resources.ErrOutOfRange,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := percent.ToRatio(tt.in.percent)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("ToRatio() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("ToRatio(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}
