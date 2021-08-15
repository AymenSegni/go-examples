/*
Copyright © 2021 Aymen Segni: segniaymen1@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will return the requested Gopher if it exist",
	Long: "This get command will browse scraly/gophers repository and fetch the desired Gopher .",
	Run: func(cmd *cobra.Command, args []string) {
		// init the Gopher with var
		var gopherName = "dr-who.png"
		
		// make sure that the gopher given by the user isn't empty
		if len (args) >= 1 && args[0] != "" {
			gopherName = args[0]
		}

		// define the gopher URL

		URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"
		fmt.Println("get'" + gopherName + "'Gopher ..")

		// get the gopher content 
		response, err := http.Get(URL)

		// error handling
		if err != nil {
			fmt.Println(err)
		}
		// close function connection
		defer response.Body.Close()

		// in case we get the desired gopher
		if response.StatusCode == 200 {
			// create the gopher file locally 
			out, err := os.Create(gopherName + ".png")
		
		    defer out.Close()

			// write body file
			_, err = io.Copy(out, response.Body)
			if err != nil {
				fmt.Println(err)
			}

			// log the succeed msg 
			fmt.Println("gopher is downloaded and saved successfly in " + out.Name())

    	} else {
            fmt.Println("Error: " + gopherName + " not exists! :-(")
        }

	},
}

// add getCmd comand
func init() {
	rootCmd.AddCommand(getCmd)
}