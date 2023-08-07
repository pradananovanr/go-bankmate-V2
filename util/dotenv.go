/*
 * Author : Pradana Novan Rianto (https://github.com/pradananovanr)
 * Created on : Thu Apr 13 2023
 * Copyright : Pradana Novan Rianto © 2023. All rights reserved
 */

package util

import (
	"os"
)

func DotEnv(key string) string {
	return os.Getenv(key)
}
