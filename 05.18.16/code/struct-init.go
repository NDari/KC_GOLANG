// Bad
flower = new(Flower)
flower.PetalCount = 6
flower.Name = "Yucca"
flower.Family = "agave"
flower.Season = "summer"
// Bad
flower = &Flower{
    PetalCount: 6,
    Name: "Yucca",
    Family: "agave",
    Season: "summer",
}
// Good
func NewFlower(petalCount int, name, family, season string) *Flower{
    return &Flower{
        PetalCount: petalCount,
        Name: name,
        Family: family,
        Season: season,
    }
}