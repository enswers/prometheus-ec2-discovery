package main

import (
	"reflect"
	"testing"

	"github.com/goamz/goamz/ec2"
)

var instances []ec2.Instance = []ec2.Instance{
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "next-production-0.6.16-work",
			},
			ec2.Tag{
				Key:   "Deployment",
				Value: "next-production",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "work",
			},
		},
		PrivateIpAddress: "172.22.2.183",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 48,
			Name: "terminated",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.16-frontend",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "frontend",
			},
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
		},
		PrivateIpAddress: "",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Type",
				Value: "frontend",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "next-production-0.6.16-frontend",
			},
			ec2.Tag{
				Key:   "Deployment",
				Value: "next-production",
			},
		},
		PrivateIpAddress: "172.22.2.57",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Deployment",
				Value: "next-production",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "work",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "next-production-0.6.16-work",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
		},
		PrivateIpAddress: "172.22.1.121",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Type",
				Value: "frontend",
			},
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.16-frontend",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
		},
		PrivateIpAddress: "172.22.1.149",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Type",
				Value: "frontend",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.18",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.18-frontend",
			},
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
		},
		PrivateIpAddress: "172.22.1.154",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "frontend",
			},
			ec2.Tag{
				Key:   "Deployment",
				Value: "next-production",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "next-production-0.6.16-frontend",
			},
		},
		PrivateIpAddress: "172.22.1.89",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Type",
				Value: "frontend",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.16-frontend",
			},
		},
		PrivateIpAddress: "172.22.2.151",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.16-work",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "work",
			},
		},
		PrivateIpAddress: "172.22.2.245",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.16-work",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "work",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
		},
		PrivateIpAddress: "172.22.2.246",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.16-work",
			},
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "work",
			},
		},
		PrivateIpAddress: "172.22.2.247",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "work",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.16-work",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
		},
		PrivateIpAddress: "172.22.2.248",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.16-work",
			},
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "work",
			},
		},
		PrivateIpAddress: "172.22.2.249",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.16-work",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "work",
			},
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
		},
		PrivateIpAddress: "172.22.1.64",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.16",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "work",
			},
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.16-work",
			},
		},
		PrivateIpAddress: "172.22.1.62",
	},
	ec2.Instance{
		State: ec2.InstanceState{
			Code: 16,
			Name: "running",
		},
		Tags: []ec2.Tag{
			ec2.Tag{
				Key:   "Name",
				Value: "www-production-0.6.18-frontend",
			},
			ec2.Tag{
				Key:   "Type",
				Value: "frontend",
			},
			ec2.Tag{
				Key:   "Version",
				Value: "0.6.18",
			},
			ec2.Tag{
				Key:   "Deployment",
				Value: "www-production",
			},
		},
		PrivateIpAddress: "172.22.2.150",
	},
}

func TestGroupByAndRender(t *testing.T) {
	want := `[
  {
    "targets": [
      "172.22.2.57:0",
      "172.22.1.89:0"
    ],
    "labels": {
      "Deployment": "next-production",
      "Type": "frontend",
      "Version": "0.6.16"
    }
  },
  {
    "targets": [
      "172.22.1.149:0",
      "172.22.2.151:0"
    ],
    "labels": {
      "Deployment": "www-production",
      "Type": "frontend",
      "Version": "0.6.16"
    }
  },
  {
    "targets": [
      "172.22.1.154:0",
      "172.22.2.150:0"
    ],
    "labels": {
      "Deployment": "www-production",
      "Type": "frontend",
      "Version": "0.6.18"
    }
  },
  {
    "targets": [
      "172.22.2.183:0",
      "172.22.1.121:0"
    ],
    "labels": {
      "Deployment": "next-production",
      "Type": "work",
      "Version": "0.6.16"
    }
  },
  {
    "targets": [
      "172.22.2.245:0",
      "172.22.2.246:0",
      "172.22.2.247:0",
      "172.22.2.248:0",
      "172.22.2.249:0",
      "172.22.1.64:0",
      "172.22.1.62:0"
    ],
    "labels": {
      "Deployment": "www-production",
      "Type": "work",
      "Version": "0.6.16"
    }
  }
]`
	targetGroups := groupByTags(instances, []string{"Type", "Deployment", "Version"})
	b := marshalTargetGroups(targetGroups)
	got := string(b)
	if want != got {
		t.Fatalf("Did not get the expected output\nwant: %#v\ngot: %#v", want, got)
	}
}

func TestAllTagKeys(t *testing.T) {
	want := []string{"Deployment", "Name", "Type", "Version"}
	got := allTagKeys(instances)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("%#v != %#v", want, got)
	}
}

func TestParseTags(t *testing.T) {
	tagKey := func(name string) Tag {
		return Tag{
			Key:         name,
			FilterName:  "tag-key",
			FilterValue: name,
		}
	}
	tagKeyVal := func(name, val string) Tag {
		return Tag{
			Key:         name,
			FilterName:  "tag:" + name,
			FilterValue: val,
		}
	}
	cases := []struct {
		tagsRaw string
		want    Tags
	}{
		{"", Tags{}},
		{"Name", Tags{tagKey("Name")}},
		{"Name,App", Tags{tagKey("Name"), tagKey("App")}},
		{"Name,App=foo", Tags{tagKey("Name"), tagKeyVal("App", "foo")}},
		{"Name,App=foo,App=bar", Tags{tagKey("Name"), tagKeyVal("App", "foo"), tagKeyVal("App", "bar")}},
	}
	for _, c := range cases {
		got := parseTags(c.tagsRaw)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("parseTags(%#v) = %#v != %#v", c.tagsRaw, c.want, got)
		}
	}
}
