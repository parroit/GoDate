package GoDate

import "testing"


func TestCreation(t *testing.T) { 
	//(days:3, hours:1, minutes:2, seconds:3, milliseconds :4)
	ts:=NewTimeSpan(3, 1, 2, 3, 4)

	if ts.Days()!=3 {
		t.Fatalf("Expected 3,was %d",ts.Days())
	}

	if ts.Hours()!=1 {
		t.Fatalf("Expected 1,was %d",ts.Hours())
	}

	
	if ts.Minutes()!=2 {
		t.Fatalf("Expected 2,was %d",ts.Minutes())
	}

	if ts.Seconds()!=3 {
		t.Fatalf("Expected 3,was %d",ts.Seconds())
	}

	if ts.Milliseconds()!=4 {
		t.Fatalf("Expected 4,was %d",ts.Milliseconds())
	}
}

func TestAdd(t *testing.T) { 
	ts1:=NewTimeSpan(3, 1, 0, 0, 0)
	ts2:=NewTimeSpan(1, 2, 2, 3, 4)

	ts:=ts1+ts2
	
	if ts.Days()!=4 {
		t.Fatalf("Expected 4,was %d",ts.Days())
	}
	if ts.Hours()!=3 {
		t.Fatalf("Expected 3,was %d",ts.Hours())
	}
	
	if ts.Minutes()!=2 {
		t.Fatalf("Expected 2,was %d",ts.Minutes())
	}

	if ts.Seconds()!=3 {
		t.Fatalf("Expected 3,was %d",ts.Seconds())
	}

	if ts.Milliseconds()!=4 {
		t.Fatalf("Expected 4,was %d",ts.Milliseconds())
	}
}


func TestSubtract(t *testing.T) { 
	ts1:=NewTimeSpan(5, 5, 4, 6, 8)
	ts2:=NewTimeSpan(1, 2, 2, 3, 4)

	ts:=ts1-ts2
	
	if ts.Days()!=4 {
		t.Fatalf("Expected 4,was %d",ts.Days())
	}
	if ts.Hours()!=3 {
		t.Fatalf("Expected 3,was %d",ts.Hours())
	}
	
	if ts.Minutes()!=2 {
		t.Fatalf("Expected 2,was %d",ts.Minutes())
	}

	if ts.Seconds()!=3 {
		t.Fatalf("Expected 3,was %d",ts.Seconds())
	}

	if ts.Milliseconds()!=4 {
		t.Fatalf("Expected 4,was %d",ts.Milliseconds())
	}
}


func TestStringFormat(t *testing.T) { 
	ts:=NewTimeSpan(1, 2,  3, 4,100)
	if ts.String()!="1.02:03:04.100" {
		t.Fatalf("Expected 1.02:03:04.100,was %s",ts.String())
	}
}

func TestEquality(t *testing.T) { 
	ts:=NewTimeSpan(1, 0,  0, 0,0)
	ts2:=NewTimeSpan(0, 24,  0, 0,0)
	if ts!=ts2 {
		t.Fatalf("Expected equals values, was not")
	}
	
}

func TestComparison(t *testing.T) { 
	ts:=NewTimeSpan(1, 0,  0, 0,0)
	ts2:=NewTimeSpan(0, 25,  0, 0,0)
	if ts2<=ts {
		t.Fatalf("Expected great value, was not")
	}
	
}


func TestFromValues(t *testing.T) { 
	ts:=NewTimeSpan(1, 12,  0, 0,0)
	if ts!=FromDays(1.5) {
		t.Fatalf("Expected equals value, was not")
	}

	if FromHours(1.5).TotalHours()!=1.5 {
		t.Fatalf("Expected equals value, was not")
	}

	if FromMinutes(1.5).TotalMinutes()!=1.5 {
		t.Fatalf("Expected equals value, was not")
	}

	if FromSeconds(1.5).TotalSeconds()!=1.5 {
		t.Fatalf("Expected equals value, was not")
	}

	if FromMilliseconds(1.5).TotalMilliseconds()!=1.5 {
		t.Fatalf("Expected equals value, was not")
	}

	
}


func TestSubtractNegative(t *testing.T) { 
	ts1:=NewTimeSpan(1, 2, 2, 3, 4)
	ts2:=NewTimeSpan(5, 5, 4, 6, 8)

	ts:=ts1-ts2
	
	if ts.Days()!=-4 {
		t.Fatalf("Expected days -4,was %d",ts.Days())
	}
	if ts.Hours()!=-3 {
		t.Fatalf("Expected hours -3,was %d",ts.Hours())
	}
	
	if ts.Minutes()!=-2 {
		t.Fatalf("Expected minutes -2,was %d",ts.Minutes())
	}

	if ts.Seconds()!=-3 {
		t.Fatalf("Expected seconds -3,was %d",ts.Seconds())
	}

	if ts.Milliseconds()!=-4 {
		t.Fatalf("Expected milliseconds -4,was %d",ts.Milliseconds())
	}
}


func TestTotalsPeriods(t *testing.T) { 
	ts:=NewTimeSpan(1, 12, 0, 0, 0)
	
	if ts.TotalDays()!=1.5 {
		t.Fatalf("Expected days 1.5,was %f",ts.TotalDays())
	}

	ts=NewTimeSpan(0, 12, 30, 0, 0)
	
	if ts.TotalHours()!=12.5 {
		t.Fatalf("Expected hours 12.5,was %f",ts.TotalHours())
	}

	ts=NewTimeSpan(0, 0, 30, 30, 0)
	
	if ts.TotalMinutes()!=30.5 {
		t.Fatalf("Expected Minutes 30.5,was %f",ts.TotalMinutes())
	}

	ts=NewTimeSpan(0, 0, 0, 30, 500)
	
	if ts.TotalSeconds()!=30.5 {
		t.Fatalf("Expected Seconds 30.5,was %f",ts.TotalSeconds())
	}

	if ts.TotalMilliseconds()!=30500 {
		t.Fatalf("Expected Milliseconds 30500,was %f",ts.TotalMilliseconds())
	}

	
}





