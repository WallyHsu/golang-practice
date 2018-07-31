package rocket

type RocketInterface interface {
	Flight()
}

func Start(r RocketInterface) {
	r.Flight()
}
