package backloggd

import (
	"os"

	"github.com/BearTS/backloggd-automation/pkg/config"
	"github.com/BearTS/backloggd-automation/pkg/utils"
	"github.com/BearTS/backloggd-go/sdk"
	"github.com/pkg/errors"
)

var Client *sdk.BackloggdSDK

func InitClient(restart bool) (*sdk.BackloggdSDK, error) {
	cookiesSave := utils.GetConfigDirectory() + "/cookies.json"
	username, err := config.GetCredentials("backloggd", "username")
	if err != nil {
		return Client, errors.Wrapf(err, "failed to get username. Kindly set it")
	}

	password, err := config.GetCredentials("backloggd", "password")
	if err != nil {
		return Client, errors.Wrapf(err, "failed to get password. Kindly set it")
	}

	if Client == nil {
		Client, err = sdk.NewBackloggdSDK(username, password, cookiesSave)
		if err != nil {
			return Client, errors.Wrap(err, "failed to initialize Backloggd client")
		}
	}

	if restart {
		err = os.Remove(cookiesSave)
		if err != nil {
			return Client, errors.Wrap(err, "failed to delete cookies")
		}

		Client, err = sdk.NewBackloggdSDK(username, password, cookiesSave)
		if err != nil {
			return Client, errors.Wrap(err, "failed to initialize Backloggd client")
		}
	}

	return Client, nil
}
