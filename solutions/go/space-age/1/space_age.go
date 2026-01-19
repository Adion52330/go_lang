package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	if planet == "Mercury" {
        age := seconds / (31557600*0.2408467)
		return age
    } else if planet =="Venus" {
        age := seconds / (31557600*0.61519726)
		return age
    } else if planet =="Earth" {
        age := seconds / (31557600*1.0)
		return age
    } else if planet =="Mars" {
        age := seconds / (31557600*1.8808158)
		return age
    } else if planet =="Jupiter" {
        age := seconds / (31557600*11.862615)
		return age
    } else if planet =="Saturn" {
        age := seconds / (31557600*29.447498)
		return age
    } else if planet =="Uranus" {
        age := seconds / (31557600*84.016846)
		return age
    } else if planet =="Neptune" {
        age := seconds / (31557600*164.79132)
		return age
    }
    return -1
}
