package waterdata

import (
	"log/slog"
	"os"
	"testing"
)

const (
	testSnoqualmieCarnation string = "12149000"
	testSnoqualmieDuvall    string = "12150400"
)

var (
	testSiteCodes = []string{
		testSnoqualmieCarnation,
		testSnoqualmieDuvall,
	}
)

// TestInstantaneousValues tests whether the service returns a value
// TODO Improve the test to check the value that's returned
func TestInstantaneousValues(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	client, err := NewClient(logger)
	if err != nil {
		t.Errorf("expected to be able to create client")
	}

	resp, err := client.GetInstantaneousValues(testSiteCodes)
	if err != nil {
		t.Errorf("expected success")
	}

	// print(resp)

	// Expected 3 TimeSeries
	{
		want := 3
		got := len(resp.Value.TimeSeries)
		if got != want {
			t.Errorf("TimeSeries got: %d; want: %d", got, want)
		}
	}

	// Expected 3 TimeSeries to contain the 2 SiteCodes
	{
		want := 2
		sitecode := map[string]struct{}{}

		for _, timeseries := range resp.Value.TimeSeries {
			sitecode[timeseries.SourceInfo.SiteCode[0].Value] = struct{}{}
		}

		got := len(sitecode)
		if got != want {
			t.Errorf("SiteCodes got: %d; want: %d", got, want)
		}
	}

	// Expected 3 TimeSeries to contain 2 GageHeightFeet measurements
	{
		want := 2
		got := 0

		for _, timeseries := range resp.Value.TimeSeries {
			if timeseries.Variable.Contains(GageHeightFeet) {
				got++
			}
		}

		if got != want {
			t.Errorf("GageHeightFeet got: %d; want: %d", got, want)
		}
	}
}
