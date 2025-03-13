package textx

import "github.com/K-Phoen/grabana/text"

func Markdown(content string) text.Option {
	return func(text *text.Text) error {
		text.Builder.TextPanel.Options.Mode = "markdown"
		text.Builder.TextPanel.Options.Content = content
		return nil
	}
}
