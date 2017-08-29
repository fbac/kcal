package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
 	RootCmd.AddCommand(hbCmd)
}

var hbCmd = &cobra.Command{
 	Use:   "hb",
  	Short: "Harris-Benedict formula",
  	Long:  `Calculate kcals and macronutrients using Harris-Benedict formula

Usage: kcal hb [options]

Basic Options
 --sex: Sex of the subject [ man | woman ]
 --weight: Current weight in kg
 --lean-mass: Current lean mass weight in kg
 --height: Current height in cm
 --age: Age in years
 --activity: Activity index. 1.1 for sedentary, 1.3 lesser exercise, 1.5 moderate exercise, 1.7 high activity
 --plan: Objetive of this diet [ bulk | cut | maint ]
 --deviation: Deficit or superavit in % to apply in kcal calculations.`,

 	Run: func(cmd *cobra.Command, args []string) {
    		fmt.Println("Harris-Benedict initialized.")
  	},
}
