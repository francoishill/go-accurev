package accurev

func Cat_WithElementIDAndVersion(depotName, elementID, elementVersion string) string {
	if elementVersion == "0/0" {
		return ""
	}

	return mustRunAccurevCommand_AsString("cat", "-e", elementID, "-p", depotName, "-v", elementVersion)
}
