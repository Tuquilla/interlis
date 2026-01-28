package v2_3

//func TestXtfReader(t *testing.T) {
//
//	// Prepare file
//	xtf, _ := os.Open("test.xtf")
//	defer xtf.Close()
//	decoder := xml.NewDecoder(xtf)
//
//	geometries := interlis2.ReadGeometry(decoder)
//	if geometries.Surfaces[0].Boundary.Polyline.Coords[0].C1 != "1" {
//		t.Errorf("C1 was not 1")
//	}
//	if geometries.Surfaces[0].Boundary.Polyline.Coords[0].C2 != "2" {
//		t.Errorf("C2 was not 2")
//	}
//	if geometries.Surfaces[0].Boundary.Polyline.Coords[0].C3 != "" {
//		t.Errorf("C3 was not empty")
//	}
//}
