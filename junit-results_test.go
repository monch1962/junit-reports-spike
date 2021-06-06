package junit_results_process

import (
	"testing"

	"github.com/joshdk/go-junit"
)

func TestJunitFileIngest(t *testing.T) {
	suites, err := junit.IngestFiles([]string{
		"junit-results-files/java-results-1.xml",
	})
	if err != nil {
		t.Fatalf("failed to ingest JUnit xml %v", err)
	}
	//fmt.Printf("Suites ingested: %v", suites)
	t.Log(suites)

	for _, suite := range suites {
		t.Logf("Suite: %s\n", suite.Name)
		for _, test := range suite.Tests {
			t.Logf("  TestName: %s\n", test.Name)
			if test.Error != nil {
				t.Logf("    Test result: %s: %s\n", test.Status, test.Error.Error())
			} else {
				t.Logf("    Test result: %s\n", test.Status)
			}
		}
	}

}
