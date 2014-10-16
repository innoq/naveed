package naveed

import "fmt"
import "os"
import "strings"
import "bufio"

// reads settings files where each line contains a key/value pair
func ReadSettings(filePath string,
	delimiter string) (settings map[string]string, err error) {
	settings = map[string]string{}

	fh, err := os.Open(filePath)
	defer fh.Close()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.SplitN(line, delimiter, 2)
		key := strings.TrimSpace(items[0])
		value := strings.TrimSpace(items[1])
		settings[key] = value
	}

	return settings, nil
}

func ReportError(err error, ctx string) {
	if err != nil {
		fmt.Printf("ERROR %s: %s\n", ctx, err)
	}
}
