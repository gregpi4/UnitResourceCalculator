package resinput

import (
	"fmt"
	"log"
)

type Product struct {
	Name   string
	Recipe map[string]int
}

var (
	VanillaProducts = map[string]Product{
		"bell":       {"bell", map[string]int{"ironPlate": 1, "cog": 1}},
		"inserter":   {"inserter", map[string]int{"iron": 1, "copper": 1}},
		"Automation": {"automate", map[string]int{"iron": 1, "copper": 1}},
		"cog":        {"cog", map[string]int{"ironPlate": 1}},
	}
)
var (
	TestModProducts = map[string]Product{
		"testmod_product_1": {"testmod_product_1", map[string]int{"atom": 100}},
		"testmod_product_2": {"testmod_product_2", map[string]int{"atom": 200}},
	}
)

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, in_set := set[item]
	return in_set
}

func productNameFromProducts(products map[string]Product) []string {
	var productNames []string
	for _, product := range products {
		productNames = append(productNames, product.Name)
	}
	return productNames
}

func GetProducts(mod string) map[string]Product {
	switch mod {
	case "vanilla":
		return VanillaProducts
	case "testmod":
		return TestModProducts
	default:
		fmt.Println("Invalid module")
		panic("Unknown module")
	}
}

func GetProductsName(mod string) []string {
	switch mod {
	case "vanilla":
		return productNameFromProducts(VanillaProducts)
	case "testmod":
		return productNameFromProducts(TestModProducts)
	default:
		fmt.Println("Invalid module")
		panic("Unknown module")
	}
}

func reprFilteredComponent(component string, filtered_component string) string {
	return component + "_" + filtered_component
}

func availableInputIntoAutomateInputs(availableInput []string) []string {
	var currentIndex int = 0
	var automateInputs []string
	for currentIndex < 6 {
		if currentIndex < len(availableInput) {
			automateInputs = append(automateInputs, availableInput[currentIndex])
		} else {
			automateInputs = append(automateInputs, "")
		}
		currentIndex++
	}
	return automateInputs
}

func undergroundBellWithAutomateInput(automateInput1 string, automateInput2 string) string {
	if automateInput1 == "" && automateInput2 == "" {
		return ""
	} else if automateInput1 == "" {
		return "undergroundBell_" + automateInput2
	} else if automateInput2 == "" {
		return "undergroundBell_" + automateInput1
	} else {
		return "undergroundBell_" + automateInput1 + "_" + automateInput2
	}
}

func inserterWithAutomateInput(automateInput1 string, automateInput2 string) string {
	if automateInput1 == "" && automateInput2 == "" {
		return ""
	} else {
		return "inserter"
	}
}

type DesignOutput struct {
	length               int
	height               int
	productOutputX       int
	productOutputY       int
	productOutputName    string
	availableInputNeeded []string
	design               [][]string
}

var wrongDesignSize = 0
var wrongDesign = DesignOutput{0, 0, 0, 0, "undefined", []string{}, [][]string{}}

func getBellWithDesignOutput(designOutput1 DesignOutput, designOutput2 DesignOutput) string {

	if designOutput1.length == wrongDesignSize && designOutput2.length == wrongDesignSize {
		return ""
	} else if designOutput1.length == wrongDesignSize {
		return "undergroundBell_" + designOutput2.design[0][designOutput2.productOutputX]
	} else if designOutput2.length == wrongDesignSize {
		return "undergroundBell_" + designOutput1.design[0][designOutput1.productOutputX]
	} else {
		return "undergroundBell_" + designOutput1.design[0][designOutput1.productOutputX] + "_" + designOutput2.design[0][designOutput2.productOutputX]
	}
}

func getInserterWithDesignOutput(designOutput1 DesignOutput, designOutput2 DesignOutput) string {
	if designOutput1.length == wrongDesignSize && designOutput2.length == wrongDesignSize {
		return ""
	} else {
		return "inserter"
	}
}

