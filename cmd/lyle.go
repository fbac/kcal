package cmd

import (
	"fmt"
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
		fmt.Printf("Lyle Formulae\n\nData input:\n- sex:\t\t%v\n- plan:\t\t%v\n- weight:\t%.2fkg\n- lean mass:\t%.2f kg\n\n", lSex, lPlan, lWeight, lLeanmass)
		lyle.InitLyle(&lWeight, &lSex, &lPlan, &lLeanmass)
  	},
}
