package typrails

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-rest-server/pkg/typrails/internal/tmpl"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/utility/runn"
	"github.com/typical-go/typical-go/pkg/utility/runner"
	"github.com/urfave/cli/v2"
)

// Module of typrails
type Module struct{}

// BuildCommands is commands to exectuce from Build-Tool
func (m *Module) BuildCommands(c typcore.Cli) []*cli.Command {
	return []*cli.Command{
		{
			Name:   "rails",
			Usage:  "Rails-like generator",
			Action: m.scaffold,
		},
	}
}

func (m *Module) scaffold(ctx *cli.Context) (err error) {
	e := Entity{
		Name:           "music",
		Table:          "musics",
		Type:           "Music",
		Cache:          "MUSIC",
		ProjectPackage: "github.com/typical-go/typical-rest-server",
		Fields: []Field{
			{Name: "ID", Column: "id", Type: "int64", StructTag: "`json:\"id\"`"},
			{Name: "Artist", Column: "artist", Type: "string", StructTag: "`json:\"artist\"`"},
			{Name: "UpdatedAt", Column: "updated_at", Type: "time.Time", StructTag: "`json:\"updated_at\"`"},
			{Name: "CreatedAt", Column: "created_at", Type: "time.Time", StructTag: "`json:\"created_at\"`"},
		},
		Forms: []Field{
			{Name: "Artist", Column: "artist", Type: "string"},
		},
	}
	return generate(e)
}

func generate(e Entity) error {
	repoPath := fmt.Sprintf("app/repository/%s_repo.go", e.Name)
	repoImplPath := fmt.Sprintf("app/repository/%s_repo_impl.go", e.Name)
	cachedRepoImplPath := fmt.Sprintf("app/repository/cached_%s_repo_impl.go", e.Name)
	servicePath := fmt.Sprintf("app/service/%s_service.go", e.Name)
	controllerPath := fmt.Sprintf("app/controller/%s_cntrl.go", e.Name)
	os.Remove(repoPath)
	os.Remove(repoImplPath)
	os.Remove(cachedRepoImplPath)
	os.Remove(servicePath)
	os.Remove(controllerPath)
	return runn.Execute(
		runner.NewWriteTemplate(repoPath, tmpl.Repo, e),
		runner.NewWriteTemplate(repoImplPath, tmpl.RepoImpl, e),
		runner.NewWriteTemplate(cachedRepoImplPath, tmpl.CachedRepoImpl, e),
		runner.NewWriteTemplate(servicePath, tmpl.Service, e),
		runner.NewWriteTemplate(controllerPath, tmpl.Controller, e),
		runner.NewGoFmt(repoPath, repoImplPath),
	)
}