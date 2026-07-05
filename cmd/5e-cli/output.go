package main

// This file defines the JSON contract 5e-cli speaks when driven as a
// non-interactive helper (e.g. by sns-companion). A request arrives on stdin
// and a view-model is written to stdout.

// Request is the context handed to a command on stdin. Both fields are
// optional; a command that needs no inputs can be run with no stdin at all.
type Request struct {
	Inputs  map[string]any `json:"inputs"`
	Session map[string]any `json:"session"`
}

// str returns a string input, or "" when absent/not a string.
func (r Request) str(key string) string {
	if s, ok := r.Inputs[key].(string); ok {
		return s
	}
	return ""
}

// num returns a numeric input. JSON numbers decode to float64.
func (r Request) num(key string) (float64, bool) {
	switch v := r.Inputs[key].(type) {
	case float64:
		return v, true
	case int:
		return float64(v), true
	default:
		return 0, false
	}
}

// Item is a single generated line, mapped by sns-companion to `:item/*` keys.
type Item struct {
	Title string   `json:"title,omitempty"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags,omitempty"`
}

// Section groups related items under an optional heading.
type Section struct {
	Heading string `json:"heading,omitempty"`
	Items   []Item `json:"items"`
}

// Action is an optional UI button (label + event vector).
type Action struct {
	Label string   `json:"label"`
	Event []string `json:"event"`
}

// ViewModel is the friendly, un-namespaced result sns-companion consumes.
type ViewModel struct {
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle,omitempty"`
	Sections []Section `json:"sections,omitempty"`
	Actions  []Action  `json:"actions,omitempty"`
}

// CommandFunc is a single loot generator: it reads the request and produces a
// view-model (or an error, which becomes a non-zero exit).
type CommandFunc func(Request) (ViewModel, error)

// sectionOf wraps items in a single, heading-less section.
func sectionOf(items ...Item) []Section {
	return []Section{{Items: items}}
}

// vmText builds a view-model with a title and a single body line.
func vmText(title, body string) ViewModel {
	return ViewModel{Title: title, Sections: sectionOf(Item{Body: body})}
}

// staticVM adapts a fixed title/body pair into a CommandFunc.
func staticVM(title, body string) CommandFunc {
	return func(Request) (ViewModel, error) { return vmText(title, body), nil }
}

// affixItem renders an affix as an item: substituted description plus its point
// value, upgrade note, and affinities as tags.
func affixItem(a Affix) Item {
	return Item{
		Body: processString(a.Description),
		Tags: append([]string{a.PointValue, a.Upgrade}, a.Affinities...),
	}
}
