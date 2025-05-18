package executor

type ExecutionConfig struct {
	Filename    string
	BinaryName  string
	CompileCmd  []string
	RunCmd      []string
	UseCompiler bool
	Interpreter string
}

var LanguageConfigs = map[string]ExecutionConfig{
	"c": {
		Filename:    "main.c",
		BinaryName:  "main",
		CompileCmd:  []string{"gcc", "main.c", "-o", "main"},
		RunCmd:      []string{"/main"},
		UseCompiler: true,
	},
	"cpp": {
		Filename:    "main.cpp",
		BinaryName:  "main",
		CompileCmd:  []string{"g++", "main.cpp", "-o", "main"},
		RunCmd:      []string{"/main"},
		UseCompiler: true,
	},
	"go": {
		Filename:    "main.go",
		BinaryName:  "main",
		CompileCmd:  []string{"go", "build", "-o", "main", "main.go"},
		RunCmd:      []string{"/main"},
		UseCompiler: true,
	},
	"java": {
		Filename:    "Main.java",
		BinaryName:  "Main.class",
		CompileCmd:  []string{"javac", "Main.java"},
		RunCmd:      []string{"java", "Main"},
		UseCompiler: true,
	},
	"python": {
		Filename:    "script.py",
		Interpreter: "python3",
		UseCompiler: false,
	},
}
