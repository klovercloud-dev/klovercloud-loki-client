package v1
func WSUrl()  interface{} {
	return tail.NewUrlBuilder().Init().Get().Label("app","klovercloud-agent").Build().Fire()
}