func glueProductWithDesignOutputs(productDesign DesignOutput, x int, y int, designOutputs map[int]DesignOutput) string {
	// this method returns for a given x y the component in the x,y tile
	// to do this, we follow this pattern :
	// if x and y are not in the zone to glue, return ""
	// else for designOutput that are odd, the bell go up on the productDesignLength + 3 to the even designOutput coordiante
	// for even designOutput, the bell go up on productDesignLength + designOutputIndex//2
	// On the y coordinate of the even design outputs, it's a bell between productDesignLength + designOutputIndex//2 and productDesignLength + 4

	// example for bell
	// {"", "", "undergroundBell_ironPlate", "", "", "", "", "", "", "", "", "undergroundBell_ironPlate", "", "", "", ""},
	// 	{"undergroundBell_bell", "inserter_bell", "automate_bell", "automate_bell", "automate_bell", "inserter", "bell_cog", "bell_cog", "bell_cog", "bell_cog", "inserter", "automate_cog", "automate_cog", "automate_cog", "", ""},
	// 	{"", "", "automate_bell", "automate_bell", "automate_bell", "", "", "", "", "", "", "automate_cog", "automate_cog", "automate_cog", "", ""},
	// 	{"", "", "automate_bell", "automate_bell", "automate_bell", "", "", "", "", "", "", "automate_cog", "automate_cog", "automate_cog", "", ""},
	// 	{"", "", "inserter", "", "", "", "", "","",  "", "", "inserter", "", "", "", ""},
	// 	{"", "", "undergroundBell_ironPlate", "", "","", "", "", "", "", "", "undergroundBell_ironPlate", "", "", "", ""},
	// }

	if productDesign.length+3 < x || x < productDesign.length {
		return ""
	}
	if x == productDesign.length {
		if y == 1 {
			return getInserterWithDesignOutput(designOutputs[0], designOutputs[1])
		} else if y == 2 {
			return getInserterWithDesignOutput(designOutputs[2], designOutputs[3])
		} else if y == 3 {
			return getInserterWithDesignOutput(designOutputs[4], designOutputs[5])
		} else {
			return ""
		}
	} else if x == productDesign.length+1 {
		if y == 1 {
			return getBellWithDesignOutput(designOutputs[0], designOutputs[1])
		} else if y == 2 {
			return getBellWithDesignOutput(designOutputs[2], designOutputs[3])
		} else if y >= 3 && y <= designOutputs[4].productOutputY {
			return getBellWithDesignOutput(designOutputs[4], designOutputs[5])
		}
	} else if x == productDesign.length+2 {
		if y == 1 {
			return getBellWithDesignOutput(designOutputs[0], designOutputs[1])
		} else if y >= 2 && y <= designOutputs[2].productOutputY {
			return getInserterWithDesignOutput(designOutputs[2], designOutputs[3])
		} else if y == designOutputs[4].productOutputY {
			return getInserterWithDesignOutput(designOutputs[4], designOutputs[5])
		}
	} else if x == productDesign.length+3 {
		if y == 1 {
			return getBellWithDesignOutput(designOutputs[0], designOutputs[1])
		} else if y >= 2 && y <= designOutputs[1].productOutputY {
			return getBellWithDesignOutput(wrongDesign, designOutputs[1])
		} else if y == designOutputs[2].productOutputY {
			return getBellWithDesignOutput(designOutputs[2], designOutputs[3])
		} else if y >= designOutputs[2].productOutputY+1 && y <= designOutputs[3].productOutputY {
			return getBellWithDesignOutput(wrongDesign, designOutputs[3])
		} else if y == designOutputs[4].productOutputY {
			return getBellWithDesignOutput(designOutputs[4], designOutputs[5])
		} else if y >= designOutputs[4].productOutputY+1 && y <= designOutputs[5].productOutputY {
			return getBellWithDesignOutput(wrongDesign, designOutputs[5])
		}
	}
	return ""

}

func glueDesignOutputsTogether(x int, y int, designOutputs map[int]DesignOutput) DesignOutput {
	//precondition : x > product output + 3

	if y < designOutputs[0].productOutputY { 

	}

func mergeDesigns(productDesign DesignOutput, designOutputs map[int]DesignOutput) {
	max_intermediate_length := 0
	max_intermediate_height := 0
	for _, designOutput := range designOutputs {
		if designOutput.productOutputX > max_intermediate_length {
			max_intermediate_length = designOutput.productOutputX
		}
		if designOutput.productOutputY > max_intermediate_height {
			max_intermediate_height = designOutput.productOutputY
		}
	}
	for x := 0; x < productDesign.length+max_intermediate_length+3; x++ {
		for y := 0; y < max_intermediate_height+3; y++ {
			if x < productDesign.length && y < productDesign.height {
				productDesign.design[x][y] = productDesign.design[x][y]
			} else if x < productDesign.length {
				productDesign.design[x][y] = ""
			} else if x > productDesign.length +3 {
				productDesign.design[x][y] = ""
			}
			} else {
				productDesign.design[x][y] = glueProductWithDesignOutputs(productDesign, x, y, designOutputs)
			}
		}
	}
}

