/*
	GoDate - A Go package for date  and time manipulation
    Copyright (C) 2010 Andrea Parodi

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package GoDate

import "testing"
import "time"

func TestDateCreation(t *testing.T) { 
	d:=NewDate(2010,8,1)
	
	if d.Year()!=2010 {
		t.Fatalf("Expected 2010,was %d",d.Year())
	}

	if d.Month()!=8 {
		t.Fatalf("Expected 8,was %d",d.Month())
	}
	
	if d.Day()!=1 {
		t.Fatalf("Expected 1,was %d",d.Day())
	}

}


func TestTimeCreation(t *testing.T) { 
	d:=NewTime(10,8,1)
	
	if d.Hour()!=10 {
		t.Fatalf("Expected 10,was %d",d.Hour())
	}

	if d.Minute()!=8 {
		t.Fatalf("Expected 8,was %d",d.Minute())
	}
	
	if d.Second()!=1 {
		t.Fatalf("Expected 1,was %d",d.Second())
	}

}


func TestDateFunc(t *testing.T) { 
	d:=NewDateTime(2010,8,1,10,8,1).Date()
	if d.Year()!=2010 {
		t.Fatalf("Expected 2010,was %d",d.Year())
	}

	if d.Month()!=8 {
		t.Fatalf("Expected 8,was %d",d.Month())
	}
	
	if d.Day()!=1 {
		t.Fatalf("Expected 1,was %d",d.Day())
	}
	
	if d.Hour()!=0 {
		t.Fatalf("Expected 0,was %d",d.Hour())
	}

	if d.Minute()!=0 {
		t.Fatalf("Expected 0,was %d",d.Minute())
	}
	
	if d.Second()!=0 {
		t.Fatalf("Expected 0,was %d",d.Second())
	}
}

func TestDateTimeCreation(t *testing.T) { 
	d:=NewDateTime(2010,8,1,10,8,1)
	if d.Year()!=2010 {
		t.Fatalf("Expected 2010,was %d",d.Year())
	}

	if d.Month()!=8 {
		t.Fatalf("Expected 8,was %d",d.Month())
	}
	
	if d.Day()!=1 {
		t.Fatalf("Expected 1,was %d",d.Day())
	}
	
	if d.Hour()!=10 {
		t.Fatalf("Expected 10,was %d",d.Hour())
	}

	if d.Minute()!=8 {
		t.Fatalf("Expected 8,was %d",d.Minute())
	}
	
	if d.Second()!=1 {
		t.Fatalf("Expected 1,was %d",d.Second())
	}

}


func TestStringer(t *testing.T) { 
	d:=NewDateTime(2010,8,2,9,8,1)
	if d.String()!="Sun Aug  2 09:08:01 +0000 2010" {
		t.Fatalf("was %s",d.String())
	}
	if d.Format("02/01/2006 15:04:05")!="02/08/2010 09:08:01" {
		t.Fatalf("was %s",d.Format("02/01/2006 15:04:05"))
	}
	
	
}

func TestTimeSpanOperation(t *testing.T) { 
	d:=NewDateTime(2010,8,2,9,8,1)+DateTime(FromHours(1)-FromDays(1))
	if d.Year()!=2010 {
		t.Fatalf("Expected 2010,was %d",d.Year())
	}

	if d.Month()!=8 {
		t.Fatalf("Expected 8,was %d",d.Month())
	}
	
	if d.Day()!=1 {
		t.Fatalf("Expected 1,was %d",d.Day())
	}
	
	if d.Hour()!=10 {
		t.Fatalf("Expected 10,was %d",d.Hour())
	}

	if d.Minute()!=8 {
		t.Fatalf("Expected 8,was %d",d.Minute())
	}
	
	if d.Second()!=1 {
		t.Fatalf("Expected 1,was %d",d.Second())
	}


}

func TestFromGoTime(t *testing.T) { 
	dGo:= time.Time{Year:2010,Month:8,Day:1,Hour:10,Minute:8,Second:1}
	d:=FromGoTime(dGo)
	if d.Year()!=2010 {
		t.Fatalf("Expected 2010,was %d",d.Year())
	}

	if d.Month()!=8 {
		t.Fatalf("Expected 8,was %d",d.Month())
	}
	
	if d.Day()!=1 {
		t.Fatalf("Expected 1,was %d",d.Day())
	}
	
	if d.Hour()!=10 {
		t.Fatalf("Expected 10,was %d",d.Hour())
	}

	if d.Minute()!=8 {
		t.Fatalf("Expected 8,was %d",d.Minute())
	}
	
	if d.Second()!=1 {
		t.Fatalf("Expected 1,was %d",d.Second())
	}

}




