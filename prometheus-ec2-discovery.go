package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
)

var (
	dest   string
	tags   []string
	region aws.Region
	port   int
)

// TargetGroup is a collection of related hosts that prometheus monitors
type TargetGroup struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

func main() {
	initFlags()

	filter := ec2.NewFilter()
	for _, t := range tags {
		filter.Add("tag-key", t)
	}

	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	e := ec2.New(auth, region)
	resp, err := e.Instances(nil, filter)
	if err != nil {
		log.Fatal(err)
	}
	instances := flattenReservations(resp.Reservations)

	if len(tags) == 0 {
		tags = allTagKeys(instances)
	}

	targetGroups := groupByTags(instances, tags)
	b := marshalTargetGroups(targetGroups)
	if dest == "-" {
		_, err = os.Stdout.Write(b)
	} else {
		err = atomicWriteFile(dest, b, ".new")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func initFlags() {
	var (
		tagsRaw   string
		regionRaw string
	)

	flag.StringVar(&dest, "dest", "-", "File to write the target group JSON. (e.g. `tgroups/target_groups.json`)")
	flag.StringVar(&tagsRaw, "tags", "Name", "Comma seperated list of tags to group by (e.g. `Environment,Application`)")
	flag.StringVar(&regionRaw, "region", "us-west-2", "AWS region to query")
	flag.IntVar(&port, "port", 80, "Port that is exposing /metrics")

	flag.Parse()
	tags = strings.Split(tagsRaw, ",")
	region = aws.Regions[regionRaw]

	if tags[0] == "" && len(tags) == 1 {
		tags = []string{}
	}
}

func groupByTags(instances []ec2.Instance, tags []string) map[string]*TargetGroup {
	targetGroups := make(map[string]*TargetGroup)

	for _, instance := range instances {
		if instance.State.Code != 16 { // 16 = Running
			continue
		}

		key := ""
		for _, tagKey := range tags {
			key = fmt.Sprintf("%s|%s=%s", key, tagKey, getTag(instance, tagKey))
		}

		targetGroup, ok := targetGroups[key]
		if !ok {
			labels := make(map[string]string)
			for _, tagKey := range tags {
				tagVal := getTag(instance, tagKey)
				if tagVal != "" {
					labels[tagKey] = tagVal
				}
			}
			targetGroup = &TargetGroup{
				Labels:  labels,
				Targets: make([]string, 0),
			}
			targetGroups[key] = targetGroup
		}

		target := fmt.Sprintf("%s:%d", instance.PrivateIpAddress, port)
		targetGroup.Targets = append(targetGroup.Targets, target)
	}

	return targetGroups
}

func marshalTargetGroups(targetGroups map[string]*TargetGroup) []byte {
	// We need to transform targetGroups into a values list sorted by key
	tgList := []*TargetGroup{}
	keys := []string{}
	for k, _ := range targetGroups {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		tgList = append(tgList, targetGroups[k])
	}

	b, err := json.MarshalIndent(tgList, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func atomicWriteFile(filename string, data []byte, tmpSuffix string) error {
	err := ioutil.WriteFile(filename+tmpSuffix, data, 0644)
	if err != nil {
		return err
	}
	err = os.Rename(filename+tmpSuffix, filename)
	if err != nil {
		return err
	}
	return nil
}

func getTag(instance ec2.Instance, key string) string {
	for _, t := range instance.Tags {
		if t.Key == key {
			return t.Value
		}
	}
	return ""
}

func flattenReservations(reservations []ec2.Reservation) []ec2.Instance {
	instances := make([]ec2.Instance, 0)
	for _, r := range reservations {
		instances = append(instances, r.Instances...)
	}
	return instances
}

func allTagKeys(instances []ec2.Instance) []string {
	tagSet := map[string]struct{}{}
	for _, instance := range instances {
		for _, t := range instance.Tags {
			tagSet[t.Key] = struct{}{}
		}
	}
	tags := []string{}
	for tag, _ := range tagSet {
		tags = append(tags, tag)
	}
	sort.Strings(tags)
	return tags
}