package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/augustohp/gh-stars/datetime"
	"github.com/google/go-github/github"
	flag "github.com/spf13/pflag"
	"golang.org/x/oauth2"
)

type option struct {
	token    string ""
	limit    int
	insecure bool
	version  bool
	help     bool
}

const Version = "1.0.1"

var opt = &option{}

func init() {
	flag.BoolVarP(&opt.version, "version", "v", false, "Displays version information")
	flag.BoolVarP(&opt.help, "help", "h", false, "Displays this help message")
	flag.BoolVarP(&opt.insecure, "insecure", "i", false, "Disables HTTPS certificate verification")
	flag.StringVarP(&opt.token, "token", "t", "", "A GitHub Personal API Access Token")
	flag.IntVarP(&opt.limit, "limit", "l", 0, "Limit the amount of starts retrieved")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "Options:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "https://github.com/augustohp/gh-stars")

	os.Exit(0)
}

func main() {
	flag.Parse()

	if opt.help {
		usage()
	}

	if opt.version {
		fmt.Fprintf(os.Stderr, "%s %s\n", os.Args[0], Version)
		os.Exit(0)
	}

	if opt.token == "" && os.Getenv("GHSTARS_GITHUB_TOKEN") == "" {
		fmt.Fprintln(os.Stderr, "Error: Please, provide a GITHUB API token.")
		os.Exit(2)
	}

	if opt.token == "" && os.Getenv("GHSTARS_GITHUB_TOKEN") != "" {
		opt.token = os.Getenv("GHSTARS_GITHUB_TOKEN")
	}

	run()
}

func run() {
	ctx := context.Background()
	gh, err := newGithubClient(ctx, opt.token)
	authenticatedUser, _, err := gh.Users.Get(ctx, "")
	if err != nil {
		log.Fatalf("Error listing authenticathed user from GitHub: %v", err)
	}

	stars, err := listLastStars(ctx, gh, authenticatedUser.GetLogin(), opt.limit)
	if err != nil {
		log.Fatalf("Error listing starred activity of %s: %v", authenticatedUser.GetLogin(), err)
	}

	lastWeek := datetime.BeginningOfWeek()
	for _, star := range stars {
		repository := star.GetRepository()
		owner := repository.GetOwner()
		starredAt := star.GetStarredAt()
		if starredAt.Before(lastWeek) {
			continue
		}
		fmt.Fprintf(os.Stdout, "%s/%s\n", owner.GetLogin(), repository.GetName())
	}
}

func listLastStars(ctx context.Context, gh *github.Client, username string, limit int) ([]*github.StarredRepository, error) {
	listStarredOptions := &github.ActivityListStarredOptions{
		Sort:      "created",
		Direction: "desc",
		ListOptions: github.ListOptions{
			PerPage: limit,
		},
	}
	starredRepositories, _, err := gh.Activity.ListStarred(ctx, username, listStarredOptions)
	if err != nil {
		return nil, err
	}

	return starredRepositories, nil
}

func newHttpClient() *http.Client {
	t := &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: opt.insecure},
	}

	return &http.Client{Transport: t}
}

func newGithubClient(ctx context.Context, token string) (*github.Client, error) {
	ctx = context.WithValue(ctx, oauth2.HTTPClient, newHttpClient())
	t := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	c := oauth2.NewClient(ctx, t)
	client := github.NewClient(c)
	var err error

	return client, err
}
