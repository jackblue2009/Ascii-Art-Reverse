package pkg

/*
marge string to be color lis and color list into a map
*/

func ColorStringMap(strcol, color []string) map[string]string {
	csmap := make(map[string]string)
	for i := 0; i < len(color); i++ {
		csmap[strcol[i]] = color[i]
	}
	return csmap
}
