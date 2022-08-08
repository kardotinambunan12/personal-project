package lib

// import (
// 	"bytes"
// 	"net/http"
// )

// // ==> Receipt Base64
// func UploadFileBase64(base_64, image_name, url_api, path string) ResponseUpload {
// 	var requestBody bytes.Buffer

// 	multiPartWriter := multipart.NewWriter(&requestBody)

// 	// check_connection_microservices
// 	microservices_connect := CheckConnection(GetEnv("HOST_CDN_IMAGE") + GetEnv("URI_CDN_CHECK_CONNECTION"))
// 	if microservices_connect == false {
// 		response := ResponseUpload{
// 			Status: ResponseStatus{
// 				Code:    503,
// 				Message: "service_unavailable",
// 			},
// 		}
// 		return response
// 	}

// 	// Image Form
// 	fieldImageName, err := multiPartWriter.CreateFormField("image_name")
// 	if err != nil {
// 		// logs.Println(err)
// 		response := ResponseUpload{
// 			Status: ResponseStatus{
// 				Code:    500,
// 				Message: "internal_server_error",
// 			},
// 		}
// 		return response
// 	}
// 	_, err = fieldImageName.Write([]byte(image_name))
// 	if err != nil {
// 		// logs.Println(err)
// 		response := ResponseUpload{
// 			Status: ResponseStatus{
// 				Code:    500,
// 				Message: "internal_server_error",
// 			},
// 		}
// 		return response
// 	}

// 	fieldBase64, err := multiPartWriter.CreateFormField("base_64")
// 	if err != nil {
// 		// logs.Println(err)
// 		response := ResponseUpload{
// 			Status: ResponseStatus{
// 				Code:    500,
// 				Message: "internal_server_error",
// 			},
// 		}
// 		return response
// 	}
// 	_, err = fieldBase64.Write([]byte(base_64))
// 	if err != nil {
// 		// logs.Println(err)
// 		response := ResponseUpload{
// 			Status: ResponseStatus{
// 				Code:    500,
// 				Message: "internal_server_error",
// 			},
// 		}
// 		return response
// 	}

// 	// Additional Form
// 	fieldPath, err := multiPartWriter.CreateFormField("path")
// 	if err != nil {
// 		// logs.Println(err)
// 		response := ResponseUpload{
// 			Status: ResponseStatus{
// 				Code:    500,
// 				Message: "internal_server_error",
// 			},
// 		}
// 		return response
// 	}
// 	_, err = fieldPath.Write([]byte(path))
// 	if err != nil {
// 		// logs.Println(err)
// 		response := ResponseUpload{
// 			Status: ResponseStatus{
// 				Code:    500,
// 				Message: "internal_server_error",
// 			},
// 		}
// 		return response
// 	}

// 	multiPartWriter.Close()

// 	// By now our original request body should have been populated, so let's just use it with our custom request
// 	req, err := http.NewRequest("POST", url_api, &requestBody)
// 	if err != nil {
// 		// logs.Println(err)
// 		response := ResponseUpload{
// 			Status: ResponseStatus{
// 				Code:    500,
// 				Message: "internal_server_error",
// 			},
// 		}
// 		return response
// 	}

// 	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

// 	client := &http.Client{}
// 	response, err := client.Do(req)
// 	if err != nil {

// 		response := ResponseUpload{
// 			Status: ResponseStatus{
// 				Code:    500,
// 				Message: "internal_server_error",
// 			},
// 		}
// 		return response
// 	}

// 	var result map[string]interface{}
// 	json.NewDecoder(response.Body).Decode(&result)
// 	file_name := fmt.Sprintf("%v", result["data"])

// 	response_data := ResponseUpload{
// 		Filename:     file_name,
// 		RelativePath: path + "/" + file_name,
// 		Status: ResponseStatus{
// 			Code:    200,
// 			Message: "OK",
// 		},
// 	}

// 	return response_data
// }
