package commands

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/spf13/cobra"
)

var webPort *int

func init() {
	webPort = webCmd.PersistentFlags().Int("port", 80, " --port 80")
	rootCmd.AddCommand(webCmd)
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "以web端运行",
	RunE: func(cmd *cobra.Command, args []string) error {
		return http.ListenAndServe(fmt.Sprintf(":%d", *webPort), &HTTPServer{})
	},
}

type HTTPServer struct {
}

func (server *HTTPServer) ServeHTTP(write http.ResponseWriter, req *http.Request) {
	tpl := template.New("main")
	tpl.Parse(web_tpl)
	tpl.Execute(write, map[string]interface{}{
		"Title": "easytime",
	})
}

var web_tpl = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
	<h1>11</h1>
	<ul>
	</ul>
	</body>
</html>
`
