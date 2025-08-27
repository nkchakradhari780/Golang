package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 status code:", resp.StatusCode)
		return
	}

	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }
	// fmt.Println("Response Data is : ", string(data))

	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Todo Item:")
	fmt.Println("Todo Item:", todo)
	fmt.Println("UserId:", todo.UserId)
	fmt.Println("Id:", todo.Id)
	fmt.Println("Title:", todo.Title)
	fmt.Println("Completed:", todo.Completed)
}

func performPostRequest() {
	todo := Todo{
		UserId:    23,
		Title:     "New Todo Item",
		Completed: false,
	}

	//convert the Todo to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myUrl, "application/json", jsonReader)

	if err != nil {
		fmt.Println("Error post Data: ", err)
		return
	}

	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response: ", string(data))
	fmt.Println("Response Status", res.Status)


}

func performPutRequest() {
	todo := Todo{
		UserId:    4546,
		Title:     "Updated Title",
		Completed: false,
	}

	//convert the Todo to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response: ", string(data))
	fmt.Println("Response Status", res.Status)
}

func performDeleteRequest() {
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response: ", string(data))
	fmt.Println("Response Status", res.Status)
}

func main() {
	fmt.Println("Learning CRUD....")
	performGetRequest()
	performPostRequest()
	performPutRequest()
	performDeleteRequest()
}
