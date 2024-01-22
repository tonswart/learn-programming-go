package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize random passwords",
	Long: `Optimize random passwords with customizable options.
For example:

passwordGen optimize -t`,

	Run: optimizePassword,
}

func init() {
	rootCmd.AddCommand(optimizeCmd)

	optimizeCmd.Flags().IntP("length", "l", 8, "Length of the generated password")
	optimizeCmd.Flags().BoolP("digits", "d", false, "Include digits in the generated password")
	optimizeCmd.Flags().BoolP("special-chars", "s", false, "Include special characters in the generated password")
}

func optimizePassword(cmd *cobra.Command, args []string) {

	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if isDigits {
		charset += "0123456789"
	}
	if isSpecialChars {
		charset += "!@#$%^&*()_+{}[]|;:,.<>?-="
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

    fmt.Println("Hello")
	fmt.Println("Optimized Password")
	fmt.Println(string(password))
}
