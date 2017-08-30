package cmd

import (
	"kcal/pkg/lyle"
	"github.com/spf13/cobra"
)

var (
	lWeight float32
	lLeanmass float32
	lPlan string
	lSex string
)

func init() {
 	RootCmd.AddCommand(lyleCmd)
	lyleCmd.Flags().StringVarP(&lSex, "sex", "s", "", "Sex of the subject [ man | woman ]")
        lyleCmd.Flags().Float32VarP(&lLeanmass, "lean-mass", "l", 0, "Current lean mass weight in kg")
        lyleCmd.Flags().StringVarP(&lPlan, "plan", "p", "", "Objetive of this diet [ bulk | cut | maint ]")
	lyleCmd.Flags().Float32VarP(&lWeight, "weight", "w", 0, "Current weight in kg")
}

var lyleCmd = &cobra.Command{
 	Use:   "lyle",
  	Short: "Lyle McDonald formula",
  	Long:  `Calculate kcals and macronutrients using Lyle McDonald formula`,

 	Run: func(cmd *cobra.Command, args []string) {
		lyle.InitLyle(&lWeight, &lSex, &lPlan, &lLeanmass)
  	},
}
