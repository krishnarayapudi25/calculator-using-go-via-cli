package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "math",
		Short: "Perform mathematical operations",
		Long:  "Perform various mathematical operations, such as addition, subtraction, multiplication, and division",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("No operation specified. Use one of the following commands: add, subtract, multiply, divide")
		},
	}

	var addCommand = &cobra.Command{
		Use:   "add",
		Short: "Add two numbers",
		Long:  "Add two numbers and print the result",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println("Error: Two numbers are required for addition")
				os.Exit(1)
			}

			var x, y float64
			var err error

			x, err = strconv.ParseFloat(args[0], 64)
			if err != nil {
				fmt.Println("Error: Invalid number:", args[0])
				os.Exit(1)
			}

			y, err = strconv.ParseFloat(args[1], 64)
			if err != nil {
				fmt.Println("Error: Invalid number:", args[1])
				os.Exit(1)
			}

			result := x + y
			fmt.Printf("%g + %g = %g\n", x, y, result)
		},
	}

	var subtractCommand = &cobra.Command{
		Use:   "subtract",
		Short: "Subtract two numbers",
		Long:  "Subtract two numbers and print the result",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println("Error: Two numbers are required for subtraction")
				os.Exit(1)
			}

			var x, y float64
			var err error

			x, err = strconv.ParseFloat(args[0], 64)
			if err != nil {
				fmt.Println("Error: Invalid number:", args[0])
				os.Exit(1)
			}

			y, err = strconv.ParseFloat(args[1], 64)
			if err != nil {
				fmt.Println("Error: Invalid number:", args[1])
				os.Exit(1)
			}

			result := x - y
			fmt.Printf("%g - %g = %g\n", x, y, result)
		},
	}

	var multiplyCommand = &cobra.Command{
		Use:   "multiply",
		Short: "Multiply two numbers",
		Long:  "Multiply two numbers and print the result",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println("Error: Two numbers are required for multiplication")
				os.Exit(1)
			}

			var x, y float64
			var err error

			x, err = strconv.ParseFloat(args[0], 64)
			if err != nil {
				fmt.Println("Error: Invalid number:", args[0])
				os.Exit(1)
			}

			y, err = strconv.ParseFloat(args[1], 64)
			if err != nil {
				fmt.Println("Error: Invalid number:", args[1])
				os.Exit(1)
			}

			result := x * y
			fmt.Printf("%g * %g = %g\n", x, y, result)
		},
	}

	var divideCommand = &cobra.Command{
		Use:   "divide",
		Short: "Divide two numbers",
		Long:  "Divide two numbers and print the result",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println("Error: Two numbers are required for division")
				os.Exit(1)
			}

			var x, y float64
			var err error

			x, err = strconv.ParseFloat(args[0], 64)
			if err != nil {
				fmt.Println("Error: Invalid number:", args[0])
				os.Exit(1)
			}

			y, err = strconv.ParseFloat(args[1], 64)
			if err != nil {
				fmt.Println("Error: Invalid number:", args[1])
				os.Exit(1)
			}

			if y == 0 {
				fmt.Println("Error: Cannot divide by zero")
				os.Exit(1)
			}

			result := x / y
			fmt.Printf("%g / %g = %g\n", x, y, result)
		},
	}

	rootCmd.AddCommand(addCommand)
	rootCmd.AddCommand(subtractCommand)
	rootCmd.AddCommand(multiplyCommand)
	rootCmd.AddCommand(divideCommand)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
