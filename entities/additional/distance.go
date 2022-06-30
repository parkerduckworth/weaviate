//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package additional

func CertaintyToDist(maybeCertainty *float64) (distPtr *float64) {
	if maybeCertainty != nil {
		dist := 1 - *maybeCertainty
		distPtr = &dist
	}
	return
}
