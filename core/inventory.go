package core

import (
	"noansible/target"
)

//独立的invetory包？
func ReadInventoryYML(groupName string, filedir string) ([]target.Hostinfo, error) {
	var ivts []target.Hostinfo
	ivtbook, err := loadrawbook(filedir)

	if imps, ok := ivtbook[groupName]; ok && err == nil {
		impsv, _ := imps.([]interface{})
		for _, v := range impsv {
			vs, _ := v.(string)
			var hinfo target.Hostinfo
			hinfo.NewHost(vs)
			ivts = append(ivts, hinfo)

		}

	}
	return ivts, err
}
