package main

import (
	"testing"
	"strings"
	"fmt"
	"encoding/json"
)


func TestParse(t *testing.T) {
	sln := `
	Microsoft Visual Studio Solution File, Format Version 12.00
	# Visual Studio Version 16
	VisualStudioVersion = 16.0.29326.143
	MinimumVisualStudioVersion = 10.0.40219.1
	Project("{9A19103F-16F7-4668-BE54-9A1E7A4F7556}") = "My.ApiName", "src\Api.My.ApiName.csproj", "{7A27BC15-7604-4C63-BB39-14FAE9933861}"
	EndProject
	Project("{9A19103F-16F7-4668-BE54-9A1E7A4F7556}") = "My.ApiNameTest", "test\Api.My.ApiNameTest.csproj", "{F884322D-22DC-4FFE-B91E-BAD5EE1D4D45}"
	EndProject
	Global
			GlobalSection(SolutionConfigurationPlatforms) = preSolution
					Debug|Any CPU = Debug|Any CPU
					Release|Any CPU = Release|Any CPU
			EndGlobalSection
			GlobalSection(ProjectConfigurationPlatforms) = postSolution
					{7A27BC15-7604-4C63-BB39-14FAE9933861}.Debug|Any CPU.ActiveCfg = Debug|Any CPU
					{7A27BC15-7604-4C63-BB39-14FAE9933861}.Debug|Any CPU.Build.0 = Debug|Any CPU
					{7A27BC15-7604-4C63-BB39-14FAE9933861}.Release|Any CPU.ActiveCfg = Release|Any CPU
					{7A27BC15-7604-4C63-BB39-14FAE9933861}.Release|Any CPU.Build.0 = Release|Any CPU
					{F884322D-22DC-4FFE-B91E-BAD5EE1D4D45}.Debug|Any CPU.ActiveCfg = Debug|Any CPU
					{F884322D-22DC-4FFE-B91E-BAD5EE1D4D45}.Debug|Any CPU.Build.0 = Debug|Any CPU
					{F884322D-22DC-4FFE-B91E-BAD5EE1D4D45}.Release|Any CPU.ActiveCfg = Release|Any CPU
					{F884322D-22DC-4FFE-B91E-BAD5EE1D4D45}.Release|Any CPU.Build.0 = Release|Any CPU
			EndGlobalSection
			GlobalSection(SolutionProperties) = preSolution
					HideSolutionNode = FALSE
			EndGlobalSection
			GlobalSection(ExtensibilityGlobals) = postSolution
					SolutionGuid = {24B5D20A-F42B-4FA1-9057-DE2A992041A7}
			EndGlobalSection
	EndGlobal`
	_ = sln
	stmt, err := NewParser(strings.NewReader(sln)).Parse()
	if(err != nil) {
		fmt.Printf("Error: %s\n",err)
		t.Fail()
	}
	j,err := json.MarshalIndent(stmt,"","  ")
	if(err != nil) {
		fmt.Printf("Error: %s\n",err)
		t.Fail()
	}
	fmt.Println(string(j))
}