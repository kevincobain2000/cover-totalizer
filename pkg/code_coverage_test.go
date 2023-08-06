package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {

}

func TestParseCoveragePercentage(t *testing.T) {
	tests := []struct {
		language        string
		coverageXMLPath string
		wantPercentage  float64
		errWant         error
	}{
		{
			language:        "php",
			coverageXMLPath: "../test_data/php_coverage.xml",
			wantPercentage:  76.85,
			errWant:         nil,
		},
		{
			language:        "go",
			coverageXMLPath: "../test_data/go_coverage.xml",
			wantPercentage:  33.03,
			errWant:         nil,
		},
		{
			language:        "js",
			coverageXMLPath: "../test_data/clover_coverage.xml",
			wantPercentage:  97.91,
			errWant:         nil,
		},
		{
			language:        "java",
			coverageXMLPath: "../test_data/java_coverage.xml",
			wantPercentage:  63.04,
			errWant:         nil,
		},
	}
	for _, test := range tests {
		t.Run(test.language, func(t *testing.T) {
			s := NewCoverageService()
			got, err := s.ParseCoveragePercentage(test.coverageXMLPath)
			assert.Equal(t, test.errWant, err)
			assert.Equal(t, test.wantPercentage, got)

		})
	}
}
