package main

func getBNK48Cap() string {
	return "Cherprang <3"
}

func getFullNameOfBnk48Member(n string) string {
	bnkMember := map[string]string{
		"CHERPRANG": "CHERPRANG AREEKUL",
		"ORN":       "PATCHANAN JIAJIRACHOTE",
		"MOBILE":    "PIMRAPAT PHADUNGWATANACHOK",
	}

	result := ""
	for nickName, fullName := range bnkMember {
		if nickName == n {
			result = fullName
		}
	}
	return result
}
