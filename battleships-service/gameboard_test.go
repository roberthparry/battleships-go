package service

import "testing"

func TestGameboardTranslateCellReference(t *testing.T) {
	type args struct {
		cellReference string
		row           *int
		column        *int
	}
	tests := []struct {
		name         string
		args         args
		want         bool
		wantedRow    int
		wantedColumn int
	}{
		{name: "valid reference a1", args: args{cellReference: "a1", row: new(int), column: new(int)}, want: true, wantedRow: 0, wantedColumn: 0},
		{name: "valid reference A1", args: args{cellReference: "A1", row: new(int), column: new(int)}, want: true, wantedRow: 0, wantedColumn: 0},
		{name: "valid reference j5", args: args{cellReference: "j5", row: new(int), column: new(int)}, want: true, wantedRow: 4, wantedColumn: 9},
		{name: "valid reference J5", args: args{cellReference: "J5", row: new(int), column: new(int)}, want: true, wantedRow: 4, wantedColumn: 9},
		{name: "invalid reference a12", args: args{cellReference: "a12", row: new(int), column: new(int)}, want: false},
		{name: "invalid reference k1", args: args{cellReference: "k1", row: new(int), column: new(int)}, want: false},
		{name: "invalid reference b1o", args: args{cellReference: "k1", row: new(int), column: new(int)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GameboardTranslateCellReference(tt.args.cellReference, tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("GameboardTranslateCellReference() = %v, want %v", got, tt.want)
			} else if got {
				if *tt.args.row != tt.wantedRow {
					t.Errorf("*row = %v, want %v", *tt.args.row, tt.wantedRow)
				}
				if *tt.args.column != tt.wantedColumn {
					t.Errorf("*column = %v, want %v", *tt.args.column, tt.wantedColumn)
				}
			}
		})
	}
}
