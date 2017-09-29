package main

func init() {
	locpool.Resources["en-xa"] = map[string]string{
		// H1 Heading
		"WelcomeMessage": "[Welcome! xxxxxxx]",

		// {NAME} is the name of the user
		"HelloUser": "[Hello, {NAME}! xxxxxxxx]",

		// {X} is the number of files, {Y} is the number of folders, {COMMAND} in the ID of the command
		"XFilesFoundInYFolders": "[{X_PLURAL:{X} file|{X} files} found in {Y_PLURAL:{Y} folder|{Y} folders}. Do you want to {COMMAND:copy|move|delete} {X:them|it|them}? xxxxxxxxxx xxxxxxxxxx xxxxxxxxxx xxxxxxxxxx xxxxxxxxxx xxxxxxxxxx]",
	}
}
