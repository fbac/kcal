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

package kcal

import (
	"fmt"
	"os"
	"errors"
	"kcal/pkg"
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

Lyle specific options
 --sex: Sex of the subject [ man | woman ]
 --weight: Current weight in kg
 --lean-mass: Current lean mass weight in kg
 --plan: Objetive of this diet [ bulk | cut | maint ]

Harris-Benefict specific options
 --sex: Sex of the subject [ man | woman ]
 --weight: Current weight in kg
 --lean-mass: Current lean mass weight in kg
 --height: Current height in cm
 --age: Age in years
 --activity: Activity index. 1.1 for sedentary, 1.3 lesser exercise, 1.5 moderate exercise, 1.7 high activity
 --plan: Objetive of this diet [ bulk | cut | maint ]
 --deviation: Deficit or superavit in % to apply in kcal calculations.
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
			lyle.InitLyle(pWeight, pSex, pPlan, pLeanmass)
		} else if pFormula == "harris-benedict" {
			fmt.Println("init hb", pHeight, pDeviation, pActivity, pAge)
		} else {
			return errors.New("Formula not recognized. Available formula: lyle, harris-benedict")
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
		return errors.New("A formula must be defined. Options: lyle | harris-benedict")
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