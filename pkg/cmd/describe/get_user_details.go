package describe

import (
	"github.com/slimdevl/gitlab/pkg/errors"
	"github.com/slimdevl/gitlab/pkg/scm"
	"github.com/slimdevl/gitlab/pkg/util"
	"github.com/urfave/cli/v2"
)

func getUserDetails(c *cli.Context) error {
	username := c.String("username")
	details, err := scm.GetUserDetails(username)
	if err != nil {
		return errors.Wrap(&err, "username %s is invalid", username)
	}
	util.PrettyPrint(details)
	return nil
}
