package convert

import (
	"encoding/json"
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

const input1 = `
locals {
	test3 = 1 + 2
	test1 = "hello"
	test2 = 5
	arr = [1, 2, 3, 4]
	hyphen-test = 3
	temp = "${1 + 2} %{if local.test2 < 3}\"4\n\"%{endif}"
	temp2 = "${"hi"} there"
		quoted = "\"quoted\""
		squoted = "'quoted'"
	x = -10
	y = -x
	z = -(1 + 4)
}

locals {
	other = {
		num = local.test2 + 5
		thing = [for x in local.arr: x * 2]
		"${local.test3}" = 4
		3 = 1
		"local.test1" = 89
		"a.b.c[\"hi\"][3].*" = 3
		loop = "This has a for loop: %{for x in local.arr}x,%{endfor}"
		a.b.c = "True"
	}
}

locals {
	heredoc = <<-EOF
		This is a heredoc template.
		It references ${local.other.3}
	EOF
	simple = "${4 - 2}"
	cond = test3 > 2 ? 1: 0
	heredoc2 = <<EOF
		Another heredoc, that
		doesn't remove indentation
		${local.other.3}
		%{if true ? false : true}"gotcha"\n%{else}4%{endif}
	EOF
}

data "terraform_remote_state" "remote" {
	backend = "s3"

	config = {
		profile = var.profile
		region  = var.region
		bucket  = "mybucket"
		key     = "mykey"
	}
}

variable "profile" {}

variable "region" {
	default = "us-east-1"
}
`

const expectedJSON1 = `{
	"data": {
		"terraform_remote_state": {
			"remote": {
				"backend": "s3",
				"config": {
					"bucket": "mybucket",
					"key": "mykey",
					"profile": "${var.profile}",
					"region": "${var.region}"
				}
			}
		}
	},
	"locals": [
		{
			"arr": [
				1,
				2,
				3,
				4
			],
			"hyphen-test": 3,
			"quoted": "\"quoted\"",
			"squoted": "'quoted'",
			"temp": "${1 + 2} %{if local.test2 \u003c 3}\"4\n\"%{endif}",
			"temp2": "hi there",
			"test1": "hello",
			"test2": 5,
			"test3": "${1 + 2}",
			"x": -10,
			"y": "${-x}",
			"z": "${-(1 + 4)}"
		},
		{
			"other": {
				"${local.test3}": 4,
				"3": 1,
				"a.b.c": "True",
				"a.b.c[\"hi\"][3].*": 3,
				"local.test1": 89,
				"loop": "This has a for loop: %{for x in local.arr}x,%{endfor}",
				"num": "${local.test2 + 5}",
				"thing": "${[for x in local.arr: x * 2]}"
			}
		},
		{
			"cond": "${test3 \u003e 2 ? 1: 0}",
			"heredoc": "This is a heredoc template.\nIt references ${local.other.3}\n",
			"heredoc2": "\t\tAnother heredoc, that\n\t\tdoesn't remove indentation\n\t\t${local.other.3}\n\t\t%{if true ? false : true}\"gotcha\"\\n%{else}4%{endif}\n",
			"simple": "${4 - 2}"
		}
	],
	"variable": {
		"profile": {},
		"region": {
			"default": "us-east-1"
		}
	}
}`

func compareTest(t *testing.T, inputStr string, expected string, options Options) {
	bytes := []byte(inputStr)
	conf, diags := hclsyntax.ParseConfig(bytes, "test", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		t.Errorf("Failed to parse config: %v", diags)
	}
	converted, err := convertFile(conf, options)
	if err != nil {
		t.Errorf("Unable to convert from hcl: %v", err)
	}

	jb, err := json.MarshalIndent(converted, "", "\t")
	if err != nil {
		t.Errorf("Failed to serialize to json: %v", err)
	}
	computedJSON := string(jb)
	if computedJSON != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, computedJSON)
	}
}

// Test that conversion works as expected
func TestConversion(t *testing.T) {
	compareTest(t, input1, expectedJSON1, Options{})
}

func TestSimplify(t *testing.T) {
	input := `locals {
		a = split("-", "xyx-abc-def")
		x = 1 + 2
		y = pow(2,3)
		t = "x=${4+abs(2-3)*parseint("02",16)}"
		j = jsonencode({
			a = "a"
			b = 5
		})
		with_vars = x + 1
	}`

	expected := `{
	"locals": {
		"a": [
			"xyx",
			"abc",
			"def"
		],
		"j": "{\"a\":\"a\",\"b\":5}",
		"t": "x=6",
		"with_vars": "${x + 1}",
		"x": 3,
		"y": 8
	}
}`

	compareTest(t, input, expected, Options{Simplify: true})
}
