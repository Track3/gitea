	"io/ioutil"
			_, err := io.Copy(ioutil.Discard, reader)
				count, err := models.Count(m)
func GetDiffRangeWithWhitespaceBehavior(gitRepo *git.Repository, beforeCommitID, afterCommitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	shortstatArgs := []string{beforeCommitID + "..." + afterCommitID}
func GetDiffCommitWithWhitespaceBehavior(gitRepo *git.Repository, commitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(gitRepo, "", commitID, maxLines, maxLineCharacters, maxFiles, whitespaceBehavior)