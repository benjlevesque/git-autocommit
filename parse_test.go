package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name       string
		branchName string
		want       string
	}{
		{
			name:       "main",
			branchName: "main",
			want:       "main",
		},
		{
			name:       "without domain",
			branchName: "feat/message",
			want:       "feat: message",
		},
		{
			name:       "with domain",
			branchName: "feat/domain/message",
			want:       "feat(domain): message",
		},
		{
			name:       "spaces in message",
			branchName: "feat/domain/message-with-spaces",
			want:       "feat(domain): message with spaces",
		},
		{
			name:       "dashes in message",
			branchName: "feat/domain/message-with--dashes",
			want:       "feat(domain): message with-dashes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parse(tt.branchName))
		})
	}
}

func TestReplaceDashes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "single word",
			input: "message",
			want:  "message",
		},
		{
			name:  "multiple words",
			input: "some-text",
			want:  "some text",
		},
		{
			name:  "dashes",
			input: "some--text",
			want:  "some-text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, replaceDashes(tt.input))
		})
	}
}
