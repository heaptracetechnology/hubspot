package hubspot

import (
	"encoding/json"
	"fmt"
	result "github.com/heaptracetechnology/hubspot/result"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//getAPIKey Hubspot
func getAPIKey() string {
	return os.Getenv("API_KEY")
}

//CreateOrUpdateContact Hubspot
func CreateOrUpdateContact(responseWriter http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	param := make(map[string]interface{})

	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	reqBody, _ := json.Marshal(param["properties"])
	email := fmt.Sprint(param["email"])
	resp, err := http.Post("https://api.hubapi.com/contacts/v1/contact/createOrUpdate/email/"+email+"/?hapikey="+getAPIKey(), "application/json", strings.NewReader(string(reqBody)))
	if err == nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(err)
	}

	result.WriteJsonResponse(responseWriter, body, http.StatusOK)
}

//GetContactByVID Hubspot
func GetContactByVID(responseWriter http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	param := make(map[string]interface{})

	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	vid := fmt.Sprint(param["vid"])

	resp, err := http.Get("https://api.hubapi.com/contacts/v1/contact/vid/" + vid + "/profile?hapikey=" + getAPIKey())
	if err == nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(err)
	}

	result.WriteJsonResponse(responseWriter, body, http.StatusOK)
}

//DeleteContactByVID Hubspot
func DeleteContactByVID(responseWriter http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	param := make(map[string]interface{})

	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	vid := fmt.Sprint(param["vid"])

	// Create client
	client := &http.Client{}

	resp, err := http.NewRequest("DELETE", "https://api.hubapi.com/contacts/v1/contact/vid/"+vid+"?hapikey="+getAPIKey(), nil)
	if err == nil {
		fmt.Println(err)
	}
	// Fetch Request
	deleteResp, err := client.Do(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer deleteResp.Body.Close()
	body, err := ioutil.ReadAll(deleteResp.Body)
	if err == nil {
		fmt.Println(err)
	}

	result.WriteJsonResponse(responseWriter, body, http.StatusOK)
}

//CreateDeal Hubspot
func CreateDeal(responseWriter http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	param := make(map[string]interface{})

	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	reqBody, _ := json.Marshal(param["properties"])

	resp, err := http.Post("https://api.hubapi.com/deals/v1/deal?hapikey="+getAPIKey(), "application/json", strings.NewReader(string(reqBody)))
	if err == nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(err)
	}

	result.WriteJsonResponse(responseWriter, body, http.StatusOK)
}

//CreateTicket Hubspot
func CreateTicket(responseWriter http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	param := make(map[string]interface{})

	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	reqBody, _ := json.Marshal(param["ticketProperties"])

	resp, err := http.Post("https://api.hubapi.com/crm-objects/v1/objects/tickets?hapikey="+getAPIKey(), "application/json", strings.NewReader(string(reqBody)))
	if err == nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(err)
	}

	result.WriteJsonResponse(responseWriter, body, http.StatusOK)
}
