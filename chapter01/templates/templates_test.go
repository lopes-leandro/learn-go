package templates

import "testing"

func TestRunTemplate(t *testing.T)  {
	tests := []struct {
		name string
		wantErr bool
	}{
		{"base-case", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			if err := RunTemplate(); (err != nil) != tt.wantErr {
				t.Errorf("RunTemplate() error = %v, watErr %v", err, tt.wantErr)
			}
		})
	}

}