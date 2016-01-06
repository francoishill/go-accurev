package accurev

func Cat_WithElementIDAndVersion(elementID, elementVersion string) string {
	if elementVersion == "0/0" {
		return ""
	}

	return mustRunAccurevCommand_AsString("cat", "-e", elementID, "-v", elementVersion)
}
