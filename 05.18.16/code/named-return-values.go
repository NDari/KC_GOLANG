// Bad
func GetLatLon(zip int) (int, int) {
    lat := latMagic(zip)
    lon := lonMagic(zip)
    return lat, lon
}
// Good
func GetLatLon(zip int) (lat int, lon int) {
    lat = latMagic(zip)
    lon = lonMagic(zip)
    return lat, lon
}
// Good
type LatLon struct {
    Lat int
    Lon int
}
func GetLatLon(zip int) *LatLon {
    return &LatLon{Lat: latMagic(zip), Lon: lonMagic(zip)}
}