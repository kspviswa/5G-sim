package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func loadNFfromFile(path string) error {
	return nil
}

func nodeListAll() error {
	table := tablewriter.NewWriter(os.Stdout)
	fmt.Println("List of Catalouged Network Functions")
	table.SetHeader([]string{"Name", "UUID", "Description", "Type", "Cardinality"})

	for _, item := range nodes {
		data := []string{
			item.name,
			item.id.String(),
			item.desp,
			string(item.category),
			fmt.Sprint(item.cardinality),
		}
		table.Append(data)
	}
	table.Render()
	return nil
}

func tableHelper(headers []string, rows []*[]string) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetAlignment(2) //Right

	for _, items := range rows {
		var data []string
		for _, item := range *items {
			data = append(data, item)
		}
		table.Append(data)
	}
	table.Render()
	return nil
}

func nodeShowDetails(arg string) error {
	fmt.Println("\nDetails of " + arg)
	item, ok := nodes[arg]

	if ok {
		fmt.Println("\nGeneric Details:")
		{
			//table := tablewriter.NewWriter(os.Stdout)
			//table.SetHeader([]string{"Name", "UUID", "Description", "Type", "Cardinality"})
			header := []string{"Name", "UUID", "Description", "Type", "Cardinality"}
			var rows []*[]string
			data := []string{
				item.name,
				item.id.String(),
				item.desp,
				string(item.category),
				fmt.Sprint(item.cardinality),
			}
			// table.Append(data)
			// table.Render()
			rows = append(rows, &data)
			tableHelper(header, rows)
		}

		fmt.Println("\nInfrastructure Requirement Details:")
		{
			//table := tablewriter.NewWriter(os.Stdout)
			//table.SetHeader([]string{"Compute", "Memory (Bytes)", "Storage (Bytes)"})
			header := []string{"Compute", "Memory (Bytes)", "Storage (Bytes)"}
			var rows []*[]string
			data := []string{
				fmt.Sprint(item.capInfra.vcpu),
				fmt.Sprint(item.capInfra.ram),
				fmt.Sprint(item.capInfra.hdd),
			}
			// table.Append(data)
			// table.Render()
			rows = append(rows, &data)
			tableHelper(header, rows)
		}

		fmt.Println("\nOther Details:")
		{
			// table := tablewriter.NewWriter(os.Stdout)
			// table.SetHeader([]string{"Compute", "Memory (Bytes)", "Storage (Bytes)"})
			header := []string{"Attribute", "Value"}
			var rows []*[]string
			rows = append(rows, &[]string{"Provisoned IP", item.nwInfra.IPs[0].String()})
			rows = append(rows, &[]string{"Category", string(item.category)})
			rows = append(rows, &[]string{"Total sessions limit", strconv.FormatUint(item.numSessions, 10)})
			rows = append(rows, &[]string{"Dataplane Element", strconv.FormatBool(item.isDataPlane)})
			tableHelper(header, rows)
		}
	}
	return nil
}
