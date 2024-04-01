package main

import (
	"github.com/gin-gonic/gin"
	"bytes"
	"encoding/json"
	"net/http"
	"io"
	"os"
	"github.com/joho/godotenv"
)

// postRequest sends a POST request to the specified URL with the provided data and headers.
func postRequest(url string, data map[string]interface{}, headers map[string]string) (*http.Response, error) {
	// Convert the data into JSON format.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Add the headers to the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request
	client := &http.Client{}
	return client.Do(req)
}


func getNearBy(c *gin.Context){
	// apiKey（GOOGLE_PLACE_API_KEY）
	godotenv.Load(".env")
	apiKey := os.Getenv("GOOGLE_PLACE_API_KEY")
	url := "https://places.googleapis.com/v1/places:searchNearby"
	data := map[string]interface{}{
		"includedTypes": []string{"restaurant"},
		"maxResultCount": 10,
		"locationRestriction": map[string]interface{}{
			"circle": map[string]interface{}{
				"center": map[string]float64{
					"latitude": 35.4428818,
					"longitude": 137.016063,
				},
				"radius": 500.0,
			},
		},
		"languageCode": "ja",
	}
	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Goog-Api-Key": apiKey,
		"X-Goog-FieldMask": "places.addressComponents,places.adrFormatAddress,places.businessStatus,places.displayName,places.formattedAddress,places.googleMapsUri,places.iconBackgroundColor,places.iconMaskBaseUri,places.id,places.location,places.name,places.photos,places.plusCode,places.primaryType,places.primaryTypeDisplayName,places.primaryTypeDisplayName,places.shortFormattedAddress,places.subDestinations,places.types,places.utcOffsetMinutes,places.viewport,places.currentOpeningHours,places.nationalPhoneNumber,places.priceLevel,places.rating,places.userRatingCount,places.websiteUri",
	}
	// response := postRequest(url, data, headers)
	// c.JSON(200, gin.H{
	// 	"data": response,
	// })
	response, err := postRequest(url, data, headers)
	if err != nil {
		// ここでエラーを処理する。例えば、クライアントにエラーメッセージを返す等
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return // エラーがあった場合、処理をここで終了させる
	}

	// レスポンスを処理する
	// 例: レスポンスボディを読み取る、デコードする等
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		// レスポンスボディの読み取りエラーを処理
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read response body",
		})
		return
	}

	// 必要に応じてレスポンスボディを使用して処理を行う
	// ここでは例としてレスポンスボディをそのままクライアントに返すことを想定
	c.JSON(http.StatusOK, gin.H{
		"data": string(body),
	})

}

