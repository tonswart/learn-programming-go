package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate random passwords",
	Long: `Generate random passwords with customizable options.
For example:

passwordGen generate -l 12 -d -s`,

	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Length of the generated password")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in the generated password")
	generateCmd.Flags().BoolP("special-chars", "s", false, "Include special characters in the generated password")
}

func generatePassword(cmd *cobra.Command, args []string) {

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

	fmt.Println("Generated Password")
	fmt.Println(string(password))

}
