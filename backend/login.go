package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")

		forceRemoveAuthentication, _ := cmd.Flags().GetBool("force")

		if forceRemoveAuthentication {
			fmt.Println("Removing access token")
		}

		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		credentialDirPath := fmt.Sprintf("%s/.interstellar/credential", homeDir)
		credentialFilePath := fmt.Sprintf("%s/access.token", credentialDirPath)

		if !dirExists(credentialDirPath) {
			err := os.MkdirAll(credentialDirPath, 0755)

			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		}

		promptLogin := true
		accessToken := ""

		if fileExists(credentialFilePath) && !forceRemoveAuthentication {
			content, err := os.ReadFile(credentialFilePath)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Using cached access token: ")
			accessToken = string(content)
			promptLogin = false
		}

		if fileExists(credentialFilePath) && forceRemoveAuthentication {
			err := os.Remove(credentialFilePath)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Access token removed")
			promptLogin = true
		}

		if promptLogin {
			tokenResult, err := handleInteractiveLogin()
			if err != nil {
				fmt.Println("An exception has occurred while using the interactive login", err)
				os.Exit(1)
			}

			accessToken = tokenResult
			writeToFile(tokenResult, credentialFilePath)
		}

		fmt.Println(accessToken)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	loginCmd.Flags().BoolP("force", "", false, "Remove existing access token and force re-authentication")
}

func handleInteractiveLogin() (string, error) {
	client, err := public.New("&client-id", public.WithAuthority("https://login.microsoftonline.com/common"))
	if err != nil {
		fmt.Println("We were unable to create a client")
	}

	result, err := client.AcquireTokenInteractive(context.Background(), []string{"user.read"})
	if err != nil {
		return "", err
	}

	return result.AccessToken, nil
}

func writeToFile(s, path string) {
	f, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(s)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func dirExists(dir string) bool {
	info, err := os.Stat(dir)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}
