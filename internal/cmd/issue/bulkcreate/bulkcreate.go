package bulkcreate

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/ankitpokhrel/jira-cli/internal/cmd/issue/create"
	"github.com/ankitpokhrel/jira-cli/internal/cmdcommon"
	"github.com/ankitpokhrel/jira-cli/internal/cmdutil"
	"github.com/ankitpokhrel/jira-cli/internal/query"
	"github.com/spf13/cobra"
)

const (
	helpText = `Create issues in bulk in a given project using files as input.`
	examples = `$ jira issue create-bulk

# Create issue in the configured project
$ jira issue bulk-create --dir ./issues"

# Create issue in another project
$ jira issue bulk-create -pPRJ --dir ./issue'

# Create issue in the configured project with JSON output
$ jira issue bulk-create --raw
`
	flagRaw = "raw"
	flagWeb = "web"
)

type BulkCreateParams struct {
	Pattern string
}

// NewCmdBulkCreate is a create command.
func NewCmdBulkCreate() *cobra.Command {
	cmd := cobra.Command{
		Use:     "bulk-create",
		Short:   "Bulk create issues in a project",
		Long:    helpText,
		Example: examples,
		Run:     bulkcreate,
	}

	cmd.Flags().Bool(flagRaw, false, "Print output in JSON format")

	return &cmd
}

func bulkcreate(cmd *cobra.Command, _ []string) {
	params := parseFlags(cmd.Flags())
	jsonFlag, err := cmd.Flags().GetBool(flagRaw)
	cmdutil.ExitIfError(err)

	web, err := cmd.Flags().GetBool(flagWeb)
	cmdutil.ExitIfError(err)

	if params.Pattern != "" {
		files, err := listFilesFromGlob(params.Pattern)
		if err != nil {
			cmdutil.ExitIfError(err)
		}
		for _, file := range files {
			createParams := parseFrontmatterFile(file)
			create.CreateIssue(createParams, jsonFlag, web)
		}
	}
}

// SetFlags sets flags supported by bulk create command.
func SetFlags(cmd *cobra.Command) {
	// TODO: is it possible to create epic like that?
	prefix := "Issue"
	cmd.Flags().SortFlags = false

	if prefix == "Epic" {
		cmd.Flags().StringP("name", "n", "", "Epic name")
	} else {
		cmd.Flags().StringP("type", "t", "", "Issue type")
		cmd.Flags().StringP("parent", "P", "", `Parent issue key can be used to attach epic to an issue.
And, this field is mandatory when creating a sub-task.`)
	}
	cmd.Flags().String("pattern", "", "Pattern containing markdown files formatted with frontmatter")
	cmd.Flags().Bool("web", false, "Open in web browser after successful creation")
}

func listFilesFromGlob(pattern string) ([]string, error) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.New("no files matched")
	}
	return files, nil
}

func parseFrontmatterFile(fileName string) *cmdcommon.CreateParams {
	params := &cmdcommon.CreateParams{}
	data, err := cmdutil.ReadFile(fileName)
	if err != nil {
		cmdutil.ExitIfError(err)
	}

	markdown, err := frontmatter.MustParse(strings.NewReader(string(data)), params)
	if err != nil {
		err = errors.New("could not parse frontmatter input")
		cmdutil.ExitIfError(err)
	}

	params.Body = string(markdown)
	params.NoInput = true
	params.Frontmatter = ""
	params.Template = ""
	return params
}

func parseFlags(flags query.FlagParser) *BulkCreateParams {
	pattern, err := flags.GetString("pattern")
	cmdutil.ExitIfError(err)

	return &BulkCreateParams{
		Pattern: pattern,
	}
}
