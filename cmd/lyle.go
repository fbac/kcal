package cmd

import (
	"kcal/pkg/lyle"
	"github.com/spf13/cobra"
)

func init() {
 	RootCmd.AddCommand(lyleCmd)
	lyleCmd.Flags().StringP("sex", "s", "", "Sex of the subject")
        lyleCmd.Flags().Float32("weight", 0, "Weight in kg")
        lyleCmd.Flags().Float32("lean-mass", 0, "Lean mass weight in kg")
        lyleCmd.Flags().StringP("plan", "p", "", "Plan: bulk, maint or cut")
}

var (
)

var lyleCmd = &cobra.Command{
 	Use:   "lyle",
  	Short: "Lyle McDonald formula",
  	Long:  `Calculate kcals and macronutrients using Lyle McDonald formula

Usage: kcal lyle [options]

Basic Options
 --sex: Sex of the subject [ man | woman ]
 --weight: Current weight in kg
 --lean-mass: Current lean mass weight in kg
 --plan: Objetive of this diet [ bulk | cut | maint ]`,

 	Run: func(cmd *cobra.Command, args []string) {
		lPlan,_  := cmd.Flags().GetString("plan")
                lWeight,_ := cmd.Flags().GetFloat32("weight")
                lLeanmass,_ := cmd.Flags().GetFloat32("lean-mass")
                lSex,_ := cmd.Flags().GetString("sex")

		lyle.InitLyle(lWeight, lSex, lPlan, lLeanmass)
  	},
}
