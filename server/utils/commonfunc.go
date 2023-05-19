package utils



func IsAuthorized(role string)bool{

	authArr := []string{"HRM","PA","HR"}

	for _,v :=range authArr{

		if v==role{
			return true
			
		}
	}

	return false;
}