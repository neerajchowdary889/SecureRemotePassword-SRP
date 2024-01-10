package client

import(
	"fmt"
	"reflect"
)

func(user *ClientDetails) checkError() (bool,string){
    val := reflect.ValueOf(user).Elem()

    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)

        // Check if the field is a pointer
        if field.Kind() == reflect.Ptr {
            // Check if the pointer is nil
            if field.IsNil() {
                fmt.Printf("Error: %s field is nil\n", val.Type().Field(i).Name)
                return false, val.Type().Field(i).Name
            }
        }
    }

    return true, ""
}

func(user *ClientDetails) SendToServer()(map[string]interface{}, error){
	status, field := user.checkError()
	if !status{
		return nil, fmt.Errorf("error: %s field is nil", field)
	}else{
		//converting struct to hashmap
		hashmap := make(map[string]interface{})

		// Get the reflect.Value of the struct
		val := reflect.ValueOf(user).Elem()
	
		// Iterate over the struct fields
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
	
			// Add the field name and value to the map
			hashmap[val.Type().Field(i).Name] = field.Interface()

			
		}
		return hashmap, nil

	}
}