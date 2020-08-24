package cliHelper

type IWriter interface{
	isError bool
	write()
}

type IWritersCollection interface{
	Writers []Writer
}

type IWritingConfiguration interface{
	isJsonLog	bool
	isWriteToFile bool
	writeToFileDirectory string
	writeToFilePath string
	isWriteToDirectory bool
}

type  IDefaultJsonInteration interface{
	Attributes	struct {
		Counter	int	`json:"counter"`
		Log	struct {
			Debug		[]map[string]string
			Display		[]string
			Error		[]map[string]string
			Info		[]map[string]string
			IsDebug		bool
			IsError		bool
			IsInfo		bool
			IsWarning	bool
			Warning		[]map[string]string
		}
		ParentTask	string
		ParentTaskID	string
		Progress	int
		Timestamp	string
	}
	ID	string
	Invoker	string
	Title	string
}
