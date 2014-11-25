package records

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestHostBySlaveId(t *testing.T) {

	slaves := []slave{
		{Id: "20140827-000744-3041283216-5050-2116-1", Hostname: "blah.com"},
		{Id: "33333333-333333-3333333333-3333-3333-2", Hostname: "blah.blah.com"},
	}

	rg := RecordGenerator{Slaves: slaves}

	for i := 0; i < len(slaves); i++ {
		host, err := rg.hostBySlaveId(slaves[i].Id)
		if err != nil {
			t.Error(err)
		}

		if host != slaves[i].Hostname {
			t.Error("wrong slave/hostname")
		}
	}

}

func TestYankPort(t *testing.T) {
	p := "[31328-31328]"

	port := yankPort(p)

	if port != "31328" {
		t.Error("not parsing port")
	}
}

func TestLeaderIP(t *testing.T) {
	l := "master@144.76.157.37:5050"

	ip := leaderIP(l)

	if ip != "144.76.157.37" {
		t.Error("not parsing ip")
	}
}

func TestStripUID(t *testing.T) {
	tname := "reviewbot.8c9b3434-615a-11e4-a088-c20493233aa5"

	name := stripUID(tname)

	if name != "reviewbot" {
		t.Error("not parsing task name")
	}
}

// ensure we are parsing what we think we are
func TestInsertState(t *testing.T) {
	var sj StateJSON

	b, err := ioutil.ReadFile("../factories/fake.json")
	if err != nil {
		t.Error("missing test data")
	}

	err = json.Unmarshal(b, &sj)
	if err != nil {
		t.Error(err)
	}

	rg := RecordGenerator{}
	rg.InsertState(sj, "mesos")

	// ensure we are only collecting running tasks
	_, ok := rg.SRVs["poseidon._tcp.marathon-0.6.0.mesos."]
	if ok {
		t.Error("should not find this not-running task - SRV record")
	}

	_, ok = rg.As["liquor-store.mesos."]
	if !ok {
		t.Error("should find this running task - A record")
	}

	_, ok = rg.As["poseidon.mesos."]
	if ok {
		t.Error("should not find this not-running task - A record")
	}

	// test for 6 SRV names
	if len(rg.SRVs) != 6 {
		t.Error("not enough SRVs")
	}

	// test for 6 A names
	if len(rg.As) != 6 {
		t.Error("not enough As")
	}

	// ensure we find this SRV
	rrs := rg.SRVs["liquor-store._tcp.marathon-0.6.0.mesos."]

	// ensure there are 2 RRDATA answers for this SRV name
	if len(rrs) != 2 {
		t.Error("not enough SRV records")
	}

	// ensure we don't find this as a SRV record
	rrs = rg.SRVs["liquor-store.marathon-0.6.0.mesos."]
	if len(rrs) != 0 {
		t.Error("not a proper SRV record")
	}

}