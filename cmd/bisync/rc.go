package bisync

import (
	"context"
	"errors"
	"log"

	"github.com/rclone/rclone/cmd/bisync/bilib"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/rc"
)

func init() {
	rc.Add(rc.Call{
		Path:         "sync/bisync",
		AuthRequired: true,
		Fn:           rcBisync,
		Title:        shortHelp,
		Help:         rcHelp,
	})
}

func GetPrefer(key string. p rc.Params) (Prefer, error) {
	str, err := p.GetString(key)
	if err != nil {
		return PreferNone, err
	}

	switch str {
	case "none":
		return PreferNone, nil
	case "path1":
		return PreferPath1, nil
	case "path2":
		return PreferPath2, nil
	case "newer":
		return PreferNewer, nil
	case "older":
		return PreferOlder, nil
	case "larger":
		return PreferLarger, nil
	case "smaller":
		return PreferSmaller, nil
	default:
		return PreferNone, ErrParamInvalid{
			fmt.Errorf("invalid prefer value %q for key %q", str, key),
		}
	}
}


func rcBisync(ctx context.Context, in rc.Params) (out rc.Params, err error) {
	opt := &Options{}
	octx, ci := fs.AddConfig(ctx)

	if dryRun, err := in.GetBool("dryRun"); err == nil {
		ci.DryRun = dryRun
		opt.DryRun = dryRun
	} else if rc.NotErrParamNotFound(err) {
		return nil, err
	}

	if maxDelete, err := in.GetInt64("maxDelete"); err == nil {
		if maxDelete < 0 || maxDelete > 100 {
			return nil, rc.NewErrParamInvalid(errors.New("maxDelete must be a percentage between 0 and 100"))
		}
		opt.MaxDelete = int(maxDelete)
	} else if rc.NotErrParamNotFound(err) {
		return nil, err
	}

	if opt.Resync, err = in.GetBool("resync"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.ResyncMode, err = GetPrefer("resyncmode", in); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.CheckAccess, err = in.GetBool("checkAccess"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.Force, err = in.GetBool("force"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.CreateEmptySrcDirs, err = in.GetBool("createEmptySrcDirs"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.RemoveEmptyDirs, err = in.GetBool("removeEmptyDirs"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.NoCleanup, err = in.GetBool("noCleanup"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.IgnoreListingChecksum, err = in.GetBool("ignoreListingChecksum"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.Resilient, err = in.GetBool("resilient"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.Recover, err = in.GetBool("recover"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.CheckFilename, err = in.GetString("checkFilename"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.FiltersFile, err = in.GetString("filtersFile"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.Workdir, err = in.GetString("workdir"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.BackupDir1, err = in.GetString("backupdir1"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.BackupDir2, err = in.GetString("backupdir2"); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.ConflictResolve, err = GetPrefer("conflictresolve", in); rc.NotErrParamNotFound(err) {
		return
	}
	if opt.ConflictSuffixFlag, err = in.GetString("conflictsuffix"); rc.NotErrParamNotFound(err) {
		return
	}

	checkSync, err := in.GetString("checkSync")
	if rc.NotErrParamNotFound(err) {
		return nil, err
	}
	if checkSync == "" {
		checkSync = "true"
	}
	if err := opt.CheckSync.Set(checkSync); err != nil {
		return nil, err
	}

	fs1, err := rc.GetFsNamed(octx, in, "path1")
	if err != nil {
		return nil, err
	}

	fs2, err := rc.GetFsNamed(octx, in, "path2")
	if err != nil {
		return nil, err
	}

	output := bilib.CaptureOutput(func() {
		err = Bisync(octx, fs1, fs2, opt)
	})
	_, _ = log.Writer().Write(output)
	return rc.Params{"output": string(output)}, err
}
