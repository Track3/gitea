	"io/ioutil"
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
		return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content))
	diffRecord := diffMatchPatch.DiffMain(highlight.Code(diffSection.FileName, diff1[1:]), highlight.Code(diffSection.FileName, diff2[1:]), true)
func ParsePatch(maxLines, maxLineCharacters, maxFiles int, reader io.Reader) (*Diff, error) {
			return diff, fmt.Errorf("Invalid first file line: %s", line)
		// TODO: Handle skipping first n files
			_, err := io.Copy(ioutil.Discard, reader)
				return diff, fmt.Errorf("Copy: %v", err)
						return diff, fmt.Errorf("Unable to ReadLine: %v", err)
				err = fmt.Errorf("Unable to ReadLine: %v", err)
			err = fmt.Errorf("Unable to ReadLine: %v", err)
					err = fmt.Errorf("Unable to ReadLine: %v", err)
			curSection = &DiffSection{}
				err = fmt.Errorf("Unexpected line in hunk: %s", string(lineBytes))
				curSection = &DiffSection{}
				curSection = &DiffSection{}
				curSection = &DiffSection{}
			err = fmt.Errorf("Unexpected line in hunk: %s", string(lineBytes))
					err = fmt.Errorf("Unable to ReadLine: %v", err)
				count, err := models.Count(m)
func GetDiffRangeWithWhitespaceBehavior(gitRepo *git.Repository, beforeCommitID, afterCommitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	var cmd *exec.Cmd
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
		cmd = exec.CommandContext(ctx, git.GitExecutable, diffArgs...)
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
		cmd = exec.CommandContext(ctx, git.GitExecutable, diffArgs...)
		return nil, fmt.Errorf("StdoutPipe: %v", err)
		return nil, fmt.Errorf("Start: %v", err)
	diff, err := ParsePatch(maxLines, maxLineCharacters, maxFiles, stdout)
		return nil, fmt.Errorf("ParsePatch: %v", err)
		return nil, fmt.Errorf("Wait: %v", err)
	shortstatArgs := []string{beforeCommitID + "..." + afterCommitID}
func GetDiffCommitWithWhitespaceBehavior(gitRepo *git.Repository, commitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(gitRepo, "", commitID, maxLines, maxLineCharacters, maxFiles, whitespaceBehavior)
		setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(c.Patch))