package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Property struct {
	RoyaAddress      string `json:"Roya_address"`
	RoyaAge          string `json:"Roya_age"`
	RoyaApartment    string `json:"Roya_apartment"`
	RoyaArea         string `json:"Roya_area"`
	RoyaBathrooms    string `json:"Roya_bathrooms"`
	RoyaBedrooms     string `json:"Roya_bedrooms"`
	RoyaCategory     string `json:"Roya_category"`
	RoyaCity         string `json:"Roya_city"`
	RoyaCrDate       string `json:"Roya_crdate"`
	RoyaDeletedBy    string `json:"Roya_deleted_by"`
	RoyaDeletedOn    string `json:"Roya_deleted_on"`
	RoyaDescription  string `json:"Roya_description"`
	RoyaFloor        string `json:"Roya_floor"`
	RoyaFloorType    string `json:"Roya_floorType"`
	RoyaID           string `json:"Roya_id"`
	RoyaIsDelete     string `json:"Roya_is_delete"`
	RoyaKitchens     string `json:"Roya_kitchens"`
	RoyaPhone1       string `json:"Roya_phone1"`
	RoyaPhone2       string `json:"Roya_phone2"`
	RoyaPhone3       string `json:"Roya_phone3"`
	RoyaPrice        string `json:"Roya_price"`
	RoyaPropertyType string `json:"Roya_propertyType"`
	RoyaQuartier     string `json:"Roya_quartier"`
	RoyaRegion       string `json:"Roya_region"`
	RoyaStatus       string `json:"Roya_status"`
	RoyaTitle        string `json:"Roya_title"`
	RoyaTransaction  string `json:"Roya_transaction"`
	RoyaUserID       string `json:"Roya_userid"`
	RoyaValidated    string `json:"Roya_validated"`
}

type APIResponse struct {
	Nxtapott string     `json:"_nxtapott"`
	Response []Property `json:"response"`
	Status   string     `json:"status"`
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		properties := []Property{
			{
				RoyaAddress:      "شارع المدينة القديمة 123",
				RoyaAge:          "10",
				RoyaApartment:    "شقة رقم 5",
				RoyaArea:         "250.5",
				RoyaBathrooms:    "3",
				RoyaBedrooms:     "4",
				RoyaCategory:     "1",
				RoyaCity:         "أكادير",
				RoyaCrDate:       "2023-06-09 16:06:18",
				RoyaDeletedBy:    "%!s(<nil>)",
				RoyaDeletedOn:    "%!s(<nil>)",
				RoyaDescription:  "فيلا فاخرة بتصميم عصري، تتميز بإطلالة رائعة على المدينة القديمة والبحر. تحتوي على حديقة خضراء ومسبح خاص.",
				RoyaFloor:        "2",
				RoyaFloorType:    "%!s(<nil>)",
				RoyaID:           "1",
				RoyaIsDelete:     "%!s(<nil>)",
				RoyaKitchens:     "1",
				RoyaPhone1:       "+212123456789",
				RoyaPhone2:       "+212987654321",
				RoyaPhone3:       "+212555555555",
				RoyaPrice:        "1500000",
				RoyaPropertyType: "فيلا",
				RoyaQuartier:     "المدينة القديمة",
				RoyaRegion:       "%!s(<nil>)",
				RoyaStatus:       "%!s(<nil>)",
				RoyaTitle:        "فيلا فخمة للبيع في المدينة القديمة",
				RoyaTransaction:  "بيع",
				RoyaUserID:       "1",
				RoyaValidated:    "%!s(<nil>)",
			},
			{
				RoyaAddress:      "شارع المدينة القديمة 123",
				RoyaAge:          "10",
				RoyaApartment:    "شقة رقم 5",
				RoyaArea:         "250.5",
				RoyaBathrooms:    "3",
				RoyaBedrooms:     "4",
				RoyaCategory:     "2",
				RoyaCity:         "أكادير",
				RoyaCrDate:       "2023-06-09 18:28:55",
				RoyaDeletedBy:    "%!s(<nil>)",
				RoyaDeletedOn:    "%!s(<nil>)",
				RoyaDescription:  "فيلا فاخرة بتصميم عصري، تتميز بإطلالة رائعة على المدينة القديمة والبحر. تحتوي على حديقة خضراء ومسبح خاص.",
				RoyaFloor:        "2",
				RoyaFloorType:    "%!s(<nil>)",
				RoyaID:           "3",
				RoyaIsDelete:     "%!s(<nil>)",
				RoyaKitchens:     "1",
				RoyaPhone1:       "+212123456789",
				RoyaPhone2:       "+212987654321",
				RoyaPhone3:       "+212555555555",
				RoyaPrice:        "1500000",
				RoyaPropertyType: "فيلا",
				RoyaQuartier:     "المدينة القديمة",
				RoyaRegion:       "%!s(<nil>)",
				RoyaStatus:       "%!s(<nil>)",
				RoyaTitle:        "فيلا فخمة للبيع في المدينة القديمة",
				RoyaTransaction:  "بيع",
				RoyaUserID:       "1",
				RoyaValidated:    "%!s(<nil>)",
			},
			{
				RoyaAddress:      "شارع المدينة القديمة 123",
				RoyaAge:          "10",
				RoyaApartment:    "شقة رقم 5",
				RoyaArea:         "250.5",
				RoyaBathrooms:    "3",
				RoyaBedrooms:     "4",
				RoyaCategory:     "3",
				RoyaCity:         "أكادير",
				RoyaCrDate:       "2023-06-09 18:28:57",
				RoyaDeletedBy:    "%!s(<nil>)",
				RoyaDeletedOn:    "%!s(<nil>)",
				RoyaDescription:  "فيلا فاخرة بتصميم عصري، تتميز بإطلالة رائعة على المدينة القديمة والبحر. تحتوي على حديقة خضراء ومسبح خاص.",
				RoyaFloor:        "2",
				RoyaFloorType:    "%!s(<nil>)",
				RoyaID:           "4",
				RoyaIsDelete:     "%!s(<nil>)",
				RoyaKitchens:     "1",
				RoyaPhone1:       "+212123456789",
				RoyaPhone2:       "+212987654321",
				RoyaPhone3:       "+212555555555",
				RoyaPrice:        "1500000",
				RoyaPropertyType: "فيلا",
				RoyaQuartier:     "المدينة القديمة",
				RoyaRegion:       "%!s(<nil>)",
				RoyaStatus:       "%!s(<nil>)",
				RoyaTitle:        "فيلا فخمة للبيع في المدينة القديمة",
				RoyaTransaction:  "بيع",
				RoyaUserID:       "1",
				RoyaValidated:    "%!s(<nil>)",
			},
		}

		apiResponse := APIResponse{
			Nxtapott: "fzfN6OXw",
			Response: properties,
			Status:   "success",
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(apiResponse)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server is listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
