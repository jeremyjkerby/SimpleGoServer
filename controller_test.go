// controller_test.go
// Jeremy J. Kerby
// Simple test cases for controller.go
// This is by no means complete test suite. Should test all possible edge cases here and server.go
package main

import (
    "testing"
    "reflect"
)

// test controller getAllStudentData() with no data
func TestGetAllStudentDataEmpty(t *testing.T) {
    students := getAllStudentData()

    var actual int = len(students)
    var expected int = 0
    if actual != expected {
       t.Errorf("Length was incorrect, got: %v, want: %v", actual, expected)
   }
}

// test controller getAllStudentData()
func TestGetAllStudentDataMany(t *testing.T) {
    // add some data
    data = append(data, Record{Exam: 1234, StudentId: "timmy.test1", Score: 0.7847112560881939})
    data = append(data, Record{Exam: 1234, StudentId: "timmy.test2", Score: 0.8847112560881939})
    data = append(data, Record{Exam: 1234, StudentId: "timmy.test3", Score: 0.9847112560881939})

    actual := getAllStudentData()
    expected := []string{"timmy.test1", "timmy.test2", "timmy.test3"}
    if !reflect.DeepEqual(actual, expected) {
       t.Errorf("Elements were not the same, got: %d, want: %d.", actual, expected)
   }
}

// test controller calculateAverage()
func TestCalculateAverage(t *testing.T) {
    cases := []struct { // build simple and common test case data
        n   string
        a   float64 
        b   float64
        e   float64
    }{{"Basic", 10, 2, 5}, {"Common", 1.6694225121763878, 2, 0.8347112560881939}}

    for _, c := range cases { // run those cases
        t.Run(c.n, func(t *testing.T) {
            actual := calculateAverage(c.a, c.b)
            expected := c.e
            if actual !=  expected {
                t.Errorf("%v average was incorrect, got: %v, want: %v", c.n, actual, expected)
            }
        })
    }
}
