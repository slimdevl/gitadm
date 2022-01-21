package describe

import (
	"github.com/slimdevl/gitadm/pkg/errors"
	"github.com/slimdevl/gitadm/pkg/scm"
	"github.com/slimdevl/gitadm/pkg/util"
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
