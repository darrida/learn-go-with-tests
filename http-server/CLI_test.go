package poker_test

import (
	"fmt"
	"poker"
	"strings"
	"testing"
	"time"
)

// var dummyBlindAlerter = &SpyBlindAlerter{}
// var dummyPlayerStore = &poker.StubPlayerStore{}
// var dummyStdIn = &bytes.Buffer{}
// var dummyStdOut = &bytes.Buffer{}

type expectedSchedule struct {
	expectedScheduleTime time.Duration
	expectedAmount       int
}

type scheduledALert struct {
	at     time.Duration
	amount int
}

func (s scheduledALert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledALert
}

func (s *SpyBlindAlerter) ScheduledAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledALert{at, amount})
}

var dummySpyAlert = &SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlert)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlert)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []expectedSchedule{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", want.expectedAmount, want.expectedScheduleTime), func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

	// t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
	// 	stdout := &bytes.Buffer{}
	// 	cli := poker.NewCLI(dummyPlayerStore, dummyStdIn, stdout, dummyBlindAlerter)
	// 	cli.PlayPoker()

	// 	got := stdout.String()
	// 	want := "Please enter the number of players: "

	// 	if got != want {
	// 		t.Errorf("got %q, want %q", got, want)
	// 	}
	// })
}

func assertScheduledAlert(t testing.TB, alert scheduledALert, want expectedSchedule) {
	t.Helper()
	amountGot := alert.amount
	if amountGot != want.expectedAmount {
		t.Errorf("got amount %d, want %d", amountGot, want.expectedAmount)
	}

	gotScheduledTime := alert.at
	if gotScheduledTime != want.expectedScheduleTime {
		t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, want.expectedScheduleTime)
	}
}
