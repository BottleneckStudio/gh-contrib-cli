package main 
import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/BottleneckStudio/gh-contrib-cli/entity"
)
	
	var (
		author
		stringsince 
		stringuntil 
		stringrepo 
		stringapiURL = "https://api.github.com/repos/a-fis"
)

// ContributionOptions ...
type ContributionOptions struct {
	Author string 
	Since  string
	Until  string
	Repo   string
}

func init() {
	flag.StringVar(&author, "author", "", "Author(user) of commit to look for.")
	flag.StringVar(&since, "since", "", "Retrieve commits after this date. Must be in YYYY-MM-DD format.")
	flag.StringVar(&until, "until", "", "Retrieve commits before this date. Must be in YYYY-MM-DD format.")
	flag.StringVar(&repo, "repo", "wikix.net", "Repository to check for commits.")
	flag.Usage = usage
}
func main() {
	flag.Parse()
	if contribOptions, err := parseOptions(); err == nil {
		request(contribOptions)
	} else {
		usage()
	}
	// log.Printf("Author: %s", author)// log.Printf("Since: %s", since)// log.Printf("Until: %s", until)// log.Printf("Repository: %s", repo)
}

func usage() {
	fmt.Fprintf(os.Stderr, `gh-contrib version: gh-contrib/1.0.0
>> HomePage: https://github.com/rbo13/gh-contrib
>> Issue   : https://github.com/rbo13/gh-contrib/issues
>> Author  : rbo13Usage: gh-contrib-cli -u <username> -sinceOptions:`
)
flag.PrintDefaults()
}

func parseOptions() (*ContributionOptions, error) {
	if author == "" {
		return nil, errors.New("Author is required")
	}
	// config  the options
	contributionOptions := &ContributionOptions{
		Author: author,
		Since:  since,
		Until:  until,
		Repo:   repo,
	}
	return contributionOptions, nil
}

func request(contribOptions *ContributionOptions) error {
	var contributions []entity.Contributions
	client := http.Client{
	Timeout: time.Duration(10 * time.Second),
	}
	url := fmt.Sprintf("%s/%s/commits?author=%s&since=%s&until=%s&page=1&access_token=11beeff32f34e6798c4cc986d173c8f7189587ad", apiURL, contribOptions.Repo, contribOptions.Author, contribOptions.Since, contribOptions.Until)
	log.Print(url)
	resp, err := http.Get(url)
	log.Print(err)
	
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	buffer, err := ioutil.ReadAll(resp.Body)
	
	log.Print(err)
	if err != nil {
		return err
	}
	json.Unmarshal(buffer, &contributions)
	log.Print(len(contributions))
	for i := 0; i < len(contributions); i++ {
		log.Print(contributions)
	}
	return nil
}