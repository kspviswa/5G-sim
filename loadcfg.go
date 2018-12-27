package main

import (
	"io/ioutil"
	"log"
	"net"

	"github.com/satori/go.uuid"

	yaml "gopkg.in/yaml.v2"
)

type infracfg struct {
	Ram  uint64
	Vcpu uint32
	Hdd  uint64
}

type networkcfg struct {
	Ip string
}

type Config struct {
	Nfs []Nfs
}

type Nfs struct {
	Name        string
	Desp        string
	Infra       infracfg
	Network     networkcfg
	Cardinality uint32
	IsDataplane bool
	Numsessions uint32
	Category    string
}

func loadFromFile(path string) (error, *map[string]*nf) {
	var config Config
	var nfs map[string]*nf = make(map[string]*nf)
	source, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, item := range config.Nfs {
		//fmt.Println("Counter : " + string(c))
		newitem := new(nf)
		newitem.capInfra.hdd = item.Infra.Hdd
		newitem.capInfra.ram = item.Infra.Ram
		newitem.capInfra.vcpu = item.Infra.Vcpu
		newitem.cardinality = item.Cardinality
		newitem.category = nfType(item.Category)
		newitem.desp = item.Desp
		newitem.isDataPlane = item.IsDataplane
		newitem.name = item.Name
		newitem.numSessions = uint64(item.Numsessions)
		newitem.nwInfra.IPs = append(newitem.nwInfra.IPs, net.ParseIP(item.Network.Ip))
		newitem.id, _ = uuid.NewV4()

		nfs[newitem.id.String()] = newitem
	}

	return nil, &nfs
}
