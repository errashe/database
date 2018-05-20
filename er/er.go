package er

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	. "fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

type Server struct {
	ID        int
	IP        string
	Priority  int
	Timestamp time.Time
}

func (s Server) String() string {
	return Sprintf("%-4d|%-15s|%d|%s", s.ID, s.IP, s.Priority, s.Timestamp.Format("15:04"))
}

type Servers []Server

func (s Servers) String() string {
	var ret []string
	for _, server := range s {
		ret = append(ret, server.String())
	}
	return strings.Join(ret, "\n")
}

func (s Servers) SortByPriority() {
	sort.Slice(s, func(i, j int) bool {
		if s[i].Priority > s[j].Priority {
			return true
		}
		if s[i].Priority < s[j].Priority {
			return false
		}
		return s[i].Timestamp.Sub(s[j].Timestamp) < 0
	})
}

func (s Servers) Save(filename string) error {
	fileJson, err := json.Marshal(s)
	if err != nil {
		return err
	}

	var fileGZ bytes.Buffer
	zipper := gzip.NewWriter(&fileGZ)

	_, err = zipper.Write(fileJson)
	if err != nil {
		return err
	}
	zipper.Close()

	err = ioutil.WriteFile("output.json.gz", fileGZ.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *Servers) Load(filename string) error {
	rzip, err := ioutil.ReadFile("output.json.gz")
	if err != nil {
		return err
	}

	r, err := gzip.NewReader(bytes.NewBuffer(rzip))
	if err != nil {
		return err
	}

	res, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(res, s)
	if err != nil {
		return err
	}

	return nil
}
