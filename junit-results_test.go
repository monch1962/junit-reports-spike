package junit_results_process

import (
	"testing"

	"github.com/joshdk/go-junit"
)

func TestJunitSingleFileIngest(t *testing.T) {
	suites, err := junit.IngestFile("junit-results-files/java-results-1.xml")
	if err != nil {
		t.Fatalf("failed to ingest JUnit xml %v", err)
	}
	if err != nil {
		t.Fatalf("failed to ingest JUnit xml %v", err)
	}
	//fmt.Printf("Suites ingested: %v", suites)
	//t.Log(suites)
	if len(suites) != 2 {
		t.Fatalf("Expected 2 suites, found %d", len(suites))
	}

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
func TestJunitMultipleFilesIngest(t *testing.T) {
	suites, err := junit.IngestFiles([]string{
		"junit-results-files/java-results-1.xml",
	})
	if err != nil {
		t.Fatalf("failed to ingest JUnit xml %v", err)
	}
	//fmt.Printf("Suites ingested: %v", suites)
	//t.Log(suites)
	if len(suites) != 2 {
		t.Fatalf("Expected 2 suites, found %d", len(suites))
	}

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

func TestJunitResultsDirectoryIngest(t *testing.T) {
	suites, err := junit.IngestDir("junit-results-files/")
	if err != nil {
		t.Fatalf("failed to ingest JUnit xml %v", err)
	}
	//t.Log(suites)
	if len(suites) != 4 {
		t.Fatalf("Expected 4 suites, found %d", len(suites))
	}

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
