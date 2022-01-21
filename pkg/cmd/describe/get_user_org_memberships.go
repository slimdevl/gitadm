package describe

import (
	"github.com/slimdevl/go-gitlab-client/pkg/scm"
	"github.com/slimdevl/go-gitlab-client/pkg/util"
	"github.com/urfave/cli/v2"
)

func getUserOrgMemberships(c *cli.Context) error {
	details, err := scm.GetUserOrganizationMemberships()
	if err != nil {
		return err
	}
	util.PrettyPrint(details)
	return nil
}
