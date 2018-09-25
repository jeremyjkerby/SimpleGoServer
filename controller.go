// controller.go
// Jeremy J. Kerby
// Process data as needed in and out of the system
package main

import (
    "log"
    "strconv"
    "strings"
)

// We know the type of data we are expecting
type Record struct {
    Exam        float64     `json: "exam"`
    StudentId   string      `json: "studentId"`
    Score       float64     `json: "score"`
}

// initialize data var as slice Record struct
var data []Record

// save record into slice as we get them
func saveRecord(r map[string]interface{}) {
    log.Print("saveRecord(): ", r)

    var sid string
    var exam, score float64
    // perform some error handling
    if str, ok := r["studentId"].(string); ok {
        sid = str
    } else {
        log.Print("Error: saveRecord(): Not given string type.")
    }
    // perform some error handling
    if eStr, ok := r["exam"].(float64); ok {
        exam = eStr
    } else {
        log.Print("Error: saveRecord(): Not given float64 type.")
    }
    // perform some error handling
    if sStr, ok := r["score"].(float64); ok {
        score = sStr
    } else {
        log.Print("Error: saveRecord(): Not given float64 type.")
    }

    data = append(data, Record{Exam: exam, StudentId: sid, Score: score}) // append new record
}

// return all current known students
func getAllStudentData() []string {
    studs := make(map[string]bool)
    for _, d := range data {
        studs[d.StudentId] = true
    }

    // make unique map of student names
    ids := make([]string, 0) // create a new slice
    for d := range studs {
        ids = append(ids, d)
    } 

    return ids
}

// return all current exams and avaerage score for given student
func getStudentData(id string) (float64, []Record) {
    log.Print("getStudentData(): id :> ", id)

    scores := make([]Record, 0) // create a new slice
    for _, d := range data {
        if d.StudentId == id {
            scores = append(scores, d)
        }
    }

    var scoresLen int = len(scores) // calculate average
    var scoresSum float64 = 0
    for _, s := range scores {
        scoresSum = scoresSum + s.Score
    }

    var scoresAvg float64 = calculateAverage(scoresSum, float64(scoresLen))

    return scoresAvg, scores // return slice and average
}

// return all current known exams
func getAllExamData() []float64 {
    studs := make(map[float64]bool)
    for _, d := range data {
        studs[d.Exam] = true
    }

    ids := make([]float64, 0) // create a new slice
    for d := range studs {
        ids = append(ids, d)
    } 

    return ids 
}

// return all current exam results from all students and avaerage score for given exam number 
func getExamData(number string) (float64, []Record){
    log.Print("getExamData(): number :> ", number)

    // trim special chars and convert to float64
    numberFloat, _ := strconv.ParseFloat(strings.TrimSpace(number), 64)

    scores := make([]Record, 0) // create a new slice
    for _, d := range data {
        if d.Exam == numberFloat {
            scores = append(scores, d)
        }
    }

    var scoresLen int = len(scores) // calculate average
    var scoresSum float64 = 0
    for _, s := range scores {
        scoresSum = scoresSum + s.Score
    }

    var scoresAvg float64 = calculateAverage(scoresSum, float64(scoresLen))

    return scoresAvg, scores // return slice and average
}

// helper function to calculate average score when needed
func calculateAverage(sSum float64, sLen float64) (float64) {
    var avg float64 = sSum / sLen
    return avg
}
