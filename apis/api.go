package apis

import (
	"ReportModule/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const reportURL = "https://thereportoftheweek-api.herokuapp.com/reports"

// FetchReportData method will fetch reports from third party API.
func FetchReportData() ([]models.Report, error) {

	reports := []models.Report{}

	resp, respErr := http.Get(reportURL)
	if respErr != nil {
		log.Println("Error getting report data", respErr)
		return reports, respErr
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		byteData, byteErr := ioutil.ReadAll(resp.Body)
		if byteErr != nil {
			log.Println("Error while reading response", byteErr)
			return reports, byteErr
		}

		unMarshalErr := json.Unmarshal(byteData, &reports)
		if unMarshalErr != nil {
			log.Println("Error while marshalling response", unMarshalErr)
			return reports, unMarshalErr
		}
		return reports, nil
	}
	// return resp.Body.Read(), nil
	return reports, errors.New("Error status code" + string(resp.StatusCode))
}
