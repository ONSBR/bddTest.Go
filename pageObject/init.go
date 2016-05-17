package pageObject

var (
	serverUrl string
)

func init() {
	serverUrl = InitTestServer()
}
