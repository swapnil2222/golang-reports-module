package main

import (
	"ReportModule/apis"
	"ReportModule/models"
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// AUTHOR: Swapnil Renge
// Created Date: 23-08-2019

// User can enter any one of the option from following.
// 1.Show All reports
// 2.Search reports by category
// 3.Search reports by date range

// Expected Input :
// ****** Enter your option : ******
// 1  OR
// 2 OR
// 3

// Expected Output : all reports based on option

var (
	port = flag.Int("port", 8081, "port on which server will be started..")
)

func main() {

	http.HandleFunc("/report", FetchReportsAPI)

	fmt.Println("Server started on", *port)
	fmt.Println("get reports by url : http://localhost" + fmt.Sprintf(":%d", *port) + "/report")

	// start server on mentioned port
	startServerErr := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if startServerErr != nil {
		fmt.Println("Error starting server.", startServerErr)
		return
	}
}

// FetchReportsAPI method : fetch reports based on option
func FetchReportsAPI(w http.ResponseWriter, r *http.Request) {

	// take user input
	reader := bufio.NewReader(os.Stdin)

	categoryName := ""
	startDate := ""
	endDate := ""

	// fmt.Println("*****************************")
	fmt.Println("****** Enter your option : ******")

	userOption, _ := reader.ReadString('\n')
	userOption = strings.Trim(userOption, "\n")

	if userOption == "2" {

		// take category name from user

		fmt.Println("Enter category name :")
		categoryName, _ = reader.ReadString('\n')

		// trim categoryName to remove \n delimeter

		categoryName = strings.Trim(categoryName, "\n")
	}

	if userOption == "3" {

		// take start and end date from user

		fmt.Println("Enter Start Date :")
		startDate, _ = reader.ReadString('\n')
		startDate = strings.Trim(startDate, "\n")

		fmt.Println("Enter End Date :")
		endDate, _ = reader.ReadString('\n')
		endDate = strings.Trim(endDate, "\n")

	}

	// fetch reports based on selected option

	reports, getReportErr := FetchReportsByOption(userOption, categoryName, startDate, endDate)

	if getReportErr != nil {
		fmt.Printf("Error getting reports %d", getReportErr)
		return
	}

	fmt.Println("Total Reports Count : ", len(reports))

	// return decoded reports data as a response..

	reportsByteData, marshalErr := json.Marshal(reports)
	if marshalErr != nil {
		fmt.Printf("Error getting reports %d", marshalErr)
		return
	}
	fmt.Fprintf(w, string(reportsByteData))
}

// FetchReportsByOption method : show reports to the user
func FetchReportsByOption(userOption, categoryName, startDate, endDate string) ([]models.Report, error) {

	// fetch all reports data from third party API
	reports, fetchErr := apis.FetchReportData()
	if fetchErr != nil {
		fmt.Println("Error fetching reports", fetchErr)
	}

	switch userOption {
	case "1":
		fmt.Println("fetching all reports......")
		return reports, nil
	case "2":
		fmt.Println("fetching categorywise reports...")
		return FetchCategortyWiseReports(reports, categoryName), nil
	case "3":
		fmt.Println("fetching datewise reports..")
		return FetchDateWiseReports(reports, startDate, endDate)
	default:
		fmt.Println("invalid option..")
		return []models.Report{}, nil
	}
}

// FetchDateWiseReports method
func FetchDateWiseReports(reports []models.Report, startDate, endDate string) ([]models.Report, error) {

	filteredReports := []models.Report{}

	for _, report := range reports {

		isWithinRange, isErr := IsDateWithinGivenRange(startDate, endDate, report)
		if isErr != nil {
			fmt.Println("Parsing date error", isErr)
			return []models.Report{}, nil
		}

		if isWithinRange {
			filteredReports = append(filteredReports, report)
		}
	}

	return filteredReports, nil
}

// IsDateWithinGivenRange method : will check if report is between given range date
func IsDateWithinGivenRange(startDate, endDate string, report models.Report) (bool, error) {

	//convert string date into time format using layout string

	parsedStartDate, _ := time.Parse("2006-1-2", startDate)
	parsedEndDate, _ := time.Parse("2006-1-2", endDate)

	//  condition to check

	if parsedStartDate.Before(report.DateReleased) && parsedEndDate.After(report.DateReleased) {
		return true, nil
	}

	return false, nil
}

// FetchCategortyWiseReports method : fetches report based on category name
func FetchCategortyWiseReports(reports []models.Report, categoryName string) []models.Report {

	// fetching categorywise reports

	filteredReports := []models.Report{}

	for _, report := range reports {
		if report.Category == categoryName {
			filteredReports = append(filteredReports, report)
		}
	}

	return filteredReports
}
