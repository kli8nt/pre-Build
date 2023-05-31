package utils

import (
	"fmt"
	"strings"

	"github.com/Nie-Mand/dfile/pkg"
)

func Build(data *Data) error {
	dockerfile := pkg.Dockerfile{}
	dockerfile.Init()
	dockerfile.SetFilename(data.ApplicationName + "/Dockerfile")
	dockerfile.
		From(data.Technology).
		ImageVersion(data.Version).
		ImageAlias("build").
		WorkDir("/app")

	for _, dep := range data.DependenciesFiles {
		dockerfile.Copy(dep, ".")
	}

	dockerfile.
		Run(data.InstallCommand).
		Copy(".", ".")

	for _, env := range data.EnvironmentVariables {
		parts := strings.Split(env, "=")
		dockerfile.Envs(parts[0], parts[1])
	}
	if !data.IsStatic {
		dockerfile.Expose(data.Port).
			Run(data.RunCommand)
	} else {
		dockerfile.Run(data.BuildCommand)
		dockerfile.NextStage()
		dockerfile.From("nginx").
			ImageVersion(data.Version).
			Copy("--from=build /app/"+data.OutputDirectory, "/usr/share/nginx/html").
			Cmd("nginx", "-g", "daemon off;").
			Expose(80)
	}

	dockerfile.Save()
	err := dockerfile.Save()
	if err != nil {
		return err
	}
	fmt.Println("Dockerfile generated successfully!")
	return nil
}
