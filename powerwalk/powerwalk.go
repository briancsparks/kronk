package powerwalk

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "io/ioutil"
  "os"
  "path/filepath"
)

type PowerDir struct {
  fulldirpath   string                /* The full path of the dir */
  dirs        []string                /* short names */
  files       []string                /* short names */
  root          string                /* The root dir from whence I came */

  dmap          map[string]bool       /* Set of full dirnames in this dir */
  fmap          map[string]bool       /* Set of full filenames in this dir */
}

func NewPowerDir(dirpath, root string) *PowerDir {
  var power PowerDir

  power.fulldirpath = dirpath
  power.root = root

  entries, _ := ioutil.ReadDir(power.fulldirpath)

  for _, entry := range entries {
    if entry.IsDir() {
      power.dirs = append(power.dirs, entry.Name())
    } else {
      power.files = append(power.files, entry.Name())
    }
  }

  return &power
}

func buildFileNameSet(fulldirpath string, entries []string) map[string]bool {
  var m map[string]bool

  for _, entry := range entries {
    m[entry] = true
    m[filepath.Join(fulldirpath, entry)] = true
  }

  return m
}

func (p *PowerDir) hasDir(d string) bool {
  if p.dmap == nil {
    p.dmap = buildFileNameSet(p.fulldirpath, p.dirs)
  }

  _, ok := p.dmap[d]
  return ok
}

func (p *PowerDir) hasFile(f string) bool {
  if p.fmap == nil {
    p.fmap = buildFileNameSet(p.fulldirpath, p.dirs)
  }

  _, ok := p.fmap[f]
  return ok
}

func PossibleContentDirs(roots []string, subss... []string) ([]string, error) {
  var allPossibles []string

  for _, subs := range subss {
    allPossibles = subDirs(allPossibles, subs)
  }

  var possibles []string
  for _, possible := range possibles {
    _, err := os.Stat(possible)

    if err == nil || os.IsExist(err) {
      possibles = append(possibles, possible)
    }
  }

  return possibles, nil
}

func subDirs(dirs, subs []string) []string {
  var result []string

  for _, dir := range dirs {
    result = append(result, dir)
    for _, sub := range subs {
      result = append(result, filepath.Join(dir, sub))
    }
  }

  return result
}
