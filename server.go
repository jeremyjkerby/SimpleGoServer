// server.go
// Jeremy J. Kerby
// Server logic routes and handlers
package main

import (
    "bufio"
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "strings"
)

func readInData() {
    var count int = 0
    var choice int = // configure
    var url string = // configure
    resp, err := http.Get(url)

    if err != nil {
        log.Fatal(err)
    } else {
        reader := bufio.NewReader(resp.Body)

        for {
            line, err := reader.ReadBytes('\n')
            if err != nil {
                log.Fatal(err)
            } else {
                if strings.Contains(string(line), "data") {
                    recordMap := make(map[string]interface{})

                    stuff := string(line[6:]) // trim line
                    json.Unmarshal([]byte(stuff), &recordMap)
                    saveRecord(recordMap) // send data to controller

                    count = count + 1
                    if count >= choice {
                        break
                    }
                }
            }
        }
    }
}

//  list all students that have received at least one test score
func getStudents (w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            students := getAllStudentData()
            payload := map[string][]string{} // form payload
            payload["students"] = students 

            log.Print("getStudents(): payload: ", payload)
            w.Header().Set("Content-Type", "application/json") // serve json not text
            json.NewEncoder(w).Encode(payload) // serve json
        default:
            var msg string = "HTTP/1.1 405 - Method Not Allowed"

            log.Print("getStudents(): ", msg)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(msg)
    }
}

// list test results for student, provides student's avg score for all exams
func getStudent (w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            vars := mux.Vars(r) // get vars
            val1, val2 := getStudentData(vars["id"])
            payload := make(map[string]interface{})
            payload["average"] = val1
            payload["scores"] = val2

            log.Print("getStudent(): payload: ", payload)
            w.Header().Set("Content-Type", "application/json") // serve json not text
            json.NewEncoder(w).Encode(payload) // serve json
        default:
            var msg string = "HTTP/1.1 405 - Method Not Allowed"

            log.Print("getStudent(): ", msg)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(msg)
    }
}

// list all exams that have been recorded
func getExams (w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            examsData := getAllExamData()
            payload := map[string][]float64{} // form payload
            payload["exams"] = examsData

            log.Print("getExams(): payload: ", payload)
            w.Header().Set("Content-Type", "application/json") // serve json not text
            json.NewEncoder(w).Encode(payload) // serve json
        default:
            var msg string = "HTTP/1.1 405 - Method Not Allowed"

            log.Print("getExams(): ", msg)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(msg)
    }
}

// list all results for exam, provides avg score across all students
func getExam (w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            vars := mux.Vars(r) // get vars
            val1, val2 := getExamData(vars["number"])
            payload := make(map[string]interface{})
            payload["average"] = val1
            payload["scores"] = val2

            log.Print("getExam(): payload: ", payload)
            w.Header().Set("Content-Type", "application/json") // serve json not text
            json.NewEncoder(w).Encode(payload) // serve json
        default:
            var msg string = "HTTP/1.1 405 - Method Not Allowed"

            log.Print("getExam(): ", msg)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(msg)
    }
}

func main() {
    r := mux.NewRouter() // initialize router
    
    log.Print("Reading in data...")
    readInData()
    log.Print("Data read in complete.")

    // register route handlers
    r.HandleFunc("/api/students", getStudents)
    r.HandleFunc("/api/students/{id}", getStudent)
    r.HandleFunc("/api/exams", getExams)
    r.HandleFunc("/api/exams/{number}", getExam)

    log.Print("Server is now listening...")
    http.ListenAndServe(":9000", r) // execute server
}
