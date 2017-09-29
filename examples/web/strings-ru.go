package main

func init() {
	locpool.Resources["ru"] = map[string]string{
		// H1 Heading
		"WelcomeMessage": "Добро пожаловать!",

		// {NAME} is the name of the user
		"HelloUser": "Здравствуйте, {NAME}!",

		// {X} is the number of files, {Y} is the number of folders, {COMMAND} in the ID of the command
		"XFilesFoundInYFolders": "В {Y_PLURAL:{Y} папке|{Y} папках|{Y} папках} {X_PLURAL:найден {X} файл|найдены {X} файла|найдено {X} файлов}. Хотите {X:его|их} {COMMAND:скопировать|переместить|удалить}?",
	}
}
