// Copyright © 2017 Borja Aranda <borja@redhat.com>
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

package lyle

import (
	"fmt"
	"os"
	"kcal/pkg/common"
)

type person struct {
	sex string
	weight float32
	plan string
}

var kcal float32

func InitLyle(weight *float32, sex *string, plan *string, leanmass *float32) {
	
	lyle := person{}

	if *weight > 0 {
		lyle.weight = *weight
	} else {
		fmt.Println("Weight not defined")
		os.Exit(1)
	}

	if *leanmass == 0 || *leanmass >= *weight {
                fmt.Println("Lean mass not defined")
                os.Exit(1)
        }

	if *sex == "man" || *sex == "woman" {
		lyle.sex = *sex
	} else {
		fmt.Println("Sex not defined")
		os.Exit(1)
	}

	if *plan == "bulk" || *plan == "maint" || *plan == "cut" {
		lyle.plan = *plan
	} else {
		fmt.Println("Plan not recognized")
		os.Exit(1)
	}
	
	execLyle(&lyle, &kcal)
	fat, prot, ch := common.CalculateMacro(&kcal, leanmass)

	fmt.Println("[ Lyle Formulae ]")
	fmt.Printf("Total kcals:\t%.1f kcals\n", kcal)
        fmt.Printf("Fat intake:\t%.1f gr\n", fat)
	fmt.Printf("Prot intake:\t%.1f gr\n", prot)
        fmt.Printf("Carbs intake:\t%.1f gr\n", ch)
}

func execLyle(dataLyle *person, kcal *float32) {

	if dataLyle.sex == "man" {
		switch dataLyle.plan {
		  case "bulk":
			  *kcal = dataLyle.weight * 40
		  case "cut":
			  *kcal = dataLyle.weight * 24
		  case "maint":
			  *kcal = dataLyle.weight * 35
		  default:
		 }
	} else if dataLyle.sex == "woman" {
		switch dataLyle.plan {    
                  case "bulk":
			  *kcal = dataLyle.weight * 35
                  case "cut":
			  *kcal = dataLyle.weight * 22
                  case "maint":
			  *kcal = dataLyle.weight * 31
                  default:
		  }
	}
}
