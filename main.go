// Package main is the main entry point for the groqmit command.
package main

import (
	_ "embed"
	"fmt"

	"github.com/conneroisu/groqmit/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// func run(ctx context.Context) error {
// 	// Get user's home directory
//
// 	prompt, err := fillTemplate(diff)
// 	if err != nil {
// 		return err
// 	}
//
// 	client, err := groq.NewClient(viper.GetString("groq_key"))
// 	if err != nil {
// 		return err
// 	}
//
// 	resp, err := client.ChatCompletion(ctx, groq.ChatCompletionRequest{
// 		Model: groq.ModelLlama3370BVersatile,
// 		Messages: []groq.ChatCompletionMessage{
// 			{
// 				Role:    "system",
// 				Content: prompt,
// 			},
// 		},
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	if len(resp.Choices) == 0 {
// 		return errors.New("no choices from chat completion")
// 	}
// 	commitMsg := resp.Choices[0].Message.Content
//
// 	tmpFile, err := os.CreateTemp("", "commitmsg")
// 	if err != nil {
// 		return err
// 	}
// 	defer os.Remove(tmpFile.Name())
//
// 	cmd := exec.Command("git", "commit", "-F", tmpFile.Name())
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	return cmd.Run()
// }
//
// var (
// 	//go:embed prompt.tmpl
// 	promptStr      string
// 	promptTemplate = template.Must(template.New("prompt").Parse(promptStr))
// )
//
// func fillTemplate(diff string) (string, error) {
// 	var (
// 		w   bytes.Buffer
// 		err error
// 	)
// 	err = promptTemplate.Execute(&w, struct {
// 		Diff string
// 	}{
// 		Diff: diff,
// 	})
// 	return w.String(), err
// }
