package file

import (
	"context"
	"io/fs"
	"naslook/pkg/hash"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
)

const (
	_BothKeep = "both keep"
)

func FileDeleteDup(ctx context.Context, paths []string, ignores []string) error {
	hashMap := map[string]string{}

	for _, path := range paths {
		filepath.Walk(path, func(path string, info fs.FileInfo, _ error) error {
			if info.IsDir() {
				return nil
			}

			for _, i := range ignores {
				if strings.Contains(path, i) {
					return nil
				}
			}

			md5, err := hash.GetFileMD5(path)
			if err != nil {
				logrus.Warnf("get file md5 error: %v", err)
				return nil
			}

			logrus.Infof("md5: %s, file: %s", md5, path)

			old, ok := hashMap[md5]
			if !ok {
				hashMap[md5] = path
				return nil
			}

			var action string
			if err := survey.AskOne(
				&survey.Select{
					Message: "find same files，please chosen delete：",
					Options: []string{old, path, _BothKeep},
				},
				&action,
			); err != nil {
				return err
			}

			if action == old {
				hashMap[md5] = path
			}

			if action != _BothKeep {
				if err := os.Remove(action); err != nil {
					logrus.Warnf("delte file error: %v", err)
				}
			}

			return nil
		})
	}

	return nil
}
