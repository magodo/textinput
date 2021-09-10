# Textinput with auto-completion

This is a simple wrapper around the `github.com/charmbracelet/bubbles/textinput`, with support for auto-completion.

The API is almost the same as the `github.com/charmbracelet/bubbles/textinput`, except users can optionally set the `.CandidateWords` to indicate the candidate words for auto-completion, which is triggered by `tab`/`shift-tab`.

See the example for a demo, and also compare it with the [bubbletea textinput example](https://github.com/charmbracelet/bubbletea/blob/master/examples/textinput/main.go) to see the usage difference.
