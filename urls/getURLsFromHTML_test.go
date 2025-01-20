package urls_test

import (
	"slices"
	"testing"

	"github.com/serux/webcrawler/urls"
)

func TestGetURL(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{{
		name:     "absolute and relative URLs",
		inputURL: "https://blog.boot.dev",
		inputBody: `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/one">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
		expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
	},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
					<html>
						<body>
							<a href="http://other.com/path/a">
								<span>Boot.dev</span>
							</a>
							<a href="https://other.com/path/one">
								<span>Boot.dev</span>
							</a>
						</body>
					</html>
					`,
			expected: []string{"https://other.com/path/a", "https://other.com/path/one"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := urls.GetURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if slices.Compare(actual, tc.expected) != 0 {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
