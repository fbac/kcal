// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type person struct {
	sex string
	age int
	height float32
	weight float32
	leanmass float32
        activity float32
	plan string
        deviation float32
}

var RootCmd = &cobra.Command{
	Use:   "kcal",
	Short: "Calculate kcal and macronutrients distribution",
	Long: `Calculate kcal income and macronutrients distribution.
Usage: kcal [options]

Where options include:

Basic Options
< --formula lyle|harris-benedict >: Formula used to calculate kcal.
[ --weight <weight> ]: Current weight in kg.
[ --sex man|woman ]: Sex of the subject.
[ --plan bulk|cut|maintenance ]: Diet's objective.
[ --deviation <deficit|superavit> ]: Deficit or superavit in % to apply in kcal calculations.
`,
	PreRunE: func(c *cobra.Command, args []string) error {
		return checkRequiredFlags(c.Flags())
	},

	RunE: func(c *cobra.Command, args []string) error {
		pFormula, err := c.Flags().GetString("formula")
		if err != nil { return err }
		pWeight, err := c.Flags().GetFloat32("weight")
		if err != nil { return err }
		pLeanmass, err := c.Flags().GetFloat32("lean-mass")
                if err != nil { return err }
		pSex, err := c.Flags().GetString("sex")
		if err != nil { return err }
		pPlan, err := c.Flags().GetString("plan")
		if err != nil { return err }
		pDeviation, err := c.Flags().GetFloat32("deviation")
		if err != nil { return err }
		pHeight, err := c.Flags().GetFloat32("height")
                if err != nil { return err }
                pActivity, err := c.Flags().GetFloat32("activity")
                if err != nil { return err }
		pAge, err := c.Flags().GetInt("age")
		if err != nil { return err }

		if pFormula == "lyle" {
			lyleDiet, err := execLyle(pWeight, pSex, pPlan)
			if err != nil { 
				return err 
			} else if pLeanmass == 0 || pLeanmass >= pWeight {
				return errors.New("Lean mass not recognized or higher than weight")
			} else {
				fat, prot, ch := calculateMacro(lyleDiet, pLeanmass)
                	        fmt.Printf("Total kcals: %.1f kcals\n", lyleDiet)
                        	fmt.Printf("Fat intake: %.1f gr\n", fat)
                        	fmt.Printf("Prot intake: %.1f gr\n", prot)
                        	fmt.Printf("Carbs intake: %.1f gr\n", ch)
			}

		} else if pFormula == "harris-benedict" {
			fmt.Println(pDeviation, pHeight, pActivity, pAge)
		} else {
			return errors.New("Formula not recognized")
		}

		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkRequiredFlags(flags *pflag.FlagSet) error {
        if len(flags.Lookup("formula").Value.String()) == 0 {
		return errors.New("A formula is needed")
        }
        return nil
}

func init() {
	// Required flags
	RootCmd.Flags().StringP("formula", "f", "", "Formula to be used")

	// Optional flags
	RootCmd.Flags().StringP("sex", "s", "", "Sex of the subject")
	RootCmd.Flags().Int("age", 0, "Age of the subject")
	RootCmd.Flags().Float32("height", 0, "Height in cm")
	RootCmd.Flags().Float32("weight", 0, "Weight in kg")
	RootCmd.Flags().Float32("lean-mass", 0, "Lean mass weight in kg")
	RootCmd.Flags().Float32("activity", 0, "Activity factor")
	RootCmd.Flags().StringP("plan", "p", "", "Plan: bulk, maint or cut")
	RootCmd.Flags().Float32("deviation",  0, "Deficit or superavit to be applied")
}

// Diet
func initLyle(w float32, s string, p string) (*person, error) {

	lyle := person{}
	
	if w > 0 {
		lyle.weight = w
	} else {
		fmt.Printf("Weight not recognized, insert weight in kg: ")
	}

	if s == "man" || s == "woman" {
		lyle.sex = s
	} else {
		fmt.Printf("Sex not recognized, insert sex [man, woman]: ")
	}

	if p == "bulk" || p == "maint" || p == "cut" {
		lyle.plan = p
	} else {
		fmt.Printf("Plan not recognized, insert plan [bulk, cut, maint]: ")
	}

	return &lyle, nil
}

func execLyle(w float32, s string, p string) (float32, error) {
	dataLyle, err := initLyle(w, s, p)
	if err != nil { return 0, err }

	var kcal float32
	if dataLyle.sex == "man" {
		switch dataLyle.plan {
		  case "bulk":
			  kcal = dataLyle.weight * 40
		  case "cut":
			  kcal = dataLyle.weight * 24
		  case "maint":
			  kcal = dataLyle.weight * 35
		  default:
		 }
	} else if dataLyle.sex == "woman" {
		switch dataLyle.plan {    
                  case "bulk":
			  kcal = dataLyle.weight * 35
                  case "cut":
			  kcal = dataLyle.weight * 22
                  case "maint":
			  kcal = dataLyle.weight * 31
                  default:
		  }
	}

	return kcal, nil
}

func calculateMacro(kcal float32, leanmass float32) (float32, float32, float32){
	var fat, prot, ch float32

	fat = (kcal * 0.25)/9
	prot = leanmass * 2.2
	ch = (kcal - (fat * 9) - (prot * 4))/4

	return fat, prot, ch
}

func isFloat(val float32) bool {
    return val == float32(int(val))
}