func recursiveDesign(mod string, to_produce string, availableInput []string) (DesignOutput, error) {
	// this function should returns a map of each tile with the name of the product on a tile
	// example :
	// bell uses automation for production. Automation produce 1 bell from 1 iron and 1 copper.

	// automation needs 3*3 tiles.
	// bell needs 1*1 tiles
	// inserter is needed to transfer input and output in automation
	// 1 inserter needs 1*1 tiles
	// bell links tiles between two inserters
	// example :
	// To produce 1 bell, the map could be
	// want := [][]string{
	// 	{"", "", "automate_cog", "automate_cog", "automate_cog", "", ""},
	// 	{"bell_cog", "inserter_cog", "automate_cog", "automate_cog", "automate_cog", "inserter_ironPlate", "bell_ironPlate"},
	// 	{"", "", "automate_cog", "automate_cog", "automate_cog", "", ""},
	// }
	// To do that, we will first find is the input needed is available
	var availableInputs []string
	var toManufactureInputs []string
	var intermediateDesigns []DesignOutput
	product := GetProducts(mod)[to_produce]
	for inputName := range product.Recipe {
		if contains(availableInput, inputName) {
			availableInputs = append(availableInputs, inputName)
		} else {
			toManufactureInputs = append(toManufactureInputs, inputName)
		}
	}
	log.Printf("%s", availableInputs)
	if len(availableInputs) > 6 {
		return wrongDesign, fmt.Errorf("too many inputs")
	}
	if len(toManufactureInputs) > 6 {
		return wrongDesign, fmt.Errorf("too many inputs, need to manufacture %s", toManufactureInputs)
	}
	automate_inputs := availableInputIntoAutomateInputs(availableInputs)
	automate_product := reprFilteredComponent("automate", to_produce)
	bell_product := reprFilteredComponent("undergroundBell", to_produce)
	inserter_product := reprFilteredComponent("inserter", to_produce)
	for index := 0; index < 6; index++ {
		if index < len(toManufactureInputs) {
			intermediate := toManufactureInputs[index]
			intermediateDesign, err := recursiveDesign(mod, intermediate, availableInput)
			if err != nil {
				return wrongDesign, err
			}
			intermediateDesigns = append(intermediateDesigns, intermediateDesign)
		} else {
			intermediateDesigns = append(intermediateDesigns, wrongDesign)
		}
	}
	main_design := [][]string{
		{"", "", undergroundBellWithAutomateInput(automate_inputs[0], automate_inputs[1]), undergroundBellWithAutomateInput(automate_inputs[2], automate_inputs[3]), undergroundBellWithAutomateInput(automate_inputs[4], automate_inputs[5])},
		{bell_product, inserter_product, automate_product, automate_product, automate_product},
		{"", "", automate_product, automate_product, automate_product},
		{"", "", automate_product, automate_product, automate_product},
		{"", "", inserterWithAutomateInput(automate_inputs[0], automate_inputs[1]), inserterWithAutomateInput(automate_inputs[2], automate_inputs[3]), inserterWithAutomateInput(automate_inputs[4], automate_inputs[5])},
		{"", "", undergroundBellWithAutomateInput(automate_inputs[0], automate_inputs[1]), undergroundBellWithAutomateInput(automate_inputs[2], automate_inputs[3]), undergroundBellWithAutomateInput(automate_inputs[4], automate_inputs[5])},
	}

	return DesignOutput{length: 1, height: 1, productOutputX: 1, productOutputY: 1, productOutputName: to_produce, availableInputNeeded: []string{}, design: main_design}, nil

}

func Design(mod string, to_produce string, available_input []string) ([][]string, error) {
	// this function should returns a map of each tile with the name of the product on a tile
	// example :
	// bell uses automation for production. Automation produce 1 bell from 1 iron and 1 copper.

	// automation needs 3*3 tiles.
	// bell needs 1*1 tiles
	// inserter is needed to transfer input and output in automation
	// 1 inserter needs 1*1 tiles
	// bell links tiles between two inserters
	// example :
	// To produce 1 bell, the map could be
	//[[nil, nil, inserter, nil, nil],
	//[nil, automation, automation, automation, nil]
	//[nil, automation, automation, automation, nil]
	//[nil, automation, automation, automation, nil]
	//[nil, nil, inserter, nil, nil]]

	// To do that, we will first find is the input needed is available
	design, err := recursiveDesign(mod, to_produce, available_input)
	return design.design, err
}
