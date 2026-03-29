package main



func Hashcode2(dec string) string {
	size := len(dec)
	hashed := ""

	for _, r := range dec {
		hash := (int(r) + size) % 126
			if hash < 32 || hash > 126{
				hash += 33
		}
		hashed +=string(hash)
	}
	return hashed
	
}
