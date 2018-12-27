# 5G Simulator - Mind Map

This is a mind dump of my experimentation.

## Preamble

Every simulation will be backed up by a experiment scenario. An Experiment Scenario is nothing but a topology of 5G usecase.

Typically a 5G usecase will contain 3 main parts - Access Network, Core Network & Data Network. Specically RAN, 5GCore & Internet.

Currently this simulator works for above specific topology, but it intends to be more abstract in future.

## An Experiment Flow

The experiments begins by letting the user expressing their ***intent*** - what is that you want to do in this experiment.

The user goes bys saying - *Hey bring up a 5G network with default settings or bring up a 5G network with specific settings.*

Once the network is up, configure different elements of network with specific values. This includes the capabilities, number of allowed slices & policies if any.
While specifying these configuration, user also specifies the ***cardinality*** of those NFs ( Network Functions )

If the cardinality is 1, then that means those NFs will be shared. If the cardinality is N, then that means, no more than N instances be instantiated.

When the experiment begins, 5G Network is up and default config is applied. Users can then attach UE to gNB and begin requesting network slices based on ***NST*** ( Network Slice Template ) . The Network then mimic the registration & slice allocation flow based on 3GPP specs. 

At any point of time, user can attach a UE, request / modify / terminate slice, remove UE, change / scale base policies of 5GC and tweak the slices. Users can also view the ***In-use / Available*** view of entire network and per NFs. 

A special command-base birds-eye-view of 5GC network and Slice based network topology will be made available for now. Future releases will have a web-based topology preview.

## What are the basic requirements be met ?

### NF Model

* A model for NF capabilities from infrastructure point of view.
* Above model should be flexible to accomdate NF specific capabilities.

### Slice Model

* A model to specify NST ( Network Slice Template)
* A model to host converted S-NSSAIs and NSSAIs ( Single Network Slice Selection Assistance Information )

### Resource Monitor

* A Global Monitor to maintain Use / Available view per slice, per NF and per Global network.
* A persistent store to make sure things are under control.

### Command Runner

* Uniform & consistent way to request system.

## Command list - mind dump

## `cloud` - All about underlying cloud

```
cloud add --file <path/to/file>
```

| Subcommand | flags | Description | Output | Metadata
|---|---|---|---|---|
add | --file | Register a cloud from file | Registered ID <br> Eg `7645-awf3-2232-1ee3` | Following is the file format `cloud1.yaml` |

`cloud1.yaml`

```
---
name: TampaDC
region: region1
infra:
    vcpu: 16
    ram: 17179869184 // 16GB
    hdd: 10737418240 // 10GB
network:
    IPs:
        - 10.10.2.1
        - fe:23:45:43
    HPAs:
        - dpdk
        - sriov
```

---

```
cloud delete <id>
```

---

```
cloud list
```
