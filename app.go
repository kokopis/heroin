package main

import (
  "encoding/json"
  "log"
  "os/exec"
  "regexp"
  "time"
  "unsafe"
)

type app struct {
  Acm                          bool         `json:"acm"`
  ArchivedAt                   string       `json:"archived_at"`
  BuildpackProvidedDescription string       `json:"buildpack_provided_description"`
  BuildStack                   buildStack   `json:"build_stack"`
  CreatedAt                    string       `json:"created_at"`
  Id                           string       `json:"id"`
  GitUrl                       string       `json:"git_url"`
  Maintenance                  bool         `json:"maintenance"`
  Name                         string       `json:"name"`
  Owner                        owner        `json:"owner"`
  Region                       region       `json:"region"`
  Organization                 organization `json:"organization"`
  Team                         team         `json:"team"`
  Space                        interface{}  `json:"space"`
  InternalRouting              interface{}  `json:"internal_routing"`
  ReleasedAt                   string       `json:"released_at"`
  RepoSize                     int          `json:"repo_size"`
  SlugSize                     int          `json:"slug_size"`
  Stack                        stack        `json:"stack"`
  UpdatedAt                    time.Time    `json:"updated_at"`
  WebUrl                       string       `json:"web_url"`
  Joined                       bool         `json:"joined"`
  LegacyId                     string       `json:"legacy_id"`
  Locked                       bool         `json:"locked"`
}

type apps []app

func getAppsByJson(str string) apps {
  var a apps
  if err := json.Unmarshal([]byte(str), &a); err != nil { log.Panic(err) }
  return a
}

func (a *app) getRedisVersionString() string {
  comm := exec.Command("heroku", "redis:info", "-a", a.Name)

  out, err := comm.Output()
  if err != nil { log.Panic(err) }

  regex := regexp.MustCompile(`(?m)^Version(?:.*)`)
  return regex.FindString(*(*string)(unsafe.Pointer(&out)))
}
