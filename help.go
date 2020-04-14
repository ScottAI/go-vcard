package go_vcard

func MatrixToString(vals [][]string) string {
	end := ""
	for _,line := range vals{
		if end != ""{
			end += ";"
		}
		for _,col := range line{
			if end != "" {
				end += "," + col
			}else {
				end = col
			}

		}

	}
	return end
}
