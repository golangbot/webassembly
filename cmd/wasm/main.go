package main

import (
    "fmt"
    "encoding/json"
    "syscall/js"
)

func jsonWrapper() js.Func {  
    jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) != 1 {
            result := map[string]interface{}{
                "error": "Invalid no of arguments passed",
            }
            return result
        }
        jsDoc := js.Global().Get("document")
        if !jsDoc.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get document object",
            }
            return result
        }
        jsonOuputTextArea := jsDoc.Call("getElementById", "jsonoutput")
        if !jsonOuputTextArea.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get output text area",
            }
            return result
        }
        inputJSON := args[0].String()
        fmt.Printf("input %s\n", inputJSON)
        pretty, err := prettyJson(inputJSON)
        if err != nil {
            errStr := fmt.Sprintf("unable to parse JSON. Error %s occurred\n", err)
            result := map[string]interface{}{
                "error": errStr,
            }
            return result
        }
        jsonOuputTextArea.Set("value", pretty)
        return nil
    })
    return jsonfunc
}

func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func main() {  
    fmt.Println("Go Web Assembly")
    js.Global().Set("formatJSON", jsonWrapper())
    <-make(chan bool)
}
