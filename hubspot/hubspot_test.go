package hubspot

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	apiKey = os.Getenv("HUBSPOT_API_KEY")
)

var _ = Describe("Create or update contact with invalid json input ", func() {

	os.Setenv("API_KEY", "66c43350-eb29-4ed9-b160-3e16611de672")

	hubspot := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(hubspot)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createOrUpdate", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateOrUpdateContact)
	handler.ServeHTTP(recorder, request)

	Describe("Create or Update contact", func() {
		Context("create or update", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create or update contact with all required param ", func() {

	os.Setenv("API_KEY", "66c43350-eb29-4ed9-b160-3e16611de672")

	hubspot := []byte(`{"email":"demot636@gmail.com","properties":{"properties":[{"property":"phone","value":"7894561230"},{"property":"firstname","value":"Nitin"},{"property":"lastname","value":"Patil"}]}}`)

	request, err := http.NewRequest("POST", "/createOrUpdate", bytes.NewBuffer(hubspot))
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateOrUpdateContact)
	handler.ServeHTTP(recorder, request)

	Describe("Create or Update contact", func() {
		Context("create or update", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Get contact with all required param ", func() {

	os.Setenv("API_KEY", "66c43350-eb29-4ed9-b160-3e16611de672")

	hubspot := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(hubspot)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/getContact", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetContactByVID)
	handler.ServeHTTP(recorder, request)

	Describe("Get contact", func() {
		Context("contact by vid", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get contact with all required param ", func() {

	os.Setenv("API_KEY", "66c43350-eb29-4ed9-b160-3e16611de672")

	hubspot := []byte(`{"vid":"651"}`)

	request, err := http.NewRequest("POST", "/getContact", bytes.NewBuffer(hubspot))
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetContactByVID)
	handler.ServeHTTP(recorder, request)

	Describe("Get contact", func() {
		Context("contact by vid", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Delete contact with invalid param ", func() {

	os.Setenv("API_KEY", "66c43350-eb29-4ed9-b160-3e16611de672")

	hubspot := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(hubspot)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/deleteContact", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteContactByVID)
	handler.ServeHTTP(recorder, request)

	Describe("delete contact", func() {
		Context("delete contact by vid", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Delete contact with all required param ", func() {

	os.Setenv("API_KEY", "66c43350-eb29-4ed9-b160-3e16611de672")

	hubspot := []byte(`{"vid":"651"}`)

	request, err := http.NewRequest("POST", "/deleteContact", bytes.NewBuffer(hubspot))
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteContactByVID)
	handler.ServeHTTP(recorder, request)

	Describe("Delete contact", func() {
		Context("Delete contact by vid", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Create deal with invalid json input ", func() {

	os.Setenv("API_KEY", "66c43350-eb29-4ed9-b160-3e16611de672")

	hubspot := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(hubspot)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createDeal", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateDeal)
	handler.ServeHTTP(recorder, request)

	Describe("Create deal", func() {
		Context("create deal", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create deal with all required param ", func() {

	os.Setenv("API_KEY", "66c43350-eb29-4ed9-b160-3e16611de672")

	hubspot := []byte(`{"associations": {"associatedVids":[601]},"properties": [{"value": "Tims Newer Deal","name": "dealname"},{"value": "appointmentscheduled","name": "dealstage"},{"value": "default","name": "pipeline"},{"value": 1409443200000,"name": "closedate"},{"value": "60000","name": "amount"},{"value": "newbusiness","name": "dealtype"}]}`)

	request, err := http.NewRequest("POST", "/createDeal", bytes.NewBuffer(hubspot))
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateDeal)
	handler.ServeHTTP(recorder, request)

	Describe("Create Deal", func() {
		Context("create deal", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Create ticket with invalid json input ", func() {

	os.Setenv("API_KEY", "66c43350-eb29-4ed9-b160-3e16611de672")

	hubspot := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(hubspot)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createTicket", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTicket)
	handler.ServeHTTP(recorder, request)

	Describe("Create ticket", func() {
		Context("create ticket", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create ticket with all required param ", func() {

	os.Setenv("API_KEY", "66c43350-eb29-4ed9-b160-3e16611de672")

	hubspot := []byte(`{"ticketProperties":[{"name": "subject","value": "This is an example ticket"},{"name": "content","value": "Here are the details of the ticket."},{"name": "hs_pipeline","value": "0"},{"name": "hs_pipeline_stage","value": "1"}]}`)

	request, err := http.NewRequest("POST", "/createTicket", bytes.NewBuffer(hubspot))
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTicket)
	handler.ServeHTTP(recorder, request)

	Describe("Create Ticket", func() {
		Context("create ticket", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
