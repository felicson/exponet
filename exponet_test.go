//Package exponet for extract exhibitions from exponet.ru
//It get exhibitions list and parse each item from list
package exponet

import (
	"io"
	"reflect"
	"testing"
)

func Test_getIndex(t *testing.T) {
	type args struct {
		rdr io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIndex(tt.args.rdr)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
